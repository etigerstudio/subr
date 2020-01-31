// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package consolidator

import (
	"io/ioutil"
	"path"
	"subr"
	"time"
)

type local struct {
	key		  string
	path      string
	prefix    string
	extension string
	filenameFunc func(c *subr.Context) string
}

func (s *local) Consolidate(c *subr.Context) error {
	var filename string

	// Preparing filename & path
	if s.filenameFunc != nil {
		filename = s.filenameFunc(c)
	} else {
		filename = s.prefix + "_" +
			c.StartTimestamp.Format(time.RFC3339) + "." + s.extension
	}
	filepath := path.Join(s.path, filename)

	// Writing file
	data, err := subr.CastBucketToBytes(c.Buckets, s.key)
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

func NewLocalWithFilenameFunc(key string, path string,
	filenameFunc func(c *subr.Context) string) *local {
	return &local{
		key:          key,
		path:         path,
		filenameFunc: filenameFunc,
	}
}
