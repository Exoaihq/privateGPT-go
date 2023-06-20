It's important to note that React is a JavaScript library for building user interfaces, and Go is a programming language mainly used for backend development. Converting a React component to Go doesn't make much sense, as they serve different purposes. However, I can provide you with a simple Go program that serves an HTML file containing the same content as your React component.

Create an `index.html` file with the following content:

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

Now, create a Go program to serve this HTML file:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			fmt.Println("Could not read index.html:", err)
			return
		}
		fmt.Fprint(w, string(content))
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

This Go program will serve the `index.html` file containing your content when accessed at `http://localhost:8080`.