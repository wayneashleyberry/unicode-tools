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
		emoji := ann.CP
		if seen[emoji] {
			continue
		}

		seen[emoji] = true
		text := strings.TrimSpace(ann.Text)
		fmt.Printf("%s", emoji)
		fmt.Printf("\t")
		for _, r := range emoji {
			fmt.Printf("U+%04X | ", r)
		}
		fmt.Printf("%s\n", text)
	}
}
