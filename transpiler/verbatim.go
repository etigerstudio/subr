// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package transpiler

import "subr"

type Verbatim struct {

}

func (t *Verbatim) Transpile(c *subr.Context) error {
	subr.Infoln("Verbatim transpiler passed")
	return nil
}

func NewVerbatim() *Verbatim {
	return &Verbatim{}
}