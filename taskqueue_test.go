// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

// taskqueue testing methods
//
package taskqueue

import (
    "fmt"
    "time"
    "sync"
    "testing"
    "math/rand"
)

const (
    // MaxTestTasks is number of tasks.
    MaxTestTasks int = 64
    // MaxWorkTime is maximum work time for a task.
    MaxWorkTime int = 32000000
    // MaxSleepTime is maximum sleep time for a task.
    MaxSleepTime int = 64000000
    // MaxWaitingTime is a working period before stop signal.
    MaxWaitingTime int = 1
)

// Task is a test structure of an user's task.
type Task struct {
    Name string
    WorkT time.Duration
    SleepT time.Duration
    GenPanic bool
}

func (t *Task) String() string {
    return fmt.Sprintf("%v [W:%v, S:%v]", t.Name, t.WorkT, t.SleepT)
}

func (t *Task) Run() {
    time.Sleep(t.WorkT)
}

func (t *Task) Sleep() {
    if t.GenPanic {
        panic("expected error")
    }
    time.Sleep(t.SleepT)
}

func gentasks(size int, wrong bool) []Tasker {
    var tasker Tasker
    rand.Seed(time.Now().UnixNano())
    if size < 1 {
        size = MaxTestTasks
    }
    delay := func(d int) time.Duration {
        r := rand.Intn(d-1)
        return time.Duration(r+1)
    }
    tasks := make([]Tasker, size)
    for i := 0; i < size; i++ {
        tasker = &Task{fmt.Sprintf("task #%v", i),
            delay(MaxWorkTime) * time.Nanosecond,
            delay(MaxSleepTime) * time.Nanosecond,
            wrong,
        }
        // fmt.Println("generated", tasker)
        tasks[i] = tasker
    }
    return tasks
}

func TestDebug(t *testing.T) {
    if (LoggerError == nil) || (LoggerDebug == nil) {
        t.Errorf("incorrect references")
    }
    Debug(false)
    if (LoggerError.Prefix() != "ERROR [taskqueue]: ") || (LoggerDebug.Prefix() != "DEBUG [taskqueue]: ") {
        t.Errorf("incorrect loggers settings")
    }
    Debug(true)
    if (LoggerError.Flags() != 19) || (LoggerDebug.Flags() != 21) {
        t.Errorf("incorrect loggers settings")
    }
}

func TestStart(t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("abonormal behavior: %v", r)
        }
    }()
    Debug(false)
    var group sync.WaitGroup
    tasks := gentasks(-1, false)
    finish := make(chan bool)
    complete := Start(tasks, &group, finish)
    time.Sleep(time.Duration(MaxWaitingTime) * time.Second)
    Stop(finish, &group, complete)
}
