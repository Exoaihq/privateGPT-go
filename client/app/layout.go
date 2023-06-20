It's important to note that the code you provided is in JavaScript (specifically, React) and deals with frontend web development. Go is a backend programming language and cannot directly replicate frontend code. However, I can provide you with a simple Go web server that can serve static files, which you can use as a starting point for your Go project.

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

Create a folder named `static` in the same directory as your Go file, and place your HTML, CSS, and JavaScript files there. This Go program will serve the files in the `static` folder when you access `http://localhost:8080` in your browser.