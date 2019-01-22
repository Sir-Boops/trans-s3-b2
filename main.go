package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"

func main() {

  // Arg 1 = ID
  // Arg 2 = Key
  // Arg 3 = bucket
  // Arg 4 is IP

  BucketID, API_URL, AUTH, DL_URL := get_keys(os.Args[1], os.Args[2])

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    bodyBytes, _ := ioutil.ReadAll(r.Body)
    path := r.URL.String()

    // Upload a file
    if r.Method == "PUT" {
      content_type := r.Header["Content-Type"][0]
      UAUTH, UURL := upload_url(BucketID, API_URL, AUTH)
      status := upload_file(UAUTH, UURL, path, content_type, bodyBytes, os.Args[3])

      if status == 200 {
        fmt.Println("Uploaded: " + path)
      } else {
        fmt.Println("Error uploading: " + path)
        w.WriteHeader(http.StatusInternalServerError)
      }
    }

    // Delete a file
    if r.Method == "DELETE" {
      status := rm_file(AUTH, DL_URL, path)

      if status == 200 {
        fmt.Println("Deleted: " + path)
      } else {
        fmt.Println("Failed to delete: " + path)
        w.WriteHeader(http.StatusInternalServerError)
      }
    }

    // Send an empty HEAD
    if r.Method == "HEAD" {
      fmt.Println("Gave some empty HEAD")
    }

  })

  http.ListenAndServe(os.Args[4] + ":9000", nil)

}
