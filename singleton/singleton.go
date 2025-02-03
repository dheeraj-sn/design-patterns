package singleton

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

type SingletonObject struct{}

var SingleObject *SingletonObject

func GetSingleObject(group *sync.WaitGroup, i int) *SingletonObject {
	if SingleObject == nil {
		lock.Lock()
		defer lock.Unlock()
		if SingleObject == nil {
			fmt.Printf("Create SingletonObject by group %d\n", i)
			SingleObject = &SingletonObject{}
		} else {
			fmt.Printf("SingletonObject already created for group %d\n", i)
		}
	} else {
		fmt.Printf("SingletonObject already created for group %d\n", i)
	}
	group.Done()
	return SingleObject
}
