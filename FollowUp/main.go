package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/reportData", getReportsHandler)
	r.HandleFunc("/report", createReportsHandler).Methods("POST")
	r.HandleFunc("/reportFill", saveReport)
	r.HandleFunc("/reportQuery", queryReport)
	r.HandleFunc("/reportDelete", deleteReport)

	return r
}

func main() {
	c := exec.Command("cmd", "/C", `start chrome http://localhost:8080/assets/`)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	r := newRouter()
	http.ListenAndServe(":8080", r)

	//c := exec.Command("cmd", "/C", "chrome.exe", "localhost:8080/assets/")
	//
	//if err := c.Run(); err != nil {
	//	fmt.Println("Error: ", err)
	//}
}
