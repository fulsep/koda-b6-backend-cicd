package di

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/services"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	db *pgx.Conn

	userRepo    *repository.UserRepo
	userService *services.UserService
	userHandler *handler.UserHandler
}

func NewContainer() *Container {
	cfg, err := pgx.ParseConfig("")
	if err != nil {
		log.Fatalln("Failed to parse DB Config")
	}
	db, err := pgx.Connect(context.Background(), cfg.ConnString())
	if err != nil {
		log.Fatalln("Failed to connect to DB")
	}

	c := &Container{
		db: db,
	}

	c.initDeps()

	return c
}

func (c *Container) initDeps() {
	c.userRepo = repository.NewUserRepo(c.db)
	c.userService = services.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)

}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func (c *Container) Close() error {
	return c.db.Close(context.Background())
}
