package singleton

import (
	"fmt"
	"sync"
)

// 实现单例模式的第一种方式
var lock = &sync.Mutex{}

type singleton struct{}

var singletonInstance *singleton

// GetInstance
func GetInstance() *singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singletonInstance = &singleton{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singletonInstance
}
