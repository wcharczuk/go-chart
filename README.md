go-chart
========
[![Build Status](https://travis-ci.org/wcharczuk/go-chart.svg?branch=master)](https://travis-ci.org/wcharczuk/go-chart)

Package `chart` is a very simple golang native charting library that supports timeseries and continuous
line charts. 

The v1.0 release has been tagged so things should be more or less stable, if something changes please log an issue.

# Installation

To install `chart` run the following:

```bash
> go get -u github.com/wcharczuk/go-chart
```

Most of the components are interchangeable so feel free to crib whatever you want. 

# Usage 

 ![](https://raw.githubusercontent.com/wcharczuk/go-chart/master/images/goog_ltm.png)


The chart code to produce the above is as follows:

```go
// note this assumes that xvalues and yvalues
// have been pulled from a pricing service.
graph := chart.Chart{
    Width:  1024,
    Height: 400,
    YAxis: chart.YAxis {
        Style: chart.Style{
            Show: true,
        },
    },
    XAxis: chart.XAxis {
        Style: chart.Style{
            Show: true,
        },
    },
    Series: []chart.Series{
        chart.TimeSeries{
            XValues: xvalues,
            YValues: yvalues,
            Style: chart.Style {
                FillColor: chart.DefaultSeriesStrokeColors[0].WithAlpha(64),
            },
        },
        chart.AnnotationSeries{
            Name: "Last Value",
            Style: chart.Style{
                Show:        true,
                StrokeColor: chart.DefaultSeriesStrokeColors[0],
            },
            Annotations: []chart.Annotation{
                chart.Annotation{
                    X:     chart.TimeToFloat64(xvalues[len(xvalues)-1]),
                    Y:     yvalues[len(yvalues)-1],
                    Label: chart.FloatValueFormatter(yvalues[len(yvalues)-1]),
                },
            },
        },
    },
}
graph.Render(chart.PNG, buffer) //thats it!
```

The key areas to note are that we have to explicitly turn on two features, the axes and add the last value label annotation series. When calling `.Render(..)` we add a parameter, `chart.PNG` that tells the renderer to use a raster renderer. Another option is to use `chart.SVG` which will use the vector renderer and create an svg representation of the chart. 

# Alternate Usage

You can alternately leave a bunch of features turned off and constrain the proportions to something like a spark line:

 ![](https://raw.githubusercontent.com/wcharczuk/go-chart/master/images/tvix_ltm.png)

The code to produce the above would be:

```go
// note this assumes that xvalues and yvalues
// have been pulled from a pricing service.
graph := chart.Chart{
    Width:  1024,
    Height: 100,
    Series: []chart.Series{
        chart.TimeSeries{
            XValues: xvalues,
            YValues: yvalues,
        },
    },
}
graph.Render(chart.PNG, buffer)
```

# 2 Y-Axis Charts 

 ![](https://raw.githubusercontent.com/wcharczuk/go-chart/master/images/two_axis.png)

It is also possible to draw series against 2 separate y-axis with their own ranges (usually good for comparison charts).
In order to map the series to an alternate axis make sure to set the `YAxis` property of the series to `YAxisSecondary`.

```go
graph := chart.Chart{
    Title: stock.Name,
    TitleStyle: chart.Style{
        Show: false,
    },
    Width:  width,
    Height: height,
    XAxis: chart.XAxis{
        Style: chart.Style{
            Show: true,
        },
    },
    YAxis: chart.YAxis{
        Style: chart.Style{
            Show: true,
        },
    },
    Series: []chart.Series{
        chart.TimeSeries{
            Name:    "vea",
            XValues: vx,
            YValues: vy,
            Style: chart.Style{
                Show: true,
                StrokeColor: chart.GetDefaultSeriesStrokeColor(0),
                FillColor:   chart.GetDefaultSeriesStrokeColor(0).WithAlpha(64),
            },
        },
        chart.TimeSeries{
            Name:    "spy",
            XValues: cx,
            YValues: cy,
            YAxis:   chart.YAxisSecondary,  // key (!)
            Style: chart.Style{
                Show: true,
                StrokeColor: chart.GetDefaultSeriesStrokeColor(1),
                FillColor:   chart.GetDefaultSeriesStrokeColor(1).WithAlpha(64),
            },
        },
        chart.AnnotationSeries{
            Name: fmt.Sprintf("%s - Last Value", "vea"),
            Style: chart.Style{
                Show:        true,
                StrokeColor: chart.GetDefaultSeriesStrokeColor(0),
            },
            Annotations: []chart.Annotation{
                chart.Annotation{
                    X:     float64(vx[len(vx)-1].Unix()),
                    Y:     vy[len(vy)-1],
                    Label: fmt.Sprintf("%s - %s", "vea", chart.FloatValueFormatter(vy[len(vy)-1])),
                },
            },
        },
        chart.AnnotationSeries{
            Name: fmt.Sprintf("%s - Last Value", "goog"),
            Style: chart.Style{
                Show:        true,
                StrokeColor: chart.GetDefaultSeriesStrokeColor(1),
            },
            YAxis: chart.YAxisSecondary, // key (!)
            Annotations: []chart.Annotation{
                chart.Annotation{
                    X:     float64(cx[len(cx)-1].Unix()),
                    Y:     cy[len(cy)-1],
                    Label: fmt.Sprintf("%s - %s", "goog", chart.FloatValueFormatter(cy[len(cy)-1])),
                },
            },
        },
    },
}
graph.Render(chart.PNG, buffer)
```

# Moving Averages

You can now also graph a moving average of a series using a special `MovingAverageSeries` that takes an `InnerSeries` as a required argument.

 ![](https://raw.githubusercontent.com/wcharczuk/go-chart/master/images/ma_goog_ltm.png)
 
 There is a helper method, `GetLastValue` on the `MovingAverageSeries` to aid in creating a last value annotation for the series.

# More Intense Technical Analysis

You can also have series that produce two values, i.e. a series that implements `BoundedValueProvider` and the `GetBoundedValue(int)(x,y1,y2 float64)` method. An example of a `BoundedValueProvider` is the included `BollingerBandsSeries`.

![](https://raw.githubusercontent.com/wcharczuk/go-chart/master/images/spy_ltm_bbs.png)

Like the `MovingAverageSeries` this series takes an `InnerSeries` argument as required, and defaults to 10 samples and a `K` value of 2.0 (or two standard deviations in either direction).

# Design Philosophy

I wanted to make a charting library that used only native golang, that could be stood up on a server (i.e. it had built in fonts).

The goal with the API itself is to have the "zero value be useful", and to require the user to not code more than they absolutely needed.

# Contributions

This library is super early but contributions are welcome.
