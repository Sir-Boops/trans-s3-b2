package main

import "os"
import "bytes"
import "strings"
import "net/http"
import "io/ioutil"
import "crypto/sha1"
import "encoding/hex"

func upload_file(UAUTH string, UURL string, PATH string, content_type string, file []byte, bucket string, UUID string) (int) {

	file_data := file

	// Check if we have to read it ourself
	if len(file) == 0 {
		base_data, _ := os.Open("gay/" + UUID)
		data, _ := ioutil.ReadAll(base_data)
		base_data.Close()
		file_data = data
	}

  // Sum the file data
  sum := sha1.New()
  sum.Write(file_data)
  sumString := hex.EncodeToString(sum.Sum(nil))

  // Create custom client
  client := &http.Client{}
  req, _ := http.NewRequest("POST", UURL, bytes.NewReader(file_data))
  req.Header.Add("Authorization", UAUTH)
  req.Header.Add("X-Bz-File-Name", strings.TrimPrefix(PATH, "/" + bucket + "/"))
  req.Header.Add("Content-Type", content_type)
  req.Header.Add("X-Bz-Content-Sha1", sumString)
  req.Header.Add("X-Bz-Info-Author", "unknown")
  resp, _ := client.Do(req)
  return resp.StatusCode
}
