package spotify

type SearchResponseItemType interface {
	SearchItemAlbum | SearchItemArtist | SearchItemPlaylist | SearchItemTrack
}

type BaseResponse struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type SearchResponse[T SearchResponseItemType] struct {
	BaseResponse
	Items []T `json:"items"`
}

type CombinedResponse struct {
	Tracks    SearchResponse[SearchItemTrack]    `json:"tracks"`
	Artists   SearchResponse[SearchItemArtist]   `json:"artists"`
	Albums    SearchResponse[SearchItemAlbum]    `json:"albums"`
	Playlists SearchResponse[SearchItemPlaylist] `json:"playlists"`
}

type SearchItemAlbum struct {
	Href        string   `json:"href"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Images      []Images `json:"images"`
	Type        string   `json:"type"`
	ReleaseDate string   `json:"release_date"`
	Artists     []struct {
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"artists"`
}

type SearchItemArtist struct {
	Href   string   `json:"href"`
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Images []Images `json:"images"`
	Type   string   `json:"type"`
}

type SearchItemPlaylist struct {
	Href   string   `json:"href"`
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Images []Images `json:"images"`
	Type   string   `json:"type"`
	Owner  struct {
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	Tracks struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"tracks"`
}

type Playlist struct {
	SearchItemPlaylist
	Description string `json:"description"`
	Tracks      struct {
		BaseResponse
		Items []struct {
			AddedAt string `json:"added_at"`
			Track   SearchItemTrack
		} `json:"items"`
	} `json:"tracks"`
}

type SearchItemTrack struct {
	Href     string             `json:"href"`
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Type     string             `json:"type"`
	Album    SearchItemAlbum    `json:"album"`
	DiscNo   int                `json:"disc_number"`
	TrackNo  int                `json:"track_number"`
	Duration int                `json:"duration_ms"`
	Artists  []SearchItemArtist `json:"artists"`
}

type Images struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
