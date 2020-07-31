package main

import "fmt"

func goToFunction() (int, string) {
	err := nil
	filename := ""
	data := ""

	err, filename = GetFileName()
	if err != nil {
		goto Done
	}

	err, data = ReadFile(filename)
	if err != nil {
		goto Done
	}

Done:
	return err, data
}
