package main

import (
	"flag"
	"log"
	"os/exec"
	"path/filepath"
)

// Some constants
const (
	defaultURL = "https://github.com/EmpiresMod/CommunityScripts/tree/master/update"
)

// Some vars
var (
	debug     bool
	update    bool
	updateURL string
	path      string
	improved  bool
	vanilla   bool
)

func init() {

	flag.BoolVar(&debug, "d", false, "debugging/verbose information")
	flag.BoolVar(&update, "u", true, "check for updates")
	flag.StringVar(&updateURL, "url", defaultURL, "url to fetch updates from")
	flag.StringVar(&path, "p", "./", "path to empires folder")
	flag.BoolVar(&improved, "i", false, "launch empires improved")
	flag.BoolVar(&vanilla, "v", true, "launch vanilla empires")
	flag.Parse()

	if improved {

		vanilla = false
	}
}

func main() {

	if improved {

		if update {

			if debug {

				log.Print("Checking for updates to manifest!")
			}

			if err := UpdateManifest(filepath.Clean(path), updateURL); err != nil {

				log.Fatal(err)
			}
		}

		if debug {

			log.Print("Launching Empires Improved!")
		}

		m, err := GetManifest(filepath.Clean(path))
		if err != nil {

			log.Fatal(err)
		}
		m.BasePath = filepath.Clean(path)
		m.BaseURL = updateURL

		if err = m.Apply(); err != nil {

			log.Fatal(err)
		}

		if err := exec.Command("cmd", "/c", "START", "steam://rungameid/17740").Start(); err != nil {

			log.Fatal(err)
		}
	}

	if vanilla {

		if debug {

			log.Print("Launching Empires Vanilla!")
		}

		m, err := GetManifest(filepath.Clean(path))
		if err != nil {

			log.Fatal(err)
		}
		m.BasePath = filepath.Clean(path)
		m.BaseURL = updateURL

		if err = m.UnApply(); err != nil {

			log.Fatal(err)
		}

		if err := exec.Command("cmd", "/c", "START", "steam://rungameid/17740").Start(); err != nil {

			log.Fatal(err)
		}
	}
}
