package buku

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateBukuService(ctx *gin.Context) (Buku, error)
	GetAllBukuService(ctx *gin.Context) (result []Buku, err error)
	// GetBioskopService(ctx *gin.Context) (buku, error)
	// UpdateBioskopService(ctx *gin.Context) (buku, error)
	// DeleteBioskopService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (s *userService) CreateBukuService(ctx *gin.Context) (Buku, error) {
	var buku Buku

	if err := ctx.ShouldBindJSON(&buku); err != nil {
		return Buku{}, err
	}

	// Validasi required fields
	if buku.TITLE_BUKU == "" {
		return Buku{}, errors.New("Judul buku harus diisi")
	}
	if buku.AUTHOR == "" {
		return Buku{}, errors.New("Author buku harus diisi")
	}
	if buku.ISBN == "" {
		return Buku{}, errors.New("ISBN buku harus diisi")
	}
	if buku.NAMA_KATEGORI == "" {
		return Buku{}, errors.New("Kategori buku harus diisi")
	}

	// Cari ID kategori berdasarkan nama
	categoryID, err := s.repository.GetCategoryIDByName(buku.NAMA_KATEGORI)
	if err != nil {
		return Buku{}, errors.New("kategori tidak ditemukan: " + err.Error())
	}

	// Buat struct buku untuk repository
	bukus := Buku{
		TITLE_BUKU:  buku.TITLE_BUKU,
		AUTHOR:      buku.AUTHOR,
		ISBN:        buku.ISBN,
		CATEGORY_ID: categoryID,
	}

	// Simpan ke database
	result, err := s.repository.CreateBukuRepository(bukus)
	if err != nil {
		return Buku{}, errors.New("gagal menambahkan buku: " + err.Error())
	}

	return result[0], nil
}

func (s *userService) GetAllBukuService(ctx *gin.Context) ([]Buku, error) {
	bukus, err := s.repository.GetAllBukuRepository()
	if err != nil {
		return nil, errors.New("gagal mengambil data buku: " + err.Error())
	}

	return bukus, nil
}

// func (s *userService) GetBioskopService(ctx *gin.Context) (result buku, err error) {
// 	idStr := ctx.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return buku{}, errors.New("ID bioskop tidak valid")
// 	}

// 	if id <= 0 {
// 		return buku{}, errors.New("ID bioskop harus lebih dari 0")
// 	}

// 	result, err = s.repository.GetBioskopRepository(id)
// 	if err != nil {
// 		return buku{}, errors.New("bioskop tidak ditemukan")
// 	}

// 	return result, nil
// }

// func (s *userService) UpdateBioskopService(ctx *gin.Context) (result Bioskop, err error) {

// 	idStr := ctx.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return Bioskop{}, errors.New("ID bioskop tidak valid")
// 	}

// 	if id <= 0 {
// 		return Bioskop{}, errors.New("ID bioskop harus lebih dari 0")
// 	}

// 	var bioskop Bioskop
// 	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
// 		return Bioskop{}, errors.New("data request tidak valid")
// 	}

// 	bioskop.ID = id

// 	if strings.TrimSpace(bioskop.Nama) == "" {
// 		return Bioskop{}, errors.New("nama bioskop harus diisi")
// 	}

// 	if strings.TrimSpace(bioskop.Lokasi) == "" {
// 		return Bioskop{}, errors.New("lokasi bioskop harus diisi")
// 	}

// 	result, err = s.repository.UpdateBioskopRepository(bioskop)
// 	if err != nil {
// 		return Bioskop{}, err
// 	}

// 	return result, nil
// }

// func (s *userService) DeleteBioskopService(ctx *gin.Context) (err error) {
// 	idStr := ctx.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return errors.New("ID bioskop tidak valid")
// 	}

// 	if id <= 0 {
// 		return errors.New("ID bioskop harus lebih dari 0")
// 	}

// 	err = s.repository.DeleteBioskopRepository(id)
// 	if err != nil {
// 		return err
// 	}

// 	return
// }
