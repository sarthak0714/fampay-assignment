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

// this will return a new sqlite struct
func NewPostgresStore(connectionString string) (*sqliteDB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &sqliteDB{
		db: db,
	}, nil
}

// Initializes the database by migrating the Video model
func (sqldb *sqliteDB) Init() error {
	err := sqldb.db.AutoMigrate(&models.Video{})
	if err != nil {
		return err
	}
	return nil
}

// Saves a video to the database, updating it if it already exists. (Upsert because the vidoe thumbnail descrption and title may be updated)
func (sqldb *sqliteDB) SaveVideo(video *models.Video) error {
	return sqldb.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "description", "thumbnail_url"}),
	}).Create(video).Error
}

// Retrieves a list of videos from the database with pagination based on the limit and offset.
func (sqldb *sqliteDB) GetVideos(limit, offset int) ([]models.Video, error) {
	var videos []models.Video
	err := sqldb.db.Order("published_at DESC").Limit(limit).Offset(offset).Find(&videos).Error
	return videos, err
}
