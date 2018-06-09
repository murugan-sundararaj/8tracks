package tags

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	set "gopkg.in/fatih/set.v0"
)

type DAL struct {
	tags           map[string]*Tag
	tagNameToID    map[string]string
	tagToPlaylists map[string]*set.Set
	playlistToTags map[string]*set.Set
}

// NewDAL creates and return a new DAL object
func NewDAL() *DAL {
	return &DAL{
		tags:           map[string]*Tag{},
		tagNameToID:    map[string]string{},
		tagToPlaylists: map[string]*set.Set{},
		playlistToTags: map[string]*set.Set{},
	}
}

func (d *DAL) storeTag(ctx context.Context, t *Tag) error {
	d.tags[t.TagID] = t
	d.tagNameToID[t.TagName] = t.TagID
	return nil
}

func (d *DAL) createTag(ctx context.Context, t *Tag) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "couldn't generate tag id")
	}
	t.TagID = string(fmt.Sprintf("%s", id))
	if err := d.storeTag(ctx, t); err != nil {
		return "", errors.Wrap(err, "couldn't store tag")
	}
	return t.TagID, nil
}

func (d *DAL) loadTag(ctx context.Context, tagIDs []string, tagNames []string) ([]*Tag, error) {
	for _, tagName := range tagNames {
		tagID, ok := d.tagNameToID[tagName]
		if !ok {
			return nil, ErrInvalidTag
		}
		tagIDs = append(tagIDs, tagID)
	}

	res := make([]*Tag, len(tagIDs))
	for i, tagID := range tagIDs {
		tag, ok := d.tags[tagID]
		if !ok {
			return nil, ErrInvalidTag
		}
		res[i] = tag
	}
	return res, nil
}

func (d *DAL) upsertTag(ctx context.Context, t *Tag) error {
	// remove the existing reference
	if tag, ok := d.tags[t.TagID]; ok {
		delete(d.tagNameToID, tag.TagName)
	}

	// add the new tag
	if err := d.storeTag(ctx, t); err != nil {
		return errors.Wrap(err, "couldn't store tag")
	}
	return nil
}

func (d *DAL) updateTag(ctx context.Context, t *Tag) error {
	if _, ok := d.tags[t.TagID]; !ok {
		return ErrInvalidTag
	}
	if t.TagName != "" {
		// remove the existing reference
		delete(d.tagNameToID, d.tags[t.TagID].TagName)
		// add
		d.tags[t.TagID].TagName = t.TagName
		d.tagNameToID[t.TagName] = t.TagID
	}
	if t.TagType != Type(0) {
		d.tags[t.TagID].TagType = t.TagType
	}
	return nil
}

func (d *DAL) removeTag(ctx context.Context, tagID string) error {
	if _, ok := d.tags[tagID]; !ok {
		return ErrInvalidTag
	}

	// remove the tags off the playlist
	if _, ok := d.tagToPlaylists[tagID]; ok {
		for _, playlistID := range d.tagToPlaylists[tagID].List() {
			if _, ok := d.playlistToTags[playlistID.(string)]; ok {
				d.playlistToTags[playlistID.(string)].Remove(tagID)
			}
		}
	}

	// remove the tag now
	delete(d.tags, tagID)
	return nil
}

func (d *DAL) tagPlayList(ctx context.Context, tagID string, playlistID string) error {
	if _, ok := d.tags[tagID]; !ok {
		return ErrInvalidTag
	}

	if _, ok := d.tagToPlaylists[tagID]; !ok {
		d.tagToPlaylists[tagID] = set.New()
	}

	d.tagToPlaylists[tagID].Add(playlistID)

	if _, ok := d.playlistToTags[playlistID]; !ok {
		d.playlistToTags[playlistID] = set.New()
	}
	d.playlistToTags[playlistID].Add(tagID)
	return nil
}

func (d *DAL) unTagPlayList(ctx context.Context, tagID string, playlistID string) error {
	if _, ok := d.tagToPlaylists[tagID]; !ok {
		return ErrInvalidTag
	}

	if _, ok := d.playlistToTags[playlistID]; ok {
		d.playlistToTags[playlistID].Remove(tagID)
	}

	d.tagToPlaylists[tagID].Remove(playlistID)
	return nil
}

func (d *DAL) loadTagPlaylistID(ctx context.Context, tagName string) (*set.Set, error) {
	tagID := d.tagNameToID[tagName]
	if _, ok := d.tagToPlaylists[tagID]; !ok {
		return set.New(), nil
	}
	return d.tagToPlaylists[tagID], nil
}

func (d *DAL) loadPlaylistTag(ctx context.Context, playlistID string) (*set.Set, error) {
	return d.playlistToTags[playlistID], nil
}
