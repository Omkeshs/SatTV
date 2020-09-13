package cmd

import (
	"std/omkesh/Practice/SatTV/db"
	"std/omkesh/Practice/SatTV/services/pkg/menu"
)

//Run **
func Run() {
	conn := db.NewDbConnection()
	repository := repository.NewRepositoryImpl(conn)
	service := service.NewServiceImpl(repository)
	handler := handler.NewHandlerImpl(service)
	menu.Menu(handler)
}
