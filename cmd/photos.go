package cmd

import (
	"bytes"
	"errors"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/disintegration/imaging"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/segmentio/ksuid"
	"gopkg.in/abiosoft/ishell.v2"

	"github.com/textileio/textile-go/core"
	"path/filepath"
)

func AddPhoto(c *ishell.Context) {
	if len(c.Args) == 0 {
		c.Err(errors.New("missing photo path"))
		return
	}
	if !core.Node.IsDatastoreConfigured() {
		c.Err(errors.New("datastore not initialized, please run 'textile init'"))
		return
	}

	// try to get path with home dir tilda
	pp, err := homedir.Expand(c.Args[0])
	if err != nil {
		pp = c.Args[0]
	}

	// open the file
	f, err := os.Open(pp)
	if err != nil {
		c.Err(err)
		return
	}
	defer f.Close()

	// try to create a thumbnail version
	th, _, err := image.Decode(f)
	if err != nil {
		c.Err(err)
		return
	}
	th = imaging.Resize(th, 300, 0, imaging.Lanczos)
	thb := new(bytes.Buffer)
	if err = jpeg.Encode(thb, th, nil); err != nil {
		c.Err(err)
		return
	}

	tp := filepath.Join(core.Node.RepoPath, "tmp", ksuid.New().String()+".jpg")
	if err = ioutil.WriteFile(tp, thb.Bytes(), 0644); err != nil {
		c.Err(err)
		return
	}

	// do the add
	f.Seek(0, 0)
	mr, err := core.Node.AddPhoto(pp, tp)
	if err != nil {
		c.Err(err)
		return
	}

	// clean up
	if err = os.Remove(tp); err != nil {
		c.Err(err)
		return
	}
	if err = os.Remove(mr.PayloadPath); err != nil {
		c.Err(err)
		return
	}

	// show user root cid
	cyan := color.New(color.FgCyan).SprintFunc()
	c.Println(cyan("added " + mr.Boundary))
}
