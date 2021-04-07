# TimeWheel

TimeWheel时间轮盘是实现延时队列的一种方式，方便定时任务的执行。本项目基于Golang实现了一个简单的单轮时间轮盘。

## 详解

TODO

## 安装

```bash
go get -u github.com/lk668/timewheel

```

## 使用

```go
package main

import (
	"fmt"
	"time"

	"github.com/lk668/timewheel"
	"github.com/lk668/timewheel/example"
)

func main() {
	//初始化一个时间间隔是1s，一共有60个齿轮的时间轮盘，默认轮盘转动一圈的时间是60s
	tw := timewheel.GetTimeWheel(1*time.Second, 60, example.TimeWheelDefaultJob)

	// 启动时间轮盘
	tw.Start()

	if tw.IsRunning() {
		// 添加一个task
		// 每隔10s执行一次
		// task名字叫task1
		// task的创建时间是time.Now()
		// task执行的任务设置为nil，所以默认执行timewheel的Job，也就是example.TimeWheelDefaultJob
		fmt.Println(fmt.Sprintf("%v Add task task-5s", time.Now().Format(time.RFC3339)))
		err := tw.AddTask(5*time.Second, "task-5s", time.Now(), -1, nil)
		if err != nil {
			panic(err)
		}

		// 该Task执行example.TaskJob
		fmt.Println(fmt.Sprintf("%v Add task task-2s", time.Now().Format(time.RFC3339)))
		err = tw.AddTask(2*time.Second, "task-2s", time.Now(), -1, example.TaskJob)
		if err != nil {
			panic(err)
		}

	} else {
		panic("TimeWheel is not running")
	}
	time.Sleep(10 * time.Second)

	// 删除task
	fmt.Println("Remove task task-5s")
	err := tw.RemoveTask("task-5s")
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("Remove task task-2s")
	err = tw.RemoveTask("task-2s")
	if err != nil {
		panic(err)
	}

	// 关闭时间轮盘
	tw.Stop()
}

```

该例子的输出结果如下
```bash
2021-04-01T18:04:40+08:00 Add task task-5s
2021-04-01T18:04:40+08:00 Add task task-2s
2021-04-01T18:04:43+08:00 This is a task job with key: task-2s
2021-04-01T18:04:45+08:00 This is a task job with key: task-2s
2021-04-01T18:04:46+08:00 This is a timewheel job with key: task-5s
2021-04-01T18:04:47+08:00 This is a task job with key: task-2s
2021-04-01T18:04:49+08:00 This is a task job with key: task-2s
Remove task task-5s
2021-04-01T18:04:51+08:00 This is a task job with key: task-2s
2021-04-01T18:04:53+08:00 This is a task job with key: task-2s
2021-04-01T18:04:55+08:00 This is a task job with key: task-2s
2021-04-01T18:04:57+08:00 This is a task job with key: task-2s
2021-04-01T18:04:59+08:00 This is a task job with key: task-2s
Remove task task-2s
```
