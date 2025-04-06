package main

import {"fmt","encoding/json","math/rand","strconv","github.com/gorilla/mux"}

type Movie struct{
	ID string `json:"id"`
	ISBN string `json:"isbn"` 
	DIRECTOR *Director
}

type Director struct{
	FirstName string
	LastName string
}

var movies []Movie