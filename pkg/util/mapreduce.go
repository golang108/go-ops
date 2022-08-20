package util

func MapReduce(mapper func(interface{}, chan interface{}),
	reducer func(chan interface{}, chan interface{}),
	input chan interface{}) interface{} {

	reduceInput := make(chan interface{})
	reduceOutput := make(chan interface{})
	workerOutput := make(chan interface{})

	workProcesses := 0

	go reducer(reduceInput, reduceOutput)

	for item := range input {
		go mapper(item, workerOutput)
		workProcesses += 1
	}

	for item := range workerOutput {
		workProcesses -= 1
		reduceInput <- item
		if workProcesses <= 0 {
			close(reduceInput)
			break
		}
	}
	return <-reduceOutput
}
