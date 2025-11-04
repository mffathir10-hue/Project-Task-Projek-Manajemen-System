package buku

import (
	"database/sql"
	"errors"
)

type Repository interface {
	CreateBukuRepository(buku Buku) (result []Buku, err error)
	GetAllBukuRepository() ([]Buku, error)
	GetCategoryIDByName(categoryName string) (int, error)
	// GetBioskopRepository(id int) (buku, error)
	// DeleteBioskopRepository(id int) (err error)
	// UpdateBioskopRepository(buku buku) (buku, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateBukuRepository(buku Buku) (result []Buku, err error) {
	query := `
        INSERT INTO buku (title_buku, author, isbn, category_id) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id_buku
    `

	var id int
	err = r.db.QueryRow(query, buku.TITLE_BUKU, buku.AUTHOR, buku.ISBN, buku.CATEGORY_ID).Scan(&id)
	if err != nil {
		return nil, err
	}

	buku.ID_BUKU = id

	allBioskop, err := r.GetAllBukuRepository()
	if err != nil {
		return nil, err
	}

	return allBioskop, nil
}

func (r *repository) GetCategoryIDByName(categoryName string) (int, error) {
	query := "SELECT id_kategori FROM kategori WHERE nama_kategori = $1"
	var categoryID int

	err := r.db.QueryRow(query, categoryName).Scan(&categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("kategori '" + categoryName + "' tidak ditemukan")
		}
		return 0, err
	}

	return categoryID, nil
}

func (r *repository) GetAllBukuRepository() ([]Buku, error) {
	query := `
		SELECT 
			b.id_buku, 
			b.title_buku, 
			b.author, 
			b.isbn, 
			b.category_id,
			k.nama_kategori
		FROM buku b
		LEFT JOIN kategori k ON b.category_id = k.id_kategori
		ORDER BY b.id_buku`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bukus []Buku
	for rows.Next() {
		var buku Buku
		err := rows.Scan(
			&buku.ID_BUKU,
			&buku.TITLE_BUKU,
			&buku.AUTHOR,
			&buku.ISBN,
			&buku.CATEGORY_ID,
			&buku.NAMA_KATEGORI,
		)
		if err != nil {
			return nil, err
		}
		bukus = append(bukus, buku)
	}

	return bukus, nil
}

// func (r *repository) GetBioskopRepository(id int) (buku, error) {
// 	sql := "SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1"

// 	var buku buku
// 	err := r.db.QueryRow(sql, id).Scan(&buku.ID, &buku.Nama, &buku.Lokasi, &buku.Rating)

// 	if err != nil {
// 		return buku{}, errors.New("gagal mengambil data bioskop: " + err.Error())
// 	}

// 	return buku, nil
// }

// func (r *repository) UpdateBioskopRepository(buku buku) (buku, error) {
// 	sql := "UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4 RETURNING id, nama, lokasi, rating"

// 	var updatedbuku buku
// 	err := r.db.QueryRow(sql, buku.Nama, buku.Lokasi, buku.Rating, buku.ID).
// 		Scan(&updatedbuku.ID, &updatedbuku.Nama, &updatedbuku.Lokasi, &updatedbuku.Rating)

// 	if err != nil {
// 		return buku{}, errors.New("gagal mengupdate bioskop: " + err.Error())
// 	}

// 	return updatedbuku, nil
// }

// func (r *repository) DeleteBioskopRepository(id int) (err error) {

// 	sql := "DELETE FROM bioskop WHERE id = $1"

// 	result, err := r.db.Exec(sql, id)
// 	if err != nil {
// 		return errors.New("gagal menghapus bioskop: " + err.Error())
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return errors.New("gagal memeriksa rows affected: " + err.Error())
// 	}

// 	if rowsAffected == 0 {
// 		return errors.New("bioskop tidak ditemukan")
// 	}

// 	return nil
// }
