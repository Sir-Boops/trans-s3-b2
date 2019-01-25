package main

import "bytes"
import "strings"
import "net/http"
import "io/ioutil"
import "crypto/sha1"
import "encoding/hex"

func b2_upload(UAUTH string, UURL string, PATH string, bucket string, uuid string) {

	// Read the file
	file_bytes := read_file(uuid)

  // Sum the file data
  sum := sha1.New()
  sum.Write(file_bytes)
  sumString := hex.EncodeToString(sum.Sum(nil))

  // Create custom client
  client := &http.Client{}
  req := create_request(UURL, file_bytes)
  req.Header.Add("Authorization", UAUTH)
  req.Header.Add("X-Bz-File-Name", strings.TrimPrefix(PATH, "/" + bucket + "/"))
  req.Header.Add("Content-Type", "b2/x-auto")
  req.Header.Add("X-Bz-Content-Sha1", sumString)
  req.Header.Add("X-Bz-Info-Author", "unknown")
  send_request(req, client)
}

func read_file(uuid string) ([]byte) {
	error_free := false
	ans := make([]byte, 0)
	for !error_free {
		file_bytes, err := ioutil.ReadFile("gay/" + uuid)
		if err == nil {
			ans = file_bytes
			error_free = true
		}
	}
	return ans
}

func create_request(UURL string, file_bytes []byte) (*http.Request) {
	error_free := false
	var ans *http.Request
	for !error_free {
		req, err := http.NewRequest("POST", UURL, bytes.NewReader(file_bytes))
		if err == nil {
			ans = req
			error_free = true
		}
	}
	return ans
}

func send_request(req *http.Request, client *http.Client) {
	error_free := false
	for !error_free {
		resp, err := client.Do(req)
		if err == nil {
			if resp.StatusCode == 200 {
				resp.Body.Close()
				error_free = true
			}
		}
	}
}
