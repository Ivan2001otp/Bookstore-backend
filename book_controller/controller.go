package bookcontroller

import (
	"bookstore_backend/model"
	utils "bookstore_backend/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newBook model.Book;

func FetchAllBook(response http.ResponseWriter,request *http.Request){
	fetchedBooks := model.GetAllBooks();

	res,_ := json.Marshal(fetchedBooks);

	response.Header().Set("Content-Type","pkglication/json");
	response.WriteHeader(http.StatusOK)
	response.Write(res);
}

func FetchBookById(response http.ResponseWriter,request *http.Request){
	vars := mux.Vars(request);
	bookId := vars["bookId"];

	Id,err := strconv.ParseInt(bookId,0,0)

	if(err!=nil){
		fmt.Println("Error while fetching Book by Id");
		panic(err);
	}

	fetchedBookDetail ,_ := model.GetBookById(Id);

	if(err!=nil){
		panic(err);
	}

	res,err := json.Marshal(fetchedBookDetail);

	response.Header().Set("Content-Type","pkglication/json");
	response.WriteHeader(http.StatusOK)
	response.Write(res);
}

func CreateBook(response http.ResponseWriter, request *http.Request){
	newBook := &model.Book{};
	utils.ParseBody(request,newBook)

	b := newBook.CreateBook();


	res,_ := json.Marshal(b);
	response.WriteHeader(http.StatusOK);
	response.Write(res);

}

func DeleteBookById(response http.ResponseWriter,request *http.Request){
	vars := mux.Vars(request);

	bookId := vars["bookId"];
	Id,err := strconv.ParseInt(bookId,0,0);

	if(err!=nil){
		fmt.Println("Parsing Exception");
		panic(err);
	}

	book := model.DeleteBookById(Id);

	res,_ := json.Marshal(book);

	response.Header().Set("Content-Type","pkglication/json");
	response.Write(res);
}

func UpdateBookById(response http.ResponseWriter,request *http.Request){
	var updatedBook = &model.Book{};

	utils.ParseBody(request,updatedBook);

	vars := mux.Vars(request);

	bookId := vars["bookId"];

	Id,err := strconv.ParseInt(bookId,0,0);

	if(err!=nil){
		fmt.Println("Error while parsing");
		panic(err);
	}

	bookDetail,db := model.GetBookById(Id);

	if(updatedBook.Name!=""){
		bookDetail.Name = updatedBook.Name;
	}

	if(updatedBook.Author!=""){
		bookDetail.Author = updatedBook.Author;
	}
	if(updatedBook.Publication!=""){
		bookDetail.Publication = updatedBook.Publication;
	}

	db.Save(&bookDetail);

	res,_ := json.Marshal(bookDetail);
	response.Header().Set("Content-Type","pkglication/json");
	response.WriteHeader(http.StatusOK);
	response.Write(res);
}