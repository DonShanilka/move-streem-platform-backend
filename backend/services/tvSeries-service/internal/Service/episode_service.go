package Service

import (
	"io"

	"github.com/DonShanilka/tvSeries-service/internal/Models"
	"github.com/DonShanilka/tvSeries-service/internal/Repository"
)

type EpisodeService struct {
	Repo *Repository.EpisodeRepository
}

func NewEpisodeService(repo *Repository.EpisodeRepository) *EpisodeService {
	return &EpisodeService{Repo: repo}
}

func (s *EpisodeService) UploadEpisode(ep *Models.Episode, file io.Reader, fileName string) error {
	return s.Repo.SaveEpisodeWithFile(ep, file, fileName)
}
