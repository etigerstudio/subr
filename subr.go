// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import "time"

type Context struct {
	StartTimestamp time.Time
	Data []byte
	LastData []byte
}

type Instance struct {
	Fetcher
	Transpiler
	Comparator
	Consolidator

	id string
	LastData []byte
}

func (i *Instance) Execute() error {
	context := &Context{
		StartTimestamp: time.Now(),
		LastData:       i.LastData,
	}

	err := i.Fetcher.Fetch(context)
	if err != nil {
		return err
	}

	err = i.Transpiler.Transpile(context)
	if err != nil {
		return err
	}

	fresh, err := i.Comparator.Compare(context)
	if err != nil {
		return err
	}
	if !fresh {
		return nil
	}

	err = i.Consolidator.Consolidate(context)
	if err != nil {
		return err
	}

	return nil
}

func NewInstance(id string, fetcher Fetcher, transpiler Transpiler,
	comparator Comparator, consolidator Consolidator) Instance {
	return Instance{
		Fetcher:      fetcher,
		Transpiler:   transpiler,
		Comparator:   comparator,
		Consolidator: consolidator,
		id:           id,
	}
}