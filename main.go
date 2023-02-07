package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide image path")
		return
	}

	if len(os.Args) == 2 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error: %s\nPress enter to exit...", err.Error())
			fmt.Scanln()
			return
		}
		newBytes := toSquare(bytes)
		format := getDataType(&newBytes)
		name := strings.Split(filepath.Base(os.Args[1]), ".")[0] + "-square"
		if format == "image/png" {
			name += ".png"
		} else {
			name += ".jpg"
		}

		os.WriteFile(name, newBytes, 0644)
		return
	}

	if len(os.Args) > 2 {
		if os.Args[2] != "png" && os.Args[2] != "jpg" {
			fmt.Printf("Error: Unsupported output format: %s\nPress enter to exit...", os.Args[1])
			fmt.Scanln()
			return
		}

		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error: %s\nPress enter tso exit...", err.Error())
			fmt.Scanln()
			return
		}

		newBytes := toSquare(bytes)
		format := getDataType(&newBytes)
		name := strings.Split(filepath.Base(os.Args[1]), ".")[0] + "-square"

		if os.Args[2] == "jpg" && format == "image/png" {
			newBytes = pngToJpeg(newBytes)
			name += ".jpg"
		} else if os.Args[2] == "png" && format == "image/jpeg" {
			newBytes = jpegToPng(newBytes)
			name += ".png"
		}

		os.WriteFile(name, newBytes, 0644)
		return
	}
}
