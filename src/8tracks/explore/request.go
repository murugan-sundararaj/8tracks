package explore

import (
	"8tracks/playlists"
	"8tracks/tags"
)

type ExploreRequest struct {
	TagNames []string
}
type ExploreResponse struct {
	Tags      []*tags.Tag           `json:"tags,omitempty"`
	Playlists []*playlists.Playlist `json:"playlists,omitempty"`
	Err       string                `json:"error,omitempty"`
}

func (e *ExploreResponse) error() string {
	return e.Err
}
