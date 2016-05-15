package main

import (
	"fmt"
	"os"
	//"github.com/JeffHask/stringutil"
	"path/filepath"
	//"flag"
)

func main() {
	//fmt.Printf("Hello, world.\n")
	error := os.Chdir("..")
	if(error != nil) {
		fmt.Printf("blarf, error", error)
		return
	}
	s := os.Args[0]
	for i := 0; s != "C:\\"; i++ {
		fmt.Printf("%d\n", i)
		fmt.Printf(s + "\n")
		s = filepath.Dir(s)
	}
	fmt.Printf(s)
	// os.Chdir("..")
	//os.Mkdir("blahz", os.ModeDir)
	//traverse(s)
	//fmt.Printf(stringutil.Reverse("!oG, olleH\n"))
}

func traverse(s string) {
	filepath.Walk(s, visit)
}
	
func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}