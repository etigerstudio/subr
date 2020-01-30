// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import "time"

type Fetcher interface {
	Fetch(c *Context) error
}

type FetchFrequency time.Duration

const (
	Faster = FetchFrequency(20 * time.Second)
	Fast   = FetchFrequency(1 * time.Minute)
	Normal = FetchFrequency(5 * time.Minute)
	Slow   = FetchFrequency(20 * time.Minute)
	Slower = FetchFrequency(1 * time.Hour)
)
