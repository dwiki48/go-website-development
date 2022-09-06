package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"website-development/entity"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("welcome page"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "wait, it's eror", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "Golang Web",
	// 	"content": "this is Golang Web",
	// }

	// data := entity.Product{ID: 1, Name: "Logitech", Price: 25000000, Stock: 10}
	data := []entity.Product{
		{ID: 1, Name: "Logitech", Price: 25000000, Stock: 2},
		{ID: 2, Name: "Rexus", Price: 25000000, Stock: 11},
		{ID: 3, Name: "DigitalAliance", Price: 25000000, Stock: 8},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "wait, it's eror", http.StatusInternalServerError)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "product page : %d", idNumb)
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "wait, it's eror", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "wait, it's eror", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, learning golang"))
}
