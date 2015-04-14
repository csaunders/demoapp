package main

import "net/http"

func deal(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, dealUrl, http.StatusMovedPermanently)
}

func magic(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, magicUrl, http.StatusMovedPermanently)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(indexHtml))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/magic", magic)
	http.HandleFunc("/deal", deal)
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
  <p><a href="/magic">See some Magic</a></p>
  <p><a href="/deal">Not a fan?</a></p>
</body>
</html>
`
