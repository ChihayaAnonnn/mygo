package main

import (
	"flag"
	"log"

	"mygo/internal/bootstrap"
)

var dryRun = flag.Bool("dry-run", false, "预览迁移变更，不实际执行（使用事务回滚）")

func main() {
	flag.Parse()

	cfg := bootstrap.MigrateConfig{
		DryRun: *dryRun,
	}

	if err := bootstrap.RunMigrate(cfg); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
}
