package user

import "github.com/BrandokVargas/api-ecommerce/model"

//PUERTO DONDE SE VAN A COMUNICAR LOS DATOS DE ENTRADA
type UseCase interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

//LO QUE NECESITAMOS PARA LOS DATOS DE SALIDA
type Storage interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
