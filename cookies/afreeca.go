package cookies

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func Afreeca(szUid, szPassword string) {
	home, _ := os.UserHomeDir()

	loginURL := "https://login.afreecatv.com/app/LoginAction.php"
	data := url.Values{
		"szWork":      {"login"},
		"szType":      {"json"},
		"szUid":       {szUid},
		"szScriptVar": {"oLoginRet"},
		"szPassword":  {szPassword},
		"isSaveId":    {"true"},
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.PostForm(loginURL, data)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	cookies := client.Jar.Cookies(resp.Request.URL)
	pdboxTicket := ""
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie.Name, "PdboxTicket") {
			pdboxTicket = cookie.Value
			break
		}
	}

	f, _ := os.OpenFile(home+"/.afreeca-downloader/cookies/afreeca", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	_, err = f.WriteString(pdboxTicket)

	if err != nil {
		log.Fatal(err)
	}

}
