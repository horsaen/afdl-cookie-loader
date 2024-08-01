package cookies

import (
	"fmt"
	"log"
	"os"

	"github.com/browserutils/kooky"
	_ "github.com/browserutils/kooky/browser/all"
)

func Panda() {
	home, _ := os.UserHomeDir()

	// please read the README.md for what you have to do
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`.pandalive.co.kr`), kooky.Name("sessKey"))

	if len(cookies) == 0 {
		fmt.Println("No sessKey detected.")
		os.Exit(0)
	}

	f, _ := os.OpenFile(home+"/.afreeca-downloader/cookies/panda", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	_, err := f.WriteString(cookies[0].Value)

	if err != nil {
		log.Fatal(err)
	}
}
