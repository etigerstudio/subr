// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package fetcher

import (
	"github.com/etigerstudio/subr"
	"io/ioutil"
	h "net/http"
)

type http struct {
	key string
	url string
}

func (f *http) Fetch(c *subr.Context) error {
	response, err := h.Get(f.url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	c.Buckets[f.key] = contents
	c.Logger.Infoln("http Fetch succeeded")

	return nil
}

func NewHTTP(key string, url string) *http {
	f := &http{
		key: key,
		url: url,
	}
	return f
}