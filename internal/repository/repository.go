package repository

import (
	"fmt"
	"sync"

	"github.com/selfharm-enj/otus_m2_hw8/internal/model"
)

var (
	Files  []model.File
	Images []model.Image
	mu     sync.Mutex
)

func AddData(data model.IDReader) {
	mu.Lock()
	defer mu.Unlock()

	switch v := data.(type) {
	case *model.File:
		Files = append(Files, *v)
	case *model.Image:
		Images = append(Images, *v)
	}
}

func FilesImagesLen() (int, int) {
	mu.Lock()
	defer mu.Unlock()
	return len(Files), len(Images)
}

func FilesImagesItems() {
	fmt.Printf("Files: %v\t", Files)
	fmt.Printf("Images: %v\n\n", Images)
}
