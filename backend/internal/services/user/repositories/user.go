package repositories

import (
	"errors"
	"travel-bot/internal/entities"
	"travel-bot/internal/services/user/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func GetUserRepository(db *gorm.DB) service.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (rep userRepository) checkUserExists(id uuid.UUID) (bool, error) {
	user := UserSchema{
		Id: id,
	}
	res := rep.db.First(&user)

	switch {
	case errors.Is(res.Error, gorm.ErrRecordNotFound):
		return false, nil
	case res.Error != nil:
		return false, res.Error
	default:
		return true, nil
	}
}


func (rep userRepository) AddUser(user *entities.User) error {
	userSchema := toSchema(user)
	res := rep.db.Create(userSchema)

	return res.Error
}

func (rep userRepository) UpdateUser(id uuid.UUID, user *entities.User) error {
	isExists, err := rep.checkUserExists(id)
	if !isExists {
		return errors.New("user not exists")
	}
	if err != nil {
		return err
	}

	res := rep.db.Save(user)
	
	return res.Error
}

func (rep userRepository) RemoveUser(id uuid.UUID) error {
	isExists, err := rep.checkUserExists(id)
	if !isExists {
		return errors.New("user not exists")
	}
	if err != nil {
		return err
	}

	res := rep.db.Delete(&UserSchema{Id: id})

	return res.Error
}

func (rep userRepository) GetUserById(id uuid.UUID) (*entities.User, error) {
	isExists, err := rep.checkUserExists(id)
	if !isExists {
		return nil, errors.New("user not exists")
	}
	if err != nil {
		return nil, err
	}

	userSchema := UserSchema{Id: id}
	res := rep.db.First(&userSchema)

	if err := res.Error; err != nil {
		return nil, err
	}
	
	user := ToEntity(&userSchema)

	return user, nil
}