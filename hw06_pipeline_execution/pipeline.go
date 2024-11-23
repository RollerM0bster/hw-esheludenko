package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = moveToNext(stage(in), done)
	}
	return in
}

func moveToNext(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out) // Ensure the output channel is always closed
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return // Input channel closed, stop processing
				}
				select {
				case out <- v: // Pass value downstream
				case <-done: // Stop processing if done signal is received
					return
				}
			case <-done: // Stop processing if done signal is received
				return
			}
		}
	}()
	return out
}
