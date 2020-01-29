// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import "time"

type Context struct {
	StartTimestamp time.Time
	Data []byte
	LastData []byte
}

func (c *Context) Abort()  {

}