package playlists

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type playlistItem struct {
	PlaylistName string
	Tracks       map[string]*Track
	Creator      Creator
	CreatedAt    time.Time
	Counter      Counter
}

type DAL struct {
	playlists        map[string]*playlistItem
	playlistNameToID map[string]string
}

// NewDAL creates and return a new DAL object
func NewDAL() *DAL {
	return &DAL{
		playlists:        map[string]*playlistItem{},
		playlistNameToID: map[string]string{},
	}
}

func (d *DAL) storePlaylist(ctx context.Context, p *Playlist) error {
	tracks := map[string]*Track{}
	for _, track := range p.Tracks {
		tracks[track.ID] = track
	}

	item := playlistItem{
		PlaylistName: p.PlaylistName,
		Creator:      p.Creator,
		Counter:      p.Counter,
		Tracks:       tracks,
		CreatedAt:    time.Now(),
	}

	d.playlists[p.PlaylistID] = &item
	d.playlistNameToID[p.PlaylistName] = p.PlaylistID

	return nil
}

func (d *DAL) createPlaylist(ctx context.Context, p *Playlist) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "couldn't generate playlist id")
	}

	p.PlaylistID = string(fmt.Sprintf("%s", id))
	if err := d.storePlaylist(ctx, p); err != nil {
		return "", errors.Wrap(err, "couldn't store playlist")
	}

	return p.PlaylistID, nil
}

func (d *DAL) loadPlaylist(ctx context.Context, playlistIDs []string, playlistNames []string) ([]*Playlist, error) {
	for _, name := range playlistNames {
		playlistID, ok := d.playlistNameToID[name]
		if !ok {
			return nil, ErrInvalidPlaylist
		}
		playlistIDs = append(playlistIDs, playlistID)
	}

	res := []*Playlist{}
	for _, playlistID := range playlistIDs {
		item, ok := d.playlists[playlistID]
		if !ok {
			return nil, ErrInvalidPlaylist
		}
		tracks := []*Track{}
		for _, track := range item.Tracks {
			tracks = append(tracks, track)
		}
		playlist := Playlist{
			PlaylistID:   playlistID,
			PlaylistName: item.PlaylistName,
			Tracks:       tracks,
			Creator:      item.Creator,
			CreatedAt:    item.CreatedAt,
			Counter:      item.Counter,
		}
		res = append(res, &playlist)
	}
	return res, nil
}

func (d *DAL) upsertPlaylist(ctx context.Context, p *Playlist) error {
	// remove the existing reference
	if _, ok := d.playlists[p.PlaylistID]; ok {
		delete(d.playlistNameToID, d.playlists[p.PlaylistID].PlaylistName)
	}

	if err := d.storePlaylist(ctx, p); err != nil {
		return errors.Wrap(err, "couldn't store playlist")
	}
	return nil
}

func (d *DAL) updatePlaylistName(ctx context.Context, playlistID, playlistName string) error {
	if _, ok := d.playlists[playlistID]; !ok {
		return ErrInvalidPlaylist
	}
	// remove the existing reference
	delete(d.playlistNameToID, d.playlists[playlistID].PlaylistName)

	d.playlists[playlistID].PlaylistName = playlistName
	d.playlistNameToID[playlistName] = playlistID
	return nil
}

func (d *DAL) removePlaylist(ctx context.Context, playlistID string) error {
	if _, ok := d.playlists[playlistID]; !ok {
		return ErrInvalidPlaylist
	}
	delete(d.playlists, playlistID)
	return nil
}

func (d *DAL) addTrack(ctx context.Context, playlistID string, track *Track) error {
	if _, ok := d.playlists[playlistID]; !ok {
		return ErrInvalidPlaylist
	}
	d.playlists[playlistID].Tracks[track.ID] = track
	return nil
}

func (d *DAL) removeTrack(ctx context.Context, playlistID, trackID string) error {
	if _, ok := d.playlists[playlistID]; !ok {
		return ErrInvalidPlaylist
	}
	delete(d.playlists[playlistID].Tracks, trackID)
	return nil
}

func (d *DAL) plays(ctx context.Context, playlistID string) error {
	if _, ok := d.playlists[playlistID]; !ok {
		return ErrInvalidPlaylist
	}
	d.playlists[playlistID].Counter.NumberOfPlays++
	return nil
}

func (d *DAL) likes(ctx context.Context, playlistID string) error {
	if _, ok := d.playlists[playlistID]; !ok {
		return ErrInvalidPlaylist
	}
	d.playlists[playlistID].Counter.NumberOfLikes++
	return nil
}

func (d *DAL) dislikes(ctx context.Context, playlistID string) error {
	if _, ok := d.playlists[playlistID]; !ok {
		return ErrInvalidPlaylist
	}
	d.playlists[playlistID].Counter.NumberOfLikes--
	if d.playlists[playlistID].Counter.NumberOfLikes < 0 {
		d.playlists[playlistID].Counter.NumberOfLikes = 0
	}
	return nil
}
