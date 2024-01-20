// // C:\Users\Utkarsh Raj\go\src\github.com\ut\go-crud

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book represents a book entity
type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// bookDB simulates a simple in-memory database
var bookDB = []Book{
	{ID: "1", Title: "XYZ", Author: "ABC", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// setupRouter configures the Gin router and defines routes
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.DELETE("/books/:id", deleteBook)

	return router
}

// getBooks handles the GET request to retrieve all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookDB)
}

// createBook handles the POST request to create a new book
func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookDB = append(bookDB, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// deleteBook handles the DELETE request to delete a book by ID
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	index := findBookIndex(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	bookDB = append(bookDB[:index], bookDB[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// findBookIndex finds the index of a book by ID in the bookDB
func findBookIndex(id string) int {
	for i, b := range bookDB {
		if b.ID == id {
			return i
		}
	}
	return -1
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
