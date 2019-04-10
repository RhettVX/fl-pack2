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

	for _, a := range os.Args[1:] {

		// Grab absolute path
		path, err := filepath.Abs(a)
		utils.Check(err)

		// Check if dir or file
		file, err := os.Stat(path)
		utils.Check(err)

		switch mode := file.Mode(); {

		case mode.IsDir():
			var p packs.Pack2
			p.LoadFromDir(a)
			p.Write()

			// for _, x := range p.Assets {
			// 	println(x.Name)
			// }

		case mode.IsRegular():
			var p packs.Pack2
			p.LoadFromFile(a)
			p.ApplyHash()
			p.Unpack("Unpacked")
		}
	}
}
