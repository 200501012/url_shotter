#!/bin/bash
set -e

# Variáveis de ambiente com valores padrão
export DB_HOST=${DB_HOST:-localhost}
export DB_PORT=${DB_PORT:-5432}
export DB_USER=${DB_USER:-postgres}
export DB_PASSWORD=${DB_PASSWORD:-postgres}
export DB_NAME=${DB_NAME:-url_shotter}
export APP_PORT=${APP_PORT:-8080}
export PGADMIN_EMAIL=${PGADMIN_EMAIL:-admin@example.com}
export PGADMIN_PASSWORD=${PGADMIN_PASSWORD:-admin}
export PGADMIN_PORT=${PGADMIN_PORT:-8081}

echo "🚀 Iniciando todos os serviços..."

# Inicializar PostgreSQL se necessário
if [ ! -s "$PGDATA/PG_VERSION" ]; then
    echo "📦 Inicializando banco de dados PostgreSQL..."
    su-exec postgres initdb -D "$PGDATA" --auth-host=trust --auth-local=trust
    
    # Configurar PostgreSQL
    echo "host all all 0.0.0.0/0 md5" >> "$PGDATA/pg_hba.conf"
    echo "listen_addresses='*'" >> "$PGDATA/postgresql.conf"
    
    # Iniciar PostgreSQL temporariamente para criar usuário e banco
    su-exec postgres pg_ctl -D "$PGDATA" -o "-c listen_addresses='*'" -w start
    
    # Configurar senha do postgres e criar usuário/banco (usando socket local, sem senha)
    export PGHOST=/var/run/postgresql
    su-exec postgres psql -c "ALTER USER postgres WITH PASSWORD '$DB_PASSWORD';"
    
    if [ "$DB_USER" != "postgres" ]; then
        su-exec postgres psql -c "CREATE USER \"$DB_USER\" WITH PASSWORD '$DB_PASSWORD';" || true
        su-exec postgres psql -c "ALTER USER \"$DB_USER\" WITH SUPERUSER;" || true
    fi
    
    su-exec postgres psql -c "CREATE DATABASE \"$DB_NAME\" OWNER \"$DB_USER\";" || true
    unset PGHOST
    
    # Parar PostgreSQL temporário
    su-exec postgres pg_ctl -D "$PGDATA" -m fast -w stop
    
    # Atualizar pg_hba.conf para usar md5 após configurar senhas
    sed -i 's/trust/md5/g' "$PGDATA/pg_hba.conf"
fi

# Iniciar PostgreSQL em background
echo "🐘 Iniciando PostgreSQL..."
su-exec postgres postgres -D "$PGDATA" &
POSTGRES_PID=$!

# Aguardar PostgreSQL estar pronto
echo "⏳ Aguardando PostgreSQL iniciar..."
for i in {1..30}; do
    if su-exec postgres pg_isready -U postgres; then
        echo "✅ PostgreSQL está pronto!"
        break
    fi
    if [ $i -eq 30 ]; then
        echo "❌ Timeout aguardando PostgreSQL"
        exit 1
    fi
    sleep 1
done

# Criar usuário e banco se não existirem (usando socket local)
export PGHOST=/var/run/postgresql
export PGPASSWORD="$DB_PASSWORD"

if [ "$DB_USER" != "postgres" ]; then
    su-exec postgres psql -U postgres -c "SELECT 1 FROM pg_user WHERE usename='$DB_USER';" | grep -q 1 || \
        su-exec postgres psql -U postgres -c "CREATE USER \"$DB_USER\" WITH PASSWORD '$DB_PASSWORD';" || true
fi

su-exec postgres psql -U postgres -c "SELECT 1 FROM pg_database WHERE datname='$DB_NAME';" | grep -q 1 || \
    su-exec postgres psql -U postgres -c "CREATE DATABASE \"$DB_NAME\" OWNER \"$DB_USER\";" || true

unset PGHOST PGPASSWORD

# Iniciar aplicação Go
echo "🚀 Iniciando aplicação Go na porta $APP_PORT..."
export DB_HOST=localhost
export DB_PORT=5432
exec /app/app

