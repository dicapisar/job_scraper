package util

import (
	"net/url"
	"strings"
)

func GetInfoJobId(urlString *string) *string {
	urlParsed, err := url.Parse(*urlString)
	if err != nil {
		return nil
	}
	path := urlParsed.Path
	pathSplit := strings.Split(path, "-")
	id := pathSplit[len(pathSplit)-1]
	return &id
}
