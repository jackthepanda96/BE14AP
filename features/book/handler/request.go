package handler

import "api/features/book"

type AddBookRequest struct {
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
}

func ToCore(data interface{}) *book.Core {
	res := book.Core{}

	switch data.(type) {
	case AddBookRequest:
		cnv := data.(AddBookRequest)
		res.Judul = cnv.Judul
		res.TahunTerbit = cnv.TahunTerbit
		res.Penulis = cnv.Penulis
	default:
		return nil
	}

	return &res
}
