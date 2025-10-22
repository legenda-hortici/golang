package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Source string
}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
	// Когда получаем первый успешный результат - отдаем его сразу. Если все SearchFunc отработали
	// с ошибкой - отдаем последнюю полученную ошибку
	var wg sync.WaitGroup

	errorCh := make(chan error, len(sfs))
	resultCh := make(chan Result, len(sfs))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, sf := range sfs {
		wg.Add(1)
		go func(sf SearchFunc) {
			defer wg.Done()
			r, err := sf(ctx, query)
			if err != nil {
				errorCh <- err
				return
			}
			select {
			case resultCh <- r:
				cancel() // отменяем остальные поисковики
			case <-ctx.Done():
			}
		}(sf)
	}

	go func() { // разблокируем главный поток
		wg.Wait()
		close(errorCh) // закрываем все каналы
		close(resultCh)
	}()

	select {
	case r := <-resultCh: // первый успешный
		return r, nil
	case <-ctx.Done(): // если был сигнал отмены или таймаута
		var last error
		for err := range errorCh {
			last = err
		}
		if last != nil {
			return Result{}, last
		}
		return Result{}, ctx.Err()
	}
}

func main() {
	sf1 := func(ctx context.Context, query string) (Result, error) {
		return Result{}, errors.New("sf1 failed")
	}

	sf2 := func(ctx context.Context, query string) (Result, error) {
		return Result{Source: "sf2 success"}, nil
	}

	sf3 := func(ctx context.Context, query string) (Result, error) {
		time.Sleep(2 * time.Second)
		return Result{Source: "sf3 late"}, nil
	}

	res, err := MultiSearch(context.Background(), "hello", []SearchFunc{sf1, sf2, sf3})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Got result:", res.Source)
}
