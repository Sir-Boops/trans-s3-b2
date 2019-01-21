package main

import "bytes"
import "net/http"
import "io/ioutil"
import "encoding/json"

type Main struct {
  UAUTH string `json:"authorizationToken"`
  UURL string `json:"uploadUrl"`
}

func upload_url(BUCKETID string, API_URL string, AUTH string) (string, string) {

  // Create custom client
  client := &http.Client{}
  req, _ := http.NewRequest("POST", (API_URL + "/b2api/v2/b2_get_upload_url"), bytes.NewReader([]byte("{\"bucketId\":\"" + BUCKETID + "\"}")))
  req.Header.Add("Authorization", AUTH)
  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
  bodyString := string(body)

  var ans Main
  json.Unmarshal([]byte(bodyString), &ans)

  return ans.UAUTH, ans.UURL
}
