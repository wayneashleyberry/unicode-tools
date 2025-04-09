package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type LDML struct {
	Annotations []Annotation `xml:"annotations>annotation"`
}

type Annotation struct {
	CP   string `xml:"cp,attr"`
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
}

func main() {
	var ldml LDML

	decoder := xml.NewDecoder(os.Stdin)
	if err := decoder.Decode(&ldml); err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "error decoding XML: %v\n", err)
		os.Exit(1)
	}

	seen := make(map[string]bool)
	for _, ann := range ldml.Annotations {
		if ann.Type == "tts" {
			continue
		}

		// Avoid printing duplicates (some emojis may appear more than once)
		if seen[ann.CP] {
			continue
		}

		seen[ann.CP] = true
		text := strings.TrimSpace(ann.Text)
		fmt.Printf("%s\t%s\n", ann.CP, text)
	}
}
