// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import (
	"errors"
	"time"
)

type Context struct {
	StartTimestamp time.Time
	Data []byte
	LastData []byte
}

type FetchFrequency time.Duration

const (
	Faster = FetchFrequency(20 * time.Second)
	Fast   = FetchFrequency(1 * time.Minute)
	Normal = FetchFrequency(5 * time.Minute)
	Slow   = FetchFrequency(20 * time.Minute)
	Slower = FetchFrequency(1 * time.Hour)
)

type InstanceID string

type Instance struct {
	Fetcher
	Transpiler
	Comparator
	Consolidator

	FetchFrequency
	LastData []byte
}

func (i *Instance) Execute() {
	//TODO: Handle errors
	Infoln("-- Execution start")

	context := &Context{
		StartTimestamp: time.Now(),
		LastData:       i.LastData,
	}
	err := i.Fetcher.Fetch(context)
	if err != nil {
		Warnln(err)
	}

	err = i.Transpiler.Transpile(context)
	if err != nil {
		Warnln(err)
	}

	fresh, err := i.Comparator.Compare(context)
	if err != nil {
		Warnln(err)
	}
	if !fresh {
		Infoln("-- Execution stale pass finished")
		return
	}

	err = i.Consolidator.Consolidate(context)
	if err != nil {
		Warnln(err)
	}

	i.LastData = context.Data

	Infoln("-- Execution fresh pass finished")
}

func NewInstance(frequency FetchFrequency, fetcher Fetcher, transpiler Transpiler,
	comparator Comparator, consolidator Consolidator) Instance {
	return Instance{
		FetchFrequency: frequency,
		Fetcher:        fetcher,
		Transpiler:     transpiler,
		Comparator:     comparator,
		Consolidator:   consolidator,
	}
}

var defaultDispatcher *Dispatcher

type Dispatcher struct {
	instances map[InstanceID]*Instance
	schedule  map[FetchFrequency][]InstanceID
}

func (d *Dispatcher) AttachInstance(id InstanceID, instance *Instance) error {
	if d.instances[id] != nil {
		return errors.New("Instance of ID: " + string(id) + "already exists.")
	}

	d.instances[id] = instance
	d.addSchedule(instance.FetchFrequency, id)

	return nil
}

func (d *Dispatcher) addSchedule(frequency FetchFrequency, id InstanceID)  {
	d.schedule[frequency] = append(d.schedule[frequency], id)
}

func (d *Dispatcher) GetInstance(id InstanceID) *Instance {
	return d.instances[id]
}

func (d *Dispatcher) Run() {
	// Run once immediately
	for _, instance := range d.instances {
		// TODO: Handle error
		instance.Execute()
	}

	// Build tickers
	//var tickers []*time.Ticker

	for frequency, ids := range d.schedule {
		ticker := time.NewTicker(time.Duration(frequency))
		//tickers = append(tickers, ticker)
		// TODO: Add halting mechanism
		go d.tickingExecute(ticker.C, ids)
	}

	select{}
}

func (d *Dispatcher) tickingExecute(c <-chan time.Time, ids []InstanceID) {
	for _ = range c {
		for _, id := range ids {
			// TODO: Handle delta time of execution
			go d.instances[id].Execute()
		}
	}
}

func Default() *Dispatcher {
	if defaultDispatcher == nil {
		defaultDispatcher = New()
	}

	return defaultDispatcher
}

func New() *Dispatcher {
	return &Dispatcher{
		instances: make(map[InstanceID]*Instance),
		schedule:  make(map[FetchFrequency][]InstanceID),
	}
}