package chart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
)

type ChartConfig struct {
	BaseConfiguration charts.BaseConfiguration
	BaseActions       charts.BaseActions
}

type XY [2]interface{}
