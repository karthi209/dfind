package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	folder := "/home/karthi209/workspace"

	hashMap := make(map[string][]string)

	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		if !info.IsDir() {
			hash, err := hashFile(path)
			if err != nil {
				fmt.Printf("Failed to hash %s: %v\n", path, err)
				return nil
			}

			hashMap[hash] = append(hashMap[hash], path)
		}

		return nil
	})

	fmt.Println("Duplicate files found:")
	for _, files := range hashMap {
		if len(files) > 1 {
			fmt.Println("-----")
			for _, f := range files {
				fmt.Println(f)
			}
		}
	}
}

// hashFile returns the SHA256 hash of a file
func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
