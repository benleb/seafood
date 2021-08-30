package utils

import (
	"log"
	"net/url"
)

func QueryEncode(q string) string {
	return url.QueryEscape(q)
}

func SafeQueryDecode(q string) string {
	qa , err := url.QueryUnescape(q)
	if err != nil {
		log.Printf("fail to decode query before: %s, after: %s error: %v\n", q, qa, err.Error())
		qa = ""
	}
	return qa
}

