package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/tvSeries-service/Middleware"

	"github.com/DonShanilka/tvSeries-service/internal/Handler"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
	"github.com/DonShanilka/tvSeries-service/internal/Routes"
	"github.com/DonShanilka/tvSeries-service/internal/Service"
	"github.com/DonShanilka/tvSeries-service/internal/db"

	episodeHandler "github.com/DonShanilka/tvSeries-service/internal/Handler"
	episodeRepo "github.com/DonShanilka/tvSeries-service/internal/Repository"
	episodeService "github.com/DonShanilka/tvSeries-service/internal/Service"
)

func main() {
	b2KeyID := "f9f45a6c989e"
	b2AppKey := "00563942506fbf1481548bd202ea51e42ec0ce19b7"
	b2BucketName := "movieStream"

	sqlDB, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to SQL DB ❌:", err)
	}

	mux := http.NewServeMux()

	tvSeriesRepo := Repository.NewTvSeriesRepository(sqlDB)
	tvSeriesService := Service.NewTvSerriesService(tvSeriesRepo)
	tvSeriesHandler := Handler.NewTvSeriesHandler(tvSeriesService)

	epRepo, err := episodeRepo.NewEpisodeRepository(sqlDB, b2KeyID, b2AppKey, b2BucketName)
	if err != nil {
		log.Fatal("Failed to create Episode Repository ❌:", err)
	}

	epService := episodeService.NewEpisodeService(epRepo)
	epHandler := episodeHandler.NewEpisodeHandler(epService)

	Routes.RegisterTvSeriesRoutes(mux, tvSeriesHandler)
	Routes.RegisterEpisodeRoutes(mux, epHandler)

	log.Println("TV Series Service running on :8081 🚀")

	err = http.ListenAndServe(":8081", Middleware.CorsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
