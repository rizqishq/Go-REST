package repositories

import (
	"context"
	"errors"
	"sync"

	"github.com/rizqishq/Go-REST/models"
)

// Respository errors
var (
	ErrNotFound = errors.New("record not found")
	ErrConflict = errors.New("record already exists")
)

// UserRepository interface to abstract storage implementation
type UserRepository interface {
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
}

// InMemoryUserRepository implements UserRepository in memory
type InMemoryUserRepository struct {
	users  map[uint]*models.User
	mutex  sync.RWMutex
	nextID uint
}

// Create new empty repository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[uint]*models.User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, ErrNotFound
	}
	return user, nil
}

func (r *InMemoryUserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, u := range r.users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, ErrNotFound
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrNotFound
}

func (r *InMemoryUserRepository) Create(ctx context.Context, user *models.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, existing := range r.users {
		if existing.Username == user.Username || existing.Email == user.Email {
			return ErrConflict
		}
	}

	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Update(ctx context.Context, user *models.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return ErrNotFound
	}

	for id, existing := range r.users {
		if id == user.ID {
			continue
		}
		if existing.Username == user.Username || existing.Email == user.Email {
			return ErrConflict
		}
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return ErrNotFound
	}
	delete(r.users, id)
	return nil
}
