package explore

import (
	"8tracks/playlists"
	"sort"
)

type playlistSlice []*playlists.Playlist

type ranker func(p playlistSlice) playlistSlice

var defaultRankingOrder ranker = func(p playlistSlice) playlistSlice {
	sort.Slice(p, func(i, j int) bool {
		// break tie. compare likes count
		if p[i].NumberOfPlays == p[j].NumberOfPlays {
			return p[i].NumberOfLikes > p[j].NumberOfLikes
		}
		return p[i].NumberOfPlays > p[j].NumberOfPlays
	})
	return p
}
