package controllers

import (
	"database/sql"
	"net/http"
	"quiz3/structs"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context, db *sql.DB) {
	var books []structs.Book
	rows, err := db.Query("SELECT id, title, description, image_url, created_at, created_by, modified_at, modified_by FROM Books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var book structs.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var book structs.Book
	query := "SELECT id, title, description, image_url, created_at, created_by, modified_at, modified_by FROM Books WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context, db *sql.DB) {
	var book structs.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	result, err := db.Exec("INSERT INTO Books (title, description, image_url, category_id, release_year, price, total_page, thickness, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), $9)", book.Title, book.Description, book.ImageURL, book.CategoryID, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book created successfully", "id": id})
}

func UpdateBook(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var book structs.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existingBook structs.Book
	query := "SELECT id, title FROM Books WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&existingBook.ID, &existingBook.Title)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	_, err = db.Exec("UPDATE Books SET title = $1, description = $2, image_url = $3, category_id = $4, release_year = $5, price = $6, total_page = $7, thickness = $8, modified_at = NOW(), modified_by = $9 WHERE id = $10", book.Title, book.Description, book.ImageURL, book.CategoryID, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.ModifiedBy, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

func DeleteBook(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var existingBook structs.Book
	query := "SELECT id, title FROM Books WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&existingBook.ID, &existingBook.Title)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("DELETE FROM Books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
