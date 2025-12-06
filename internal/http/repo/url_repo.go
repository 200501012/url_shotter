package repo

import (
	"errors"
	"url_shotter/internal/http/domain"

	"gorm.io/gorm"
)

type URLRepository interface {
	Create(url *domain.URL) error
	FindByShortCode(shortCode string) (*domain.URL, error)
	FindByID(id uint) (*domain.URL, error)
	Update(url *domain.URL) error
	Delete(id uint) error
	GetAll() ([]domain.URL, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{
		db: db,
	}
}

func (instance *urlRepository) Create(url *domain.URL) error {

	if err := instance.db.Create(url).Error; err != nil {
		return err
	}

	return nil
}

func (instance *urlRepository) FindByShortCode(shortCode string) (*domain.URL, error) {
	var url domain.URL
	if err := instance.db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("URL não encontrada")
		}
		return nil, err
	}
	return &url, nil
}

func (instance *urlRepository) FindByID(id uint) (*domain.URL, error) {
	var url domain.URL

	if err := instance.db.First(&url, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("URL não encontrada")
		}
		return nil, err
	}
	return &url, nil
}

func (instance *urlRepository) Update(url *domain.URL) error {
	if err := instance.db.Save(url).Error; err != nil {
		return err
	}
	return nil
}

func (instance *urlRepository) Delete(id uint) error {
	if err := instance.db.Delete(&domain.URL{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (instance *urlRepository) GetAll() ([]domain.URL, error) {
	var urls []domain.URL
	if err := instance.db.Find(&urls).Error; err != nil {
		return nil, err
	}

	return urls, nil
}
