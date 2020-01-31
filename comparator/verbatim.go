// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package comparator

import (
	"bytes"
	"subr"
)

type Verbatim struct {

}

func (t *Verbatim) Compare(c *subr.Context) (fresh bool, err error) {
	if bytes.Compare(c.Data, c.LastData) == 0 {
		c.Logger.Infoln("Comparator result stale")
		return false, nil
	}
	c.Logger.Infoln("Comparator result fresh")
	return true, nil
}

func NewVerbatim() *Verbatim {
	return &Verbatim{}
}