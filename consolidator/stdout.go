// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package consolidator

import (
	"errors"
	"fmt"
	"github.com/etigerstudio/subr"
)

type stdout struct {
	key string
}

func (s *stdout) Consolidate(c *subr.Context) error {
	bytes, ok := c.Buckets[s.key].([]byte)
	if !ok {
		return errors.New("cannot read bucket: " + s.key +" as bytes")
	}
	c.Logger.Infoln("stdout consolidator '" + s.key + "':")
	fmt.Print(string(bytes))
	return nil
}

func NewStdout(key string) *stdout {
	return &stdout{key: key}
}