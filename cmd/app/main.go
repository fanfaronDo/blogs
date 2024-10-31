package main

import (
	"fmt"
	"github.com/fanfaronDo/blogs/internal/config"
	"github.com/fanfaronDo/blogs/internal/handler"
	"github.com/fanfaronDo/blogs/internal/repository"
	"github.com/fanfaronDo/blogs/internal/server"
	"github.com/fanfaronDo/blogs/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.ConfigLoad()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
	server := &server.Server{}
	connectToDB, err := repository.NewMySql(&cfg.Mysql)

	if err != nil {
		fmt.Println(err)
	}

	repo := repository.NewRepository(connectToDB)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	route := handler.InitRoutes()
	go func() {
		if err = server.Run(cfg.HttpServer.Address, cfg.HttpServer.Port, route); err != nil {
			fmt.Printf("error occured while running http server: %s\n", err.Error())
			return
		}
	}()

	fmt.Printf("Server started on %s\n", "http://"+cfg.HttpServer.Address+":"+cfg.HttpServer.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("Shutting down server...")

	defer server.Stop(nil)

}
