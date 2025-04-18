package main

import (
	"fmt"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"log"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	ISBN string `json:"isbn"` 
	Title string `json:"title"`
	Director *Director
}

type Director struct{
	FirstName string
	LastName string
}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
	return 
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params:= mux.Vars(r)

	for index,item :=range movies{
		
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			break
		}
		
	}
}

func getMovie(w http.ResponseWriter,r * http.Request){
	w.Header().Set("Content-Type","application/json")

	params:= mux.Vars(r)

	for _,item := range movies{
		if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
	}}
}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies,movie)

	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
	params:= mux.Vars(r)

	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	movies = append(movies,movie)

	json.NewEncoder(w).Encode(movie)
}


func main(){
	r:= mux.NewRouter()

	movies =append(movies,Movie{
		ID:"1",
		ISBN:"213421",
		Title:"Movie One",
		Director:&Director{
			FirstName:"Jhon",
			LastName:"Smith",
			}})

	movies=append(movies,Movie{
		ID:"2",
		ISBN:"52323",
		Title:"Movie Two",
		Director:&Director{
			FirstName:"Maria",
			LastName:"Hill",
		}})
	
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}