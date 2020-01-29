// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

type Transpiler interface {
	Transpile(c *Context) error
}
