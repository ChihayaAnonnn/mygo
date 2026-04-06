package domain

import "time"

// PickKind 表示收藏条目的内容类型
type PickKind string

const (
	PickKindWebsite PickKind = "website"
	PickKindArticle PickKind = "article"
	PickKindEssay   PickKind = "essay"
	PickKindTool    PickKind = "tool"
	PickKindVideo   PickKind = "video"
)

// PickStatus 表示收藏条目的发布状态
type PickStatus string

const (
	PickStatusDraft     PickStatus = "draft"
	PickStatusPublished PickStatus = "published"
	PickStatusArchived  PickStatus = "archived"
)

// Pick 收藏条目领域模型
type Pick struct {
	ID    int64
	Title string
	URL   string

	Kind         PickKind
	Category     string
	CollectionID *int64
	Source       string

	Note        string
	RevisitHint string

	IsFeatured bool
	SortOrder  int
	Status     PickStatus

	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt *time.Time
}
