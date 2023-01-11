package data

import (
	"api/features/book"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Judul       string
	TahunTerbit int
	Penulis     string
	UserID      uint
}

func ToCore(data Books) book.Core {
	return book.Core{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}

func CoreToData(data book.Core) Books {
	return Books{
		Model:       gorm.Model{ID: data.ID},
		Judul:       data.Judul,
		Penulis:     data.Penulis,
		TahunTerbit: data.TahunTerbit,
	}
}
