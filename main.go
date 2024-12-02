package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get current date and time
	now := time.Now()

	// Greeting message
	greeting := fmt.Sprintf("Hello this is a second test, the current time is %s.<br>", now.Format("15:04"))

	// Hours passed today
	hoursPassedToday := float64(now.Hour()) + float64(now.Minute())/60
	greeting += fmt.Sprintf("Hours passed today: %.2f hours.<br>", hoursPassedToday)

	// Calculate Mondays passed since the beginning of the year
	startOfYear := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	daysPassed := int(now.Sub(startOfYear).Hours() / 24)
	mondaysPassed := 0

	// Count how many Mondays have passed
	for i := 0; i <= daysPassed; i++ {
		date := startOfYear.AddDate(0, 0, i)
		if date.Weekday() == time.Monday {
			mondaysPassed++
		}
	}

	greeting += fmt.Sprintf("Hello this is a little Test with Action so... Mondays passed since the beginning of the year: %d.", mondaysPassed)

	// Output the result as HTML
	fmt.Fprintf(w, "<html><head><title>Greeting Page</title></head><body>%s</body></html>", greeting)
}

func main() {
	// Set up the route and start the server
	http.HandleFunc("/", handler)

	// Run the server on port 5000
	fmt.Println("Server started on http://localhost:5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
