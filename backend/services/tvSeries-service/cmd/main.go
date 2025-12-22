package main

import (
	"log"
	"net/http"

	"github.com/DonShanilka/tvSeries-service/Middleware"

	// TV Series
	"github.com/DonShanilka/tvSeries-service/internal/Handler"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
	"github.com/DonShanilka/tvSeries-service/internal/Routes"
	"github.com/DonShanilka/tvSeries-service/internal/Service"
	"github.com/DonShanilka/tvSeries-service/internal/db"

	// Episodes (Backblaze B2)
	episodeHandler "github.com/DonShanilka/tvSeries-service/internal/Handler"
	episodeRepo "github.com/DonShanilka/tvSeries-service/internal/Repository"
	episodeService "github.com/DonShanilka/tvSeries-service/internal/Service"
)

func main() {

	// ------------------ EPISODES ------------------
	// Backblaze B2 Config
	b2KeyID := "f9f45a6c989e"
	b2AppKey := "00563942506fbf1481548bd202ea51e42ec0ce19b7"
	b2BucketName := "movieStream"

	// ------------------ DATABASE ------------------
	sqlDB, err := db.InitDB() // MySQL/Postgres for Episodes
	if err != nil {
		log.Fatal("Failed to connect to SQL DB ❌:", err)
	}

	// ------------------ ROUTES ------------------
	mux := http.NewServeMux()

	// ------------------ TV SERIES ------------------
	tvSeriesRepo := Repository.NewTvSeriesRepository(sqlDB)
	tvSeriesService := Service.NewTvSerriesService(tvSeriesRepo)
	tvSeriesHandler := Handler.NewTvSeriesHandler(tvSeriesService)

	epRepo, err := episodeRepo.NewEpisodeRepository(sqlDB, b2KeyID, b2AppKey, b2BucketName)
	if err != nil {
		log.Fatal("Failed to create Episode Repository ❌:", err)
	}

	epService := episodeService.NewEpisodeService(epRepo)
	epHandler := episodeHandler.NewEpisodeHandler(epService)

	// TV Series routes
	Routes.RegisterTvSeriesRoutes(mux, tvSeriesHandler)
	Routes.RegisterEpisodeRoutes(mux, epHandler)

	// ------------------ SERVER + CORS ------------------
	log.Println("TV Series Service running on :8080 🚀")

	err = http.ListenAndServe(":8080", Middleware.CorsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
