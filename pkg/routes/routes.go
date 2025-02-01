package routes

import (
	"github.com/marcfranquesa/bla/pkg/handlers"
	"net/http"
)

func SetupRoutes() {
	http.Handle("/", handlers.ServeStaticFiles())
}

//    conn, _ := db.Connect()
// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	results, _ := db.UrlById(conn, "43")
// 	fmt.Fprint(w, results)
// })
//
// http.ListenAndServe(":" + os.Getenv("PORT"), nil)
