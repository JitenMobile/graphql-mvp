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

	defer resp.Body.Close()

	var tracks []*model.Track
	if err := json.NewDecoder(resp.Body).Decode(&tracks); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (service *TrackService) GetTrackByID(id string) (*model.Track, error) {
	url := service.baseURL + "track/" + id
	resp, err := service.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var track model.Track
	if err := json.NewDecoder(resp.Body).Decode(&track); err != nil {
		return nil, err
	}
	return &track, nil
}

func (service *TrackService) GetAuthor(authorId string) (*model.Author, error) {
	url := service.baseURL + "author/" + authorId
	resp, err := service.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var author model.Author
	if err := json.NewDecoder(resp.Body).Decode(&author); err != nil {
		return nil, err
	}
	return &author, nil
}

func (service *TrackService) GetModuleContents(trackId string) ([]*model.Module, error) {
	url := service.baseURL + "track/" + trackId + "/modules"
	resp, err := service.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	var modules []*model.Module
	if err := json.NewDecoder(resp.Body).Decode(&modules); err != nil {
		return nil, err
	}
	return modules, nil
}

func (service *TrackService) IncrementTrackViews(id string) (*model.IncrementTrackViewsResponse, error) {
	url := service.baseURL + "track/" + id + "/numberOfViews"
	req, err := http.NewRequest(http.MethodPatch, url, nil)

	if err != nil {
		return &model.IncrementTrackViewsResponse{
			Code:    404,
			Success: false,
			Message: err.Error(),
			Track:   nil,
		}, err
	}

	// execute the request
	resp, err := service.httpClient.Do(req)
	if err != nil {
		return &model.IncrementTrackViewsResponse{
			Code:    404,
			Success: false,
			Message: err.Error(),
			Track:   nil,
		}, err
	}
	defer resp.Body.Close()

	var track model.Track
	if err := json.NewDecoder(resp.Body).Decode(&track); err != nil {
		return &model.IncrementTrackViewsResponse{
			Code:    404,
			Success: false,
			Message: err.Error(),
			Track:   nil,
		}, err
	}
	return &model.IncrementTrackViewsResponse{
		Code:    200,
		Success: true,
		Message: "Successfully incremented number of views for track" + id,
		Track:   &track,
	}, nil
}
