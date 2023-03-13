package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/HondaAo/snippet/src/pkg/driver/repositories"
	"github.com/HondaAo/snippet/src/pkg/handle"
	idiomUsecase "github.com/HondaAo/snippet/src/pkg/services/idioms"
	videoUsecase "github.com/HondaAo/snippet/src/pkg/services/videos"
	wordUsecase "github.com/HondaAo/snippet/src/pkg/services/word"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
	db   *gorm.DB
}

func NewServer(echo *echo.Echo, db *gorm.DB) *Server {
	return &Server{
		echo,
		db,
	}
}

func (s Server) Run() error {
	vRepository := repositories.NewVideoRepository(s.db)
	sRepository := repositories.NewScriptRepository(s.db)
	swRepository := repositories.NewScriptIdiomsRepository(s.db)
	iRepository := repositories.NewIdiomsRepository(s.db)
	wRepository := repositories.NewWordRepository(s.db)
	vUsecase := videoUsecase.NewVideoUsecase(vRepository, sRepository, swRepository, iRepository)
	iUsecase := idiomUsecase.NewIdiomUsecase(iRepository)
	wUsecase := wordUsecase.NewWordUsecase(wRepository)
	vHandler := handle.NewVideoHandler(vUsecase)
	iHandler := handle.NewIdiomHandler(iUsecase)
	wHandler := handle.NewWordHandler(wUsecase)

	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      s.echo,
	}
	s.echo.Use(middleware.Secure())
	s.echo.Use(middleware.Logger())

	// s.echo.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "joe" && password == "secret" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	api := s.echo.Group("/api/v1")
	videoGroup := api.Group("/video")
	idiomGroup := api.Group("/idiom")
	wordGroup := api.Group("/words")

	handle.RegisterVideoRoute(videoGroup, vHandler)
	handle.RegisterIdiomRoute(idiomGroup, iHandler)
	handle.RegisterWordRoute(wordGroup, wHandler)

	go func() {
		if err := s.echo.StartServer(server); err != nil {
			s.echo.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}
