# go-worker-pool
A small Go example demonstrating a worker-pool pattern using goroutines, channels and sync.WaitGroup.

-> Worker : Go routine that processes taska from pool
-> channel1 : medium through which tasks are sent to the worker (receive only channel)
-> channel2 : send response from each worker on this channel (send only channel)


Steps:
-> Create task channel
-> Spawn workers
-> send tasks to the pool (we will send task to channel 1 and then which ever worker will be free will take that task)
-> wait for the workers to complete


Sample Output: 
Task 2 is picked up by the worker 3
Task 3 is picked up by the worker 2
Task 1 is picked up by the worker 1
Task 4 is picked up by the worker 2
Task 5 is picked up by the worker 3
Task 6 is picked up by the worker 1
Task 7 is picked up by the worker 3
9
4
1
25
16
36
49
**Note: Result ordering is not guaranteed because workers run concurrently.


Important Concepts Used:
-> Goroutine: A lightweight thread managed by Go runtime.
   Started using: go functionName()
   
-> Channel: A communication mechanism between goroutines.
   ch := make(chan int)

-> Types:
Unbuffered channel → make(chan int)
Buffered channel → make(chan int, N)
Buffered channels allow storing N values before blocking.

-> Directional Channels: Used in function signatures for safety:
   tasks <-chan int    // receive-only
   results chan<- int  // send-only
This prevents misuse inside worker functions.

-> sync.WaitGroup: Used to wait for multiple goroutines to finish execution.
  Methods:
    wg.Add(n) → increase counter
    wg.Done() → decrease counter
    wg.Wait() → block until counter becomes zero
  Typical usage:
    wg.Add(1)
    go func() {
      defer wg.Done()
    }()
    wg.Wait()
    
-> Why Close Channels?
  close(tasks) → signals workers that no more tasks will come.
  close(results) → done after all workers finish.
  Closing a channel allows range loops to exit safely.
