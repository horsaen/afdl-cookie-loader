package cookies

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type FlexData struct {
	Token struct {
		AccessToken string `json:"accessToken"`
	} `json:"token"`
}

func Flex(username string, password string) {
	home, _ := os.UserHomeDir()

	url := "https://api.flextv.co.kr/v2/api/auth/signin"

	payload := strings.NewReader("{\"loginId\":\"" + username + "\",\"password\":\"" + password + "\",\"loginKeep\":true,\"saveId\":true,\"device\":\"PCWEB\"}")

	client := &http.Client{}

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Referer", "https://www.flextv.co.kr/")
	req.Header.Add("Origin", "https://www.flextv.co.kr")
	req.Header.Add("DNT", "1")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Connection", "keep-alive")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	bodyText, _ := io.ReadAll(resp.Body)

	var flexData FlexData

	err = json.Unmarshal(bodyText, &flexData)

	if err != nil {
		log.Fatal(err)
	}

	f, _ := os.OpenFile(home+"/.afreeca-downloader/cookies/flex", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	_, err = f.WriteString(flexData.Token.AccessToken)

	if err != nil {
		log.Fatal(err)
	}

}
