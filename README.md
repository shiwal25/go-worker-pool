<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <title>go-worker-pool — README</title>
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <style>
    body { font-family: system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial; line-height: 1.5; padding: 20px; color: #111827; }
    code { background:#f3f4f6; padding:2px 6px; border-radius:4px; }
    pre { background:#111827; color:#f8fafc; padding:12px; border-radius:6px; overflow:auto; }
    h1 { font-size:1.6rem; margin-bottom:.25rem; }
    h2 { margin-top:1.1rem; }
    ul, ol { margin-left:1.2rem; }
    .note { background:#fffbeb; border-left:4px solid #f59e0b; padding:8px 12px; border-radius:4px; display:inline-block; }
  </style>
</head>
<body>

  <h1>go-worker-pool</h1>

  <p>A small Go example demonstrating a worker-pool pattern using goroutines, channels and <code>sync.WaitGroup</code>.</p>

  <p>→ <strong>Worker</strong> : Go routine that processes tasks from the pool.<br />
  → <code>channel1</code> : medium through which tasks are sent to the worker (receive-only channel).<br />
  → <code>channel2</code> : send response from each worker on this channel (send-only channel).</p>

  <h2>Steps</h2>
  <ol>
    <li>Create task channel.<br /></li>
    <li>Spawn workers.<br /></li>
    <li>Send tasks to the pool (we will send tasks to <code>channel1</code> and whichever worker is free will take that task).<br /></li>
    <li>Wait for the workers to complete.<br /></li>
  </ol>

  <h2>Sample Output</h2>
  <pre>
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
  </pre>
  <p><em>Note: Result ordering is not guaranteed because workers run concurrently.</em></p>

  <h2>Important Concepts Used</h2>
  <ul>
    <li><strong>Goroutine</strong>: A lightweight thread managed by Go runtime.<br />
      Started using: <code>go functionName()</code>.</li>

    <li><strong>Channel</strong>: A communication mechanism between goroutines.<br />
      Example: <code>ch := make(chan int)</code>.<br />
      <em>Unbuffered channel</em> → <code>make(chan int)</code><br />
      <em>Buffered channel</em> → <code>make(chan int, N)</code><br />
      Buffered channels allow storing N values before blocking.</li>

    <li><strong>Directional Channels</strong>: Used in function signatures for safety:<br />
      <code>tasks <-chan int</code>  <em>// receive-only</em><br />
      <code>results chan<- int</code> <em>// send-only</em><br />
      This prevents misuse inside worker functions.</li>

    <li><strong>sync.WaitGroup</strong>: Used to wait for multiple goroutines to finish execution.<br />
      Methods:<br />
      <code>wg.Add(n)</code> → increase counter<br />
      <code>wg.Done()</code> → decrease counter<br />
      <code>wg.Wait()</code> → block until counter becomes zero<br />
      Typical usage:<br />
      <pre>
wg.Add(1)
go func() {
  defer wg.Done()
  // work...
}()
wg.Wait()
      </pre>
    </li>

    <li><strong>Why close channels?</strong><br />
      <code>close(tasks)</code> → signals workers that no more tasks will come.<br />
      <code>close(results)</code> → done after all workers finish.<br />
      Closing a channel allows <code>range</code> loops to exit safely.</li>
  </ul>

  <div class="note">
    <strong>Tip:</strong> Use directional channels in worker function signatures to make intent explicit and reduce bugs.
  </div>

</body>
</html>
