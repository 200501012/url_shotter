#!/bin/bash
# Script para executar a imagem do Docker Hub

docker run -d \
  --name url_shotter_from_hub \
  -p 8080:8080 \
  -e DB_HOST=postgres \
  -e DB_PORT=5432 \
  -e DB_USER="eltoncavele8@gmail.com" \
  -e DB_PASSWORD="URLShortenerPass" \
  -e DB_NAME="URL_Shortener" \
  --network url_shotter_db_net \
  elton971/url_shotter:latest
