package main

import (
	utctime "github.com/nirvana-labs/go-analyzer-utctime"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(utctime.Analyzer)
}
