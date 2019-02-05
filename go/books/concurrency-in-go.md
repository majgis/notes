# Concurrency in Go 
[By Katherine Cox-Buday][1]

- Understand how Go addresses fundamental problems that make concurrency 
difficult to do correctly
- Learn the key differences between concurrency and parallelism
- Dig into the syntax of Go’s memory synchronization primitives
- Form patterns with these primitives to write maintainable concurrent code
- Compose patterns into a series of practices that enable you to write large, 
distributed systems that scale
- Learn the sophistication behind goroutines and how Go’s runtime stitches 
everything together 

## 1 - An Introduction to Concurrency

- **Concurrent** - referring to a process that happens simultaneously with one 
or more processes 

### Moore's Law, Web Scale, and the Mess We're In

- Moore's Law (Gordon Moore, 1965)
  - Held true until around 2012
  - This threshold gave increased motive for multi-core

- Amdahl's Law (Gene Amdahl)
  - How to model gains from parallel solution to problem
  - Gains bounded by how much must be written sequentially

- Examples
  - serial
    - UI user interaction
  - parallel
    - calculating digits of Pi
      - spigot algorithms
      - embarrassingly parallel

- Scale horizontally 
  - Take instances of program and run on more CPUs or machines
  - Able to send chunks of problem to different instances of application
  - Became easier with cloud computing due to compute resources being on-demand
  - Also a challenge to aggregate/store results and model code concurrently
    - **web scale** - brand for software solving these issues

- Herb Sutter
  - [The free lunch is over: A fundamental turn toward concurrency in 
  software][2]
  - "We desperately need a higher-level programming model for concurrency than 
  languages offer today"

### Why is Concurrency Hard?

- Concurrent code is notoriously difficult to get right
  - Issue may lay dormant until change affects timing

#### Race Conditions

- **race condition** - when multiple operations must execute in a certain order 
but program is not written to guarantee that order
- **data race** - one operation attempts to read variable while another 
attempts to write
- Use *go* keyword to run a function concurrently
  - Doing so creates a *goroutine*
- [example][3]
- Prevention
  - Do not assume that if a line of code executes before another it will execute
first
  - Meticulously iterate through possible scenarios
  - Avoid trap of using sleeps to make less likely but not solve the issue
  - Always target logical correctness
- Changes to environment can evoke bug of untouched code

#### Atomicity

- **atomic** - within the context that is operating, it is indivisible
- Context is critical as atomicity of operation can change depending on scope
- Combining atomic operations does not necessarily produce a larger atomic 
operation
- If something is atomic, implicitly it is safe within concurrent contexts

#### Memory Access Synchronization

- **critical section** - section of program needing exclusive access to a 
shared resource
- Synchronizing access to the memory between critical sections is one way to 
solve problem
```go
package main
import (
	"sync"
)

func main (){
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()
}
```
- If developers don't follow this convention, there is no guarantee
of exclusive access.
- Doesn't automatically solve data races or logical correctness and 
potentially introduces maintenance and performance issues
- Two questions in context of program is an art:
    - Are my critical sections entered and exited repeatedly?
    - What size should my critical sections be?

#### Deadlocks, Livelocks, and Starvation

##### Deadlock

##### Livelock

##### Starvation

#### Determining Concurrency Safety

### Simplicity in the Face of Complexity


## 2 - Modeling Your Code: Communicating Sequential Processes




[1]: http://shop.oreilly.com/product/0636920046189.do
[2]: http://gotw.ca/publications/concurrency-ddj.htm
[3]: concurrency-in-go-example1.go