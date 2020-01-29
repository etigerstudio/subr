// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package comparator

import (
	"bytes"
	"subr"
)

type Verbatim struct {

}

func (t *Verbatim) Compare(c *subr.Context) error {
	if bytes.Compare(c.Data, c.LastData) == 0 {
		c.Abort()
	}
	return nil
}