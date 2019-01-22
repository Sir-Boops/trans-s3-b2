package main

import "bytes"
import "strings"
import "net/http"
import "io/ioutil"
import "encoding/json"

type hide_resp struct {
  FILEID string `json:"fileId"`
}

func hide_file(AUTH string, API_URL string, BUCKETID string, FILE_NAME string, BucketNAME string) (string, int) {

  fileName := strings.TrimPrefix(FILE_NAME, "/" + BucketNAME + "/")

  client := &http.Client{}
  req, _ := http.NewRequest("POST", (API_URL + "/b2api/v2/b2_hide_file"), bytes.NewReader([]byte("{\"bucketId\":\"" + BUCKETID + "\", \"fileName\":\"" + fileName + "\"}")))
  req.Header.Set("Authorization", AUTH)
  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
  bodyString := string(body)

  var ans hide_resp
  json.Unmarshal([]byte(bodyString), &ans)

  return ans.FILEID, resp.StatusCode

}
