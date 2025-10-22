package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//===========================================================
//Задача 5
//1. Релизовать ручку так, чтобы она выполнялась быстрее чем за одну секунду
//2. Теперь допустим, что запрашивается температура в каком-то location_id. Опиши, как это реализовать.
//===========================================================

// Есть функция getWeather, которая через нейронную сеть вычисляет температуру за ~1 секунду
// Есть highload ручка /weather/highload с нагрузкой 3k RPS
// Необходимо реализовать код этой ручки

var weatherCache sync.Map // хранит последнее значение температуры

func getWeather() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func updater(locationID string) {
	for {
		temp := getWeather()
		weatherCache.Store(locationID, temp)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	locations := []string{"moscow", "london", "tokyo"}
	for _, loc := range locations {
		go updater(loc)
	}

	http.HandleFunc("/weather/highload", func(resp http.ResponseWriter, req *http.Request) {
		loc := req.URL.Query().Get("location_id")
		if loc == "" {
			http.Error(resp, "missing location_id", http.StatusBadRequest)
			return
		}

		if temp, ok := weatherCache.Load(loc); ok {
			fmt.Fprintf(resp, "Location: %s, Temperature: %d°C", loc, temp.(int))
		} else {
			http.Error(resp, "location not found", http.StatusNotFound)
		}
	})

	fmt.Println("Listening on port 8080")
}
