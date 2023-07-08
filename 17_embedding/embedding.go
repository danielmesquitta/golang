//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package embedding

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Bytes int
type Celsius float32

type Number interface {
	constraints.Integer | constraints.Float
}

type BandwidthUsage struct {
	amount []Bytes
}

func (b BandwidthUsage) BandwidthUsageAverage() Bytes {
	return average(b.amount...)
}

type CpuTemp struct {
	amount []Celsius
}

func (c CpuTemp) CpuTempAverage() Celsius {
	return average(c.amount...)
}

type MemoryUsage struct {
	amount []Bytes
}

func (m MemoryUsage) MemoryUsageAverage() Bytes {
	return average(m.amount...)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func average[T Number](amounts ...T) T {
	var total T

	for _, amount := range amounts {
		total += amount
	}

	return total / T(len(amounts))
}

func Embedding() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celsius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{bandwidth, temp, memory}

	fmt.Printf("Bandwidth Usage Average: %v\n", dashboard.BandwidthUsageAverage())
	fmt.Printf("Cpu Temp Average: %v\n", dashboard.CpuTempAverage())
	fmt.Printf("Memory Usage Average: %v\n", dashboard.MemoryUsageAverage())
}
