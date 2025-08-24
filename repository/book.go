package repository

import (
	"Quiz/structs"
	"database/sql"
	"errors"
	"time"
)

func GetAllBook(db *sql.DB) (result []structs.Book, err error) {
	sql := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var books structs.Book

		err = rows.Scan(&books.ID, &books.Title, &books.Description, &books.ImageURL, &books.ReleaseYear,
			&books.Price, &books.TotalPage, &books.Thickness, &books.CategoryID,
			&books.CreatedAt, &books.CreatedBy, &books.ModifiedAt, &books.ModifiedBy)
		if err != nil {
			return nil, err
		}

		result = append(result, books)
	}

	return

}

func GetOneBook(db *sql.DB, id int) (book structs.Book, err error) {
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id,
	                 created_at, created_by, modified_at, modified_by
	          FROM books WHERE id = $1`

	err = db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CategoryID,
		&book.CreatedAt,
		&book.CreatedBy,
		&book.ModifiedAt,
		&book.ModifiedBy,
	)
	return
}

func InsertBook(db *sql.DB, book structs.Book) error {
	// Konversi thickness berdasarkan total_page
	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	_, err := db.Exec(`INSERT INTO books 
		(title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price,
		book.TotalPage, book.Thickness, book.CategoryID, time.Now(), book.CreatedBy,
	)
	return err
}

func UpdateBook(db *sql.DB, book structs.Book) error {
	// Konversi thickness berdasarkan total_page
	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	res, err := db.Exec(`UPDATE books SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, category_id=$8, modified_at=$9, modified_by=$10 
		WHERE id=$11`,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price,
		book.TotalPage, book.Thickness, book.CategoryID, time.Now(), book.ModifiedBy, book.ID,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("book not found")
	}
	return nil
}

func DeleteBook(db *sql.DB, book structs.Book) error {
	res, err := db.Exec(`DELETE FROM books WHERE id=$1`, book.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("book not found")
	}
	return nil
}
