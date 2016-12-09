package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"os"
	"os/user"
	"github.com/fatih/color"
	"path/filepath"
	"golang.org/x/sys/unix"
)

var dotFilesPath string
var shouldPark bool

var usr *user.User

func init() {
	var err error
	usr, err = user.Current()
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.StringVar(&dotFilesPath, "p", ".", "dotmgmt -p /path/to/dir")
	flag.BoolVar(&shouldPark, "park", false, "dotmgmt --park")
	flag.Parse()

	list := getFiles()

	if shouldPark {
		park(list)
		return
	}

	deploy(list)
}

func getFiles() []os.FileInfo {
	list, err := ioutil.ReadDir(dotFilesPath)
	if err != nil {
		panic(err)
	}

	return list
}

func park(list []os.FileInfo) {
	// Remove symlinks
	color.Green("%s", "Removing Symlinks...")
	fmt.Println("____")
	defer fmt.Println("____")
	for _, v := range list {
		newPath := filepath.Join(usr.HomeDir, v.Name())
		fmt.Println("Removing", color.RedString("%s", newPath))
		err := unix.Unlink(newPath)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func deploy(list []os.FileInfo) {
	// Symlink the files
	color.Green("%s", "Creating symlinks...")
	fmt.Println("____")
	defer fmt.Println("____")
	for _, v := range list {
		oldPath := filepath.Join(getPwd(), v.Name())
		newPath := filepath.Join(usr.HomeDir, v.Name())

		fmt.Println("Symlinking", color.RedString("%s", oldPath), color.BlueString("%s", "=>"), color.GreenString("%s", newPath))
		err := os.Symlink(oldPath, newPath)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}
