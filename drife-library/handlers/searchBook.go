package handlers

import (
	"drife-library/models"
	"fmt"
)

func SearchBook(query string) []models.Book {
	var result []models.Book
	for _, book := range Library {
		if book.Title == query || book.Author == query {
			result = append(result, book)
		}
	}
	if len(result) == 0 {
		fmt.Println("No books found")
	}
	return result
}
