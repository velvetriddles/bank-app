package di

import (
	"github.com/velvetriddles/bank-app/config"
	"github.com/velvetriddles/bank-app/internal/delivery/http/handler"
	"github.com/velvetriddles/bank-app/internal/repository"
	"github.com/velvetriddles/bank-app/internal/repository/memory"
	"github.com/velvetriddles/bank-app/internal/usecase"
	"github.com/velvetriddles/bank-app/pkg/logger"
)

type Container struct {
	Config            config.Config
	Logger            logger.Logger
	AccountRepository repository.AccountRepository
	AccountUseCase    *usecase.AccountUseCase
	AccountHandler    *handler.AccountHandler
}

func NewContainer(cfg config.Config) *Container {
	return &Container{
		Config: cfg,
	}
}

func (c *Container) Init() error {
	c.Logger = logger.NewLogger()
	switch c.Config.Database.Type {
	case "memory":
		c.AccountRepository = memory.NewInMemoryAccountRepository()
	default:
		c.AccountRepository = memory.NewInMemoryAccountRepository()
	}

	c.AccountUseCase = usecase.NewAccountUseCase(c.AccountRepository, c.Logger)

	c.AccountHandler = handler.NewAccountHandler(c.AccountUseCase, c.Logger)

	return nil
}
