// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package comparator

import "github.com/etigerstudio/subr"

type comparator struct {
	subr.CompareFunc
}

func (p *comparator) Compare(c *subr.Context) (fresh bool, err error) {
	return p.CompareFunc(c)
}

func New(compareFunc subr.CompareFunc) *comparator {
	return &comparator{compareFunc}
}