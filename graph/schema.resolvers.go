package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	pb "graphqhhowto/client"
	"graphqhhowto/database"
	"graphqhhowto/graph/model"

	"github.com/google/uuid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var newUser database.User
	newUser.Name = input.Name
	userGRPC := pb.Cl.CreateUserInDb(&newUser)

	return &model.User{
		ID:   userGRPC.Id,
		Name: userGRPC.Name,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	uuID, _ := uuid.Parse(input.ID)

	var user database.User
	userUpdate, err := user.UpdateUserById(ctx, uuID, input.Name)
	if err != nil {
		return nil, err
	}
	userUpdate, _ = user.GetByID(ctx, uuID)

	return &model.User{
		ID:   userUpdate.ID.String(),
		Name: userUpdate.Name,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, input *model.DeleteUser) (bool, error) {
	id, _ := uuid.Parse(input.ID)
	var user database.User
	res, err := user.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	if res != 0 {
		return false, err
	}

	return true, nil
}

// CreateCar is the resolver for the createCar field.
func (r *mutationResolver) CreateCar(ctx context.Context, input *model.NewCar) (*model.Cars, error) {
	panic(fmt.Errorf("not implemented: CreateCar - createCar"))
}

// UpdateCar is the resolver for the updateCar field.
func (r *mutationResolver) UpdateCar(ctx context.Context, id string, carName string, model *string) (*model.Cars, error) {
	panic(fmt.Errorf("not implemented: UpdateCar - updateCar"))
}

// DeleteCar is the resolver for the deleteCar field.
func (r *mutationResolver) DeleteCar(ctx context.Context, id string) (*model.Cars, error) {
	panic(fmt.Errorf("not implemented: DeleteCar - deleteCar"))
}

// AllUsers is the resolver for the allUsers field.
func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	usersgRPC := pb.Cl.ListAllUsers()

	for _, u := range usersgRPC {
		users = append(users, &model.User{
			ID:   u.Id,
			Name: u.Name,
		})
	}
	return users, nil
}

// GetUserByID is the resolver for the getUserByID field.
func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	idUUID, _ := uuid.Parse(id)
	userModel := database.User{}
	userDb, _ := userModel.GetByID(ctx, idUUID)

	return &model.User{
		ID:   userDb.ID.String(),
		Name: userDb.Name,
	}, nil
}

// AllCars is the resolver for the allCars field.
func (r *queryResolver) AllCars(ctx context.Context, last *int) (*model.Cars, error) {
	panic(fmt.Errorf("not implemented: AllCars - allCars"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
