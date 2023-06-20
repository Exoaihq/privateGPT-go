It's important to note that React and Go serve different purposes. React is a JavaScript library for building user interfaces, while Go is a statically typed, compiled programming language designed for systems programming. Converting a React component to Go doesn't make sense as they are not directly interchangeable.

However, if you want to create a Go server that handles the same API calls as the React component, you can use the following example:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ingestData() {
	resp, err := http.Get("http://localhost:5000/ingest")
	if err != nil {
		log.Println("Error ingesting data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Ingest Data:", string(body))
}

func downloadModel() {
	resp, err := http.Get("http://localhost:5000/download_model")
	if err != nil {
		log.Println("Error downloading model:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Download Model:", string(body))
}

func uploadDocument(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	resp, err := http.Post("http://localhost:5000/upload_doc", "multipart/form-data", file)
	if err != nil {
		log.Println("Error uploading document:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Upload Document:", string(body))
}

func main() {
	// Replace "path/to/your/file" with the actual file path you want to upload
	uploadDocument("path/to/your/file")
	ingestData()
	downloadModel()
}
```

This Go code provides functions for ingesting data, downloading a model, and uploading a document, similar to the functionality in the provided React component. However, this code runs on the server-side, not as a user interface.