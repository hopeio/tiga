package tiny_engine

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestBaseEngine(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	engine := NewEngine[int, int, int](10)

	tasks := make([]*Task[int, int], 100)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = baseTaskGen(strconv.Itoa(i), engine)
	}
	engine.Run(tasks...)
}

func baseTaskGen(id string, engine *Engine[int, int, int]) *Task[int, int] {
	return &Task[int, int]{TaskFunc: func(ctx context.Context) {
		log.Println("task", id)
		n := rand.Intn(10)
		//log.Println("rand", n)
		if n < 3 {
			for i := 0; i < n; i++ {
				engine.AddTasks(baseTaskGen(id+"_"+strconv.Itoa(i), engine))
			}
		}
		if n == 3 {
			panic(n)
		}
	}}
}

func TestBaseEngineOneTask(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	engine := NewEngine[int, int, int](10)
	ch := make(chan string)
	go func() {
		for {
			id := <-ch
			log.Println("rand", id)
			for i := 0; i < 3; i++ {
				engine.AddTasks(taskgen2(id+"_"+strconv.Itoa(i), ch))
			}
		}
	}()
	engine.Run(&Task[int, int]{
		TaskFunc: func(ctx context.Context) {
			ch <- "1"
		},
	})
}

func taskgen2(id string, ch chan string) *Task[int, int] {
	return &Task[int, int]{
		TaskFunc: func(ctx context.Context) {
			log.Println("task", id)
			n := rand.Intn(10)
			//log.Println("rand", n)
			if n == 5 {
				panic("5")
			}
			if n < 3 {
				ch <- id
			}
		},
	}
}
