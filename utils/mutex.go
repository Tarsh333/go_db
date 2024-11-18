package utils

import (
	"sync"
)

var Mutexes map[string]*sync.Mutex
var mutex sync.Mutex

func CheckAndInitMutex() {
	if Mutexes == nil {
		Mutexes = make(map[string]*sync.Mutex)
	}
}

func GetOrCreateMutex(path string) *sync.Mutex {

	mutex.Lock()
	defer mutex.Unlock()
	m, ok := Mutexes[path]

	if !ok {
		m = &sync.Mutex{}
		Mutexes[path] = m
	}

	return m
}
