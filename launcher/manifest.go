package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Manifest struct {

	// base URL
	BaseURL string

	// base path
	BasePath string

	// array holding the files in the manifest
	Files []File
}

type File struct {

	// relative to the base path
	Path string

	// is this item a directory and not a file?
	Dir bool

	// sha512 sum
	HashSum string

	// where to download vpk from? Relative to base URL
	URL string

	// can this file be disabled?
	CanDisable bool

	// should a backup be made before modifying this file?
	Backup bool
}

func (m *Manifest) Apply() (err error) {

	for _, v := range m.Files {

		p := fmt.Sprintf("%s/%s", m.BasePath, v.Path)
		u := fmt.Sprintf("%s/%s", m.BaseURL, v.URL)

		if v.Dir {

			if !FileExists(p) {

				if debug {

					log.Printf("Making new directory: %s", p)
				}

				if err = os.Mkdir(p, 0755); err != nil {

					return err
				}
			}

			continue
		}

		if !FileExists(p) {

			if debug {

				log.Printf("File does not exist: %s", p)
			}

			f := fmt.Sprintf("%s.disable", p)

			if FileExists(f) {

				if err = os.Rename(f, p); err != nil {

					return err
				}
			}

			b, err := FetchRemoteFile(u)
			if err != nil {

				return err
			}

			err = ioutil.WriteFile(p, b, 0644)
			if err != nil {

				return err
			}

			continue
		}

		h, err := HashFile(p)
		if err != nil {

			return err
		}
		if h != v.HashSum {

			if debug {

				log.Printf("File is an older version: %s", p)
			}

			b, err := FetchRemoteFile(u)
			if err != nil {

				return err
			}

			if v.Backup {

				f := fmt.Sprintf("%s.bck", p)

				if FileExists(f) {

					if err = os.Remove(f); err != nil {

						return err
					}
				}

				if err = os.Rename(p, f); err != nil {

					return err
				}
			}

			if err = ioutil.WriteFile(p, b, 0644); err != nil {

				return err
			}
		}
	}

	return
}

func (m *Manifest) UnApply() (err error) {

	for _, v := range m.Files {

		p := fmt.Sprintf("%s/%s", m.BasePath, v.Path)

		if v.Dir {

			continue
		}

		if v.CanDisable && FileExists(p) {

			f := fmt.Sprintf("%s.disable", p)

			if FileExists(f) {

				if err = os.Remove(f); err != nil {

					return err
				}
			}

			if err = os.Rename(p, f); err != nil {

				return err
			}
		}
	}

	return
}

func UpdateManifest(basepath, baseurl string) (err error) {

	p := fmt.Sprintf("%s/%s", basepath, "manifest.json")
	u := fmt.Sprintf("%s/%s", baseurl, "manifest.json")

	if !FileExists(p) {

		b, err := FetchRemoteFile(u)
		if err != nil {

			return err
		}

		err = ioutil.WriteFile(p, b, 0644)
		if err != nil {

			return err
		}
	}

	lsize, err := GetFileSize(p)
	if err != nil {

		return
	}

	rsize, err := GetRemoteFileSize(u)
	if err != nil {

		return
	}
	if rsize != lsize {

		b, err := FetchRemoteFile(u)
		if err != nil {

			return err
		}

		err = ioutil.WriteFile(p, b, 0644)
		if err != nil {

			return err
		}
	}

	return
}

func GetManifest(basepath string) (m *Manifest, err error) {

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", basepath, "manifest.json"))
	if err != nil {

		return
	}

	err = json.Unmarshal(b, &m)
	if err != nil {

		return
	}

	return
}
