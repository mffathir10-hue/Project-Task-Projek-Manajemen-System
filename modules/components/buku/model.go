package buku

type Buku struct {
	ID_BUKU       int
	TITLE_BUKU    string
	AUTHOR        string
	ISBN          string
	CATEGORY_ID   int
	NAMA_KATEGORI string
}

var (
	Dummybuku []Buku
)
