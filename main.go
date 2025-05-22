package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	for i := 0; i < 20; i++ {
		videoUrl := fmt.Sprintf("/video%d", i)
		videoPath := fmt.Sprintf("video%d.mp4", i)
		handleHttpRequest(videoUrl, videoPath)
	}

	log.Println("Serving video on http://0.0.0.0:8080/video")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHttpRequest(videoUrl string, videoPath string) {
	http.HandleFunc(videoUrl, func(w http.ResponseWriter, r *http.Request) {
		videoPath := videoPath

		file, err := os.Open(videoPath)
		if err != nil {
			http.Error(w, "Video not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			http.Error(w, "Unable to read file info", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Accept-Ranges", "bytes")
		http.ServeContent(w, r, filepath.Base(videoPath), stat.ModTime(), file)
	})
}
