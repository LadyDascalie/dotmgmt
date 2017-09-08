package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/fatih/color"
	"golang.org/x/sys/unix"
)

var (
	dotFilesPath string
	shouldMake   string // the file that should be made
	shouldRemove string // the file that should be removed
	shouldReset  bool
	usr          *user.User // the current user
)

func init() {
	var err error
	usr, err = user.Current()
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.StringVar(&dotFilesPath, "p", ".", "dotmgmt -p /path/to/dir")
	flag.StringVar(&shouldMake, "make", "", "dotmgmt -make /path/to/dir")
	flag.StringVar(&shouldRemove, "del", "", "dotmgmt -del /path/to/file")
	flag.BoolVar(&shouldReset, "reset", false, "dotmgmt -reset")
	flag.Parse()

	safetyCheck()

	// move and symlink the file
	if shouldMake != "" {
		moveAndSymlink()
		return
	}

	if shouldRemove != "" {
		undoSymlink()
		return
	}

	list := getFiles()

	if shouldReset {
		removeSymlink(list)
		return
	}

	makeSymlinks(list)
}

// make sure we aren't going to break things
func safetyCheck() {
	pwd := getPwd()
	danger := []string{
		"/", "/home",
		"/usr", "/usr/bin", "/usr/local/bin",
		"/etc", "/private/etc",
		"/tmp", "/private/tmp",
		"/var", "/private/var",
		"/sbin", "/private",
	}

	for _, p := range danger {
		if pwd == p {
			color.Red("%s", "dotmgmt: Dangerous folder. Refusing to run.")
			os.Exit(1)
		} else if pwd == usr.HomeDir {
			color.Red("%s", "dotmgmt: Cannot run on the home directory")
			os.Exit(1)
		}
	}
}

//getFiles gets the list of files in the dotfiles folder
func getFiles() []os.FileInfo {
	list, err := ioutil.ReadDir(derive(dotFilesPath))
	if err != nil {
		color.Red("Could not get list of files")
		os.Exit(1)
	}

	// filter out .git and .gitignore
	for k, f := range list {
		if f.Name() == ".git" || f.Name() == ".gitignore" {
			list[k] = list[len(list)-1]
			list = list[:len(list)-1]
		}
	}

	return list
}

func moveAndSymlink() {
	var list []os.FileInfo

	file := store(shouldMake)
	i, err := os.Stat(file)

	if err != nil {
		color.Red("No such file or directory")
		os.Exit(1)
	}
	list = append(list, i)

	makeSymlinks(list)
}

func undoSymlink() {
	var list []os.FileInfo

	fname := filepath.Join(dotFilesPath, filepath.Base(shouldRemove))

	i, err := os.Stat(fname)
	if err != nil {
		color.Red("No such file or directory")
		os.Exit(1)
	}
	list = append(list, i)

	removeSymlink(list)
}

func derive(fname string) string {
	var err error
	fname, err = filepath.Abs(fname)
	if err != nil {
		color.Red("Could not derive path for file %s", fname)
		os.Exit(1)
	}

	return fname
}

func store(fname string) (newname string) {
	oldname := derive(fname)
	newname = filepath.Join(dotFilesPath, filepath.Base(fname))

	err := os.Rename(oldname, newname)
	if err != nil {
		color.Red("Could nor rename file")
		os.Exit(1)
	}

	return newname
}

// removeSymlink removes the symlinks from the user's home folder
func removeSymlink(list []os.FileInfo) {
	// Notify user and defer closing statement
	color.Green("%s", "Removing Symlinks...")
	fmt.Println("____")
	defer fmt.Println("____")

	// Iterate over files...
	for _, v := range list {
		oldPath := filepath.Join(getPwd(), v.Name())
		newPath := filepath.Join(usr.HomeDir, v.Name())

		if oldPath != newPath {
			fmt.Println("Removing", color.RedString("%s", newPath))
			err := unix.Unlink(newPath)
			if err != nil {
				fmt.Println(newPath, err)
				os.Exit(1)
			}

		}
	}
}

func makeSymlinks(list []os.FileInfo) {
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
		color.Red("Could not get current working directory path")
		os.Exit(1)
	}
	return pwd
}
