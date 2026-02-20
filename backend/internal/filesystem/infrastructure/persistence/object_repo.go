package persistence

import (
	"context"
	"errors"
	"time"

	"mygo/internal/filesystem/domain"
	"mygo/internal/infra"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ObjectRepository 对象元数据仓储实现（fs_objects）。
type ObjectRepository struct {
	db *infra.GormDB
}

func NewObjectRepository(res *infra.Resources) (*ObjectRepository, error) {
	if res == nil {
		return nil, errors.New("fs object repo: resources is nil")
	}
	if res.DB == nil {
		return nil, errors.New("fs object repo: resources db is nil")
	}
	return &ObjectRepository{db: res.DB}, nil
}

func (r *ObjectRepository) UpsertMeta(ctx context.Context, meta *domain.ObjectMeta) error {
	if r.db == nil {
		return errors.New("fs object repo: db is nil")
	}
	if meta == nil {
		return errors.New("fs object repo: meta is nil")
	}
	if meta.Namespace == "" || meta.Key == "" {
		return domain.ErrInvalidInput
	}

	now := time.Now()
	po := &ObjectPO{
		Namespace:   meta.Namespace.String(),
		Key:         meta.Key.String(),
		Backend:     meta.Backend,
		ContentType: meta.ContentType,
		SizeBytes:   meta.SizeBytes,
		ETag:        meta.ETag,
		Checksum:    meta.Checksum,
		Metadata:    meta.Metadata,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Upsert 且清空 deleted_at（写入即复活）。
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "namespace"},
			{Name: "key"},
		},
		DoUpdates: clause.Assignments(map[string]any{
			"backend":      po.Backend,
			"content_type": po.ContentType,
			"size_bytes":   po.SizeBytes,
			"etag":         po.ETag,
			"checksum":     po.Checksum,
			"metadata":     po.Metadata,
			"updated_at":   now,
			"deleted_at":   nil,
		}),
	}).Create(po).Error
}

func (r *ObjectRepository) Head(ctx context.Context, namespace domain.Namespace, key domain.Key) (*domain.ObjectMeta, error) {
	if namespace == "" || key == "" {
		return nil, domain.ErrInvalidInput
	}

	var po ObjectPO
	if err := r.db.WithContext(ctx).
		Where("namespace = ? AND key = ?", namespace.String(), key.String()).
		First(&po).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}

	return objectMetaFromPO(&po), nil
}

func (r *ObjectRepository) ListByPrefix(ctx context.Context, namespace domain.Namespace, prefix string, limit, offset int) ([]*domain.ObjectMeta, error) {
	if namespace == "" {
		return nil, domain.ErrInvalidInput
	}
	if prefix != "" {
		if err := validateKeyPrefix(prefix); err != nil {
			return nil, err
		}
	}

	db := r.db.WithContext(ctx).Model(&ObjectPO{}).Where("namespace = ?", namespace.String())
	if prefix != "" {
		db = db.Where(`key LIKE ?`, prefix+"%")
	}

	if limit <= 0 {
		return nil, domain.ErrInvalidInput
	}
	db = db.Limit(limit)

	if offset < 0 {
		return nil, domain.ErrInvalidInput
	}
	if offset > 0 {
		db = db.Offset(offset)
	}
	db = db.Order("updated_at DESC")

	var pos []ObjectPO
	if err := db.Find(&pos).Error; err != nil {
		return nil, err
	}

	out := make([]*domain.ObjectMeta, 0, len(pos))
	for i := range pos {
		out = append(out, objectMetaFromPO(&pos[i]))
	}
	return out, nil
}

// SoftDelete 软删除对象元数据。
//
// ObjectPO 含 gorm.DeletedAt，GORM 会将 Delete() 转换为
// UPDATE fs_objects SET deleted_at = NOW() WHERE ...
// 而非物理删除。payload 文件不受影响，对象可通过 Undelete 恢复。
func (r *ObjectRepository) SoftDelete(ctx context.Context, namespace domain.Namespace, key domain.Key) error {
	if namespace == "" || key == "" {
		return domain.ErrInvalidInput
	}

	result := r.db.WithContext(ctx).
		Where("namespace = ? AND key = ?", namespace.String(), key.String()).
		Delete(&ObjectPO{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrObjectNotFound
	}
	return nil
}

func (r *ObjectRepository) Undelete(ctx context.Context, namespace domain.Namespace, key domain.Key) error {
	if namespace == "" || key == "" {
		return domain.ErrInvalidInput
	}

	now := time.Now()
	result := r.db.WithContext(ctx).
		Unscoped().
		Model(&ObjectPO{}).
		Where("namespace = ? AND key = ?", namespace.String(), key.String()).
		Updates(map[string]any{
			"deleted_at": nil,
			"updated_at": now,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrObjectNotFound
	}
	return nil
}

func objectMetaFromPO(p *ObjectPO) *domain.ObjectMeta {
	if p == nil {
		return nil
	}
	ns, _ := domain.ParseNamespace(p.Namespace)
	k, _ := domain.ParseKey(p.Key)

	var deletedAt *time.Time
	if p.DeletedAt.Valid {
		t := p.DeletedAt.Time
		deletedAt = &t
	}

	return &domain.ObjectMeta{
		Namespace:   ns,
		Key:         k,
		Backend:     p.Backend,
		ContentType: p.ContentType,
		SizeBytes:   p.SizeBytes,
		ETag:        p.ETag,
		Checksum:    p.Checksum,
		Metadata:    p.Metadata,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func validateKeyPrefix(prefix string) error {
	// prefix 允许以 "/" 结尾（用于模拟目录前缀），但其他规则与 Key 一致。
	if prefix == "" {
		return nil
	}
	if prefix[0] == '/' {
		return domain.ErrInvalidInput
	}
	if prefix[len(prefix)-1] == '/' {
		prefix = prefix[:len(prefix)-1]
	}
	if prefix == "" {
		return nil
	}
	_, err := domain.ParseKey(prefix)
	if err != nil {
		return domain.ErrInvalidInput
	}
	return nil
}

var _ domain.ObjectRepository = (*ObjectRepository)(nil)
