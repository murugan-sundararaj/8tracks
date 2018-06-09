package playlists

import (
	"8tracks/tags"
	"time"
)

// Counter represents the playlist attributes that changes frequently
type Counter struct {
	NumberOfPlays int64 `json:"number_of_plays"`
	NumberOfLikes int64 `json:"number_of_likes"`
}

type Track struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Creator struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PlayList represents the play list object
type Playlist struct {
	PlaylistID   string      `json:"playlist_id,omitempty"`
	PlaylistName string      `json:"playlist_name"`
	Tags         []*tags.Tag `json:"tags,omitempty"`
	Tracks       []*Track    `json:"tracks,omitempty"`
	Creator      Creator     `json:"creator"`
	CreatedAt    time.Time   `json:"created_at"`
	Counter
}

type CreatePlaylistRequest struct {
	*Playlist
}
type CreatePlaylistResponse struct {
	PlaylistID string `json:"playlist_id,omitempty"`
	Err        string `json:"error,omitempty"`
}

func (e *CreatePlaylistResponse) error() string {
	return e.Err
}

type LoadPlaylistRequest struct {
	PlaylistIDs   []string `json:"playlist_ids"`
	PlaylistNames []string `json:"playlist_names"`
}
type LoadPlaylistResponse struct {
	Playlists []*Playlist `json:"playlists,omitempty"`
	Err       string      `json:"error,omitempty"`
}

func (e *LoadPlaylistResponse) error() string {
	return e.Err
}

type UpsertPlaylistRequest struct {
	*Playlist
}
type UpsertPlaylistResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *UpsertPlaylistResponse) error() string {
	return e.Err
}

type UpdatePlaylistNameRequest struct {
	PlaylistID   string `json:"playlist_id"`
	PlaylistName string `json:"playlist_name"`
}
type UpdatePlaylistNameResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *UpdatePlaylistNameResponse) error() string {
	return e.Err
}

type RemovePlaylistRequest struct {
	PlaylistID string `json:"playlist_id"`
}
type RemovePlaylistResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *RemovePlaylistResponse) error() string {
	return e.Err
}

type AddTrackRequest struct {
	PlaylistID string `json:"playlist_id"`
	*Track
}
type AddTrackResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *AddTrackResponse) error() string {
	return e.Err
}

type RemoveTrackRequest struct {
	PlaylistID string `json:"playlist_id"`
	TrackID    string `json:"track_id"`
}
type RemoveTrackResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *RemoveTrackResponse) error() string {
	return e.Err
}

type PlaysRequest struct {
	PlaylistID string `json:"playlist_id"`
}
type PlaysResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *PlaysResponse) error() string {
	return e.Err
}

type LikesRequest struct {
	PlaylistID string `json:"playlist_id"`
}
type LikesResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *LikesResponse) error() string {
	return e.Err
}

type DislikesRequest struct {
	PlaylistID string `json:"playlist_id"`
}
type DislikesResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *DislikesResponse) error() string {
	return e.Err
}
