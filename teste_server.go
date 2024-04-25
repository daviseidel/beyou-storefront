
package main

import (
    "net/http"
)

func regionHandler(w http.ResponseWriter, r *http.Request) {
    // Return "br" as the response
    w.Write([]byte("br"))
}

func main() {
    // Register the regionHandler function to handle requests to "/stores/region"
    http.HandleFunc("/stores/region", regionHandler)

    // Start the HTTP server on port 8080
    http.ListenAndServe(":8080", nil)
}
