package main

import (
	"log"
	"net/http"

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
	// ------------------ DATABASE ------------------
	database, err := db.InitDB() // MongoDB for TV Series
	if err != nil {
		log.Fatal("Failed to connect to DB ❌:", err)
	}

	sqlDB, err := db.InitDB() // MySQL/Postgres for Episodes
	if err != nil {
		log.Fatal("Failed to connect to SQL DB ❌:", err)
	}

	// ------------------ TV SERIES ------------------
	tvSeriesRepo := Repository.NewTvSeriesRepository(database)
	tvSeriesService := Service.NewTvSerriesService(tvSeriesRepo)
	tvSeriesHandler := Handler.NewTvSeriesHandler(tvSeriesService)

	// ------------------ EPISODES ------------------
	// Backblaze B2 config
	b2KeyID := "f9f45a6c989e"
	b2AppKey := "00563942506fbf1481548bd202ea51e42ec0ce19b7"
	b2BucketName := "movieStream"

	//ctx := context.Background()
	//b2Client, err := b2.NewClient(ctx, b2KeyID, b2AppKey)
	//if err != nil {
	//	log.Fatal("Failed to connect to Backblaze B2 ❌:", err)
	//}
	//
	//b2Bucket, err := b2Client.Bucket(ctx, b2BucketName)
	//if err != nil {
	//	log.Fatal("Failed to get B2 bucket ❌:", err)
	//}

	epRepo, err := episodeRepo.NewEpisodeRepository(sqlDB, b2KeyID, b2AppKey, b2BucketName)
	if err != nil {
		log.Fatal("Failed to create Episode Repository ❌:", err)
	}

	epService := episodeService.NewEpisodeService(epRepo)
	epHandler := episodeHandler.NewEpisodeHandler(epService)

	// ------------------ ROUTES ------------------
	mux := http.NewServeMux()

	// TV Series routes
	Routes.RegisterTvSeriesRoutes(mux, tvSeriesHandler)

	// Episode routes
	mux.HandleFunc("/api/episodes/create", epHandler.UploadEpisode)

	// ------------------ SERVER + CORS ------------------
	log.Println("TV Series Service running on :8080 🚀")

	err = http.ListenAndServe(":8080", corsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}

// ------------------ CORS MIDDLEWARE ------------------
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
