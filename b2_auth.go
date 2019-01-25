package main

import "net/http"
import "io/ioutil"
import "encoding/base64"
import "encoding/json"

type Allowed struct {
  BucketID string `json:"bucketId"`
}

type Base struct {
  Allowed Allowed `json:"allowed"`
  API_URL string `json:"apiUrl"`
  AUTH string `json:"authorizationToken"`
  DLURL string `json:"downloadUrl"`
}

func b2_auth(id string, key string) (string, string, string) {

  // Create custom client
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://api.backblazeb2.com/b2api/v2/b2_authorize_account", nil)
  req.Header.Set("Authorization", ("Basic " + base64.StdEncoding.EncodeToString([]byte(id + ":" + key))))
  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
  bodyString := string(body)

  var ans Base
  json.Unmarshal([]byte(bodyString), &ans)

  return ans.Allowed.BucketID, ans.API_URL, ans.AUTH

}
