package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}
	w.Write([]byte(indexHtml))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

const magicUrl string = "https://dl.dropboxusercontent.com/u/113966/gifs/glitter-magic.gif"
const dealUrl string = "https://dl.dropboxusercontent.com/u/113966/gifs/dealwithitice.gif"
const indexHtml string = `
<!DOCTYPE html>
<html>
<head>
  <title>Demo Go App</title>
</head>
<body>
  <p><a href="/magic.gif">See some Magic</a></p>
  <p><a href="/deal.gif">Not a fan?</a></p>
</body>
</html>
`
