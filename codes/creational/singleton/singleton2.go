package singleton

import (
	"fmt"
	"sync"
)

// 实现单例模式的第二种方式
var once sync.Once

type single struct{}

var singleInstance *single

func GetInstance2() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single Instance already created-2")
	}

	return singleInstance
}
