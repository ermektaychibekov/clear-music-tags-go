// file_processor.go
package main

import (
	"log"
	"os"
	"path/filepath"
)

type FileProcessor struct {
	tagProcessor *TagProcessor
}

func NewFileProcessor(config *Config) *FileProcessor {
	return &FileProcessor{
		tagProcessor: NewTagProcessor(config),
	}
}

func (fp *FileProcessor) ProcessPaths() {
	for _, path := range fp.tagProcessor.config.Paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Printf("Path not found: %s\n", path)
			continue
		}
		fp.processPath(path)
	}
}

func (fp *FileProcessor) processPath(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return
	}

	if fileInfo.IsDir() {
		filepath.Walk(path, func(subPath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fp.processFile(subPath)
			}
			return nil
		})
	} else {
		fp.processFile(path)
	}
}

func (fp *FileProcessor) processFile(path string) {
	switch filepath.Ext(path) {
	case ".mp3":
		fp.tagProcessor.ProcessFile(path)
	}
}
