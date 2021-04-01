package example

import (
	"fmt"
	"time"
)

func TimeWheelDefaultJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a timewheel job with key: %v", time.Now().Format(time.RFC3339), key))
}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(time.RFC3339), key))
}
