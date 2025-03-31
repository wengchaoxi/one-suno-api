package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/wengchaoxi/one-suno-api/internal/model"
	"github.com/wengchaoxi/one-suno-api/internal/provider"
	"github.com/wengchaoxi/one-suno-api/internal/provider/acedata"
	"github.com/wengchaoxi/one-suno-api/internal/repository"
	"github.com/wengchaoxi/one-suno-api/internal/service"
	"github.com/wengchaoxi/one-suno-api/pkg/logger"
)

var (
	HOST    = getEnv("HOST", "0.0.0.0")
	PORT    = getEnv("PORT", "8080")
	API_KEY = getEnv("API_KEY", "")
)

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func main() {
	logger.Init()

	db, dberr := gorm.Open(sqlite.Open(getEnv("DB_PATH", "one-suno-api.db")), &gorm.Config{})
	if dberr != nil {
		logger.Panicf("failed to connect database: %v", dberr)
	}
	if err := db.AutoMigrate(&model.ApiKey{}, &model.User{}); err != nil {
		logger.Panicf("failed to migrate database: %v", err)
	}
	userRepo := repository.NewUserRepository(db)
	apikeyRepo := repository.NewApiKeyRepository(db)

	providerManager := provider.ProviderManager{
		Balancer: provider.NewWeightedRoundRobinBalancer(),
	}
	providerManager.Register(provider.ProviderMeta{
		Id:     "acedata",
		Weight: 6,
		Provider: acedata.NewAcedataProvider(&acedata.AcedataProviderOptions{
			BaseUrl:       getEnv("ACEDATA_BASE_URL", "https://api.acedata.cloud"),
			PlatformToken: os.Getenv("ACEDATA_PLATFORM_TOKEN"),
			AppId:         os.Getenv("ACEDATA_APP_ID"),
			AppToken:      os.Getenv("ACEDATA_APP_TOKEN"),
		}),
	})

	app := service.New(&service.ServiceOptions{
		ProviderManager:  providerManager,
		ApiKey:           API_KEY,
		UserRepository:   userRepo,
		ApiKeyRepository: apikeyRepo,
	})
	engine := gin.Default()
	app.Init(engine)

	var err error
	var wg sync.WaitGroup
	addr := net.JoinHostPort(HOST, PORT)
	srv := &http.Server{Addr: addr, Handler: engine}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Panicf("server listen error", err)
		}
	}()
	logger.Infof("Running on http://%s\n", addr)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 360*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Infof("server shutdown err: %v\n", err)
	}
	wg.Wait()
}
