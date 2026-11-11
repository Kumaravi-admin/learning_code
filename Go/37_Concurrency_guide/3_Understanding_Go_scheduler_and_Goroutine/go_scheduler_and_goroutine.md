# Goroutines

*   **Baseline Requirement**: Every Go program contains at least one main goroutine.
*   **Execution Initiation**: The `go` keyword spins up new, independent concurrent goroutines.

# Threads

*   **Execution Units**: Threads represent the literal fundamental units of execution.
*   **Process Containers**: Threads exist within and belong entirely to a parent process.
*   **OS Orchestrated**: Operating system kernels directly manage thread creation and context switching.

# Green Threads

*   **User Space Managed**: Green threads are managed by user-level runtimes instead of the OS.
*   **Highly Lightweight**: Consume significantly less memory and overhead compared to native threads.
*   **Cooperative Scheduling**: Rely on code paths or execution cycles to yield control voluntarily.

# Coroutines

*   **Cooperative Subroutines**: Allow multiple entry points for suspending and resuming execution cooperatively.
*   **Resource Efficient**: Provide lightweight concurrency primitives without traditional kernel thread weight.
*   **Structural Assistance**: Help organize complex asynchronous program flow into readable linear code.

# Benefits of Coroutines

*   **Simpler Concurrency**: Eliminate the mental burden of traditional multi-threaded programming patterns.
*   **Reduced Errors**: Decrease the incidence of race conditions and synchronization errors.
*   **Highly Composable**: Allow developers to easily chain and combine modular concurrent behaviors.

# Traditional Concurrency

Direct mapping of native OS threads directly to hardware processor cores:

```text
 ┌──────────────┐                    ┌──────────────┐
 │              │                    │              │
 │   Thread 1   │                    │   Thread 2   │
 │              │                    │              │
 └──────────────┘                    └──────────────┘
         ▲                                   ▲
         │                                   │
         │                                   │
         │            ╭────────╮             │
         └────────────┤  Core  ├─────────────┘
                      ╰────────╯
```

# Go Scheduler

### Single Thread Multiplexing
The Go runtime schedules multiple lightweight goroutines onto a single OS thread using an internal queue:

```text
 ┌──────────────┐     ┌─────┐     ┌─────┐     ┌─────┐
 │              │ ◄── │  1  │ ◄── │  2  │ ◄── │  3  │
 │   Thread 1   │     └─────┘     └─────┘     └─────┘
 │              │
 └──────────────┘            Goroutine Queue
         ▲
         │
         │
     ╭────────╮
     │  Core  │
     ╰────────╯
```

### Multi-Thread Work Distribution (M:N Scheduler)
Goroutines are dynamically distributed across multiple OS threads operating on hardware cores:

```text
    ┌─────┐           ┌─────┐
    │  1  │           │  2  │
    └─────┘           └─────┘
       ▲                 ▲
       │                 │
 ┌────────────┐    ┌────────────┐     ┌─────┐     ┌─────┐     ┌─────┐
 │            │    │            │ ◄── │  3  │ ◄── │  4  │ ◄── │  5  │
 │  Thread 1  │    │  Thread 2  │     └─────┘     └─────┘     └─────┘
 │            │    │            │
 └────────────┘    ┌────────────┘          Goroutine Queue
       ▲                 ▲
       │                 │
       └──────┐   ┌──────┘
              ╭────────╮
              │  Core  │
              ╰────────╯
```
