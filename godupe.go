package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/corona10/goimagehash"
)

func main() {

	dt := flag.Int("d", 10, "Distance threshold to mark as duplicate")
	s_size := flag.Int("s", 16, "Image sample size")
	flag.Parse()

	imgs := make(map[string]*goimagehash.ExtImageHash)
	dupes := make(map[string][]string)
	allc := make(map[string]bool)
	dir, err := os.Getwd()
	files, err := filepath.Glob(dir + "/*")
	if err != nil {
		log.Fatal(err)
	}

	// build a map of all hashes
	for _, f := range files {
		file1, _ := os.Open(f)
		img1, err := jpeg.Decode(file1)
		if err != nil {
		} else {
			imgs[f], _ = goimagehash.ExtAverageHash(img1, *s_size, *s_size)
		}
		file1.Close()
	}
	// compare all of our collected hashes
	for img1, hash1 := range imgs {
		for img2, hash2 := range imgs {
			if img2 != img1 {
				if _, ok := allc[img2]; ok {
				} else {
					distance, _ := hash2.Distance(hash1)
					if distance < *dt {
						dupes[img1] = append(dupes[img1], img2)
					}
				}
			}
		}
		allc[img1] = true
		// we dont need to compare the children to others because the parent
		// would have picked up any similarities itself...we're just looking
		// for duplicates, not similar images
		for _, img := range dupes[img1] {
			allc[img] = true
		}
	}

	// print all of our dupes, 1 per line
	// use external image viewer to confirm dupes:
	// godupe . | sxiv -
	for d, s := range dupes {
		fmt.Println(d)
		for _, img := range s {
			fmt.Println(img)
		}
	}
}
