package bootstrap

import (
	"log"

	"mygo/internal/config"
	"mygo/internal/infra"
	knowledgeApp "mygo/internal/knowledge/application"
	knowledgePersistence "mygo/internal/knowledge/infrastructure/persistence"
	knowledgeHttp "mygo/internal/knowledge/interfaces/http"
	"mygo/internal/server"
	userApp "mygo/internal/user/application"
	userCache "mygo/internal/user/infra/cache"
	userPersistence "mygo/internal/user/infra/persistence"
	userHttp "mygo/internal/user/interfaces/http"

	"github.com/gin-gonic/gin"
)

// App 应用容器，持有所有依赖
type App struct {
	Config    *config.Config
	Resources *infra.Resources

	// User 模块
	UserHandler *userHttp.Handler

	// Knowledge 模块
	KnowledgeHandler *knowledgeHttp.Handler
}

// NewApp 创建并初始化应用
func NewApp() (*App, error) {
	app := &App{}

	// 1. 加载配置
	app.Config = config.Load()
	log.Printf("Config loaded: port=%s, mode=%s", app.Config.Server.Port, app.Config.Server.Mode)

	// 设置 gin 模式
	gin.SetMode(app.Config.Server.Mode)

	// 2. 初始化基础设施
	resources, err := infra.NewResources(app.Config.Infra)
	if err != nil {
		return nil, err
	}
	app.Resources = resources
	log.Println("Infrastructure initialized")

	// 3. 初始化各模块
	if err := app.initUserModule(); err != nil {
		return nil, err
	}

	if err := app.initKnowledgeModule(); err != nil {
		return nil, err
	}

	return app, nil
}

// initUserModule 初始化 User 模块
func (app *App) initUserModule() error {
	// Repository
	userRepo, err := userPersistence.NewUserRepository(app.Resources)
	if err != nil {
		return err
	}

	sessionCache, err := userCache.NewSessionCache(app.Resources)
	if err != nil {
		return err
	}

	// Application Service
	userAppService := userApp.NewAppService(userRepo, sessionCache)

	// HTTP Handler
	app.UserHandler = userHttp.NewHandler(userAppService)

	log.Println("User module initialized")
	return nil
}

// initKnowledgeModule 初始化 Knowledge 模块
func (app *App) initKnowledgeModule() error {
	// Repository
	nodeRepo, err := knowledgePersistence.NewNodeRepository(app.Resources)
	if err != nil {
		return err
	}

	versionRepo, err := knowledgePersistence.NewVersionRepository(app.Resources)
	if err != nil {
		return err
	}

	// Domain Services
	knowledgeSvc := knowledgeApp.NewKnowledgeService(nodeRepo)
	versionSvc := knowledgeApp.NewVersionService(versionRepo, nodeRepo)

	// Application Service (P1 依赖暂传 nil，后续补全)
	appSvc := knowledgeApp.NewAppService(
		knowledgeSvc,
		versionSvc,
		nil, // renderSvc - P2
		nil, // chunkSvc - P1
		nil, // embeddingSvc - P1
		nil, // retrievalSvc - P1
		nodeRepo,
		versionRepo,
		nil, // chunkRepo - P1
		nil, // embeddingRepo - P1
	)

	// HTTP Handler
	app.KnowledgeHandler = knowledgeHttp.NewHandler(
		knowledgeSvc,
		versionSvc,
		nil, // chunkSvc - P1
		nil, // retrievalSvc - P1
		appSvc,
	)

	log.Println("Knowledge module initialized")
	return nil
}

// RouterConfig 返回路由配置
func (app *App) RouterConfig() server.RouterConfig {
	return server.RouterConfig{
		UserHandler:      app.UserHandler,
		KnowledgeHandler: app.KnowledgeHandler,
	}
}

// Close 关闭应用资源
func (app *App) Close() error {
	if app.Resources != nil {
		return app.Resources.Close()
	}
	return nil
}
