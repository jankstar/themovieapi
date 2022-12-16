package themovieapi

import (
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetGengre(t *testing.T) {
	type args struct {
		apiKey string
		langu  string
	}
	type tyTests struct {
		name    string
		args    args
		want    TypGengres
		wantErr bool
	}
	var tests []tyTests
	{
		var myArgs args
		var myTests tyTests
		myArgs.apiKey = ""
		myArgs.langu = "de-DE"
		myTests.name = "Liefert die Gengre"
		myTests.args = myArgs
		myTests.want.Genres = []EleGenres{
			{28, "Action"}, {12, "Abenteuer"}, {16, "Animation"}, {35, "Komödie"}, {80, "Krimi"},
			{99, "Dokumentarfilm"}, {18, "Drama"}, {10751, "Familie"},
			{14, "Fantasy"}, {36, "Historie"}, {27, "Horror"}, {10402, "Musik"}, {9648, "Mystery"},
			{10749, "Liebesfilm"}, {878, "Science Fiction"}, {10770, "TV-Film"}, {53, "Thriller"},
			{10752, "Kriegsfilm"}, {37, "Western"}}
		tests = append(tests, myTests)
	}

	//apikey for the movie DB
	godotenv.Load()
	if os.Getenv("apikey") != "" {
		tests[0].args.apiKey = os.Getenv("apikey")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGengre(tt.args.apiKey, tt.args.langu)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGengre() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGengre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ChangeUmlauteAll(t *testing.T) {
	type args struct {
		iStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Testet german umlaute",
			args{"Die Schöne und das Biest"},
			"Die Schone und das Biest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeUmlauteAll(tt.args.iStr); got != tt.want {
				t.Errorf("changeUmlauteAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMovieDetail(t *testing.T) {
	type args struct {
		apiKey  string
		langu   string
		movieID int
	}
	tests := []struct {
		name    string
		args    args
		want    MovieDetailResponse
		wantErr bool
	}{
		{name: "Test FilmDetail",
			args: args{"",
				"de-DE",
				335797},
			want: MovieDetailResponse{
				OriginalTitle: "Sing",
			},
			wantErr: false,
		},
	}

	//apikey for the movie DB
	godotenv.Load()
	if os.Getenv("apikey") != "" {
		tests[0].args.apiKey = os.Getenv("apikey")
	}

	//tests.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMovieDetail(tt.args.apiKey, tt.args.langu, tt.args.movieID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovieDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.OriginalTitle != tt.want.OriginalTitle {
				t.Errorf("GetMovieDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSearchMovie(t *testing.T) {
	type args struct {
		apiKey string
		langu  string
		query  string
		page   int
	}
	type tyResults []struct {
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
	}
	tests := []struct {
		name    string
		args    args
		want    SearchMovieResponse
		wantErr bool
	}{
		{name: "Suche Film",
			args: args{
				apiKey: "",
				langu:  "de-DE",
				query:  "Sing 2",
				page:   1,
			},
			want: SearchMovieResponse{
				Results: tyResults{{
					OriginalTitle: "Sing 2",
				}},
			},
			wantErr: false,
		},
	}

	//apikey for the movie DB
	godotenv.Load()
	if os.Getenv("apikey") != "" {
		tests[0].args.apiKey = os.Getenv("apikey")
	}

	tests[0].want.Results = append(tests[0].want.Results)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSearchMovie(tt.args.apiKey, tt.args.langu, tt.args.query, "", tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Results[0].OriginalTitle != tt.want.Results[0].OriginalTitle {
				t.Errorf("GetSearchMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}
