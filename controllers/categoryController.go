package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"Quiz/repository"
	"Quiz/structs"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	DB *sql.DB
}

// Create Category
func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var category structs.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repository.InsertCategory(cc.DB, category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

// Get all categories
func (cc *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := repository.GetAllCategories(cc.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": categories})
}

// Get category by ID
func (cc *CategoryController) GetCategoryByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := repository.GetOneCategory(cc.DB, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": category})
}

// Delete category
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category := structs.Category{ID: id}
	err := repository.DeleteCategory(cc.DB, category)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"deleted": id})
}

// Get books by category
func (cc *CategoryController) GetBooksByCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	books, err := repository.GetBooksByCategory(cc.DB, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": books})
}
