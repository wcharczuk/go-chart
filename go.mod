module github.com/ebudan/go-chart/v2

go 1.15

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/wcharczuk/go-chart/v2 v2.1.0
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
)

// Replacing original's import path to avoid editing all imports here.
replace github.com/wcharczuk/go-chart/v2 => ./
