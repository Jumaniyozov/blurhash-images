package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/buckket/go-blurhash"
)

func HashHandler(w http.ResponseWriter, r *http.Request) {
	imageFile, _ := os.Open("test.png")
	loadedImage, err := png.Decode(imageFile)
	str, _ := blurhash.Encode(4, 3, loadedImage)
	if err != nil {
		log.Println(err)
	}

	img, err := blurhash.Decode(str, 300, 500, 1)
	if err != nil {
		log.Println(err)
	}
	f, _ := os.Create("test_blur.png")
	_ = png.Encode(f, img)

	// Convert the png image to webp using cwebp
	cmd := exec.Command("cwebp", "test_blur.png", "-o", "test_blur.webp")
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
	

	// Get the x and y components used for encoding a given BlurHash
	x, y, err := blurhash.Components(str)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("xComponents: %d, yComponents: %d", x, y)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	mux.HandleFunc("GET /hash", HashHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
