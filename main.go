package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type JSON struct {
	Title  string
	Auteur string
	About  string
}

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("public/*.html")
	//the diffenrent possible path
	http.HandleFunc("/Home", getFormeHandler)
	http.Handle("/style/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
func getFormeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Home")
	jsonfile, _ := ioutil.ReadFile("info.json")

	var data []JSON
	// unmarshall (tres sombre --> aller voir)
	json.Unmarshal(jsonfile, &data)

	choice := r.FormValue("choiceInfo")
	if choice == "choix 2" {
		tpl.ExecuteTemplate(w, "index.html", data[1])
	} else {
		tpl.ExecuteTemplate(w, "index.html", data[0])
	}

}
