package main

import (
    "encoding/json"
    "fmt"
    "log"
)


type Document struct {
    Title  string `json:"title"`
    Author string `json:"author"`
    DOI    string `json:"doi"`
}

func main() {

    jsonData := []byte(`{
        "title": "Sample Title",
        "author": "John Doe",
        "doi": "https://doi.org/10.1000/xyz123"
    }`)


    var doc Document


    err := json.Unmarshal(jsonData, &doc)
    if err != nil {
        log.Fatalf("Error unmarshalling JSON: %v", err)
    }


    fmt.Printf("DOI: %s\n", doc.DOI)
}

