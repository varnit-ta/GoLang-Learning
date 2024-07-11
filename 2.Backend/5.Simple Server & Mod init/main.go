package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//To bring third party library the syntax is go get "package name"

//"go mod verify" to verify the dependencies using the hash value in go.sum file

//"go mod tidy" to remove the unused dependencies

//"go mod init" to create a new module

//"go mod download" to download all the dependencies

//"go list -m all" to list all the dependencies

//"go list -m -version <package name>" to list the version of the package

//"go mod why -m <package name>" to list the reason why the package is being used

//"go mod graph" to list the dependency graph

//"go mod vendor" to create a vendor folder with all the dependencies i.e brings the dependencies into the project from the cache

//"go run -mod=vendor" to run the project using the dependencies from the vendor folder instead of the cache

func main() {
	fmt.Println("Create a simple server and mod init")
	greeter()

	//Syntax for creating a router: mux.NewRouter()
	r := mux.NewRouter()

	//In the router we pass reference to the function that will be called when the route is hit
	//Syntax for creating a route: router.HandleFunc("route", function).Methods("GET/POST/PUT/DELETE")
	r.HandleFunc("/", serverHome).Methods("GET")

	//Syntax for running the server: http.ListenAndServe(":port", router)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page"))
}
