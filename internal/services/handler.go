package services

import (
	"fmt"
	"kssyncservice_go/pkg/res"
	"net/http"
)

type ServicesHandlerDeps struct{
	ServicesRepository *ServicesRepository
}

type ServicesHandler struct{
	ServicesRepository *ServicesRepository
}

func NewServicesHandler(router *http.ServeMux, deps ServicesHandlerDeps) {
	handler := &ServicesHandler{
		ServicesRepository: deps.ServicesRepository,
	}
	router.HandleFunc("/services", handler.getAllServices())
}

func (handler *ServicesHandler) getAllServices() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Получение данных из БД")
		services, err := handler.ServicesRepository.GetAllServices()
		fmt.Println("Данные получены")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := GetAllServicesResponse{
			Count: len(*services),
			Text: "Тестовый текст",
			Data: *services,
		}
		fmt.Println("Отправка данных в ответ")
		res.Json(w, data, 200)
	}
}
