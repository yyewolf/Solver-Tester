package main

import (
	"sort"
	"time"
)

type Counters struct {
	Name        string
	WinCount    int
	Total       int
	TotalTime   int
	AverageTry  float64
	AverageTime float64
}

var statCounters map[string]*Counters

func init() {
	statCounters = make(map[string]*Counters)
}

func (c *Counters) UpdateAverageTries(tries int) {
	if c.AverageTry == 0 {
		c.AverageTry = float64(tries)
	} else {
		c.AverageTry = c.AverageTry*float64(c.Total-1)/float64(c.Total) + float64(tries)/float64(c.Total)
	}
}

func (c *Counters) UpdateAverageTime(duration time.Duration) {
	c.TotalTime++
	d := float64(duration)
	if c.AverageTime == 0 {
		c.AverageTime = d
	} else {
		c.AverageTime = c.AverageTime*float64(c.TotalTime-1)/float64(c.TotalTime) + d/float64(c.TotalTime)
	}
}

func OrderCountersByAverage() []*Counters {
	var ordered []*Counters
	for _, counter := range statCounters {
		ordered = append(ordered, counter)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].WinCount > ordered[j].WinCount
	})
	return ordered
}
