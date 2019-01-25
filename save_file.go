package main

import "io/ioutil"
import uuid_gen "github.com/gofrs/uuid"

func save_file(bodybytes []byte) (string) {

	uuid := gen_uuid()
	save_file_inner(bodybytes, uuid)
	return uuid
}

func gen_uuid() (string) {
	error_free := false
	uuid := ""
	for !error_free {
		UUID, err := uuid_gen.NewV4()
		if err == nil {
			uuid = UUID.String()
			error_free = true
		}
	}
	return uuid
}

func save_file_inner(bodybytes []byte, uuid string) {
	error_free := false
	for !error_free {
		err := ioutil.WriteFile("gay/" + uuid, bodybytes, 0600)
		if err == nil {
			error_free = true
		}
	}
}
