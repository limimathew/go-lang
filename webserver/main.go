package main

import (
	"fmt"
	"io/ioutil"

	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"log"
// 	"mime/multipart"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// )

// // Creates a new file upload http request with optional extra params
// func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = io.Copy(part, file)

// 	for key, val := range params {
// 		_ = writer.WriteField(key, val)
// 	}
// 	err = writer.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", uri, body)
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	return req, err
// }

// func main() {
// 	path, _ := os.Getwd()
// 	path += "/test.pdf"
// 	extraParams := map[string]string{
// 		"title":       "My Document",
// 		"author":      "Matt Aimonetti",
// 		"description": "A document with all the Go programming language secrets",
// 	}
// 	request, err := newfileUploadRequest("https://google.com/upload", extraParams, "temp-images", "/tmp/doc.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	client := &http.Client{}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		body := &bytes.Buffer{}
// 		_, err := body.ReadFrom(resp.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		resp.Body.Close()
// 		fmt.Println(resp.StatusCode)
// 		fmt.Println(resp.Header)

// 		fmt.Println(body)
// 	}
// }
