package main

import "bytes"
import "strings"
import "net/http"
import "io/ioutil"
import "encoding/json"

type hide_resp struct {
  FILEID string `json:"fileId"`
}

func b2_hide_file(AUTH string, API_URL string, BUCKETID string, FILE_NAME string, BucketNAME string) (string, int) {

  fileName := strings.TrimPrefix(FILE_NAME, "/" + BucketNAME + "/")

  client := &http.Client{}
	req := b2_hide_file_create_request(API_URL, BUCKETID, fileName)
  req.Header.Set("Authorization", AUTH)
	bodyString, status := b2_hide_file_request(req, client)

  var ans hide_resp
  json.Unmarshal([]byte(bodyString), &ans)

  return ans.FILEID, status

}

func b2_hide_file_create_request(api_url string, bucketid string, filename string) (*http.Request) {
	error_free := false
	var ans *http.Request
	for !error_free {
		req, err := http.NewRequest("POST", (api_url + "/b2api/v2/b2_hide_file"), bytes.NewReader([]byte("{\"bucketId\":\"" + bucketid + "\", \"fileName\":\"" + filename + "\"}")))
		if err == nil {
			ans = req
			error_free = true
		}
	}
	return ans
}

func b2_hide_file_request(req *http.Request, client *http.Client) (string, int) {
	error_free := false
	ans := ""
	int := 0
	for !error_free {
		resp, err := client.Do(req)
		if err == nil {
			if resp.StatusCode == 200 || resp.StatusCode == 404 {
				body, _ := ioutil.ReadAll(resp.Body)
				ans = string(body)
				int = resp.StatusCode
				resp.Body.Close()
				error_free = true
			}
		}
	}
	return ans, int
}
