<h1>go-worker-pool</h1>

A small Go example demonstrating a worker-pool pattern using goroutines, channels and sync.WaitGroup.<br>
<br>
<p>
-> Worker : Go routine that processes taska from pool<br>
-> channel1 : medium through which tasks are sent to the worker (receive only channel)<br>
-> channel2 : send response from each worker on this channel (send only channel)<br>
</p>
<br>
<p>
Steps:<br>
-> Create task channel<br>
-> Spawn workers<br>
-> send tasks to the pool (we will send task to channel 1 and then which ever worker will be free will take that task)<br>
-> wait for the workers to complete<br>
</p>
<br>
<p>
Sample Output: <br>
Task 2 is picked up by the worker 3<br>
Task 3 is picked up by the worker 2<br>
Task 1 is picked up by the worker 1<br>
Task 4 is picked up by the worker 2<br>
Task 5 is picked up by the worker 3<br>
Task 6 is picked up by the worker 1<br>
Task 7 is picked up by the worker 3<br>
9<br>
4<br>
1<br>
25<br>
16<br>
36<br>
49<br>
**Note: Result ordering is not guaranteed because workers run concurrently.<br>
</p>
<br>
<p>
Important Concepts Used:<br>
-> Goroutine: A lightweight thread managed by Go runtime.<br>
   Started using: go functionName()<br>
   <br>
-> Channel: A communication mechanism between goroutines.<br>
   ch := make(chan int)<br>
<br>
-> Types:<br>
Unbuffered channel → make(chan int)<br>
Buffered channel → make(chan int, N)<br>
Buffered channels allow storing N values before blocking.<br>
<br>
-> Directional Channels: Used in function signatures for safety:<br>
   tasks <-chan int    // receive-only<br>
   results chan<- int  // send-only<br>
This prevents misuse inside worker functions.<br>
<br>
-> sync.WaitGroup: Used to wait for multiple goroutines to finish execution.<br>
  Methods:<br>
    wg.Add(n) → increase counter<br>
    wg.Done() → decrease counter<br>
    wg.Wait() → block until counter becomes zero<br>
  Typical usage:<br>
    wg.Add(1)<br>
    go func() {<br>
      defer wg.Done()<br>
    }()<br>
    wg.Wait()<br>
<br>
-> Why Close Channels?<br>
  close(tasks) → signals workers that no more tasks will come.<br>
  close(results) → done after all workers finish.<br>
  Closing a channel allows range loops to exit safely.<br>
</p>
