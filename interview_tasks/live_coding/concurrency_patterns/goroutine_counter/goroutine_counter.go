// Необходимо реализовать счетчик горутин
// Пользователь выполняет запрос на создание файлов и в любой момент времени хочет
// узнать сколько горутин еще выполняется

package main

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services       *Service // какой то сервис
	activeRoutines int32    // Счетчик активных горутин
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		services:       service,
		activeRoutines: 0, // присваиваем 0 в конструкторе
	}
}

func (h *Handler) InitRoutes() http.Handler {
	routers := gin.Default()
	routers.POST("files/create", h.CreateFile)
	routers.GET("files/check", h.CheckFiles)
}

func (h *Handler) CreateFiles(ctx *gin.Context) {
	// Берется из БД для определенного пользователя

	countReportFiles := h.services.Files.GetFiles(userID)

	// Здесь формируются файлы от 5 до 60 минут

	for i := 0; i < len(countReportFiles); i++ {
		atomic.AddInt32(&h.activeRoutines, 1) // при помощи атомиков добавляем активную горутину к счетчику

		go func() {
			// какой то скрипт создания файлов
			_ = GenerateReport(countReportFiles[i])

			atomic.AddInt32(&h.activeRoutines, -1) // при помощи атомиков уменьшаем счетчик горутин
		}()
	}
}

func (h *Handler) CheckFiles(ctx *gin.Context) {
	// Получает количество оставшихся файлов к формированию

	atomic.LoadInt32(&h.activeRoutines)                                 // выгружаем активные горутины
	ctx.JSON(http.StatusOK, gin.H{"active_routines": h.activeRoutines}) // отдаем ответ
}

func GenerateReport(url string) error {
	// Время выполнения рандомное от 5 до 60 минут
	return nil
}
