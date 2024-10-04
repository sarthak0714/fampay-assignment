package store

import (
	"github.com/sarthak0714/fampay-assignment/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// created an interface for abstaction
type Store interface {
	Init() error
	SaveVideo(*models.Video) error
	GetVideos(int, int) (*models.Video, error)
}

// the db struct is made private and only the storage interface is accessable
type postgresDB struct {
	db *gorm.DB
}

func NewPostgresStore(connectionString string) (*postgresDB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &postgresDB{
		db: db,
	}, nil
}

func (pgdb *postgresDB) Init() error {
	err := pgdb.db.AutoMigrate(&models.Video{})
	if err != nil {
		return err
	}
	return nil
}

func (pgdb *postgresDB) SaveVideo(video *models.Video) error {
	return pgdb.db.Create(video).Error
}

func (pgdb *postgresDB) GetVideos(limit, offset int) ([]models.Video, error) {
	var videos []models.Video
	err := pgdb.db.Order("published_at DESC").Limit(limit).Offset(offset).Find(&videos).Error
	return videos, err
}
