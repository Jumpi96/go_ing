package main

import (
	cmd "./cmd"
	internal "./internal"
)

var r = &internal.BoltDBRepository{}

func main() {
	cmd.Init(r)
}
