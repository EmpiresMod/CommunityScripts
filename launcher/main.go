package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// URL which updates will be downloaded from
const defaultURL = "http://apollo.firebit.co.uk/~dc0/update"

// Some vars
var (
	Debug     bool
	Update    bool
	UpdateURL string
	DirPath   string
	Community bool
	Vanilla   bool
)

func LaunchCommunity() (err error) {

	if Debug {

		log.Print("Launching Empires with Community Scripts!")
	}

	m, err := GetManifest(fmt.Sprintf("%s/%s", DirPath, "manifest.json"))
	if err != nil {

		return
	}
	m.BasePath = DirPath
	m.BaseURL = UpdateURL

	if err = m.Apply(); err != nil {

		return
	}

	if err = exec.Command("cmd", "/c", "START", "steam://rungameid/17740").Start(); err != nil {

		return
	}

	return
}

func LaunchVanilla() (err error) {

	if Debug {

		log.Print("Launching Empires Vanilla!")
	}

	m, err := GetManifest(fmt.Sprintf("%s/%s", DirPath, "manifest.json"))
	if err != nil {

		return
	}
	m.BasePath = DirPath
	m.BaseURL = UpdateURL

	if err = m.UnApply(); err != nil {

		return
	}

	if err = exec.Command("cmd", "/c", "START", "steam://rungameid/17740").Start(); err != nil {

		return
	}

	return
}

func init() {

	flag.BoolVar(&Debug, "d", false, "debugging/verbose information")
	flag.BoolVar(&Update, "u", true, "check for updates")
	flag.StringVar(&UpdateURL, "url", defaultURL, "url to fetch updates from")
	flag.StringVar(&DirPath, "p", "./", "path to empires folder")
	flag.BoolVar(&Community, "c", false, "launch empires with community scripts")
	flag.BoolVar(&Vanilla, "v", false, "launch vanilla empires")
	flag.Parse()

	if Community {

		Vanilla = false
	}
}

func main() {

	if Update {

		if Debug {

			log.Print("Checking for updates to manifest!")
		}

		err := UpdateManifest(
			fmt.Sprintf("%s/%s", DirPath, "manifest.json"),
			fmt.Sprintf("%s/%s", UpdateURL, "manifest.json"))
		if err != nil {

			return
		}
	}

	if Community {

		if err := LaunchCommunity(); err != nil {

			log.Fatal(err)
		}

		return

	} else if Vanilla {

		if err := LaunchVanilla(); err != nil {

			log.Fatal(err)
		}

		return
	}

	if err := ShowUI(); err != nil {

		log.Fatal(err)
	}
}
