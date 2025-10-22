package main

import (
	"context"
	"errors"
	"sync"
)

// Написать асинхронный обработчик задач в виде библиотеки.
//
// Клиент библиотеки передаёт на вход некоторый объект (Task), у которого
// есть метод Do().
// Клиент может передать несколько задач обработчику, задачи передаются
// по одной; задачи могут передаваться из разных рутин.
// Как только клиент передаст все задачи обработчику, он закрывает
// обработчик и ждёт завершения выполнения всех переданных задач.
//
// Обработчик внутри имеет свою очередь, откуда берёт задачи на выполнение.
// Очередь обработчика ограничена размером X; если при добавлении задачи
// она заполнена, обработчик сразу возвращает ошибку.
// Обработчик выполняет не более N задач одновременно.

// 1. Интерфейс асинхронного обработчика

// 2. Реализация

type Task interface {
	Do()
}

type TaskProccessor struct {
	queue      chan Task
	maxQueue   int
	maxWorkers int
	mu         sync.Mutex
	wg         sync.WaitGroup
	taskCtx    context.Context
	cancel     context.CancelFunc
	isClosed   bool
}

func NewTaskProccessor(ctx context.Context, maxQueue, maxWorkers int) *TaskProccessor {
	ctx, cancel := context.WithCancel(ctx)
	return &TaskProccessor{
		queue:      make(chan Task, maxQueue),
		maxQueue:   maxQueue,
		maxWorkers: maxWorkers,
		taskCtx:    ctx,
		cancel:     cancel,
	}
}

func (tp *TaskProccessor) Start() {
	for i := 0; i < tp.maxWorkers; i++ {
		tp.wg.Add(1)
		go tp.worker()
	}
}

func (tp *TaskProccessor) worker() {
	defer tp.wg.Done()
	for {
		select {
		case <-tp.taskCtx.Done():
			return
		case task, ok := <-tp.queue:
			if !ok {
				return
			}
			task.Do()
		}
	}
}

func (tp *TaskProccessor) Add(task Task) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	if tp.isClosed {
		tp.mu.Unlock()
		return errors.New("proccessor already closed")
	}

	select {
	case tp.queue <- task:
		return nil
	default:
		return errors.New("queue is full")
	}

}

func (tp *TaskProccessor) Stop() error {
	tp.mu.Lock()
	if tp.isClosed {
		tp.mu.Unlock()
		return errors.New("already closed")
	}
	tp.isClosed = true
	tp.mu.Unlock()
	tp.cancel()

	tp.wg.Wait()
	close(tp.queue)

	return nil
}
