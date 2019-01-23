# Concurrency in Go 
[By Katherine Cox-Buday][1]

- Understand how Go addresses fundamental problems that make concurrency difficult to do correctly
- Learn the key differences between concurrency and parallelism
- Dig into the syntax of Go’s memory synchronization primitives
- Form patterns with these primitives to write maintainable concurrent code
- Compose patterns into a series of practices that enable you to write large, distributed systems that scale
- Learn the sophistication behind goroutines and how Go’s runtime stitches everything together 

## 1 - An Introduction to Concurrency

- **Concurrent** - referring to a process that happens simultaneously with one or more processes 

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
  - [The free lunch is over: A fundamental turn toward concurrency in software][2]
  - "We desperately need a higher-level programming model for concurrency than languages offer today"

### Why is Concurrency Hard?

### Atomicity

### Memory Access Synchronization

### Deadlocks, Livelocks, and Starvation

#### Deadlock

#### Livelock

#### Starvation

### Determining Concurrency Safety

### Simplicity in the Face of Complexity


## 2 - Modeling Your Code: Communicating Sequential Processes




[1]: http://shop.oreilly.com/product/0636920046189.do
[2]: http://gotw.ca/publications/concurrency-ddj.htm
