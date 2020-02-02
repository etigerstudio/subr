// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package consolidator

import (
	"github.com/etigerstudio/subr"
	"io/ioutil"
	"path"
	"time"
)

type local struct {
	key		  string
	path      string
	prefix    string
	extension string
	filenameFunc func(c *subr.Context) string
}

func (l *local) Consolidate(c *subr.Context) error {
	var filename string

	// Preparing filename & path
	if l.filenameFunc != nil {
		filename = l.filenameFunc(c)
	} else {
		filename = l.prefix + "_" +
			c.StartTimestamp.Format(time.RFC3339) + "." + l.extension
	}
	filepath := path.Join(l.path, filename)

	// Writing file
	data, err := subr.CastBucketToBytes(c.Buckets, l.key)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath, data, 0644)

	c.Logger.Infoln("local file consolidated")
	return err
}

func NewLocal(key string, path string, prefix string, extension string) *local {
	return &local{
		key:       key,
		path:      path,
		prefix:    prefix,
		extension: extension,
	}
}

// TODO: Refactor either using local as argument or
//       reading name from specific bucket
func NewLocalWithFilenameFunc(key string, path string,
	filenameFunc func(c *subr.Context) string) *local {
	return &local{
		key:          key,
		path:         path,
		filenameFunc: filenameFunc,
	}
}
