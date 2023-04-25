// Package version modelo de version
package version

type Version struct {
	Title      string `json:"title"`
	Version    string `json:"version"`
	Build      string `json:"build"`
	EntryPoint Entrys `json:"entryPoint"`
}

type Entry struct {
	Get    []Uris
	Post   []Uris
	Delete []Uris
}

type Entrys []Entry

type Uris struct {
	Uri string
}
