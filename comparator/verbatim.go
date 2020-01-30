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
		return false, nil
	}
	return true, nil
}