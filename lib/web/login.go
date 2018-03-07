package web

import "net/http"

func (a *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 300)
}
