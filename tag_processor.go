// tag_processor.go
package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
)

type TagProcessor struct {
	config *Config
}

func NewTagProcessor(config *Config) *TagProcessor {
	return &TagProcessor{config: config}
}

func (tp *TagProcessor) ProcessFile(path string) {
	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		log.Printf("Error opening %s: %v\n", path, err)
		return
	}
	defer tag.Close()

	changed := false
	keepFrames := make(map[string][]id3v2.Frame)

	for frameID, frames := range tag.AllFrames() {
		for _, frame := range frames {
			if text := extractFrameText(frame); text != "" {
				trimmed := strings.TrimSpace(text)
				shouldRemove := false
				for _, s := range tp.config.RemoveStrings {
					if s == trimmed {
						shouldRemove = true
						changed = true
						break
					}
				}
				if !shouldRemove {
					keepFrames[frameID] = append(keepFrames[frameID], frame)
				}
			} else {
				keepFrames[frameID] = append(keepFrames[frameID], frame)
			}
		}
	}

	if !changed {
		return
	}

	tag.DeleteAllFrames()
	for id, frames := range keepFrames {
		for _, frame := range frames {
			tag.AddFrame(id, frame)
		}
	}

	if err := tag.Save(); err != nil {
		log.Printf("Error saving %s: %v\n", path, err)
	} else {
		log.Printf("Processed: %s\n", path)
	}
}

func extractFrameText(frame id3v2.Frame) string {
	switch f := frame.(type) {
	case id3v2.TextFrame:
		return f.Text
	case id3v2.CommentFrame:
		return f.Text
	case id3v2.UnsynchronisedTextFrame:
		return f.Text
	default:
		return ""
	}
}
