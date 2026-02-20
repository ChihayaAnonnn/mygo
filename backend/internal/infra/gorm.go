package infra

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// GormDB 作为 infra 暴露给上层的 ORM 依赖类型别名，
// repo 可以只依赖 infra.Resources，而不是直接依赖 gorm 包路径。
type GormDB = gorm.DB

// ErrRecordNotFound 透出给上层做 errors.Is 判断。
var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewGormPG(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("failed to open gorm: dsn is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DryRun: true, // 开启时打印 SQL 但不执行
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm: %w", err)
	}

	// 连接池配置（与 pgxpool 的默认思路保持一致）
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get gorm sql db: %w", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	return db, nil
}
