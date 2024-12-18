package spotify

type ResponseItemType interface {
	ResponseItemAlbum | ResponseItemArtist | ResponseItemPlaylist | ResponseItemTrack
}

type SearchResponse[T ResponseItemType] struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []T    `json:"items"`
}

type CombinedResponse struct {
	Tracks    SearchResponse[ResponseItemTrack]    `json:"tracks"`
	Artists   SearchResponse[ResponseItemArtist]   `json:"artists"`
	Albums    SearchResponse[ResponseItemAlbum]    `json:"albums"`
	Playlists SearchResponse[ResponseItemPlaylist] `json:"playlists"`
}

type ResponseItemAlbum struct {
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

type ResponseItemArtist struct {
	Href   string   `json:"href"`
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Images []Images `json:"images"`
	Type   string   `json:"type"`
}

type ResponseItemPlaylist struct {
	Href   string   `json:"href"`
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Images []Images `json:"images"`
	Type   string   `json:"type"`
	Owner  struct {
		DisplayName string `json:"display_name"`
	} `json:"owner"`
}

type ResponseItemTrack struct {
	Href    string `json:"href"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Artists []struct {
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"artists"`
}

type Images struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
