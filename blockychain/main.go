package main

import (
	cmd "./cmd"
	internal "./internal"
)

var r *internal.BoltDBRepository

func main() {
	r := internal.InitDB()
	cmd.Init(r)
}
