package main

import (
    "log"
    "net/http"
)


type userHandle struct{}

func main(){

    uh := userHandle{}

    http.Handle("/user", uh)
    http.HandleFunc("/group", Group)

    err := http.ListenAndServe(":8000", nil)

    if err != nil {
		log.Fatal("Server error ", err.Error())
	}else{
        log.Println("Server exec ")
    }
}


func(uh userHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
        case "GET":
            uh.getUser(w, r)
        case "POST":
            uh.setUser(w, r)
        default:
            http.Error(w, "Method Allow", http.StatusMethodNotAllowed)
    }
}


func(uh userHandle) getUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("this is getUser"))
}

func(uh userHandle) setUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("this is setUser"))
}

func Group(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("this is group"))
}