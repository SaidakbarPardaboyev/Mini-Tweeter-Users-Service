package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/static"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
	"github.com/google/uuid"
)

type users struct {
	db  *db.Postgres
	cff *config.Config
	log logger.Logger
}

func NewUsersRepo(db *db.Postgres, log logger.Logger, cff *config.Config) repo.UsersRepo {
	return &users{db: db, log: log, cff: cff}
}

func (u *users) CreateUser(ctx context.Context, in *pb.User) (*pb.CreateUserResponse, error) {
	var (
		id = uuid.New().String()

		// Insert the new user into the database
		insertQuery = static.InsertUserQuery
	)

	_, err := u.db.Db.Exec(ctx, insertQuery,
		id,
		in.Name,
		in.Surname,
		in.Bio,
		in.ProfileImage,
		in.Username,
		in.PasswordHash,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return &pb.CreateUserResponse{
		Id: id,
	}, nil
}

func (u *users) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	query := static.GetUserByIDQuery

	var (
		user                 pb.User
		createdAt, updatedAt sql.NullTime
	)

	err := u.db.Db.QueryRow(ctx, query, in.Id).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Bio,
		&user.ProfileImage,
		&user.Username,
		&user.PasswordHash,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	// Formatting timestamps
	if createdAt.Valid {
		user.CreatedAt = createdAt.Time.Format(time.RFC3339)
	}
	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &user, nil
}

func (u *users) GetUserWithLogin(ctx context.Context, in *pb.GetUserWithUsernameRequest) (*pb.User, error) {
	query := static.GetUserWithLoginQuery

	var (
		user                 pb.User
		createdAt, updatedAt sql.NullTime
	)

	err := u.db.Db.QueryRow(ctx, query, in.Username).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Bio,
		&user.ProfileImage,
		&user.Username,
		&user.PasswordHash,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	// Formatting timestamps
	if createdAt.Valid {
		user.CreatedAt = createdAt.Time.Format(time.RFC3339)
	}
	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &user, nil
}

func (u *users) GetListUser(ctx context.Context, in *pb.GetListUserRequest) (*pb.UserList, error) {
	var (
		whereClauses []string
		args         []interface{}
		argIndex     = 1

		queryBase  = static.GetListUserQuery
		queryCount = static.GetListUserQueryCount
	)

	// Full-text search on description, title, fullname, and surname
	if in.Search != nil {
		whereClauses = append(whereClauses, fmt.Sprintf(
			" (name ILIKE $%d OR surname ILIKE $%d OR username ILIKE $%d)",
			argIndex, argIndex+1, argIndex+2,
		))
		args = append(args, "%"+*in.Search+"%", "%"+*in.Search+"%", "%"+*in.Search+"%")
		argIndex += 3
	}

	// Append WHERE clause if filters exist
	if len(whereClauses) > 0 {
		queryBase += " WHERE " + strings.Join(whereClauses, " AND ")
		queryCount += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Sorting
	sortBy := "created_at"
	if in.SortBy != nil {
		sortBy = *in.SortBy
	}
	order := "DESC"
	if in.Order != nil && (*in.Order == "asc" || *in.Order == "ASC") {
		order = "ASC"
	}
	queryBase += fmt.Sprintf(" ORDER BY %s %s", sortBy, order)

	// Pagination
	queryBase += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, in.Limit, (in.Page-1)*in.Limit)

	// Count
	var count int64
	err := u.db.Db.QueryRow(ctx, queryCount, args[:len(args)-2]...).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error fetching user count: %v", err)
	}

	rows, err := u.db.Db.Query(ctx, queryBase, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching user list: %v", err)
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var (
			user                 pb.User
			createdAt, updatedAt sql.NullTime
		)

		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Surname,
			&user.Bio,
			&user.ProfileImage,
			&user.Username,
			&user.PasswordHash,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning merchant: %v", err)
		}

		if createdAt.Valid {
			user.CreatedAt = createdAt.Time.Format(time.RFC3339)
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
		}

		users = append(users, &user)
	}

	// Check for row iteration errors
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating merchants: %v", err)
	}

	return &pb.UserList{
		Users: users,
		Count: count,
	}, nil
}

func (u *users) UpdateUser(ctx context.Context, in *pb.User) error {
	// Query to update user details
	query := static.UpdateUserQuery

	_, err := u.db.Db.Exec(ctx, query,
		&in.Name,
		&in.Surname,
		&in.Bio,
		&in.ProfileImage,
		&in.Username,
		&in.PasswordHash,
		in.Id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	return nil
}

func (u *users) DeleteUser(ctx context.Context, id string) error {
	// Query to delete the user permanently
	query := static.DeleteUserQuery

	result, err := u.db.Db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
