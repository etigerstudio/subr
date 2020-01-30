// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

type Fetcher interface {
	Fetch(c *Context) error
}
