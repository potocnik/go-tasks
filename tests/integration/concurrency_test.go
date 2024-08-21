package integration_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	tasks "tasks/pkg/task_list"
	"testing"
	"time"
)

var TaskList = []string{}

func TestConcurrency(t *testing.T) {
	t.Run("push with empty list", func(t *testing.T) {
		taskData := bytes.NewBuffer([]byte("[\"Task 1\",\"Task 2\",\"Task 3\",\"Task 4\",\"Task 5\",\"Task 6\",\"Task 7\",\"Task 8\",\"Task 9\",\"Task 10\"]"))
		TaskList = tasks.LoadState(taskData)
		var wg sync.WaitGroup
		for i := 1; i <= 100; i++ {
			delay := time.Duration(rand.Intn(i * 100))
			count := rand.Intn(5)
			for j := 1; j <= count; j++ {
				wg.Add(1)
				go runConcurrent(i, &wg, delay)
			}
		}
		wg.Wait()
	})
}

func runConcurrent(iteration int, wg *sync.WaitGroup, delay time.Duration) {
	fmt.Println("Concurrent " + fmt.Sprintf("%d", iteration) + " " + fmt.Sprintf("%d", delay) + " ms")
	time.Sleep(time.Millisecond * delay)
	identity := fmt.Sprintf("%d", iteration)
	operation := rand.Intn(4)
	switch operation {
	case 0:
		fmt.Println("[DEBUG]: PUSH (" + identity + ")")
		TaskList = tasks.Push(TaskList, "Item "+fmt.Sprintf("%d", delay))
	case 1:
		fmt.Println("[DEBUG]: POP (" + identity + ")")
		TaskList, _ = tasks.Pop(TaskList)
	case 2:
		fmt.Println("[DEBUG]: PUT (" + identity + ")")
		position := rand.Intn(10)
		TaskList = tasks.Set(TaskList, position, "Item "+identity)
	}
	defer wg.Done()
}
