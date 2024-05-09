package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var prefix = "templates/"

type ViewData struct {
	Title           string
	Message         string
	CardItemInfo    string
	CardItemAnimate string
	CardItemDate    string
	CardItemMap     string
	CardItemNotify  string
}

func main() {

	// ctrltcp.Fun()

	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{}
		tmpl, err := template.ParseFiles(prefix + "menu.html")
		if err != nil {
			fmt.Println("error open html file")
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/date", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:           "World Cup",
			Message:         "ead",
			CardItemInfo:    "<p>INfo</p>",
			CardItemAnimate: "<p>anim</p>",
			CardItemDate:    "<p>date</p>",
			CardItemMap:     "<p>map</p>",
			CardItemNotify:  "<p>notify</p>",
		}
		tmpl, err := template.ParseFiles(prefix + "datetime.html")
		if err != nil {
			fmt.Println("error open html file")
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:           "World Cup",
			Message:         "ead",
			CardItemInfo:    "<p>INfo</p>",
			CardItemAnimate: "<p>anim</p>",
			CardItemDate:    "<p>date</p>",
			CardItemMap:     "<p>map</p>",
			CardItemNotify:  "<p>notify</p>",
		}
		tmpl, err := template.ParseFiles(prefix + "notify.html")
		if err != nil {
			fmt.Println("error open html file")
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:           "World Cup",
			Message:         "ead",
			CardItemInfo:    "<p>INfo</p>",
			CardItemAnimate: "<p>anim</p>",
			CardItemDate:    "<p>date</p>",
			CardItemMap:     "<p>map</p>",
			CardItemNotify:  "<p>notify</p>",
		}
		tmpl, err := template.ParseFiles(prefix + "map.html")
		if err != nil {
			fmt.Println("error open html file")
		}
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
