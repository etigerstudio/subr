// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package consolidator

import (
	"io/ioutil"
	"path"
	"subr"
	"time"
)

type Local struct {
	Path      string
	Prefix    string
	Extension string
}

func (s *Local) Consolidate(c *subr.Context) error {
	filename := path.Join(s.Path, s.Prefix + "_" +
		c.StartTimestamp.Format(time.RFC3339) + "." + s.Extension)
	err := ioutil.WriteFile(filename, c.Data, 0644)

	c.Logger.Infoln("Local file consolidated")
	return err
}

func NewLocal(path string, prefix string, extension string) *Local {
	return &Local{
		Path:      path,
		Prefix:    prefix,
		Extension: extension,
	}
}
