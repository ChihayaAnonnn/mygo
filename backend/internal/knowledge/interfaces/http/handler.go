package http

import (
	"net/http"
	"time"

	"mygo/internal/knowledge/application"
	"mygo/internal/knowledge/domain"

	"github.com/gin-gonic/gin"
)

// Handler 知识 HTTP 处理器
type Handler struct {
	knowledgeSvc  domain.KnowledgeService
	versionSvc    domain.KnowledgeVersionService
	appSvc        application.KnowledgeApplicationService
	chunkRepo     domain.ChunkRepository
	embeddingRepo domain.EmbeddingRepository
}

// NewHandler 构造函数
func NewHandler(
	knowledgeSvc domain.KnowledgeService,
	versionSvc domain.KnowledgeVersionService,
	appSvc application.KnowledgeApplicationService,
	chunkRepo domain.ChunkRepository,
	embeddingRepo domain.EmbeddingRepository,
) *Handler {
	return &Handler{
		knowledgeSvc:  knowledgeSvc,
		versionSvc:    versionSvc,
		appSvc:        appSvc,
		chunkRepo:     chunkRepo,
		embeddingRepo: embeddingRepo,
	}
}

// Response 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 响应辅助函数
func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func fail(c *gin.Context, httpCode int, code int, message string) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
	})
}

// ==================== Knowledge Handlers ====================

// CreateKnowledge 创建知识
// POST /api/knowledge
func (h *Handler) CreateKnowledge(c *gin.Context) {
	var req CreateKnowledgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	cmd := domain.CreateKnowledgeCmd{
		NodeType: domain.NodeType(req.NodeType),
		Title:    req.Title,
		Summary:  req.Summary,
	}

	id, err := h.knowledgeSvc.CreateKnowledge(c.Request.Context(), cmd)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, map[string]string{"id": string(id)})
}

// GetKnowledge 获取知识
// GET /api/knowledge/:id
func (h *Handler) GetKnowledge(c *gin.Context) {
	id := domain.KnowledgeID(c.Param("id"))

	knowledge, err := h.knowledgeSvc.GetKnowledge(c.Request.Context(), id)
	if err != nil {
		if err == domain.ErrKnowledgeNotFound {
			fail(c, http.StatusNotFound, 404, "knowledge not found")
		} else {
			fail(c, http.StatusInternalServerError, 500, "internal server error")
		}
		return
	}

	success(c, knowledgeToResponse(knowledge))
}

// UpdateKnowledgeMeta 更新知识元信息
// PUT /api/knowledge/:id
func (h *Handler) UpdateKnowledgeMeta(c *gin.Context) {
	id := domain.KnowledgeID(c.Param("id"))

	var req UpdateKnowledgeMetaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	cmd := domain.UpdateKnowledgeMetaCmd{
		ID:         id,
		Title:      req.Title,
		Summary:    req.Summary,
		Status:     (*domain.NodeStatus)(req.Status),
		Confidence: req.Confidence,
	}

	if err := h.knowledgeSvc.UpdateKnowledgeMeta(c.Request.Context(), cmd); err != nil {
		if err == domain.ErrKnowledgeNotFound {
			fail(c, http.StatusNotFound, 404, "knowledge not found")
		} else {
			fail(c, http.StatusInternalServerError, 500, "internal server error")
		}
		return
	}

	success(c, nil)
}

// ListKnowledge 列出知识
// GET /api/knowledge
func (h *Handler) ListKnowledge(c *gin.Context) {
	var req ListKnowledgeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid query params")
		return
	}

	query := domain.KnowledgeQuery{
		Offset: req.Offset,
		Limit:  req.Limit,
	}
	if req.NodeType != "" {
		nt := domain.NodeType(req.NodeType)
		query.NodeType = &nt
	}
	if req.Status != "" {
		st := domain.NodeStatus(req.Status)
		query.Status = &st
	}

	list, err := h.knowledgeSvc.ListKnowledge(c.Request.Context(), query)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	items := make([]*KnowledgeResponse, 0, len(list))
	for _, k := range list {
		items = append(items, knowledgeToResponse(k))
	}
	success(c, items)
}

// ArchiveKnowledge 归档知识
// POST /api/knowledge/:id/archive
func (h *Handler) ArchiveKnowledge(c *gin.Context) {
	id := domain.KnowledgeID(c.Param("id"))

	if err := h.knowledgeSvc.ArchiveKnowledge(c.Request.Context(), id); err != nil {
		if err == domain.ErrKnowledgeNotFound {
			fail(c, http.StatusNotFound, 404, "knowledge not found")
		} else {
			fail(c, http.StatusInternalServerError, 500, "internal server error")
		}
		return
	}

	success(c, nil)
}

// ==================== Version Handlers ====================

// CreateVersion 创建版本
// POST /api/knowledge/:id/versions
func (h *Handler) CreateVersion(c *gin.Context) {
	knowledgeID := domain.KnowledgeID(c.Param("id"))

	var req CreateVersionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	cmd := domain.CreateVersionCmd{
		KnowledgeID: knowledgeID,
		ContentMd:   req.ContentMd,
	}

	versionID, err := h.versionSvc.CreateVersion(c.Request.Context(), cmd)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, map[string]string{"id": string(versionID)})
}

// GetLatestVersion 获取最新版本
// GET /api/knowledge/:id/versions/latest
func (h *Handler) GetLatestVersion(c *gin.Context) {
	knowledgeID := domain.KnowledgeID(c.Param("id"))

	version, err := h.versionSvc.GetLatestVersion(c.Request.Context(), knowledgeID)
	if err != nil {
		if err == domain.ErrVersionNotFound {
			fail(c, http.StatusNotFound, 404, "version not found")
		} else {
			fail(c, http.StatusInternalServerError, 500, "internal server error")
		}
		return
	}

	success(c, versionToResponse(version))
}

// ListVersions 列出所有版本
// GET /api/knowledge/:id/versions
func (h *Handler) ListVersions(c *gin.Context) {
	knowledgeID := domain.KnowledgeID(c.Param("id"))

	versions, err := h.versionSvc.ListVersions(c.Request.Context(), knowledgeID)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	items := make([]*VersionResponse, 0, len(versions))
	for _, v := range versions {
		items = append(items, versionToResponse(v))
	}
	success(c, items)
}

// ==================== Chunk Handlers ====================

// BatchCreateChunks 批量创建分块（Agent 端预切分后写入）
// POST /api/knowledge/:id/chunks
func (h *Handler) BatchCreateChunks(c *gin.Context) {
	nodeID := domain.KnowledgeID(c.Param("id"))

	var req BatchCreateChunksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	now := time.Now()
	chunks := make([]*domain.Chunk, 0, len(req.Chunks))
	for _, item := range req.Chunks {
		chunks = append(chunks, &domain.Chunk{
			ChunkID:     item.ChunkID,
			NodeID:      string(nodeID),
			Version:     req.Version,
			HeadingPath: item.HeadingPath,
			Content:     item.Content,
			TokenCount:  item.TokenCount,
			ChunkIndex:  item.ChunkIndex,
			CreatedAt:   now,
		})
	}

	if err := h.chunkRepo.BatchCreate(c.Request.Context(), chunks); err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, map[string]int{"count": len(chunks)})
}

// ListChunks 列出分块
// GET /api/knowledge/:id/chunks?version=1
func (h *Handler) ListChunks(c *gin.Context) {
	nodeID := domain.KnowledgeID(c.Param("id"))

	var query struct {
		Version int `form:"version" binding:"required"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		fail(c, http.StatusBadRequest, 400, "version is required")
		return
	}

	chunks, err := h.chunkRepo.ListByNodeVersion(c.Request.Context(), nodeID, query.Version)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	items := make([]*ChunkResponse, 0, len(chunks))
	for _, chunk := range chunks {
		items = append(items, chunkToResponse(chunk))
	}
	success(c, items)
}

// DeleteChunks 删除分块
// DELETE /api/knowledge/:id/chunks
func (h *Handler) DeleteChunks(c *gin.Context) {
	nodeID := domain.KnowledgeID(c.Param("id"))

	var req DeleteChunksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	if err := h.chunkRepo.DeleteByNodeVersion(c.Request.Context(), nodeID, req.Version); err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, nil)
}

// ==================== Embedding Handlers ====================

// BatchCreateEmbeddings 批量创建向量（Agent 端预计算后写入）
// POST /api/knowledge/embeddings
func (h *Handler) BatchCreateEmbeddings(c *gin.Context) {
	var req BatchCreateEmbeddingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	now := time.Now()
	embeddings := make([]*domain.Embedding, 0, len(req.Embeddings))
	for _, item := range req.Embeddings {
		embeddings = append(embeddings, &domain.Embedding{
			ChunkID:   item.ChunkID,
			Embedding: item.Embedding,
			Model:     item.Model,
			CreatedAt: now,
		})
	}

	if err := h.embeddingRepo.BatchCreate(c.Request.Context(), embeddings); err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, map[string]int{"count": len(embeddings)})
}

// ==================== Search Handlers ====================

// Search 向量搜索（Agent 端传入预计算的 query 向量）
// POST /api/knowledge/search
func (h *Handler) Search(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, 400, "invalid request body")
		return
	}

	topK := req.TopK
	if topK <= 0 {
		topK = 10
	}

	// 向量相似度搜索
	chunkIDs, err := h.embeddingRepo.SearchSimilar(c.Request.Context(), domain.EmbeddingVector(req.Vector), topK)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	if len(chunkIDs) == 0 {
		success(c, []*ChunkResponse{})
		return
	}

	// 获取 Chunk 详情
	type chunkByIDsGetter interface {
		GetByIDs(ctx interface{}, chunkIDs []domain.ChunkID) ([]*domain.Chunk, error)
	}

	getter, ok := h.chunkRepo.(chunkByIDsGetter)
	if !ok {
		fail(c, http.StatusInternalServerError, 500, "chunk repo does not support GetByIDs")
		return
	}

	chunks, err := getter.GetByIDs(c.Request.Context(), chunkIDs)
	if err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	items := make([]*ChunkResponse, 0, len(chunks))
	for _, chunk := range chunks {
		items = append(items, chunkToResponse(chunk))
	}
	success(c, items)
}

// ==================== Application Handlers ====================

// PublishKnowledge 发布知识
// POST /api/knowledge/:id/publish
func (h *Handler) PublishKnowledge(c *gin.Context) {
	id := domain.KnowledgeID(c.Param("id"))

	if err := h.appSvc.PublishKnowledge(c.Request.Context(), id); err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, nil)
}

// RebuildIndex 重建索引（仅删除旧数据，Agent 负责重新生成）
// POST /api/knowledge/:id/rebuild-index
func (h *Handler) RebuildIndex(c *gin.Context) {
	id := domain.KnowledgeID(c.Param("id"))

	if err := h.appSvc.RebuildIndex(c.Request.Context(), id); err != nil {
		fail(c, http.StatusInternalServerError, 500, "internal server error")
		return
	}

	success(c, nil)
}

// ==================== Helper Functions ====================

func knowledgeToResponse(node *domain.Node) *KnowledgeResponse {
	if node == nil {
		return nil
	}
	return &KnowledgeResponse{
		ID:             node.NodeID,
		NodeType:       string(node.NodeType),
		Title:          node.Title,
		Summary:        node.Summary,
		Status:         string(node.Status),
		Confidence:     node.Confidence,
		CurrentVersion: node.CurrentVersion,
		CreatedAt:      node.CreatedAt,
		UpdatedAt:      node.UpdatedAt,
	}
}

func versionToResponse(version *domain.Version) *VersionResponse {
	if version == nil {
		return nil
	}
	return &VersionResponse{
		ID:        version.VersionID,
		NodeID:    version.NodeID,
		Version:   version.Version,
		ContentMd: version.ContentMd,
		CreatedAt: version.CreatedAt,
	}
}

func chunkToResponse(chunk *domain.Chunk) *ChunkResponse {
	if chunk == nil {
		return nil
	}
	return &ChunkResponse{
		ID:          chunk.ChunkID,
		NodeID:      chunk.NodeID,
		Version:     chunk.Version,
		HeadingPath: chunk.HeadingPath,
		Content:     chunk.Content,
		TokenCount:  chunk.TokenCount,
		ChunkIndex:  chunk.ChunkIndex,
		CreatedAt:   chunk.CreatedAt,
	}
}
