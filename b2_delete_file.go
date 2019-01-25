package main

import "bytes"
import "strings"
import "net/http"

func b2_delete_file(AUTH string, API_URL string, PATH string, BucketID string, BucketNAME string) {

  FILEID, HIDE_STATUS := b2_hide_file(AUTH, API_URL, BucketID, PATH, BucketNAME)
  fileName := strings.TrimPrefix(PATH, "/" + BucketNAME + "/")

  if HIDE_STATUS == 200 {
    client := &http.Client{}
		req := b2_delete_file_request(API_URL, fileName, FILEID)
    req.Header.Set("Authorization", AUTH)
    b2_delete_file_send_request(req, client)
  }
}

func b2_delete_file_request(api_url string, filename string, fileid string) (*http.Request) {
	error_free := false
	var ans *http.Request
	for !error_free {
		req, err := http.NewRequest("POST", (api_url + "/b2api/v2/b2_delete_file_version"), bytes.NewReader([]byte("{\"fileName\":\"" + filename + "\", \"fileId\":\"" + fileid + "\"}")))
		if err == nil {
			ans = req
			error_free = true
		}
	}
	return ans
}

func b2_delete_file_send_request(req *http.Request, client *http.Client) {
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
