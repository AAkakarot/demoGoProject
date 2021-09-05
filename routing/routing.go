package routing

import (
	"example.com/Users/akashkumar/go/demo_project/controllers"
	"example.com/Users/akashkumar/go/demo_project/utility"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter()
	log.Print("Inside handlers")
	router.HandleFunc("/api/", controllers.HomePage)
	router.Handle("/api", utility.IsAuthorized(controllers.HomePage))
	router.HandleFunc("/api/addbook", controllers.AddBook).Methods("POST")
	router.HandleFunc("/api/login", utility.Login).Methods("POST")
	router.HandleFunc("/api/register", utility.Register).Methods("POST")
	router.Handle("/api/showbooks", utility.IsAuthorized(controllers.GetBooks)).Methods("GET")
	router.HandleFunc("/api/showbooks/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/searchbook/{name}", controllers.SearchoneBook).Methods("POST")
	router.HandleFunc("/api/searchbook/{name}", controllers.SearchBooks).Methods("POST")
	router.HandleFunc("/api/updatebook/{name}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/removeonebook/{name}", controllers.DeleteoneBook).Methods("DELETE")
	router.HandleFunc("/api/removebooks/{name}", controllers.DeleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}
