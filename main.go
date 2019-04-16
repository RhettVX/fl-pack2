package main

import (
	"fl-library/files/packs"
	"fl-library/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	argc := len(os.Args)

	if argc < 2 {
		log.Fatal("Missing arguments")
	}

	for _, arg := range os.Args[1:] {

		// Grab absolute path
		path, err := filepath.Abs(arg)
		utils.Check(err)

		// Check if dir or file
		file, err := os.Stat(path)
		utils.Check(err)

		switch mode := file.Mode(); {

		// Pack folder to .pack2
		case mode.IsDir():
			var p packs.Pack2
			p.LoadFromDir(path)
			p.SortAssets()
			p.WritePack2("Packed", "Custom_x64_0")

		// Unpack .pack2
		case mode.IsRegular():
			var p packs.Pack2
			p.LoadFromFile(path)
			p.ApplyHash()
			p.Unpack("Unpacked")
		}
	}
}
