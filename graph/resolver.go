package graph

import (
	"github.com/JitenMobile/graphql-mvp/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TrackService *services.TrackService
}

func NewResolver() *Resolver {
	return &Resolver{
		TrackService: services.NewTrackService(),
	}
}
