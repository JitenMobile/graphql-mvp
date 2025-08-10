package services

import (
	"encoding/json"
	"net/http"

	"github.com/JitenMobile/graphql-mvp/graph/model"
)

type ModuleSercve struct {
	baseURL    string
	httpClient *http.Client
}

func (service *ModuleSercve) GetModuleContents(trackId string) (*model.Module, error) {
	url := service.baseURL + "module/" + trackId
	resp, err := service.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	var module *model.Module
	if err := json.NewDecoder(resp.Body).Decode(&module); err != nil {
		return nil, err
	}
	return module, nil
}

func NewModuleService() *ModuleSercve {
	return &ModuleSercve{
		baseURL:    "https://odyssey-lift-off-rest-api.herokuapp.com/",
		httpClient: &http.Client{},
	}
}
