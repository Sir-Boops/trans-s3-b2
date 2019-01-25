package main

import "os"
import "strconv"
import "io/ioutil"

func finish_multi_part(UUID string, PATH string) {

	FILE := make([]byte, 0)

	done := false
	if _, err := os.Stat("gay/" + UUID); err != nil {
		for i := 1; !done; i++ {
			if _, err := os.Stat("gay/" + UUID + "_" + strconv.Itoa(i)); err == nil {
				base_data, _ := os.Open("gay/" + UUID + "_" + strconv.Itoa(i))
				data, _ := ioutil.ReadAll(base_data)
				base_data.Close()
				FILE = append(FILE, data...)
				os.Remove("gay/" + UUID + "_" + strconv.Itoa(i))
			} else {
				done = true
			}
		}
		ioutil.WriteFile("gay/" + UUID, FILE, 0600)
	}
}
