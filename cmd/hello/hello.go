package main

import (
"fmt"
"log"
"net/http"
)

// HelloServer responds to requests with the given URL path.
func HelloServer(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello!! This is Pyxis, you requested: %s", r.URL.Path)
log.Printf("Received request for path: %s", r.URL.Path)
}

func main() {
var addr string = ":80"
handler := http.HandlerFunc(HelloServer)
if err := http.ListenAndServe(addr, handler); err != nil {
log.Fatalf("Could not listen on port %s %v", addr, err)
}
}