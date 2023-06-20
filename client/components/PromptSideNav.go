It's important to note that React is a JavaScript library for building user interfaces, and Go is a programming language primarily used for backend development. Converting a React component to Go directly is not possible, as they serve different purposes.

However, I can provide you with a simple Go program that serves an HTML file containing the same content as the given React component. First, create an `index.html` file with the following content:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Prompt Side Nav</title>
</head>
<body>
    <div class='prompt-bg p-3 mx-4'>
        <span>Prompt</span>
        <div style="color:gray">You are privateGPT which built with Langchain and GPT4All and Llamacpp</div>
    </div>
</body>
</html>
```

Now, create a Go program that serves this HTML file:

```go
package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Could not read file", http.StatusInternalServerError)
			return
		}
		w.Write(content)
	})

	http.ListenAndServe(":8080", nil)
}
```

This Go program will serve the `index.html` file on `http://localhost:8080`. Note that this is not a direct conversion of the React component, but a way to serve the same content using Go.