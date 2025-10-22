//===========================================================
//Задача 15
//===========================================================

// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).

// Нужно написать обертку predictableFunc,
// которая будет работать с заданным фиксированным таймаутом (например, 1 секунду).

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Есть функция, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	start := time.Now()
	resCh := make(chan int64)

	go func() {
		resCh <- unpredictableFunc()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timed out, total time:", time.Since(start))
		return 0, ctx.Err()
	case result := <-resCh:
		fmt.Println("total time:", time.Since(start))
		return result, nil
	}
}

func main() {
	fmt.Println("started")

	ctx := context.Background()

	fmt.Println(predictableFunc(ctx))
}
