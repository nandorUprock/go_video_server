package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		videoPath := "./test_file.mp4"

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

	log.Println("Serving video on http://0.0.0.0:8080/video")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
