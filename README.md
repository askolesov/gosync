GoSync: Concurrent Synchronization Library for Go
=================================================

GoSync is a lightweight and user-friendly library designed to manage concurrent synchronization in Go applications. It simplifies the experience for users familiar with async-await constructs.

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

Contributing
------------

If you find any issues or have suggestions for improvements, feel free to contribute to the project.

License
-------

GoSync is released under the [MIT License](https://opensource.org/licenses/MIT).