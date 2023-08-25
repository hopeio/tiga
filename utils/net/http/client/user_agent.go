package client

import "math/rand"

const (
	UserAgent1 = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"
	UserAgent2 = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36"
)

var userAgent = []string{UserAgent1}

func GetRandUserAgent() string {
	return userAgent[rand.Intn(len(userAgent))]
}
