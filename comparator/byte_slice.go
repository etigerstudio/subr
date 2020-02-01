// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package comparator

import (
	"bytes"
	"github.com/etigerstudio/subr"
)

type byteSlice struct {
	key string
}

func (b *byteSlice) Compare(c *subr.Context) (fresh bool, err error) {
	// Preparing data
	data, err := subr.CastBucketToBytes(c.Buckets, b.key)
	if err != nil {
		return false, err
	}

	// Inspecting last buckets
	if len(c.LastBuckets) == 0 {
		c.Logger.Infoln("[]byte Comparator: fresh")
		return true, nil
	}

	lastData, err := subr.CastBucketToBytes(c.LastBuckets, b.key)
	if err != nil {
		return false, err
	}

	// Comparing bytes
	if bytes.Compare(data, lastData) == 0 {
		c.Logger.Infoln("[]byte Comparator: stale")
		return false, nil
	}

	c.Logger.Infoln("[]byte Comparator: fresh")
	return true, nil
}

func NewByteSlice(key string) *byteSlice {
	return &byteSlice{
		key: key,
	}
}