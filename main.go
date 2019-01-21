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

  BucketID, API_URL, AUTH := get_keys(os.Args[1], os.Args[2])

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    bodyBytes, _ := ioutil.ReadAll(r.Body)
    path := r.URL.String()
    content_type := r.Header["Content-Type"][0]
    UAUTH, UURL := upload_url(BucketID, API_URL, AUTH)
    upload_file(UAUTH, UURL, path, content_type, bodyBytes, os.Args[3])
    fmt.Println("Uploaded: " + path)
  })

  http.ListenAndServe(os.Args[4] + ":9000", nil)

}
