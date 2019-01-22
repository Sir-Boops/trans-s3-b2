package main

import "bytes"
import "strings"
import "net/http"

func rm_file(AUTH string, API_URL string, PATH string, BucketID string, BucketNAME string) (int) {

  FILEID, HIDE_STATUS := hide_file(AUTH, API_URL, BucketID, PATH, BucketNAME)
  fileName := strings.TrimPrefix(PATH, "/" + BucketNAME + "/")

  ans := HIDE_STATUS

  if HIDE_STATUS == 200 {
    client := &http.Client{}
    req, _ := http.NewRequest("POST", (API_URL + "/b2api/v2/b2_delete_file_version"), bytes.NewReader([]byte("{\"fileName\":\"" + fileName + "\",\"fileId\":\"" + FILEID + "\"}")))
    req.Header.Set("Authorization", AUTH)
    resp, _ := client.Do(req)
    ans = resp.StatusCode
  }

  return ans
}
