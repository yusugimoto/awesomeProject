package main

import (
	"fmt"
	"net/http"
)

type Server struct {}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("input_value")
	h := `
<html><head><title>HTMLのフォーム</title></head><body>
  <form action="/" method="post">
	<input type="text" name="input_value">
	<input type="submit" name="送信"><br>
	入力値：` + v + `
  </form>
</body></html>
`
	fmt.Fprint(w, h)
}

func main() {
	http.ListenAndServe(":4000", Server{})
}
