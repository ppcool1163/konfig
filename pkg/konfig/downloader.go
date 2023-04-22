package konfig

import (
	"fmt"
	"github.com/zellyn/kooky"
	"io/ioutil"
	"net/http"
)

const url = "https://%s/api/manifest/%d"

func download(endpoint string, manifestId int) string {
	cookies := kooky.ReadCookies(kooky.DomainHasSuffix(endpoint))
	url := fmt.Sprintf(url, endpoint, manifestId)
	method := "GET"
	requestCookies := ""
	for _, cookie := range cookies {
		addCookies(&requestCookies, cookie.Name, cookie.Value)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cookie", requestCookies)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

func addCookies(cookie *string, cookieName string, cookieValue string) {
	cookie = cookie
	if "SESSION" == cookieName {
		*cookie = *cookie + "SESSION=" + cookieValue + ";"
	}

	if "XSRF-TOKEN" == cookieName {
		*cookie = *cookie + "XSRF-TOKEN=" + cookieValue + ";"
	}
}
