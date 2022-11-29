package main

import (
	"html/template"
	"io"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track, w io.Writer) {
	var templ = template.Must(template.New("tracklist").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>テスト</title>
<style type="text/css">
th, tr, td {
	border: 1px solid black
}
</style>
</head>
<body>
<table>
<tr>
	<th><a href="/Title">Title</a></th>
	<th><a href="/Artist">Artist</a></th>
	<th><a href="/Album">Album</a></th>
	<th><a href="/Year">Year</a></th>
	<th><a href="/Length">Length</a></th>
</tr>
{{range .}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

	templ.Execute(w, tracks)
}

type customSort struct {
	t         []*Track
	firstKey  string
	secondKey string
}

func (x customSort) Len() int { return len(x.t) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x customSort) Less(i, j int) bool {
	firstKeyFunc := selectSortFunc(x.firstKey)
	res := firstKeyFunc(x.t[i], x.t[j])

	if res {
		secondKeyFunc := selectSortFunc(x.secondKey)
		res = secondKeyFunc(x.t[i], x.t[j])
	}

	return res
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/Title" {
		sort.Sort(customSort{tracks, "Title", "Artist"})
	}
	if r.URL.Path == "/Artist" {
		sort.Sort(customSort{tracks, "Artist", "Album"})
	}
	if r.URL.Path == "/Album" {
		sort.Sort(customSort{tracks, "Album", "Year"})
	}
	if r.URL.Path == "/Year" {
		sort.Sort(customSort{tracks, "Year", "Length"})
	}
	if r.URL.Path == "/Length" {
		sort.Sort(customSort{tracks, "Length", "Year"})
	}
	printTracks(tracks, w)
}

func selectSortFunc(funcName string) func(x, y *Track) bool {
	if funcName == "Title" {
		return func(x, y *Track) bool {
			return x.Title <= y.Title
		}
	}
	if funcName == "Artist" {
		return func(x, y *Track) bool {
			return x.Artist <= y.Artist
		}
	}
	if funcName == "Album" {
		return func(x, y *Track) bool {
			return x.Album <= y.Album
		}
	}
	if funcName == "Year" {
		return func(x, y *Track) bool {
			return x.Year <= y.Year
		}
	}
	if funcName == "Length" {
		return func(x, y *Track) bool {
			return x.Length <= y.Length
		}
	}

	return func(x, y *Track) bool {
		panic("invalid function name")
	}
}
