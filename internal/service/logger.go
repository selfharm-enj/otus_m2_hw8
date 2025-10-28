package service

import (
	"context"
	"fmt"
	"time"

	"github.com/selfharm-enj/otus_m2_hw8/internal/repository"
)

func LogChanges(ctx context.Context) {
	var (
		lastFiles  int
		lastImages int
	)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			currentFiles, currentImages := repository.FilesImagesLen()
			if currentFiles != lastFiles {
				fmt.Printf("Added %v Files\t", currentFiles-lastFiles)
				lastFiles = currentFiles
			}
			if currentImages != lastImages {
				fmt.Printf("Added %d Images\n", currentImages-lastImages)
				lastImages = currentImages
			}
			repository.FilesImagesItems()
		}
	}
}
