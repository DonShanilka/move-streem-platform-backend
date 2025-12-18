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

func (s *EpisodeService) CreateEpisode(ep *Models.Episode, file io.Reader, fileName string) error {
	return s.Repo.SaveEpisodeWithFile(ep, file, fileName)
}

func (s *EpisodeService) UpdateEpisode(
	ep *Models.Episode,
	file io.Reader,
	fileName string,
) error {
	return s.Repo.UpdateEpisodeWithFile(ep, file, fileName)
}

func (s *EpisodeService) DeleteEpisode(id int) error {
	return s.Repo.DeleteEpisode(id)
}

func (s *EpisodeService) GetAllEpisodes() ([]Models.Episode, error) {
	return s.Repo.GetAllEpisode()
}

// Business logic: get episode by ID
func (s *EpisodeService) GetEpisodeByID(id int) (*Models.Episode, error) {
	return s.Repo.GetEpisodeByID(id)
}
