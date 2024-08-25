package main

import (
	"fmt"
	"net/http"
)

type Activity struct {
	name string
	timeframe string
	time int
}

func main() {
	// Define a handler function

	activities := []Activity{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "./application")
		}

	})

	http.HandleFunc("/addActivity", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

		activityName := r.FormValue("activityName")
		activityTimeFrame := r.FormValue("timeFrame")

		activities = append(activities, Activity{activityName, activityTimeFrame, 0})

		w.Write([]byte(""))

	})


	http.HandleFunc("/getActivityForm", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w,
			`
			<form hx-target="this" hx-post="/addActivity">
			<p>what's the name of the activity you'd like to track?</p>
			<input type="text" name="activityName" placeholder="reading">
			<p>on what timeframe would you like to track this activity</p>
			<input type="radio" id="Daily" name="timeFrame" value="Daily">
			<label for="Daily">Daily</label><br>
			<input type="radio" id="Weekly" name="timeFrame" value="Weekly">
			<label for="Weekly">Weekly</label><br>
			<input type="radio" id="Monthly" name="timeFrame" value="Monthly">
			<label for="Monthly">Monthly</label><br>
			<button type="submit">create activity</button>
			</form>
		`)
	})

	http.HandleFunc("/displayActivities", func(w http.ResponseWriter, r *http.Request){
		if len(activities) > 0 {
			for _, activity := range activities {
				fmt.Fprintf(w, `<p>Activity: %s, Hours spent: 0, time frame: %s</p>`, activity.name, activity.timeframe)
			}
		} else {
			w.Write([]byte("<p>your activities will show up here when you add them</p>"))
		}
	})

	// Start an HTTP server listening on port 8080
	fmt.Println("Server is running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
