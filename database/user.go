package database

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Name string    `json:"name"`
}

func (u *User) Save(ctx context.Context) (*User, error) {
	u.ID = uuid.New()
	result := DB.WithContext(ctx).Create(&u)

	if result.Error != nil {
		return &User{}, result.Error
	}
	return u, nil
}
func (u *User) GetAllUsers() (*[]User, error) {
	var users []User

	result := DB.Model(&u).Find(&users)
	if result.Error != nil {
		return &[]User{}, result.Error
	}

	return &users, nil
}
func (u *User) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	result := DB.WithContext(ctx).Model(&User{}).Where("id = ?", id).Take(&u)

	if result.Error != nil {
		return &User{}, result.Error
	}

	return u, nil
}
