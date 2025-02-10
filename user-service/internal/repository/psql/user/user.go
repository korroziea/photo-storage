package user

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/korroziea/photo-storage/internal/domain"
)

const (
	userTable = "users"
)

var (
	userColumns                = "first_name, email, password, created_at"
	userColumnsWithoutPassword = "first_name, email, created_at"
)

type Repo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	r := &Repo{
		db: db,
	}

	return r
}

func (r *Repo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	query, args, err := sq.
		Insert(userTable).
		Columns(userColumns).
		Values(
			user.FirstName,
			user.Email,
			user.Password,
		).
		Suffix("RETURNING " + userColumnsWithoutPassword).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return domain.User{}, errors.Join(err, domain.ErrInternal)
	}
	
	return r.doQueryRow(ctx, query, args)
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	query, args, err := sq.
		Select(userColumnsWithoutPassword).
		From(userTable).
		Where(sq.Eq{
			"email": email,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return domain.User{}, errors.Join(err, domain.ErrInternal)
	}

	return r.doQueryRow(ctx, query, args...)
}

func (r *Repo) FindByEmailAndPassword(ctx context.Context, user domain.User) (domain.User, error) {
	query, args, err := sq.
		Select(userColumnsWithoutPassword).
		From(userTable).
		Where(sq.Eq{
			"email":    user.Email,
			"password": user.Password,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return domain.User{}, errors.Join(err, domain.ErrInternal)
	}

	return r.doQueryRow(ctx, query, args...)
}

func (r *Repo) doQueryRow(ctx context.Context, query string, args ...any) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(ctx, query, args).
		Scan(
			&user.FirstName,
			&user.Email,
			&user.CreatedAt,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, errors.Join(err, domain.ErrNotFound)
		}

		return domain.User{}, errors.Join(err, domain.ErrInternal)
	}

	return user, nil
}
