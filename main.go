package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

// Benchmark represents a row in the benchmarks table.
type Benchmark struct {
	ID                int
	BName             string
	BType             string
	Kind              string
	Enabled           bool
	Branches          string // Assuming the branches field is of type `_text` in PostgreSQL.
	Priority          int
	BCycle            int
	BOffset           int
	BURL              string
	NotificationEmail string
	Issues            string
	Notes             string
}

var db *sql.DB

// Initialize database connection
func initDB() {
	var err error
	connStr := "user=postgres password=postgres dbname=polu sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	fmt.Println("Database connected successfully!")
}

// Fetch data from the benchmarks table
func fetchBenchmarks() ([]Benchmark, error) {
	rows, err := db.Query(`
		SELECT id, bname, btype, kind, enabled, branches, priority, bcycle, boffset, burl, notification_email, issues, notes 
		FROM public.benchmarks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var benchmarks []Benchmark
	for rows.Next() {
		var b Benchmark
		if err := rows.Scan(&b.ID, &b.BName, &b.BType, &b.Kind, &b.Enabled, &b.Branches, &b.Priority, &b.BCycle, &b.BOffset, &b.BURL, &b.NotificationEmail, &b.Issues, &b.Notes); err != nil {
			return nil, err
		}
		benchmarks = append(benchmarks, b)
	}
	return benchmarks, nil
}

// HTTP handler for the main page
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	benchmarks, err := fetchBenchmarks()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, benchmarks); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func main() {
	initDB()
	defer db.Close()
	port := ":8081"
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8081")
	log.Fatal(http.ListenAndServe(port, nil))
}
