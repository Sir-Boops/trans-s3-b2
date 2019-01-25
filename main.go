package main

import "os"
import "fmt"
import "strings"
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

		// Upload a file
		if r.Method == "PUT" {
			// Check if it's normal or multi
			if len(r.URL.Query()["uploadId"]) > 0 {
				if len(r.URL.Query()["partNumber"]) > 0 {
					upload_part(r.URL.Query()["uploadId"][0], r.URL.Query()["partNumber"][0], bodyBytes) // Save the uploaded parts
				}
			} else {
				// Normal upload
				content_type := r.Header["Content-Type"][0]
				UAUTH, UURL := upload_url(BucketID, API_URL, AUTH)
				status := upload_file(UAUTH, UURL, path, content_type, bodyBytes, os.Args[3], "")

				if status == 200 {
					fmt.Println("Uploaded: " + path)
				} else {
					fmt.Println("Error uploading: " + path)
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}

		// Multi part upload
		if r.Method == "POST" {

			if len(r.URL.Query()["uploadId"]) > 0 {
				finish_multi_part(r.URL.Query()["uploadId"][0], path)
				UAUTH, UURL := upload_url(BucketID, API_URL, AUTH)
				status := upload_file(UAUTH, UURL, strings.Split(path, "?")[0], "b2/x-auto", make([]byte, 0), os.Args[3], r.URL.Query()["uploadId"][0])
				if status == 200 {
					fmt.Println("Uploaded Multipart: " + path)
				} else {
					fmt.Println("Error uploading Multipart: " + path)
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				UUID := start_multi_part()
				ans := gen_xml_multi(UUID)
				w.Write([]byte(ans))
			}
		}

		// Delete a file
		if r.Method == "DELETE" {
			status := rm_file(AUTH, API_URL, path, BucketID, os.Args[3])

			if status == 200 || status == 400 || status == 404 {
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

	http.ListenAndServe(os.Args[4]+":9000", nil)

}
