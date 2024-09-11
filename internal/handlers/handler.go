package handlers

import "net/http"

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Items are 1.Bike 2.Car 3.Phone"))
}
