package main

import "fmt"
import "path/filepath"
import "log"
import "github.com/corona10/goimagehash"
import "image/jpeg"
import "os"

func main() {
    imgs := make(map[string]*goimagehash.ExtImageHash)
    dupes := make(map[string][]string)
    allc := make(map[string]bool)
    files, err := filepath.Glob(os.Args[1] + "/*")
    if err != nil {
        log.Fatal(err)
    }
    
    // build a map of all hashes
    for _, f := range files {
        //fmt.Println("Hashing", f)
        file1, _ := os.Open(f)
        img1, err := jpeg.Decode(file1)
        if err != nil {
        } else {
            imgs[f], _ = goimagehash.ExtAverageHash(img1,16,16)
        }
        file1.Close()
        //fmt.Println("Hashed %s", f)
    }

    // compare all of our collected hashes
    for img1, hash1 := range imgs {
        for img2, hash2 := range imgs {
            if img2 != img1 {
                if _, ok := allc[img2]; ok {
                } else {
                    distance, _ := hash2.Distance(hash1)
                    if distance < 10 {
                        // check if the dupe is already in the map
                        if _, ok := dupes[img2]; ok {
                        } else {
                            dupes[img1] = append(dupes[img1], img2)
                        }
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

    // print all of our dupes
    for d, s := range dupes {
        fmt.Println(d)
        for _, img := range s {
            fmt.Println(img)
        }
    }
}

