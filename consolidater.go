// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

type Consolidator interface {
	Consolidate(c *Context) error
}