package users

import "context"

//go:generate mockgen -destination=../mocks/mock_users/mock_repository.go -package=mock_users github.com/alanyeung95/GoProjectDemo/pkg/users Repository

// Repository is the user repo
type Repository interface {
	Upsert(ctx context.Context, id string, user User) (*User, error)
	Find(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
