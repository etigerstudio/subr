// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"subr"
)

type HTTP struct {
	frequency subr.FetchFrequency
	url string
	lastData []byte
}

func (f *HTTP) Fetch(c *subr.Context) error {
	response, err := http.Get(f.url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(contents)
	return nil
}

func NewHTTP(frequency subr.FetchFrequency, url string) HTTP {
	f := HTTP{
		frequency: frequency,
		url:       url,
	}
	return f
}