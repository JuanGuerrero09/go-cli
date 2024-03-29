package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type config struct {
	//extenstion to filter out
	ext string
	// min file size
	size int64
	//list files
	list bool
}

func main() {
	root := flag.String("root", "", "Root directory to start")
	ext := flag.String("ext", "", "File extension to filter out")
	list := flag.Bool("list", false, "List files only")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	c := &config{
		ext:  *ext,
		list: *list,
		size: *size,
	}

	if err := run(*root, os.Stdout, *c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}
			// If list was explicitly set, don't do anything else
			if cfg.list {
				return listFile(path, out)
			}
			// List is the default option if nothing else was set
			return listFile(path, out)
		})
}
