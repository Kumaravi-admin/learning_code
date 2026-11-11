# Concurrency & Parallelism

- **Context Dependent**: Concurrency is highly context-dependent.
- **Distinct Concepts**: Parallelism is fundamentally different from concurrency.
- **Broader Scope**: Concurrency encompasses parallelism.

# Challenges & CSP

- **Encapsulation Limits**: Traditional encapsulation makes it difficult to write concurrent programs.
- **CSP Foundation**: Go's concurrency model is heavily inspired by Communicating Sequential Processes (CSP).

# Goroutines & Channels

- **Lightweight Execution**: Goroutines function as lightweight, managed threads.
- **Communication Pipelines**: Channels act as data pipelines for safe communication between routines.
- **Runtime Managed**: The Go runtime dynamically manages most underlying concurrency complexity.

# More About Channels

- **Data Exchange**: Safely transfer ownership of data between concurrent routines.
- **Combining Outputs**: Fan-in multiple channel streams into a unified data destination.
- **Control Flow**: Provide structural mechanics for implementing timeouts and operations cancellations.

# The Select Statement

- **Event Handling**: Blocks until one of several communication operations is ready to proceed.
- **Multi-Channel Multiplexing**: Directly selects from multiple active channel streams.

# Traditional Primitives

- **Sync Package**: Low-level synchronization primitives are fully available via the `sync` package.
- **Targeted Usage**: They should be used deliberately when specialized performance or memory sharing is needed.
