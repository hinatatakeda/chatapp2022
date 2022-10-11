package chat


import (
	"encoding/json"
	"net/http"
	"strings"
)





package main
 
import (
	"net/http"
 
	"github.com/julienschmidt/httprouter"
)

 
func (app *application) routes() http.Handler {
	router := httprouter.New()
	
	router.HandlerFunc(http.MethodGet, "/general", app.getGeneralTalks)
	router.HandlerFunc(http.MethodGet, "/random", app.getRandomTalks)
	router.HandlerFunc(http.MethodPost, "/login", app.loginHandler)

	router.HandlerFunc(http.MethodPost, "/post/general", app.postGeneral)
	router.HandlerFunc(http.MethodPost, "/delete/general", app.deleteGeneral)
	router.HandlerFunc(http.MethodPost, "/edit/general", app.editGeneral)
 
	return app.enableCORS(router)






// type Message struct {
// 	Msg string `json:"message"`
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError)+"(status code = 500)", http.StatusInternalServerError)
// 		}
// 	}()

// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(http.StatusBadRequest)+"(status code = 400)", http.StatusBadRequest)
// 		return
// 	}
// 	name := r.FormValue("name")
// 	if name == "" {
// 		http.Error(w, http.StatusText(http.StatusBadRequest)+"(status code = 400)", http.StatusBadRequest)
// 		return
// 	}
// 	m := Message{
// 		"Hello, " + name + "!",
// 	}
// 	w.Header().Set("Content-Type", "application/json;charset=utf8")
// 	enc := json.NewEncoder(w)
// 	enc.SetIndent("", strings.Repeat(" ", 2))
// 	if err := enc.Encode(m); err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError)+"(status code = 500)", http.StatusInternalServerError)
// 	}
// }

// func main() {
// 	http.HandleFunc("/hello", handler)
// 	http.ListenAndServe(":8080", nil)
// }