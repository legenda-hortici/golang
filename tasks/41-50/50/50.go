package main

import (
	"net/http"
	"sync"
)

//===========================================================
//Задача 13
//1. Запросить параллельно данные из источников. Если все таки где-то произошла ошибка, то вернуть ошибку, иначе вернуть nil.
//2. Представим, что теперь функция должна возвращать результат int. Есть функция resp.Size(), для каждого url надо проссумировать и вернуть, если ошибок не было.
//   Просто описать подход к решению
//3. Что делать, если урлов у нас миллионы?
//
//===========================================================

func main() {
	_, err := download([]string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
	})

	if err != nil {
		panic(err)
	}
}

func download(urls []string) (int64, error) {
	var wg sync.WaitGroup
	errCh := make(chan error, len(urls))
	sizeCh := make(chan int64, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				errCh <- err
				return
			}
			defer func() {
				_ = resp.Body.Close()
			}()

			sizeCh <- resp.ContentLength
		}(url)
	}

	go func() { // разблокируем главный поток
		wg.Wait()
		close(errCh) // закрываем все каналы
		close(sizeCh)
	}()

	var lastErr error
	for err := range errCh {
		if err != nil {
			return -1, err
		}
		lastErr = err
	}

	sum := int64(0)
	for size := range sizeCh {
		sum += size
	}

	return sum, lastErr
}
