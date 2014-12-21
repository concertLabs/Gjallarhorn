package main

import "net/http"

func statichandler(
	w http.ResponseWriter,
	r *http.Request) {
	//Nur für den fall das wir mal irgendwas
	//limitieren müssen
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	port := ":2650"

	http.HandleFunc("/static/", statichandler)
	http.ListenAndServe(":8080", nil)
}
