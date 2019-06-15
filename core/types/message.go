package types

import "time"

type message struct {
	id  int64
	msgHash []byte
	sender  string
	Timestamp	time.Time
}