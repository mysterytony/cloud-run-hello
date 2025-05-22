package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Region string
	UnicornImagePath string
}

func main() {
	http.HandleFunc("/", rootHandler)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	region := os.Getenv("REGION")
	if region == "" {
		region = "unknown"
	}

	// Assume the image is at /assets/happy-unicorn.svg
	// The actual serving of this static asset will be handled by another handler.
	data := PageData{
		Region: region,
		UnicornImagePath: "/assets/happy-unicorn.svg",
	}

	// Create a simple HTML template
	tmplStr := `
<!DOCTYPE html>
<html>
<head>
    <title>Service Health</title>
    <style>
        body { font-family: sans-serif; display: flex; flex-direction: column; align-items: center; justify-content: center; height: 100vh; margin: 0; background-color: #f0f0f0; }
        img { max-width: 200px; margin-bottom: 20px; }
        h1 { color: #333; }
    </style>
</head>
<body>
    <img src="{{.UnicornImagePath}}" alt="Happy Unicorn">
    <h1>Served from region: {{.Region}}</h1>
</body>
</html>
`
	tmpl, err := template.New("index").Parse(tmplStr)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}
