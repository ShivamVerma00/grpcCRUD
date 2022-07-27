package crudimp

import (
	"context"
	"errors"
	"grpcCRUD/model"
	pb "grpcCRUD/protoFile"

	"grpcCRUD/database"
)

//Implemetation of proto interfaces
type BookManagementServiceServer struct {
	pb.UnimplementedBookManagementServiceServer
}

//Create Book
func (s *BookManagementServiceServer) CreateBook(ctx context.Context, req *pb.ReqCreateBook) (*pb.ResCreatebook, error) {
	book := req.Book

	//request Validation
	if book.BookTitle == "" || book.BookAuthor == "" {
		return nil, errors.New("invalid Data")
	}

	bookData := model.Book{
		Title:  book.BookTitle,
		Author: book.BookAuthor,
	}

	id, err := database.CreateBook(ctx, bookData)

	if err != nil {
		return nil, err
	}

	book.Id = uint64(id)

	return &pb.ResCreatebook{Book: &pb.Book{Id: book.Id, BookTitle: book.BookTitle, BookAuthor: book.BookAuthor}}, nil
}

//Update Book
func (s *BookManagementServiceServer) UpdateBook(ctx context.Context, req *pb.ReqUpdateBook) (*pb.ResUpdateBook, error) {
	bookUpdates := req.Book
	oldTitle := req.Title

	//request Validation
	if bookUpdates.BookTitle == "" || bookUpdates.BookAuthor == "" || oldTitle == "" {
		return nil, errors.New("invalid Data")
	}

	bookData := model.Book{
		Title:  bookUpdates.BookTitle,
		Author: bookUpdates.BookAuthor,
	}

	id, err := database.UpdateBook(ctx, bookData, oldTitle)
	if err != nil {
		return nil, err
	}

	bookUpdates.Id = uint64(id)

	return &pb.ResUpdateBook{Book: &pb.Book{Id: bookUpdates.Id, BookTitle: bookUpdates.BookTitle, BookAuthor: bookUpdates.BookAuthor}}, nil
}

//Delete Book
func (s *BookManagementServiceServer) DeleteBook(ctx context.Context, req *pb.ReqDeleteBook) (*pb.ResDeleteBook, error) {
	//Validation
	del := req.BookTitle
	if del == "" {
		return nil, errors.New("invalid Delete Request")
	}

	err := database.DeleteBook(del)
	if err != nil {
		return nil, err
	}

	return &pb.ResDeleteBook{Delete: true}, nil
}

//Get All Books
func (s *BookManagementServiceServer) GetAllBooks(req *pb.ReqGetAllBooks, stream pb.BookManagementService_GetAllBooksServer) error {

	list, err := database.GetAllBooks()
	if err != nil {
		return err

	}

	for _, book := range list {
		err = stream.Send(&pb.ResGetAllBooks{Book: &pb.Book{Id: uint64(book.ID), BookTitle: book.Title, BookAuthor: book.Author}})

		if err != nil {
			return err
		}
	}

	return nil

}

//Search Book
func (s *BookManagementServiceServer) SearchBook(req *pb.ReqSearchBook, stream pb.BookManagementService_SearchBookServer) error {

	title := req.GetBookTitle()
	author := req.GetBookAuthor()

	if title == "" && author == "" {
		return errors.New("nothing to search, empty argment")
	}

	search, err := database.SearchBook(title, author)

	if err != nil {
		return err
	}

	for _, book := range search {
		err = stream.Send(&pb.ResSearchBook{Book: &pb.Book{Id: uint64(book.ID), BookTitle: book.Title, BookAuthor: book.Author}})

		if err != nil {
			return err
		}
	}

	return nil
}
