// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package consolidator

import (
	"errors"
	"fmt"
	"github.com/etigerstudio/subr"
	"os"
	"strconv"
)

type localrm struct {
	key		      string
	path          string
	shouldPrepend bool
}

func (l *localrm) Consolidate(c *subr.Context) error {
	fmt.Println("kk")
	files, ok := c.Buckets[l.key].([]string)
	if !ok {
		return errors.New("bucket '" + l.key + "' is not a []string")
	}

	prompt := strconv.Itoa(len(files)) + " files were removed"
	if len(files) > 0 {
		prompt += ":"
	}

	for _, file := range files {
		if l.shouldPrepend {
			file = l.path + file
		}
		err := os.Remove(file)
		if err != nil {
			return err
		}
		prompt += "\n" + file
	}

	c.Logger.Infoln(prompt)
	return nil
}

func NewLocalrm(key string) *localrm {
	return &localrm{key: key}
}

func NewLocalrmInPath(key string, path string) *localrm {
	return &localrm{
		key:           key,
		path:          path,
		shouldPrepend: true,
	}
}