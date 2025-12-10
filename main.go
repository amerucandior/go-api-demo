package main

import (
    "encoding/json"
    "net/http"
)

// helloHandler responds with a JSON {"hello":"world"} message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"hello": "world"}               // key-value to send
    json.NewEncoder(w).Encode(response)                           // encode map to JSON and write to response
}

func main() {
    http.HandleFunc("/hello", helloHandler)                       // register handler for "/hello" URL path
    http.ListenAndServe(":8080", nil)                             // start server on port 8080 (blocking call)
}
