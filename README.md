# themovieapi
An GO API for accessing The Movie database

You need the APIKEY of The-Movie-Db to access it; get it via the URL "https://www.themoviedb.org/documentation/api" and put it into the .env file 
```
apikey="<APIKEY>"
```
or give them as parameters to the functions. 

This library uses the API 3 "https://developers.themoviedb.org/3/getting-started/introduction".

The following functions are implemented:
```
func GetSearchMovie(apiKey string, langu string, query string, page int) (SearchMovieResponse, error) {} 
func GetSearchTV(apiKey string, langu string, query string, page int) (SearchTVResponse, error) {}
func GetMovieDetail(apiKey string, langu string, movieID int) (MovieDetailResponse, error) {}
func GetGengre(apiKey string, langu string) (TypGengres, error) {}
```
also auxiliary functions e.g. for converting umlauts
```
func ChangeUmlauteAll(iStr string) string {}
func ChangeUmlauteSingle(iStr string, fromChar string, toChar string, runeLen int) string {}
func GetImageURL(imagePath string, size string) string {}

```
