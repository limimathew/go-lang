package main

import (
    "log"
    "net/http"
)

// temporary directory location
const TmpDir = "temp-images/";


func main() {

    // return a `.pdf` file for `/pdf` route
    http.HandleFunc( "/pdf", func( res http.ResponseWriter, req *http.Request ) {
        http.ServeFile( res, req, TmpDir + "/temp-images/upload-513598501.png");
    } )

    // start HTTP server with `http.DefaultServeMux` handler
    log.Fatal(http.ListenAndServe( ":9000", nil ))

}