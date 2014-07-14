package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	YB
	ZB
)
