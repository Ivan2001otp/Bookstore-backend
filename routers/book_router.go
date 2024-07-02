package routers

import ("github.com/gorilla/mux"
controller "bookstore_backend/book_controller";
)


var RegisteredRouters = func(router *mux.Router) {
	router.HandleFunc("/book/",controller.CreateBook).Methods("POST");
	router.HandleFunc("/book/",controller.FetchAllBook).Methods("GET");
	router.HandleFunc("/book/{bookId}",controller.FetchBookById).Methods("GET");
	router.HandleFunc("/book/{bookId}",controller.UpdateBookById).Methods("PUT");
	router.HandleFunc("/book/{bookId}",controller.DeleteBookById).Methods("DELETE");
}