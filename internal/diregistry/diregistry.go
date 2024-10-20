package diregistry

import (
	"example/config"
	"example/internal/common/helper/dihelper"
	"example/internal/common/helper/validatehelper"
	"example/internal/controller"
	"example/internal/repository"
	"example/internal/usecase"

	"github.com/sarulabs/di"
)

const (
	ConfigDIName            string = "Config"
	ValidateDIName          string = "Validate"
	UserRepositoryDIName    string = "UserRepository"
	ProductRepositoryDIName string = "ProductRepository"
	UserUseCaseDIName       string = "UserUseCase"
	UserControllerDIName    string = "UserController"
)

func BuildDIContainer() {
	initBuilder()
	dihelper.BuildLibDIContainer()
}

func GetDependency(name string) interface{} {
	return dihelper.GetLibDependency(name)
}

func initBuilder() {
	dihelper.ConfigsBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ConfigDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg, err := config.LoadEnvironment()
				if err != nil {
					return nil, err
				}
				return cfg, nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.HelpersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ValidateDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return validatehelper.NewValidate(), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.RepositoriesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  UserRepositoryDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(ConfigDIName).(*config.Config)
				return repository.NewUserRepository(cfg), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  ProductRepositoryDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(ConfigDIName).(*config.Config)
				return repository.NewProductRepository(cfg), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.UseCasesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  UserUseCaseDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userRepository := ctn.Get(UserRepositoryDIName).(repository.UserRepository)
				return usecase.NewUserUseCase(userRepository), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.ControllersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  UserControllerDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userUseCase := ctn.Get(UserUseCaseDIName).(usecase.UserUseCase)
				return controller.NewUserController(userUseCase), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}
}
