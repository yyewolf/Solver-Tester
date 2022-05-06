package main

import (
	"io/ioutil"
	"strings"
)

var wordleDictionary []string

func openDict() error {
	// On ouvre juste le fichier dico là
	data, err := ioutil.ReadFile(*dictionary)
	if err != nil {
		return err
	}

	content := string(data)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if len(line) == *size {
			wordleDictionary = append(wordleDictionary, line)
		}
	}
	return nil
}
