// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package fetcher

import (
	"github.com/etigerstudio/subr"
	"io/ioutil"
)

type local struct {
	key string
	filename string
}

func (l *local) Fetch(c *subr.Context) error {
	bytes, err := ioutil.ReadFile(l.filename)
	if err != nil {
		return err
	}

	c.Buckets[l.key] = bytes
	return nil
}

func NewLocal(key string, filename string) *local {
	return &local{
		key:      key,
		filename: filename,
	}
}