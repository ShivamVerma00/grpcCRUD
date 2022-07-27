package database

import (
	"context"
	"errors"

	"grpcCRUD/model"
)

//Create Book
func CreateBook(ctx context.Context, newBook model.Book) (uint, error) {
	user := model.Book{Title: newBook.Title, Author: newBook.Author}
	create := C.Create(&user)
	if create.Error != nil {
		return 0, create.Error
	}

	return user.ID, nil
}

//Update Book
func UpdateBook(ctx context.Context, updateBook model.Book, title string) (uint, error) {

	var book model.Book

	err := C.Where(&model.Book{Title: title}).Find(&book).Error
	if err != nil {
		return 0, err
	}

	err = C.Model(&book).Updates(updateBook).Error
	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

//Delete Book
func DeleteBook(s string) error {
	tx := C.Delete(s)
	if tx != nil {
		return tx.Error
	}

	return nil
}

//GetAll Books
func GetAllBooks() ([]model.Book, error) {
	var listOfBooks []model.Book
	err := C.Find(&listOfBooks).Error
	if err != nil {
		return nil, err
	}

	if len(listOfBooks) == 0 {
		return nil, errors.New("empty list")
	}

	return listOfBooks, nil
}

// Search Books
func SearchBook(title string, author string) ([]model.Book, error) {

	if title == "" && author == "" {
		return nil, errors.New("nothing to search, empty argment")
	}

	if title == "" {
		err := C.Find(&author).Error
		if err != nil {
			return nil, err
		}
	}
	if author == "" {
		err := C.First(&title).Error
		if err != nil {
			return nil, err
		}
	}

	return nil, errors.New("Invalid")

}
