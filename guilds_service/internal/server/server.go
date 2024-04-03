package server

import (
	"context"
	"errors"
	"github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1/channelsv1connect"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/gorm"

	"github.com/X3ne/ds_ms/api/gen/guilds_service/guilds/v1/guildsv1connect"
	"github.com/X3ne/ds_ms/api/gen/users_service/users/v1/usersv1connect"
	"github.com/X3ne/ds_ms/guilds_service/config"
	"github.com/X3ne/ds_ms/guilds_service/internal/handlers"
	"github.com/X3ne/ds_ms/guilds_service/internal/interceptors"
	"github.com/X3ne/ds_ms/guilds_service/internal/repositories"
)

type Server struct{}

func LaunchServer(cfg *config.Config, db *gorm.DB) {
	errorsInterceptor := connect.WithInterceptors(interceptors.NewErrorInterceptor())
	api := http.NewServeMux()

	api.Handle(guildsv1connect.NewGuildsServiceHandler(&handlers.GuildsServer{
		Repository:     repositories.NewGuildRepository(db),
		UsersClient:    usersv1connect.NewUsersServiceClient(http.DefaultClient, "http://127.0.0.1:8080"),
		ChannelsClient: channelsv1connect.NewChannelsServiceClient(http.DefaultClient, "http://127.0.0.1:8082"),
	}, errorsInterceptor))

	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(
		"guilds.v1.GuildsService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle("/", api)

	srv := &http.Server{
		Addr:              cfg.SERVER.Host + ":" + cfg.SERVER.Port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024,
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	log.Printf("Starting HTTP server on %s", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server: %v", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err)
	}
}
