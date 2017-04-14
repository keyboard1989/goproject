package application

import (
	"net"
	"net/http"
	"time"

	"github.com/jinzhu/now"
)

func initHTTPClient(timeOut time.Duration, maxIdleConnsPerHost int) *http.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, timeOut)
				if err != nil {
					return nil, err
				}
				return conn, nil
			},
			MaxIdleConnsPerHost:   maxIdleConnsPerHost,
			ResponseHeaderTimeout: timeOut,
		},
	}
	return httpClient
}

func GetTime() time.Time {
	return now.BeginningOfMinute()
}
