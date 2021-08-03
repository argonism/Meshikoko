
package main

import (
    "flag"
    "app/config"
    "app/database"
    "app/server"
    // "app/models"
)

func main() {

    env := flag.String("e", "development", "")
    flag.Parse()

    config.Init(*env)
    // database.Init(true, models.GetModels()...)
    defer database.Close()
    if err := server.Init(); err != nil {
        panic(err)
    }
}


// [START import]
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// // [END import]
// // [START main_func]

// func main() {
// 	http.HandleFunc("/", indexHandler)

// 	// [START setting_port]
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 		log.Printf("Defaulting to port %s", port)
// 	}

// 	log.Printf("Listening on port %s", port)
// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		log.Fatal(err)
// 	}
// 	// [END setting_port]
// }

// // [END main_func]

// // [START indexHandler]

// // indexHandler responds to requests with our greeting.
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Fprint(w, "Hello, World!")
// }