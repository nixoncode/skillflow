package main

import "github.com/nixoncode/skillflow/app"

func main() {
	sf := app.New()

	if err := sf.Bootstrap(); err != nil {
		panic(err)
	}

	if err := sf.Start(); err != nil {
		panic(err)
	}
}
