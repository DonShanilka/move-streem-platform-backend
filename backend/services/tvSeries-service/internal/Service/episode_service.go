package Service

import (
	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
	"github.com/DonShanilka/tvSeries-service/internal/cloudflare"
)

type EpisodeService struct {
	Repo       *Repository.EpisodeRepository
	Cloudflare *cloudflare.StreamClient
}

func (s *EpisodeService) CreateEpisode(ep *Models.Episode) error {
	// 1️⃣ Ask Cloudflare to create a video
	uid, err := s.Cloudflare.CreateVideo()
	if err != nil {
		return err
	}

	// 2️⃣ Save UID in EpisodeURL field
	ep.EpisodeURL = uid

	// 3️⃣ Save episode metadata to MySQL
	return s.Repo.Save(ep)
}
