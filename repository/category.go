package repository

import (
	"Quiz/structs"
	"database/sql"
	"errors"
)

func GetAllCategories(db *sql.DB) ([]structs.Category, error) {
	rows, err := db.Query(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var c structs.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy, &c.ModifiedAt, &c.ModifiedBy); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func GetOneCategory(db *sql.DB, id int) (structs.Category, error) {
	var c structs.Category
	err := db.QueryRow(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id=$1`, id).
		Scan(&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy, &c.ModifiedAt, &c.ModifiedBy)
	if err == sql.ErrNoRows {
		return c, errors.New("not found")
	}
	return c, err
}

func InsertCategory(db *sql.DB, category structs.Category) error {
	_, err := db.Exec(`
    INSERT INTO categories (name, created_at, created_by, modified_at, modified_by) 
    VALUES ($1, NOW(), $2, NOW(), '')
`, category.Name, category.CreatedBy)

	return err
}

func DeleteCategory(db *sql.DB, category structs.Category) error {
	res, err := db.Exec(`DELETE FROM categories WHERE id=$1`, category.ID)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("not found")
	}
	return nil
}

func GetBooksByCategory(db *sql.DB, categoryID int) ([]structs.Book, error) {
	query := `
		SELECT b.id, b.title, b.description, b.image_url, b.release_year, 
		       b.price, b.total_page, b.thickness, b.category_id,
		       b.created_at, b.created_by, b.modified_at, b.modified_by,
		       c.name AS category_name
		FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE b.category_id = $1
	`

	rows, err := db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var b structs.Book
		var categoryName string

		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.ImageURL,
			&b.ReleaseYear,
			&b.Price,
			&b.TotalPage,
			&b.Thickness,
			&b.CategoryID,
			&b.CreatedAt,
			&b.CreatedBy,
			&b.ModifiedAt,
			&b.ModifiedBy,
			&categoryName, // tambahan dari JOIN
		)
		if err != nil {
			return nil, err
		}

		// Kalau mau bisa set ke field baru di struct Book
		// contoh: b.CategoryName = categoryName
		books = append(books, b)
	}

	return books, nil
}
