package controllers

import (
	"database/sql"
	"net/http"
	"quiz3/structs"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context, db *sql.DB) {
	var categories []structs.Category
	rows, err := db.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM Categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category structs.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, category)
	}
	c.JSON(http.StatusOK, categories)
}

func GetCategoryByID(c *gin.Context, db *sql.DB) {
	var category structs.Category
	id := c.Param("id")
	err := db.QueryRow("SELECT id, name, created_at, created_by, modified_at, modified_by FROM Categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context, db *sql.DB) {
	var category structs.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO Categories (name, created_at, created_by) VALUES ($1, NOW(), $2)", category.Name, category.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	category.ID = int(id)
	c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c *gin.Context, db *sql.DB) {
	var category structs.Category
	id := c.Param("id")
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCategory structs.Category
	query := "SELECT id, name FROM categories WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&existingCategory.ID, &existingCategory.Name)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	_, err = db.Exec("UPDATE Categories SET name = $1, modified_at = NOW(), modified_by = $2 WHERE id = $3", category.Name, category.ModifiedBy, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func DeleteCategory(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var existingCategory structs.Category
	query := "SELECT id, name FROM categories WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&existingCategory.ID, &existingCategory.Name)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	_, err = db.Exec("DELETE FROM Categories WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func GetBooksByCategoryID(c *gin.Context, db *sql.DB) {
	var books []structs.Book
	categoryID := c.Param("id")
	rows, err := db.Query("SELECT id, title, description, image_url, category_id, created_at, created_by, modified_at, modified_by FROM Books WHERE category_id = $1", categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book structs.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}
