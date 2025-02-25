package main

// will automatically import after you save
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a struct to store the books for our library

type book struct {
	// Our fields start with a capital letter. This makes it an exported field name. Means it can be viewed in modules outside of our own.
	// If you don't, everytime you try to return books, it will return an empty JSON object
	// Followed by the type fo that field
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	// Want to convert it to JSON (standard language for APIs) so add another entry at the end for the json fieldname in backticks for our entries in our struct
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// gin context is all the info about the request and allows you to return a response
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // sending http response code, then the books as JSON with the proper indentation. S
	// serializes the book struct and sends it as a json object.
	// can return files or other artifacts, doesnt need to be json.
	// test with curl localhost:8080/books
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return //if you dont return, youll bind the json to the newBook return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
