GoSync: Bringing Async/Await to Go
=================================================

![Gopher](gopher.png)

GoSync is a lightweight and user-friendly library designed to manage concurrent synchronization in Go applications. It simplifies the experience for users familiar with async-await constructs.

[![test](https://github.com/askolesov/gosync/actions/workflows/test.yaml/badge.svg)](https://github.com/askolesov/gosync/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/askolesov/gosync/graph/badge.svg?token=TLMHQW5TF5)](https://codecov.io/gh/askolesov/gosync)
[![Go Reference](https://pkg.go.dev/badge/github.com/askolesov/gosync.svg)](https://pkg.go.dev/github.com/askolesov/gosync)
[![Go Report Card](https://goreportcard.com/badge/github.com/askolesov/gosync)](https://goreportcard.com/report/github.com/askolesov/gosync)

Installation
------------

First, use `go get` to install the latest version of the library:

```go
go get -u github.com/askolesov/gosync@latest
```

After installation, import the package into your code:

```go
import "github.com/askolesov/gosync"
```

Features
--------

GoSync offers the following features:

- Launching concurrent tasks using `Go()` and `GoRes()` functions. 
- Waiting for all tasks to complete using `WaitAll()` and `WaitAllRes()` functions.
- Waiting for any task to complete using `WaitAny()` and `WaitAnyRes()` functions.
- Utilizing context for waiting using `WaitCtx()` and `WaitResCtx()` functions.
- Setting a timeout for waiting using `WaitTimeout()` and `WaitResTimeout()` functions.

Basic Usage
-----------

### Launching Concurrent Tasks

GoSync provides two functions for launching concurrent tasks:

-   The `Go()` function is used to launch tasks without expecting any result.
-   The `GoRes()` function is used to launch tasks that are expected to return a result.

These functions take a function as an argument and return a `gosync.Task` or `gosync.TaskRes` object, respectively.

```go
ts := gosync.Go(func() {
    // Your task logic here
})

tr := gosync.GoRes(func() int {
    // Your task logic here
    return someResult
})
```

### Waiting for Tasks to Complete

You can wait for task completion using the `Wait()` method. Furthermore, you have the option to employ a context or establish a timeout using the `WaitCtx(context)` and `WaitTimeout(timeout)` methods.

These waiting methods are accessible for both `Task` and `TaskRes` objects, and they are applicable across all variants of the wait functions.

```go
ts.Wait() // Wait for the task to complete

err := ts.WaitCtx(ctx) // Wait with a context

err := ts.WaitTimeout(timeoutDuration) // Wait with a timeout
```

### Waiting for All Tasks to Complete

To wait for all tasks to complete, you can utilize the `WaitAll()` and `WaitAllRes()` functions. These functions take multiple `Task` or `TaskRes` objects as arguments and block execution until all tasks are done.

```go
gosync.WaitAll(ts1, ts2, ts3) // Wait for all tasks to complete

results := gosync.WaitAllRes(tr1, tr2, tr3) // Wait for all tasks to complete and get results
```

### Waiting for Any Task to Complete

GoSync empowers you to wait for any of the tasks to complete using the `WaitAny()` and `WaitAnyRes()` functions.

```go
gosync.WaitAny(ts1, ts2, ts3) // Await completion of any task

result := gosync.WaitAnyRes(tr1, tr2, tr3) // Await completion of any task and retrieve the result
```

Examples
--------

### Launching Tasks and Waiting for All

```go
ts1 := gosync.Go(func() {
    // Task 1 logic
})
ts2 := gosync.Go(func() {
    // Task 2 logic
})

// Wait for both tasks to complete
gosync.WaitAll(ts1, ts2)
```

### Launching Tasks with Results and Waiting for All

```go
tr1 := gosync.GoRes(func() int {
    // Task 1 logic
    return someResult
})
tr2 := gosync.GoRes(func() int {
    // Task 2 logic
    return anotherResult
})

// Wait for all tasks to complete and retrieve results
results := gosync.WaitAllRes(tr1, tr2) // results = [someResult, anotherResult]
```

### Launching Tasks and Waiting for Any

```go
tr1 := gosync.GoRes(func() int {
    // Task 1 logic
    return someResult
})
tr2 := gosync.GoRes(func() int {
    // Task 2 logic
    return anotherResult
})
tr3 := gosync.GoRes(func() int {
    // Task 3 logic
    return anotherResult
})

// Wait for any task to complete
gosync.WaitAny(ts1, ts2, ts3)
```

Limitations
-----------

In Go, there is no direct way to terminate a running goroutine from outside of it. Goroutines can only be terminated from within by returning from the function or using channels to signal termination. As a result:

- Calling `Wait()` on an infinitely running task will block the calling routine indefinitely.
- Calling `WaitTimeout()` or `WaitCtx()` on an infinitely running goroutine will result in a new locked routine being spawned.
- If `WaitTimeout()` or `WaitCtx()` is called and then exited due to timeout or cancellation, the internally spawned routine will continue to stay alive until the task itself is finished.

FAQ (Frequently Asked Questions)
--------------------------------

**Q: Is GoSync a replacement for goroutines?**

**A:** No, GoSync provides shortcuts and helpers to maintain cleaner and simpler code. However, it's important to have a foundational understanding of goroutines at a lower level and their limitations to avoid misusing the library.

**Q: Why did you create another library?**

**A:** While there are numerous tutorials on implementing similar functionality, I often found myself duplicating code across different projects. Although there is an existing implementation like [go-asynctask](https://github.com/Azure/go-asynctask), GoSync is designed to be cleaner and simpler, following idiomatic Go practices.

Contributing
------------

If you find any issues or have suggestions for improvements, feel free to contribute to the project.

License
-------

Cobra is released under the Apache 2.0 license. See [LICENSE.txt](LICENSE.txt)
