//Package themovieapi Includes functions to access TheMovie API
// for movies and TV
package themovieapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	urlMovie       string = "https://api.themoviedb.org/3/search/movie"
	urlMovieDetail string = "https://api.themoviedb.org/3/movie/"
	urlTV          string = "https://api.themoviedb.org/3/search/tv"
	urlGenre       string = "https://api.themoviedb.org/3/genre/movie/list"
	urlImage       string = "https://image.tmdb.org/t/p"
)

//SearchMovieResponse the return structure from the get /search/movie
type SearchMovieResponse struct {
	Page         int `json:"page,omitempty"`
	TotalResults int `json:"total_results,omitempty"`
	TotalPages   int `json:"total_pages,omitempty"`
	Results      []struct {
		Popularity       float32 `json:"popularity,omitempty"`
		ID               int     `json:"id,omitempty"`
		Video            bool    `json:"video,omitempty"`
		VoteCount        int     `json:"vote_count,omitempty"`
		VoteAverage      float32 `json:"vote_average,omitempty"`
		Title            string  `json:"title,omitempty"`
		ReleaseDate      string  `json:"release_date,omitempty"`
		OriginalLanguage string  `json:"original_language,omitempty"`
		OriginalTitle    string  `json:"original_title,omitempty"`
		GenreIds         []int   `json:"genre_ids,omitempty"`
		BackdropPath     string  `json:"backdrop_path,omitempty"`
		Adult            bool    `json:"adult,omitempty"`
		Overview         string  `json:"overview,omitempty"`
		PosterPath       string  `json:"poster_path,omitempty"`
	} `json:"results,omitempty"`
}

//MovieDetailResponse the return values to the detail of a movie
type MovieDetailResponse struct {
	Adult               bool   `json:"adult,omitempty"`
	BackdropPath        string `json:"backdrop_path,omitempty"`
	BelongsToCollection struct {
		ID           int    `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		PosterPath   string `json:"poster_path,omitempty"`
		BackdropPath string `json:"backdrop_path,omitempty"`
	} `json:"belongs_to_collection,omitempty"`
	Budget int `json:"budget,omitempty"`
	Genres []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"genres,omitempty"`
	Homepage            string  `json:"homepage,omitempty"`
	ID                  int     `json:"id,omitempty"`
	ImdbID              string  `json:"imdb_id,omitempty"`
	OriginalLanguage    string  `json:"original_language,omitempty"`
	OriginalTitle       string  `json:"original_title,omitempty"`
	Overview            string  `json:"overview,omitempty"`
	Popularity          float32 `json:"popularity,omitempty"`
	PosterPath          string  `json:"poster_path,omitempty"`
	ProductionCompanies []struct {
		ID            int    `json:"id,omitempty"`
		LogoPath      string `json:"logo_path,omitempty"`
		Name          string `json:"name,omitempty"`
		OriginCountry string `json:"origin_country,omitempty"`
	} `json:"production_companies,omitempty"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"production_countries,omitempty"`
	ReleaseDate     string `json:"release_date,omitempty"`
	Revenue         int    `json:"revenue,omitempty"`
	Runtime         int    `json:"runtime,omitempty"`
	SpokenLanguages []struct {
		Iso6391 string `json:"iso_639_1,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"spoken_languages,omitempty"`
	Status      string  `json:"status,omitempty"`
	Tagline     string  `json:"tagline,omitempty"`
	Title       string  `json:"title,omitempty"`
	Video       bool    `json:"video,omitempty"`
	VoteAverage float32 `json:"vote_average,omitempty"`
	VoteCount   int     `json:"vote_count,omitempty"`
}

//SearchTVResponse the return structure from the get /search/tv
type SearchTVResponse struct {
	Page         int `json:"page,omitempty"`
	TotalResults int `json:"total_results,omitempty"`
	TotalPages   int `json:"total_pages,omitempty"`
	Results      []struct {
		OriginalName     string   `json:"original_name,omitempty"`
		ID               int      `json:"id,omitempty"`
		Name             string   `json:"name,omitempty"`
		Popularity       float32  `json:"popularity,omitempty"`
		VoteCount        int      `json:"vote_count,omitempty"`
		VoteAverage      float32  `json:"vote_average,omitempty"`
		FirstAirDate     string   `json:"first_air_date,omitempty"`
		PosterPath       string   `json:"poster_path,omitempty"`
		GenreIds         []int    `json:"genre_ids,omitempty"`
		OriginalLanguage string   `json:"original_language,omitempty"`
		BackdropPath     string   `json:"backdrop_path,omitempty"`
		Overview         string   `json:"overview,omitempty"`
		OriginCountry    []string `json:"origin_country,omitempty"`
	} `json:"results,omitempty"`
}

//EleGenres  Genre element
type EleGenres struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//TypGengres Genre table
type TypGengres struct {
	Genres []EleGenres `json:"genres,omitempty"`
}

//ChangeUmlauteSingle replaces in a string an umlaut
func ChangeUmlauteSingle(iStr string, fromChar string, toChar string, runeLen int) string {
	rStr := iStr
	for strings.Contains(rStr, fromChar) {
		offset := strings.Index(rStr, fromChar)
		offsetP1 := offset + runeLen //umlaute are 2 characters
		if offset == 0 {
			rStr = toChar + rStr[offsetP1:]
		} else if offset > 0 {
			rStr = rStr[:offset] + toChar + rStr[offsetP1:]
		} else {
			break //Error and out
		}
	}
	return rStr
}

//ChangeUmlauteAll replaces all umlauts in a string
func ChangeUmlauteAll(iStr string) string {
	var convert = []struct {
		fromChar string
		toChar   string
		runeLen  int
	}{
		{"Ä", "Ae", 2},
		{"Ö", "Oe", 2},
		{"Ü", "Ue", 2},
		{"ä", "ae", 2},
		{"ö", "oe", 2},
		{"ü", "ue", 2},
		{"ß", "ss", 2},
		{"\x41\xcc\x88", "Ae", 3},
		{"\x61\xcc\x88", "ae", 3},
		{"\x4f\xcc\x88", "Oe", 3},
		{"\x6f\xcc\x88", "oe", 3},
		{"\x55\xcc\x88", "Ue", 3},
		{"\x75\xcc\x88", "ue", 3}}

	rStr := iStr

	for _, element := range convert {
		rStr = ChangeUmlauteSingle(rStr, element.fromChar, element.toChar, element.runeLen)
	}
	return rStr
}

//GetMovieDetail liefert Detailinfos zum Film
func GetMovieDetail(apiKey string, langu string, movieID int) (MovieDetailResponse, error) {
	var myurl string
	var dst MovieDetailResponse

	if langu == "" {
		langu = "de-DE"
	}

	myurl = urlMovieDetail + strconv.Itoa(movieID) + "?api_key=" + apiKey + "&language=" + langu

	res, err1 := http.Get(myurl)
	if err1 != nil || res.StatusCode != 200 {
		return dst, err1
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	//dec.DisallowUnknownFields()

	err2 := dec.Decode(&dst)
	if err2 != nil {
		return dst, err2
	}

	return dst, nil
}

//GetSearchMovie returns the search result for query in the form SearchMovieResponse
//the apikey and the language must be specified
func GetSearchMovie(apiKey string, langu string, query string) (SearchMovieResponse, error) {
	var myurl string
	var dst SearchMovieResponse
	dst.Page = 0
	dst.TotalPages = 0
	dst.TotalResults = 0
	dst.Results = nil

	if langu == "" {
		langu = "de-DE"
	}

	myurl = urlMovie + "?api_key=" + apiKey + "&language=" + langu + "&query=" + url.QueryEscape(ChangeUmlauteAll(query)) + "&page=1"

	res, err1 := http.Get(myurl)
	if err1 != nil || res.StatusCode != 200 {
		return dst, err1
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	//dec.DisallowUnknownFields()

	err2 := dec.Decode(&dst)
	if err2 != nil || dst.Page == 0 {
		return dst, err2
	}

	return dst, nil
}

//GetSearchTV returns the search result for query in the form SearchTVResponse
//you have to specify the apikey and the language
func GetSearchTV(apiKey string, langu string, query string) (SearchTVResponse, error) {
	var myurl string
	var dst SearchTVResponse
	dst.Page = 0
	dst.TotalPages = 0
	dst.TotalResults = 0
	dst.Results = nil

	if langu == "" {
		langu = "de-DE"
	}
	myurl = urlTV + "?api_key=" + apiKey + "&language=" + langu + "&query=" + url.QueryEscape(ChangeUmlauteAll(query)) + "&page=1"

	res, err1 := http.Get(myurl)
	if err1 != nil || res.StatusCode != 200 {
		return dst, err1
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()

	err2 := dec.Decode(&dst)
	if err2 != nil || dst.Page == 0 {
		return dst, err2
	}

	return dst, nil
}

//GetGengre provide the genre table
func GetGengre(apiKey string, langu string) (TypGengres, error) {
	var myurl string
	var dst TypGengres

	if langu == "" {
		langu = "de-DE"
	}
	myurl = urlGenre + "?api_key=" + apiKey + "&language=" + langu

	res, err1 := http.Get(myurl)
	if err1 != nil || res.StatusCode != 200 {
		return dst, err1
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	//dec.DisallowUnknownFields()

	err2 := dec.Decode(&dst)
	if err2 != nil || len(dst.Genres) == 0 {
		return dst, err2
	}

	return dst, nil
}

//GetImageURL returns the complete URL from the image path and the resolution
func GetImageURL(imagePath string, size string) string {
	if imagePath == "" {
		return ""
	}
	if !(size == "w200" || size == "w300" || size == "w400" || size == "w500") {
		size = "w500"
	}

	return urlImage + "/" + size + "/" + imagePath

}
