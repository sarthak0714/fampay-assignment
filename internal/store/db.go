package store

import (
	"github.com/sarthak0714/fampay-assignment/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// created an interface for abstaction
type Store interface {
	Init() error
	SaveVideo(*models.Video) error
	GetVideos(int, int) ([]models.Video, error)
}

// the db struct is made private and only the storage interface is accessable
type sqliteDB struct {
	db *gorm.DB
}

func NewPostgresStore(connectionString string) (*sqliteDB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &sqliteDB{
		db: db,
	}, nil
}

func (sqldb *sqliteDB) Init() error {
	err := sqldb.db.AutoMigrate(&models.Video{})
	if err != nil {
		return err
	}
	return nil
}

func (sqldb *sqliteDB) SaveVideo(video *models.Video) error {
	return sqldb.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "description", "published_at", "thumbnail_url"}),
	}).Create(video).Error
}

func (sqldb *sqliteDB) GetVideos(limit, offset int) ([]models.Video, error) {
	var videos []models.Video
	err := sqldb.db.Order("published_at DESC").Limit(limit).Offset(offset).Find(&videos).Error
	return videos, err
}
