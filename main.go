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

	BucketID, API_URL, AUTH := b2_auth(os.Args[1], os.Args[2])

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
				uuid := save_file(bodyBytes)
				UAUTH, UURL := b2_upload_url(BucketID, API_URL, AUTH)
				b2_upload(UAUTH, UURL, path, os.Args[3], uuid)
				fmt.Println("Uploaded: " + path)
				w.WriteHeader(http.StatusOK)
			}
		}

		// Multi part upload
		if r.Method == "POST" {

			if len(r.URL.Query()["uploadId"]) > 0 {
				finish_multi_part(r.URL.Query()["uploadId"][0], path)
				//UAUTH, UURL := b2_upload_url(BucketID, API_URL, AUTH)
				//status := b2_upload(UAUTH, UURL, strings.Split(path, "?")[0], "b2/x-auto", make([]byte, 0), os.Args[3], r.URL.Query()["uploadId"][0])
			} else {
				//ans := gen_xml_multi(UUID)
				//w.Write([]byte(ans))
			}
		}

		// Delete a file
		if r.Method == "DELETE" {
			b2_delete_file(AUTH, API_URL, path, BucketID, os.Args[3])
			fmt.Println("Deleted: " + path)
		}

		// Send an empty HEAD
		if r.Method == "HEAD" {
			fmt.Println("Gave some empty HEAD")
		}

	})

	http.ListenAndServe(os.Args[4]+":9000", nil)

}
