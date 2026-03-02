<body>
  <header>
    <h1>Go Worker Pool</h1>
    <p class="meta">A small Go example demonstrating a worker-pool pattern using goroutines, channels and sync.WaitGroup.</p>
  </header>

  <section>
    <h2>Overview</h2>
    <p>This page shows a compact, runnable Go example that creates a pool of workers which receive tasks from a channel, process them concurrently, and send results back on another channel. The example highlights directional channels (receive-only / send-only), waiting for workers using <code>sync.WaitGroup</code>, and safe channel closing.</p>
  </section>

  <section>
    <h2>Components</h2>
    <ul>
      <li><strong>Worker</strong> — Go routine that processes taska from pool.
</li>
      <li><strong>tasks (channel)</strong> — medium through which tasks are sent to the worker (receive only channel).</li>
      <li><strong>results (channel)</strong> — send response from each worker on this channel (send only channel).</li>
      <li><strong>sync.WaitGroup</strong> — ensures the main goroutine knows when all workers finished, so it can safely close.</li>
    </ul>
  </section>

  <section>
    <h2>Steps</h2>
    <ol>
      <li>Create task channel.</li>
      <li>Spawn workers.</li>
      <li>send tasks to the pool (we will send task to channel 1 and then which ever worker will be free will take that task).</li>
      <li>wait for the workers to complete.</li>
      <li>Collect/print results by ranging over <code>results</code>.</li>
    </ol>
  </section>

  <section>
    <h2>Sample Output (possible)</h2>
    <pre><code>Task 2 is picked up by the worker 3
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
49</code></pre>
    <p class="note"><strong>Note:</strong> The ordering of which worker picks a task and the order results are printed is <em>not</em> guaranteed. Workers run concurrently so task assignment and result arrival interleave nondeterministically.</p>
  </section>

  <section>
    <h2>Important Concepts Used</h2>
    <ul>
      <li><strong>Goroutine</strong>: A lightweight thread managed by Go runtime, started with <code>go functionName()</code>.</li>
      <li><strong>Channel</strong>: A typed communication mechanism between goroutines.<code>ch := make(chan T)</code> Examples:
        <ul>
          <li>Unbuffered channel: <code>make(chan T)</code></li>
          <li>Buffered channel: <code>make(chan T, N)</code></li>
        </ul>
      </li>
      <li><strong>Directional channels</strong> (for safety): declare in function signature as <code>tasks &lt;-chan int</code> (receive-only) and <code>results chan&lt;- int</code> (send-only).</li>
      <li><strong>sync.WaitGroup</strong>: Used to manage lifecycle for multiple goroutines; use <code>wg.Add(n)  (increase counter)</code>, <code>wg.Done()	(decrease counter)</code>, and <code>wg.Wait()  (block until counter becomes 0)</code>.</li>
      <li><strong>Closing channels</strong>: After workers finish, close <code>results</code> so the receiver can range over it safely.</li>
    </ul>
  </section>

  <footer style="margin-top:28px; color:#475569; font-size:14px">
    <p>Created for quick reference — feel free to copy the code into a <code>.go</code> file and run with <code>go run</code>.</p>
  </footer>
</body>
</html>
