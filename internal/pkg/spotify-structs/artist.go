package spotify_structs

import pb "hihoak/api-service/pkg/pb/hihoak/api-service"

type SpotifyItemsArtist struct {
	Items []SpotifyArtist `json:"items"`
}

type SpotifyArtist struct {
	Followers artistsFollowers `json:"followers"`
	Genres []string `json:"genres"`
	Id string `json:"id"`
	Images []artistsImage `json:"images"`
	Name string `json:"name"`
	Popularity int64 `json:"popularity"`
	Uri string `json:"uri"`
}

type artistsFollowers struct {
	Href string `json:"href"`
	Total int64 `json:"total"`
}

type artistsImage struct {
	Url string `json:"url"`
	Height int64 `json:"height"`
	Width int64 `json:"width"`
}

func FromStructToPb(artist SpotifyArtist) *pb.SpotifyArtist {
	return &pb.SpotifyArtist{
		Followers: artist.Followers.Total,
		Genres: artist.Genres,
		Name: artist.Name,
		Uri: artist.Uri,
	}
}
