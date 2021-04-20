package snowflake

import "sync"

type Worker struct {
	mu           sync.Mutex
	lastStamp    int64
	workerID     int64
	dataCenterID int64
	sequence     int64
}
