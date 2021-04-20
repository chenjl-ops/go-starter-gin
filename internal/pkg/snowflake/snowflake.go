package snowflake

import (
	"errors"
	"time"
)

const (
	workerIDBits     = uint64(5)
	dataCenterIDBits = uint64(5)
	sequenceBits     = uint64(12)

	workerIDMax     = int64(-1) ^ (int64(-1) << workerIDBits)
	sequenceMax     = int64(-1) ^ (int64(-1) << sequenceBits)
	dataCenterIDMAX = int64(-1) ^ (int64(-1) << dataCenterIDBits)

	timeShift   = workerIDBits + sequenceBits
	dataShift   = dataCenterIDBits + sequenceBits
	workerShift = sequenceBits

	startTime = int64(1525705533000) // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

var SnowWorker *Worker

func NewWorker(workerId int64, dataCenterId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerIDMax {
		return nil, errors.New("worker ID excess of quantity")
	}
	if dataCenterId < 0 || dataCenterId > dataCenterIDMAX {
		return nil, errors.New("dataCenter ID excess of quantity")
	}

	// 生成一个新节点
	return &Worker{
		lastStamp:    0,
		workerID:     workerId,
		dataCenterID: dataCenterId,
		sequence:     0,
	}, nil
}

func InitSnowWorker(workerId int64, dataCenterId int64) error {
	worker, err := NewWorker(workerId, dataCenterId)

	if err != nil {
		return err
	}

	SnowWorker = worker

	return nil
}

func (w *Worker) getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *Worker) NextID() (uint64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.nextID()
}

func (w *Worker) nextID() (uint64, error) {
	timeStamp := w.getMilliSeconds()

	if timeStamp < w.lastStamp {
		return 0, errors.New("time is moving backwards, waiting until")
	}

	if w.lastStamp == timeStamp {
		w.sequence = (w.sequence + 1) & sequenceMax

		if w.sequence == 0 {
			for timeStamp <= w.lastStamp {
				timeStamp = w.getMilliSeconds()
			}
		}
	} else {
		w.sequence = 0
	}

	w.lastStamp = timeStamp

	id := ((timeStamp - startTime) << timeShift) |
		(w.dataCenterID << dataShift) |
		(w.workerID << workerShift) |
		w.sequence

	return uint64(id), nil
}
