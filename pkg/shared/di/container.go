package di

import (
	"basesvc_v2/internal/config"
	"basesvc_v2/pkg/adapter/repository/db"
	"basesvc_v2/pkg/infrastructure/rest/service/dto"
	"basesvc_v2/pkg/usecase/author"

	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()

	_ = builder.Add([]di.Def{
		{Name: "ausvc", Build: auUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func auUsecase(_ di.Container) (interface{}, error) {
	configuration := config.GetConfig()
	authorDb := db.NewAuthor(configuration.Database)
	crBuilder := &dto.AuthorkBuilder{}
	return author.NewAuthorUsecase(authorDb, crBuilder), nil
}
