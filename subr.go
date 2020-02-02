// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import (
	"errors"
	"fmt"
	"time"
)

type Context struct {
	*Logger

	StartTimestamp time.Time
	Buckets        map[string]interface{}
	LastBuckets    map[string]interface{}
}

type FetchFrequency time.Duration

const (
	Faster = FetchFrequency(20 * time.Second)
	Fast   = FetchFrequency(1 * time.Minute)
	Normal = FetchFrequency(5 * time.Minute)
	Slow   = FetchFrequency(20 * time.Minute)
	Slower = FetchFrequency(1 * time.Hour)
)

//type FetchFunc func(c *Context) error
type TranspileFunc func(c *Context) error
type CompareFunc func(c *Context) (fresh bool, err error)
//type ConsolidateFunc func(c *Context) error

type Fetcher interface {
	Fetch(c *Context) error
}

type Transpiler interface {
	Transpile(c *Context) error
}

type Comparator interface {
	Compare(c *Context) (fresh bool, err error)
}

type Consolidator interface {
	Consolidate(c *Context) error
}

type InstanceID string

type Instance struct {
	fetchers []Fetcher
	transpilers []Transpiler
	comparator Comparator
	consolidators []Consolidator

	logger *Logger

	FetchFrequency
	LastBuckets map[string]interface{}
}

func (i *Instance) UseFetcher(fetcher Fetcher) *Instance {
	i.fetchers = append(i.fetchers, fetcher)
	return i
}

func (i *Instance) UseTranspiler(transpiler Transpiler) *Instance {
	i.transpilers = append(i.transpilers, transpiler)
	return i
}

func (i *Instance) SetComparator(comparator Comparator) *Instance {
	i.comparator = comparator
	return i
}

func (i *Instance) UseConsolidator(consolidator Consolidator) *Instance {
	i.consolidators = append(i.consolidators, consolidator)
	return i
}

func (i *Instance) Execute() {
	i.logger.Infoln(GetColoredText(" Execution start ", magentaControlText))

	context := &Context{
		StartTimestamp: time.Now(),
		Logger:			i.logger,
		Buckets:		make(map[string]interface{}),
		LastBuckets:    i.LastBuckets,
	}

	// Fetching
	if len(i.fetchers) == 0 {
		i.logger.Errorln("No attached fetchers: One fetcher at least is required to bootstrap execution")
		return
	}
	for _, fetcher := range i.fetchers {
		err := fetcher.Fetch(context)
		if err != nil {
			i.logger.Errorln("Error occurred while fetching:", err)
			return
		}
	}

	// Transpiling
	if len(i.transpilers) == 0 {
		i.logger.Infoln("No attached transpilers: No transpiling is applied")
	}
	for _, transpiler := range i.transpilers {
		err := transpiler.Transpile(context)
		if err != nil {
			i.logger.Errorln("Error occurred while transpiling:", err)
			return
		}
	}

	// Comparing
	if i.comparator == nil {
		i.logger.Infoln("No attached comparator: Fresh pass is assumed")
	} else {
		fresh, err := i.comparator.Compare(context)
		if err != nil {
			i.logger.Errorln("Error occurred while comparing:", err)
			return
		}
		if !fresh {
			i.logger.Infoln(GetColoredText(" Execution stale pass finished ", whiteControlText))
			return
		}
	}

	// Consolidating
	if len(i.consolidators) == 0 {
		i.logger.Warnln("No attached consolidators: No buckets are saved")
	}
	for _, consolidator := range i.consolidators {
		err := consolidator.Consolidate(context)
		if err != nil {
			i.logger.Errorln("Error occurred while consolidating:", err)
			return
		}
	}

	// Buckets passing
	i.LastBuckets = context.Buckets
	i.logger.Infoln(GetColoredText(" Execution fresh pass finished ", cyanControlText))
}

func NewInstance(frequency FetchFrequency) *Instance {
	return &Instance{
		FetchFrequency: frequency,
		LastBuckets:    make(map[string]interface{}),
	}
}

var defaultDispatcher *Dispatcher

type Dispatcher struct {
	instances map[InstanceID]*Instance
	schedule  map[FetchFrequency][]InstanceID
}

func (d *Dispatcher) AttachInstance(id InstanceID, instance *Instance) error {
	if d.instances[id] != nil {
		return errors.New("Instance of ID: " + string(id) + " already exists")
	}

	d.instances[id] = instance
	instance.logger = getLoggerForInstance(id)
	d.addSchedule(instance.FetchFrequency, id)
	instance.logger.Infoln(GetColoredText(" Instance attached, executing every " +
		fmt.Sprintf("%.1f", time.Duration(instance.FetchFrequency).Seconds()) + " seconds ", whiteControlText))

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