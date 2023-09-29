package uniqueProcess

import (
	"errors"
	"sync"
)

var ErrProcessRunning = errors.New("process is already running")

type UniqueProcess struct {
	data map[string]struct{}
	sync.Mutex
}

func NewUniqueProcess() *UniqueProcess {
	return &UniqueProcess{data: map[string]struct{}{}}
}

func (u *UniqueProcess) Set(key string) {
	u.data[key] = struct{}{}
}

func (u *UniqueProcess) Get(key string) bool {
	_, ok := u.data[key]
	if ok {
		return true
	}
	return false
}

func (u *UniqueProcess) Wait(key string) error {
	u.Lock()
	defer u.Unlock()
	exist := u.Get(key)
	if exist {
		return ErrProcessRunning
	}
	u.Set(key)
	return nil
}
