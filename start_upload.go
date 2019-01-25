package main

import "github.com/gofrs/uuid"

func start_multi_part() (string) {
	UUID, _ := uuid.NewV4()
	UUID_STRING := UUID.String()
	return UUID_STRING
}
