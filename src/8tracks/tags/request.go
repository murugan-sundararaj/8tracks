package tags

// Type represents type of the tag
type Type int

const (
	// Artist tag
	Artist Type = 1 << iota
	// Mood tag
	Mood
	// Genre tag
	Genre
)

var typeNames = []string{"", "Artist", "Mood", "Genre"}

func (t Type) String() string {
	return typeNames[t]
}

type TypeName struct {
	TagType  Type   `json:"tag_type"`
	TypeName string `json:"type_name"`
}

func loadTypes() []TypeName {
	res := []TypeName{}
	for i := 1; i < len(typeNames); i++ {
		res = append(res, TypeName{
			TagType:  Type(i),
			TypeName: typeNames[i],
		})
	}
	return res
}

// Tag represents the tag object
type Tag struct {
	TagID   string `json:"tag_id,omitempty"`
	TagName string `json:"tag_name"`
	TagType Type   `json:"tag_type"`
}

type CreateTagRequest struct {
	*Tag
}
type CreateTagResponse struct {
	TagID string `json:"tag_id,omitempty"`
	Err   string `json:"error,omitempty"`
}

func (e *CreateTagResponse) error() string {
	return e.Err
}

type LoadTagRequest struct {
	TagIDs   []string `json:"tag_ids"`
	TagNames []string `json:"tag_names"`
}
type LoadTagResponse struct {
	Tags []*Tag `json:"tags,omitempty"`
	Err  string `json:"error,omitempty"`
}

func (e *LoadTagResponse) error() string {
	return e.Err
}

type UpsertTagRequest struct {
	*Tag
}
type UpsertTagResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *UpsertTagResponse) error() string {
	return e.Err
}

type UpdateTagRequest struct {
	*Tag
}
type UpdateTagResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *UpdateTagResponse) error() string {
	return e.Err
}

type RemoveTagRequest struct {
	TagID string `json:"tag_id"`
}
type RemoveTagResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *RemoveTagResponse) error() string {
	return e.Err
}

type LoadTagTypesRequest struct{}
type LoadTagTypesResponse struct {
	TagTypes []TypeName `json:"tag_types"`
	Err      string     `json:"error,omitempty"`
}

func (e *LoadTagTypesResponse) error() string {
	return e.Err
}

// AssignTagToPlaylistRequest tags a playlist_id to tag_id
type AssignTagToPlaylistRequest struct {
	TagID      string `json:"tag_id"`
	PlaylistID string `json:"playlist_id"`
}
type AssignTagToPlaylistResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *AssignTagToPlaylistResponse) error() string {
	return e.Err
}

// UnAssignTagFromPlaylistRequest untags a playlist_id to tag_id
type UnAssignTagFromPlaylistRequest struct {
	TagID      string `json:"tag_id"`
	PlaylistID string `json:"playlist_id"`
}
type UnAssignTagFromPlaylistResponse struct {
	Err string `json:"error,omitempty"`
}

func (e *UnAssignTagFromPlaylistResponse) error() string {
	return e.Err
}
