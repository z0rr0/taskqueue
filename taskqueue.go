// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

// Package taskqueue is a library to run/stop a queue of periodic tasks.
//
// A user's task should implement methods of Tasker interface.
//
//     // tasks = []Tasker
//     var group sync.WaitGroup
//     finish := make(chan bool)
//     complete := Start(tasks, &group, finish)
//     // some actions
//     Stop(finish, &group, complete)
//
// Error logger is always activated,
// use Debug method to turn on debug mode:
//
//     Debug(true)
//
package taskqueue

import (
    "io/ioutil"
    "log"
    "os"
    "sync"
)

var (
    // LoggerError implements error logger.
    LoggerError = log.New(os.Stderr, "ERROR [taskqueue]: ", log.Ldate|log.Ltime|log.Lshortfile)
    // LoggerDebug implements debug logger, it's disabled by default.
    LoggerDebug = log.New(ioutil.Discard, "DEBUG [taskqueue]: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
)

// Tasker is an interface of a task.
type Tasker interface {
    String() string
    Run()
    Sleep()
}

// Start is a method to start cycles of the tasks.
func Start(tasks []Tasker, g *sync.WaitGroup, finish chan bool) chan Tasker {
    LoggerDebug.Println("called Start()")
    workers := len(tasks)
    pending, complete := make(chan Tasker), make(chan Tasker)
    for i := 0; i < workers; i++ {
        go Poll(pending, complete, g)
    }
    go func() {
        for _, t := range tasks {
            pending <- t
        }
        LoggerDebug.Println("all tasks are running")
    }()
    stopped := make(chan bool)
    go func() {
        <-finish
        LoggerDebug.Println("finish signal is gotten")
        close(stopped)
        close(pending)
    }()
    go func() {
        for task := range complete {
            go Sleep(task, pending, stopped)
        }
    }()
    return complete
}

// Stop is a method to stop the tasks.
// It's a blocked method, waiting time is related with tasks' implementation.
func Stop(finish chan bool, g *sync.WaitGroup, complete chan Tasker) {
    LoggerDebug.Println("called Stop()")
    close(finish)
    LoggerDebug.Println("wait a completion of tasks")
    g.Wait()
    close(complete)
    LoggerDebug.Println("all taks are completed")
}

// Poll is a task handler.
// It reads Tasker from "in" channel and write it to "out" one.
func Poll(in chan Tasker, out chan Tasker, g *sync.WaitGroup) {
    g.Add(1)
    defer g.Done()
    for {
        t, ok := <-in
        if !ok {
            return
        }
        LoggerDebug.Printf("start task %v\n", t)
        t.Run()
        LoggerDebug.Printf("stop task %v\n", t)
        out <- t
    }
}

// It will be running when a task is completed and should sleep.
// After that, a task will be again sent to "pending" channel,
// if "stopped" one is empty or not closed.
func Sleep(t Tasker, pending chan Tasker, stopped chan bool) {
    LoggerDebug.Printf("%v is sleeping\n", t)
    t.Sleep()
    LoggerDebug.Printf("%v waked up\n", t)
    select {
        case <-stopped:
        default:
            LoggerDebug.Printf("%v task is sent to input\n", t)
            pending <- t
    }
}

// Debug turns on debug mode.
func Debug(debugmode bool) {
    debugHandle := ioutil.Discard
    if debugmode {
        debugHandle = os.Stdout
    }
    LoggerDebug = log.New(debugHandle, "DEBUG [taskqueue]: ",
        log.Ldate|log.Lmicroseconds|log.Lshortfile)
}
