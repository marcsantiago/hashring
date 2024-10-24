package go_hashring

import "errors"

var (
	ErrNoNodes = errors.New("there are no nodes in the ring")
)
