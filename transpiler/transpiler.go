// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package transpiler

import "github.com/etigerstudio/subr"

type transpiler struct {
	subr.TranspileFunc
}

func (t *transpiler) Transpile(c *subr.Context) error {
	return t.TranspileFunc(c)
}

func New(transpileFunc subr.TranspileFunc) *transpiler {
	return &transpiler{transpileFunc}
}