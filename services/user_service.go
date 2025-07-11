package services

import (
	"context"
	"errors"
	"time"

	"github.com/rizqishq/Go-REST/models"
	"github.com/rizqishq/Go-REST/repositories"
	"github.com/rizqishq/Go-REST/utils"
)

type UserService struct {
	userRepo repositories.UserRepository
}

// Create new UserService
func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]models.UserResponse, len(users))
	for i, user := range users {
		res[i] = user.ToResponse()
	}
	return res, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := user.ToResponse()
	return &result, nil
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	if u, _ := s.userRepo.FindByUsername(ctx, req.Username); u != nil {
		return nil, errors.New("username already exists")
	}
	if u, _ := s.userRepo.FindByEmail(ctx, req.Email); u != nil {
		return nil, errors.New("email already exists")
	}

	now := time.Now()
	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  utils.HashPassword(req.Password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	res := user.ToResponse()
	return &res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, req models.UpdateUserRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Username != "" && req.Username != user.Username {
		if u, _ := s.userRepo.FindByUsername(ctx, req.Username); u != nil && u.ID != id {
			return nil, errors.New("username already exists")
		}
	}
	if req.Email != "" && req.Email != user.Email {
		if u, _ := s.userRepo.FindByEmail(ctx, req.Email); u != nil && u.ID != id {
			return nil, errors.New("email already exists")
		}
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = utils.HashPassword(req.Password)
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	user.UpdatedAt = time.Now()

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}
	res := user.ToResponse()
	return &res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}
