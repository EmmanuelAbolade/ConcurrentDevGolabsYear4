Concurrent Development Go Labs - Year 4

Author: Emmanuel Abolade  
Course: Concurrent Development, Year 4  
Institution: South East Technological University (SETU)  
Instructor: Dr. Joseph Kehoe  
Academic Year: 2025/26

Table of Contents

- [Overview](#overview)
- [Repository Structure](#repository-structure)
- [Labs Included](#labs-included)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Programs](#running-the-programs)
- [Lab Descriptions](#lab-descriptions)
- [Key Concepts](#key-concepts)
- [License](#license)
- [Acknowledgments](#acknowledgments)

Overview

This repository contains laboratory exercises exploring concurrent programming concepts in Go (Golang). The labs demonstrate various synchronization primitives, classical concurrency problems, and their solutions.

Learning Objectives:
- Understanding goroutines and channels
- Implementing synchronization using mutexes, semaphores, and barriers
- Solving classical concurrency problems (Dining Philosophers, etc.)
- Preventing deadlocks and race conditions
- Using atomic operations for lock-free programming

Repository Structure


ConcurrentDevGolabsYear4/
├── LICENSE                    # GPL v3 License
├── README.md                  # This file
├── go.mod                     # Go module file
├── go.sum                     # Go dependencies
│
├── mutex/
│   ├── mutex.go              # fixed version with improvement
│   └── mutex_fixed.go        # Go module file
├── atomic/
│   ├── atomic.go             # Fixed version
│   └── go.mod                # Go module file
│
├── semaphore/
│   ├── semaphore.go          # Struct based implementation
│   └── go.mod                # Go module file
│
├── barrier/
│   ├── barrier(2).go         # Barrier synchronization
│   ├── barrier2.go           # Struct based implementation
│   └── go.mod                # Go module file
│
├── rendezvous/
│   ├── rendezvous.go         # Rendezvous pattern
│
├── signalling/
│   ├── signalling.go         # signalling pattern
│   └── go.mod                # Go module file
├── dining-philosophers/  
│   ├── dinPhil.go            # Asymmetric solution
│   └── go.mod                # Go module file
│
└── fibonacci/
    ├── fib.go                # Parallel Fibonacci fixed version with threshold
    └── go.mod                # Go module file
```

Labs Included

 1. Mutex and Atomic Operations
- Basic mutex usage for protecting shared variables
- Atomic operations as an alternative to mutexes
- Performance comparison

 2. Semaphores
- Binary and counting semaphores
- Worker pool pattern
- Rate limiting with semaphores

 3. Barriers
- Barrier synchronization primitive
- Reusable vs single-use barriers
- Multiple implementation strategies

 4. Rendezvous
- Two-thread synchronization
- Channel-based signaling
- Coordination patterns

 5. Dining Philosophers Problem
- Classical deadlock problem
- Three deadlock-free solutions:
  - Footman (limiting concurrency)
  - Asymmetric (breaking symmetry)
  - Resource hierarchy (ordering resources)

 6. Parallel Fibonacci
- Recursive parallel computation
- Threshold-based parallelization
- Goroutine overhead analysis

 Prerequisites

- Go: Version 1.21 or higher
- Git: For cloning the repository
- GoLand IDE (optional but recommended)

 Required Go Packages

bash
go get golang.org/x/sync/semaphore


Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/EmmanuelAbolade/ConcurrentDevGolabsYear4.git
   cd ConcurrentDevGolabsYear4
   ```

2. Initialize Go module:
   ```bash
   go mod init lab
   go mod tidy
   ```

3. Download dependencies:
   ```bash
   go get golang.org/x/sync/semaphore
   ```

Running the Programs

#General Format
```bash
go run path/to/file.go
```

#Examples

Mutex example:
```bash
go run mutex/mutex_fixed.go
```

Dining Philosophers (Asymmetric solution):
```bash
go run dining-philosophers/dinPhil.go
```

Barrier example:
```bash
go run barrier2.go
```

Parallel Fibonacci:
```bash
go run fib.go
```

# Running All Tests
```bash
# From project root
go run ./...
```

Lab Descriptions

# Lab 1: Mutex Synchronization

Problem: Multiple goroutines incrementing a shared counter cause race conditions.

Solution: Use `sync.Mutex` to protect critical sections.

Files: `mutex/mutex.go`.

Key Learning: Proper use of locks to prevent data races.

---

# Lab 2: Atomic Operations

Problem: Mutex overhead for simple operations.

Solution: Use `sync/atomic` for lock-free operations.

Files: `atomic/atomic.go`.

Key Learning: When to use atomic operations vs mutexes.

---

# Lab 3: Semaphores

Problem: Limiting concurrent access to resources.

Solution: Implement counting semaphores using buffered channels or `golang.org/x/sync/semaphore`.

Files: `semaphore/semaphore.go`.

Key Learning: Resource management and worker pools.

---

# Lab 4: Barriers

Problem: Synchronizing multiple goroutines at a meeting point.

Solution: Implement barriers using WaitGroups, channels, or sync.Cond.

Files: `barrier/barrier2.go`.

Key Learning: Phase-based synchronization.

---

# Lab 5: Rendezvous

Problem: Two goroutines need to meet and exchange information.

Solution: Use unbuffered channels for synchronization.

Files: `rendezvous/rendezvous.go`, `signalling/signalling.go`

Key Learning: Two-way synchronization patterns.

---

# Lab 6: Dining Philosophers

Problem: Classical deadlock scenario with 5 philosophers and 5 forks.

Requirements:
- No deadlock
- No starvation  
- Multiple philosophers can eat simultaneously

Solutions Implemented:

Asymmetric Solution: One philosopher picks up forks in opposite order
   - File: `dinPhil_asymmetric.go`
   - Breaks: Circular wait condition

Key Learning: Deadlock prevention strategies.

---

# Lab 7: Parallel Fibonacci

Problem: Naive parallel Fibonacci creates exponential goroutines.

Solution: Use threshold to limit parallelization depth.

Files: `fibonacci/fib.go`, `fibonacci/fib_fixed.go`

Key Learning: When parallelization hurts performance.

---

 Key Concepts

# Synchronization Primitives

| Primitive       | Use Case              | Go Implementation     |
|-----------      |----------             |-------------------    |
| Mutex           | Mutual exclusion      | `sync.Mutex`          |
| RWMutex         | Read-heavy workloads  | `sync.RWMutex`        |
| WaitGroup       | Wait for goroutines   | `sync.WaitGroup`      |
| Channel         | Communication         | `make(chan T)`        |
| Semaphore       | Resource limiting     | `semaphore.Weighted`  |
| Atomic          | Lock-free ops         | `sync/atomic`         |

# Deadlock Conditions (Coffman)

Deadlock requires ALL four:
1. Mutual Exclusion: Resources can't be shared
2. Hold and Wait: Holding resources while waiting
3. No Preemption: Can't force resource release
4. Circular Wait: Cycle in resource graph

Prevention: Break ANY one condition!

# Race Conditions

Detection:
```bash
go run -race program.go
```

Prevention:
- Use mutexes
- Use channels
- Use atomic operations
- Avoid shared mutable state

# License

This project is licensed under the **GNU General Public License v3.0**.

See [LICENSE](LICENSE) file for full details.

Summary:
- ✓ Free to use, modify, and distribute
- ✓ Must disclose source code
- ✓ Must use same license (copyleft)
- ✓ Changes must be documented

# Acknowledgments

- Dr. Joseph Kehoe - Course instructor and original template code
- South East Technological University (SETU) - Educational institution
- Go Team at Google - For the excellent Go programming language
- The Little Book of Semaphores by Allen B. Downey - Inspiration for problems

# Contact

Student: Emmanuel Abolade  
GitHub: [@EmmanuelAbolade](https://github.com/EmmanuelAbolade)  
Repository: [ConcurrentDevGolabsYear4](https://github.com/EmmanuelAbolade/ConcurrentDevGolabsYear4)

# Additional Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [The Go Memory Model](https://go.dev/ref/mem)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example: Goroutines](https://gobyexample.com/goroutines)

---

Last Updated: October 2025  
Status: Active Development
