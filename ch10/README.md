# Chapter 10. Concurrency in Go

Concurrency - breaking up single processes into independent components and specifying how those components safely share data.

Communicating Sequential Processes (CSP)

## When to Use Concurrency

-> It's WAY harder than advertised. Especially if people are moving into Mutexes (bottlenecks)

Amdahl's Law - cost of synchronization!!! Bottlenecks again

> Use concurrency when you want to combine data from multiple operations that can operate independently.

There is a cost of concurrency -> This is why it's mostly used for I/O!
Write a sequential program first.

## Goroutines

