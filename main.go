package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    // Define a handler function
	
	activities := map[string]string{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "./application")
		}
		
    })

    http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>the current time is: %s</p>", time.Now().Format("3:04"))
    })
	http.HandleFunc("/activities", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST"{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

		activityName := r.FormValue("activity")
		activityHours := r.FormValue("hours")
		activities[activityName] = activityHours

		fmt.Fprintf(w, "<p> %s was added to your activities</p>", activityName)

	})

    // Start an HTTP server listening on port 8080
    fmt.Println("Server is running on http://localhost:8080/")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}