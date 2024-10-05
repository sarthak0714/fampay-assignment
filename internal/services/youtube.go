package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sarthak0714/fampay-assignment/internal/models"
	"github.com/sarthak0714/fampay-assignment/internal/store"
	"github.com/sarthak0714/fampay-assignment/pkg/utils"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// interface for abstraction
type YouTubeVideoService interface {
	SearchVideos(string) ([]*models.Video, error)
	FetchVideosWorker(string, time.Duration, store.Store)
}

type youTubeVideoService struct {
	youtubeApiService *youtube.Service
	store             store.Store
	apiKeys           []string
	currKeyIdx        int
}

// Constructs a new YouTubeVideoService instance with the provided API keys and store.
func NewService(keys []string, store store.Store) (*youTubeVideoService, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("no API keys found")
	}
	svc, err := youtube.NewService(context.Background(), option.WithAPIKey(keys[0]))
	if err != nil {
		return nil, err
	}

	return &youTubeVideoService{
		youtubeApiService: svc,
		store:             store,
		apiKeys:           keys,
		currKeyIdx:        0,
	}, nil
}

/*
Searches for videos on YouTube based on the provided query string and returns a list of videos.
refrence https://developers.google.com/youtube/v3/docs/search/list#go
*/

func (s *youTubeVideoService) SearchVideos(q string) ([]*models.Video, error) {
	apiCall := s.youtubeApiService.Search.List([]string{"id,snippet"}).
		Q(q).
		Type("video").
		Order("date").
		PublishedAfter(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)).
		MaxResults(50)

	res, err := apiCall.Do()
	if err != nil {
		if gapierr, ok := err.(*googleapi.Error); ok && gapierr.Code == 403 {
			s.nextApiIdx()
			return s.SearchVideos(q)
		}
		return nil, err
	}
	var videos []*models.Video
	for _, item := range res.Items {
		timeStr, err := utils.ConvertStrToTime(item.Snippet.PublishedAt)
		if err == nil {
			vid := &models.Video{
				Id:           item.Id.VideoId,
				Title:        item.Snippet.Title,
				Description:  item.Snippet.Description,
				PublishedAt:  timeStr,
				ThumbnailURL: item.Snippet.Thumbnails.Default.Url,
			}
			videos = append(videos, vid)
		}

	}
	return videos, nil
}

// Background Worker to periodically fetches videos based on the search query and saves them to the store.
func (s *youTubeVideoService) FetchVideosWorker(query string, interval time.Duration, store store.Store) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		videos, err := s.SearchVideos(query)
		if err != nil {
			log.Printf("Error while fetching video: %v", err)
		}
		for _, video := range videos {
			if err := store.SaveVideo(video); err != nil {
				log.Printf("Error while Saving video: %v", err)
			}
		}
		utils.FetchLogger()
	}
}

// changes the api key to next in the list
func (s *youTubeVideoService) nextApiIdx() {
	s.currKeyIdx = (s.currKeyIdx + 1) % len(s.apiKeys)
	s.youtubeApiService.BasePath = fmt.Sprintf("https://www.googleapis.com/youtube/v3?key=%s", s.apiKeys[s.currKeyIdx])
}
