package controllers

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path"
)

func ChangeValue(w http.ResponseWriter, r *http.Request) {

	var (
		file        = path.Join("views", "index.html")
		templ, err  = template.ParseFiles(file)
		waterStatus string
		windStatus  string
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	water := rand.Intn(100)
	wind := rand.Intn(100)

	if water < 5 {
		waterStatus = "aman"
	} else if water > 5 && water <= 15 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	if wind < 6 {
		windStatus = "aman"
	} else if wind > 6 && wind <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	message1 := fmt.Sprintf("Water %d\n", water)
	message2 := fmt.Sprintf("Wind %d\n", wind)

	data := map[string]interface{}{
		"water":        message1,
		"wind":         message2,
		"status_water": waterStatus,
		"status_wind":  windStatus,
	}

	err = templ.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
