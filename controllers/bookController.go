package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"Quiz/repository"
	"Quiz/structs"

	"github.com/gin-gonic/gin"
)

// BookController struct
type BookController struct {
	DB *sql.DB
}

// CreateBook - buat data buku baru
func (bc *BookController) CreateBook(ctx *gin.Context) {
	var book structs.Book

	// Validasi input
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan ke database lewat repository
	err := repository.InsertBook(bc.DB, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

// GetAllBooks - ambil semua data buku
func (bc *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := repository.GetAllBook(bc.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": books})
}

// GetBookByID - ambil detail buku by ID
func (bc *BookController) GetBookByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	book, err := repository.GetOneBook(bc.DB, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": book})
}

// UpdateBook - update data buku
func (bc *BookController) UpdateBook(ctx *gin.Context) {
	var book structs.Book
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.ID = id

	err = repository.UpdateBook(bc.DB, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// DeleteBook - hapus data buku
func (bc *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	book := structs.Book{ID: id}
	err = repository.DeleteBook(bc.DB, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": book.ID})
}
