go-chart
========

Package `chart` is a very simple golang native charting library that supports timeseries and continuous
line charts. 

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
                    X:     float64(xvalues[len(xvalues)-1].Unix()), //todo: helpers for this.
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

You can alternately turn a bunch of features off and constrain the proportions to something like a spark line:

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

# Design Philosophy

I wanted to make a charting library that used only native golang, that could be stood up on a server (i.e. it had built in fonts).

The goal with the API itself is to have the "zero value be useful", and to require the user to not code more than they absolutely needed.

# Contributions

This library is super early but contributions are welcome.
