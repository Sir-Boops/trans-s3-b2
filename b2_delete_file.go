package main

import "net/http"

func rm_file(AUTH string, DL_URL string, PATH string) (int) {

  client := &http.Client{}
  req, _ := http.NewRequest("GET", (DL_URL + "/file/" + PATH), nil)
  req.Header.Set("Authorization", AUTH)
  resp, _ := client.Do(req)
  return resp.StatusCode

}
