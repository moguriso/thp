package thp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const DEVICE_NAME string = "WS-USB01-THP"
const BASE_URL string = "https://svcipp.planex.co.jp/api/get_data.php"

func BuildEndpoint(mac, token string) string {
	const t_layout = "2006-01-02+15:04:05"
	t := time.Now().In(time.UTC)
	to := t.Format(t_layout)
	f := t.Add(-1 * time.Duration(5) * time.Minute)
	from := url.PathEscape(f.Format(t_layout))

	baseUrl, _ := url.Parse(BASE_URL)
	endpoint := "?type=\"" + DEVICE_NAME + "\"&mac=\"" + mac + "\"" +
		"&from=\"" + from + "\"" + "&to=\"" + to + "\"" + "&token=\"" + token + "\""

	return baseUrl.String() + endpoint
}

func GetRequest(endpoint string) string {
	req, _ := http.NewRequest("GET", endpoint, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(byteArray)
}
