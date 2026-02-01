package persistence

import (
	"context"
	"errors"

	"mygo/internal/infra"
	"mygo/internal/user/domain"

	"gorm.io/gorm/clause"
)

// UserRepository 用户仓储实现
type UserRepository struct {
	db *infra.GormDB
}

// NewUserRepository 构造函数
func NewUserRepository(res *infra.Resources) (*UserRepository, error) {
	if res == nil {
		return nil, errors.New("user repo: resources is nil")
	}
	if res.DB == nil {
		return nil, errors.New("user repo: resources db is nil")
	}
	return &UserRepository{db: res.DB}, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	if r.db == nil {
		return errors.New("user repo: db is nil")
	}
	if user == nil {
		return errors.New("user repo: user is nil")
	}
	if user.UserID == 0 {
		return errors.New("user repo: user_id is required")
	}

	p := FromDomain(user)
	if err := r.db.WithContext(ctx).Create(p).Error; err != nil {
		return err
	}
	*user = *p.ToDomain()
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	if r.db == nil {
		return nil, errors.New("user repo: db is nil")
	}

	var p UserPO
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		if errors.Is(err, infra.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return p.ToDomain(), nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	if r.db == nil {
		return nil, errors.New("user repo: db is nil")
	}

	var p UserPO
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&p).Error; err != nil {
		if errors.Is(err, infra.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return p.ToDomain(), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	if r.db == nil {
		return nil, errors.New("user repo: db is nil")
	}

	var p UserPO
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&p).Error; err != nil {
		if errors.Is(err, infra.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return p.ToDomain(), nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	if r.db == nil {
		return errors.New("user repo: db is nil")
	}
	if user == nil {
		return errors.New("user repo: user is nil")
	}
	if user.ID == 0 {
		return errors.New("user repo: user id is required")
	}
	if user.UserID == 0 {
		return errors.New("user repo: user_id is required")
	}
	if user.Username == "" {
		return errors.New("user repo: username is required")
	}
	if user.Email == "" {
		return errors.New("user repo: email is required")
	}
	if user.Password == "" {
		return errors.New("user repo: password is required")
	}

	updates := map[string]any{
		"user_id":  user.UserID,
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"avatar":   user.Avatar,
	}

	// 使用 Postgres RETURNING，一次往返拿到更新后的行（含 updated_at）
	p := UserPO{ID: user.ID}
	tx := r.db.WithContext(ctx).
		Model(&p).
		Clauses(clause.Returning{}).
		Updates(updates)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	*user = *p.ToDomain()
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if r.db == nil {
		return errors.New("user repo: db is nil")
	}
	if id == 0 {
		return errors.New("user repo: id is required")
	}

	tx := r.db.WithContext(ctx).Delete(&UserPO{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

// 确保 UserRepository 实现了 domain.UserRepository 接口
var _ domain.UserRepository = (*UserRepository)(nil)
