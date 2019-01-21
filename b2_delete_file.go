package main

import "net/http"

func rm_file(AUTH string, DL_URL string, PATH string) {

  client := &http.Client{}
  req, _ := http.NewRequest("GET", (DL_URL + "/file/" + PATH), nil)
  req.Header.Set("Authorization", AUTH)
  client.Do(req)

}
