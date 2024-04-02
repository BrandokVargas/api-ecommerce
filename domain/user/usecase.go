package user

import (
	"fmt"
	"time"

	"github.com/BrandokVargas/api-ecommerce/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// LOGICA QUE NECESITO
// DOMINIO , NECESITA UN PUERTO
type User struct {
	storage Storage
}

// CONSTRUCTOR
func New(s Storage) *User {
	return &User{storage: s}
}

// CREANDO EL USUARIO
func (u User) Create(m *model.User) error {

	//1. DARLE UN IDENTIFICADOR ÚNICO
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w ", "uuid.NewUUID()", err)
	}
	//PASANDO EL ID A LA ESTRUCTURA USER
	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w ", "bcrypt.GenerateFromPassword()", err)
	}

	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte("{}")
	}

	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w ", "storage.Create()", err)
	}

	m.Password = ""
	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w ", "storage.GetByEmail()", err)
	}
	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return model.Users{}, fmt.Errorf("%s %w ", "storage.GetAll()", err)
	}
	return users, nil
}
