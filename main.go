package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"casualdb.com/m/controllers"
	"casualdb.com/m/models"
)

func main() {
	fmt.Println("Welcome to CasualDB!")

	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/submit", handleFormSubmission)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/client.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	blockSizeInitial := r.FormValue("blockSize")
	fileDirectory := r.FormValue("fileDirectory")
	fileName := r.FormValue("fileName")
	data := r.FormValue("data")

	blockSize, err := strconv.Atoi(blockSizeInitial)
	if err != nil {
		http.Error(w, "Invalid block size", http.StatusBadRequest)
		return
	}

	fileManager := controllers.NewFileManager(blockSize, fileDirectory)
	fileController := &controllers.FileController{
		FileManager: fileManager,
	}

	page := controllers.NewPage(blockSize)
	pageController := &controllers.PageController{
		Page: page,
	}

	block := &models.Block{
		FileName: fileName,
		Identity: 0,
	}

	fmt.Printf("Block Size: %d\n", blockSize)
	fmt.Printf("File Directory: %s\n", fileDirectory)
	fmt.Printf("File Name: %s\n", fileName)
	fmt.Printf("Data: %s\n", data)

	dataBytes := []byte(data)
	n, err := pageController.Write(0, dataBytes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing to page: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Wrote %d bytes to page\n", n)

	bytesWritten, err := fileController.Write(block, page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing to file: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Wrote %d bytes to file\n", bytesWritten)

	response := fmt.Sprintf(`
		<h2>Operation Successful!</h2>
		<p>Block Size: %d bytes</p>
		<p>File Directory: %s</p>
		<p>File Name: %s</p>
		<p>Data Written: %s</p>
		<p>Bytes Written to Page: %d</p>
		<p>Bytes Written to File: %d</p>
		<br>
		<a href="/">Go Back</a>
	`, blockSize, fileDirectory, fileName, data, n, bytesWritten)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))

	fileController.Close()
}
