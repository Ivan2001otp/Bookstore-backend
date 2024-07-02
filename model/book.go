package model

import (
	"bookstore_backend/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name 	string `gorm:""json:"name"`;
	Author  string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect();
	db = config.GetDb();
	db.AutoMigrate(&Book{});
}


//bare bone methods of CRUD
func (b *Book)CreateBook() *Book{
	db.NewRecord(b);
	db.Create(&b);
	return b;
}

func GetAllBooks() []Book{
	var bookList []Book
	db.Find(&bookList);
	return bookList;
}

func GetBookById(Id int64)(*Book,*gorm.DB){
	var bookById Book;
	db := db.Where("ID=?",Id).Find(&bookById);
	return &bookById,db;
}

func DeleteBookById(Id int64) Book{
	var book Book;
	db.Where("ID=?",Id).Delete(&book);
	return book;
}