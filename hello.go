package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"go-id3-master/src/id3"
)

var fileMap map[string]int

var count = 0

func main() {
	fileMap = make(map[string]int)
	error := os.Chdir("..")
	if(error != nil) {
		fmt.Printf("%s", error)
		return
	}
	s := os.Args[1]
	traverse(s)
	fmt.Println(fileMap)
}

func traverse(s string) {
	filepath.Walk(s, visit)
	fmt.Printf("%d Songs\n", count)
	for i, x := range fileMap {
		var percent float64 = float64(x) / float64(count)
		fmt.Printf("%d Songs in %s (%.3f%%)\n", x, i, percent * float64(100))
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if strings.Contains(filepath.Base(path),".mp3") {
		fd, _ := os.Open(path)
		defer fd.Close()
		file := id3.Read(fd)
		if file != nil {
			count++
			v := ""
			if len(os.Args) > 2 {
				v = os.Args[2]
			}
			switch v {
			case "Album" :
				v = file.Album
				fileMap[file.Album] = fileMap[file.Album] + 1
			case "Artist":
				v = file.Artist
				fileMap[file.Artist] = fileMap[file.Artist] + 1
			case "Genre" :
				v = file.Genre
				fileMap[file.Genre] = fileMap[file.Genre] + 1
			case "Year" :
				v = file.Year
				fileMap[file.Year] = fileMap[file.Year] + 1
			default:
				//TODO figure out default case
				v = "REKT"
			}
			fmt.Println(v)
		}
	}
	return nil
}