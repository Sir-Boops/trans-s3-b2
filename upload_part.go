package main

import "io/ioutil"

func upload_part(ID string, PART string, bodyBytes []byte) {
	ioutil.WriteFile("gay/" + ID + "_" + PART, bodyBytes, 0600)
}
