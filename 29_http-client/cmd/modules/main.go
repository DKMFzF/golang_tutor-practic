package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
		Timeout: time.Second * 1,
		Jar:     jar,
	}

	req, err := http.NewRequest(http.MethodGet, "http://130.193.40.67:8081", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	//req.AddCookie(&http.Cookie{
	//Name:  "ID",
	//Value: "3675",
	////Expires: time.Time{},
	//MaxAge: (int(time.Second) * 60) * 60,
	////Secure: true,
	//HttpOnly: true,
	//})

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Printf("%s", body)
}
