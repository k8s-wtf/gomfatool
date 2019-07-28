package main

import (
	"fmt"
	"github.com/t0mmyt/mfa/keystorage"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	log "github.com/sirupsen/logrus"
	"github.com/t0mmyt/mfa/aesjson"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	// TODO - Dynamic?  Read $HOME maybe
	FILENAME = "/home/tom/.mfa"
)

var (
	doCommand  = kingpin.Command("do", "do an mfa")
	doName     = doCommand.Arg("doName", "doName of mfa").Required().String()
	addCommand = kingpin.Command("add", "add an mfa")
	addName    = addCommand.Arg("name", "name of mfa").Required().String()
	addUrl     = addCommand.Arg("url", "url of totp").Required().String()
	allCommand = kingpin.Command("all", "show all keys")
)

func main() {
	var storage keystorage.Backend
	storage = aesjson.NewAesJson()
	
	keys, err := storage.Read(FILENAME)
	if err != nil {
		log.Fatal(err)
	}

	switch kingpin.Parse() {
	case doCommand.FullCommand():
		if url, ok := keys[*doName]; ok {
			k, err := otp.NewKeyFromURL(url)
			if err != nil {
				log.Fatal(err)
			}
			code, err := totp.GenerateCode(k.Secret(), time.Now())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s: %s\n", *doName, code)
		} else {
			log.Fatalf("Key %s not found")
		}
	case addCommand.FullCommand():
		err := storage.Add(*addName, *addUrl)
		if err != nil {
			log.Fatal(err)
		}
		err = storage.Write(FILENAME)
		if err != nil {
			log.Fatal(err)
		}
	case allCommand.FullCommand():
		{
			for k, v := range keys {
				key, err := otp.NewKeyFromURL(v)
				if err != nil {
					log.Fatal(err)
				}
				code, err := totp.GenerateCode(key.Secret(), time.Now())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s: %s\n", k, code)

			}
		}
	}
}
