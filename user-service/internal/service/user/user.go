package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/korroziea/photo-storage/internal/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByEmailAndPassword(ctx context.Context, user domain.User) (domain.User, error)
}

type Hasher interface {
	Generate(password string) (string, error)
	Verify(password, hash string) (bool, error)
}

type Service struct {
	repo   Repo
	hasher Hasher
}

func New(hasher Hasher, repo Repo) *Service {
	s := &Service{
		repo:   repo,
		hasher: hasher,
	}

	return s
}

func (s *Service) SignUp(ctx context.Context, user domain.User) (domain.User, error) {
	if err := s.isUserExist(ctx, user.Email); err != nil {
		return domain.User{}, fmt.Errorf("isUserExist: %w", err)
	}

	hashedPassword, err := s.hasher.Generate(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hasher.Generate: %w", err)
	}

	user.Password = hashedPassword

	return s.repo.Create(ctx, user)
}

func (s *Service) SignIn(ctx context.Context, user domain.User) (domain.User, error) {
	// TODO hashing

	return s.repo.FindByEmailAndPassword(ctx, user)
}

func (s *Service) isUserExist(ctx context.Context, email string) error {
	_, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil
		}

		return fmt.Errorf("repo.FindByEmail: %w", err)
	}

	return domain.ErrUserAlreadyExists
}
