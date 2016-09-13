package main

import (
	"fmt"
	"log"

	"github.com/nicksnyder/go-i18n/i18n"
)

var (
	availableLanuages = []string{
		"de-DE",
		"en-US",
	}
)

func init() {
	for _, lang := range availableLanuages {
		filename := fmt.Sprintf("i18n/%s.all.yaml", lang)
		langContent, err := Asset(filename)
		if err != nil {
			log.Fatalf("Could not read source of language file %s: %s", lang, err)
		}
		if err := i18n.ParseTranslationFileBytes(filename, langContent); err != nil {
			log.Fatalf("Could not parse language file %s: %s", lang, err)
		}
	}
}
