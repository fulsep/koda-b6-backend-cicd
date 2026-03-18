package di

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/services"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db *pgxpool.Pool

	userRepo    *repository.UserRepo
	userService *services.UserService
	userHandler *handler.UserHandler

	authService *services.AuthService
	authHandler *handler.AuthHandler
}

func NewContainer() *Container {
	cfg, err := pgx.ParseConfig("")
	if err != nil {
		log.Fatalln("Failed to parse DB Config")
	}
	// db, err := pgx.Connect(context.Background(), cfg.ConnString())
	db, err := pgxpool.New(context.Background(), cfg.ConnString())
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

	c.authService = services.NewAuthService(c.userRepo)
	c.authHandler = handler.NewAuthHandler(c.authService)

}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func (c *Container) AuthHandler() *handler.AuthHandler {
	return c.authHandler
}

func (c *Container) Close() {
	c.db.Close()
}
