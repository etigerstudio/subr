// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

type Comparator interface {
	Compare(c *Context) (fresh bool, err error)
}