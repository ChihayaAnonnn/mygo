package bootstrap

import (
	"errors"
	"log"

	"mygo/internal/config"
	filesystemPersistence "mygo/internal/filesystem/infrastructure/persistence"
	"mygo/internal/infra"
	knowledgePersistence "mygo/internal/knowledge/infrastructure/persistence"
	userPersistence "mygo/internal/user/infra/persistence"

	"gorm.io/gorm"
)

// MigrateConfig 迁移配置
type MigrateConfig struct {
	DryRun bool // 预览模式，不实际执行
}

// 所有需要迁移的 PO 模型
var migrateModels = []any{
	// User 模块
	&userPersistence.UserPO{},

	// Filesystem 模块
	&filesystemPersistence.ObjectPO{},

	// Knowledge 模块
	&knowledgePersistence.NodePO{},
	&knowledgePersistence.VersionPO{},
	&knowledgePersistence.ChunkPO{},
	&knowledgePersistence.EdgePO{},
	// &knowledgePersistence.EmbeddingPO{}, // TODO: 需要先安装 pgvector 扩展
	&knowledgePersistence.AITaskPO{},
	&knowledgePersistence.TagPO{},
	&knowledgePersistence.KnowledgeNodeTagPO{},
}

// errDryRunRollback 用于 dry-run 模式触发回滚
var errDryRunRollback = errors.New("dry-run: rollback")

// RunMigrate 执行数据库迁移
func RunMigrate(cfg MigrateConfig) error {
	if cfg.DryRun {
		log.Println("🔍 Dry-run mode: 将在事务中执行迁移，完成后回滚")
	}
	log.Println("🚀 Starting database migration...")

	// 1. 加载配置
	appCfg := config.Load()

	// 2. 连接数据库
	db, err := infra.NewGormPG(appCfg.Infra.PGDSN)
	if err != nil {
		return err
	}

	// 开启 Debug 模式以打印 SQL
	if cfg.DryRun {
		db = db.Debug()
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	log.Println("✅ Database connected")

	// 3. 执行迁移
	if cfg.DryRun {
		// Dry-run: 在事务中执行后回滚
		err = db.Transaction(func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(migrateModels...); err != nil {
				return err
			}
			// 迁移成功，但返回错误以触发回滚
			return errDryRunRollback
		})

		if errors.Is(err, errDryRunRollback) {
			log.Println("✅ Dry-run completed! 事务已回滚，数据库未做任何改动")
			err = nil
		}
	} else {
		err = db.AutoMigrate(migrateModels...)
	}

	if err != nil {
		return err
	}

	if !cfg.DryRun {
		log.Printf("✅ Migration completed! Migrated %d model(s)", len(migrateModels))
	}

	log.Println("📋 Registered models:")
	for _, m := range migrateModels {
		log.Printf("   - %T", m)
	}

	return nil
}

// GetMigrateModels 返回所有迁移模型（供外部使用）
func GetMigrateModels() []any {
	return migrateModels
}
