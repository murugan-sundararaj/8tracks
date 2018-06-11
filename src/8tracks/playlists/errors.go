package playlists

import "errors"

var (
	ErrInvalidPlaylist = errors.New("one or more invalid playlist found")
	ErrPlaylistExist   = errors.New("playlist name already exist")
)
