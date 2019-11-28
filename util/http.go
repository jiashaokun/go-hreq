package util

import "net/url"

func GetPostUrlPath(path string) url.Values {

	ul := url.Values{}
	pth, _ := url.ParseQuery(path)
	for k, v := range pth {
		ul.Add(k, v[0])
	}

	return ul
}
