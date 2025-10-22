package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// есть компонент, который ходит по урлам, делает гет запросы и выводит в stdout ответ
// проблемы: 1) работает очень медленно
//             2) при первой же ошибке программа падает
// хотим:     1) чтобы работало быстро
//             2) чтобы программа не падала, а в stdout выводилась информативная ошибка
// дополнение: посчитать общее количество байт во всех ответах

func main() {
	start := time.Now()
	urls := []string{
		"https://mail.ru",
		"https://yandex.ru",
		"https://drive.google.com/drive/my-drive",
		"https://avito.ru",
		"https://yandex.ru/maps/51/samara/?ll=50.194198%2C53.195417&z=13",
		"https://github.com/",
	}

	channel := make(chan string)

	var wg sync.WaitGroup
	// var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			content, err := fetchURL(url)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			channel <- content
			// fmt.Println(content)
			//mu.Lock()
			//count += len(content)
			//mu.Unlock()
		}(url)
	}

	wg.Wait()

	close(channel)

	count := 0
	for c := range channel {
		fmt.Printf("Content: %s", c)
		count += len(c)
	}

	fmt.Print(count/1024/1024, "MB")
	fmt.Println(time.Since(start))
}

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
