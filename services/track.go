package services

import (
	"encoding/json"
	"net/http"

	"github.com/JitenMobile/graphql-mvp/graph/model"
)

type TrackService struct {
	baseURL    string
	httpClient *http.Client
}

func NewTrackService() *TrackService {
	return &TrackService{
		baseURL:    "https://odyssey-lift-off-rest-api.herokuapp.com/",
		httpClient: &http.Client{},
	}
}

func (service *TrackService) GetTracksForHome() ([]*model.Track, error) {
	url := service.baseURL + "tracks"
	resp, err := service.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	var tracks []*model.Track
	if err := json.NewDecoder(resp.Body).Decode(&tracks); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (service *TrackService) GetAuthor(authorId string) (*model.Author, error) {
	url := service.baseURL + "author/" + authorId
	resp, err := service.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	var author model.Author
	if err := json.NewDecoder(resp.Body).Decode(&author); err != nil {
		return nil, err
	}
	return &author, nil
}
