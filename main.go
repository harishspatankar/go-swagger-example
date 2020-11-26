package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var books = []Book{
	{ID: "1", Name: "Ikigai", Price: 100.00},
	{ID: "2", Name: "Panchatantra", Price: 200.00},
	{ID: "3", Name: "Raja Ram", Price: 400.00},
}

func main() {

	router := mux.NewRouter()

	// swagger:route POST /books books books CreateBook
	// Creates new book
	// Creates new book
	router.HandleFunc("/books", CreateBook).Methods("POST")

	// swagger:route GET /books books books GetBook
	// Returns list of all books
	// Returns list of all books
	router.HandleFunc("/books", GetBook).Methods("GET")

	// swagger:route GET /books/{id} books books GetBookById
	// Returns list of all books
	// Returns list of all books
	router.HandleFunc("/books/{id}", GetBookById).Methods("GET")

	// swagger:route PUT /books/{id} books books UpdateBook
	// Returns list of all books
	// Returns list of all books
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")

	// swagger:route DELETE /books/{id} books books DeleteBook
	// Returns list of all books
	// Returns list of all books
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

// swagger:operation POST /books books CreateBook
// ---
// summary: Creates new book.
// description: Method creates new book.
// requestBody:
//   content:
//     application/json:
//       schema:
//         type: object
//         properties:
//           name:
//             type: string
//           price:
//             type: number
//         example:
//           name: Rich Dad Poor Dad
//           price: 300.00
// responses:
//   "200":
//     description: OK
func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateBook")
}

// swagger:operation GET /books books GetBook
// ---
// summary: Returns list of all books.
// description: Method returns list all books.
// responses:
//   "200":
//     description: OK
func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetBook")
}

// swagger:operation GET /books/{id} books GetBookById
// ---
// summary: Returns specific book according to param id.
// description: Method returns specific book according to param id.
// parameters:
// - name: id
//   in: path
//   description: id of book
//   type: integer
//   required: true
// responses:
//   "200":
//     description: OK
func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetBookById")
}

// swagger:operation PUT /books/{id} books UpdateBook
// ---
// summary: Updates specific book according to param id.
// description: Method Updates specific book according to param id.
// parameters:
// - name: id
//   in: path
//   description: id of book
//   type: integer
//   required: true
// responses:
//   "200":
//     description: OK
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateBook")
}

// swagger:operation DELETE /books/{id} books DeleteBook
// ---
// summary: Deletes specific book according to param id.
// description: Method deletes specific book according to param id.
// parameters:
// - name: id
//   in: path
//   description: id of book
//   type: integer
//   required: true
// responses:
//   "200":
//     description: OK
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteBook")
}
