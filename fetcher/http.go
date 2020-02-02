// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package fetcher

import (
	"github.com/etigerstudio/subr"
	"io/ioutil"
	nethttp "net/http"
)

type http struct {
	key string
	url string
}

func (h *http) Fetch(c *subr.Context) error {
	response, err := nethttp.Get(h.url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	c.Buckets[h.key] = contents
	c.Logger.Infoln("http Fetch succeeded")

	return nil
}

// TODO: Add support for custom dynamic url func
func NewHTTP(key string, url string) *http {
	f := &http{
		key: key,
		url: url,
	}
	return f
}