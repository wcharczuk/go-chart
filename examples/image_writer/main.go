package main

import (
	"fmt"
	"github.com/d-Rickyy-b/go-chart-x/v2/pkg/chart"
	"log"
)

func main() {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}
	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	image, err := collector.Image()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final Image: %dx%d\n", image.Bounds().Size().X, image.Bounds().Size().Y)
}
