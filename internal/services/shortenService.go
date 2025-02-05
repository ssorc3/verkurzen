package services

import (
	"ssorc3/verkurzen/internal/data"
	"ssorc3/verkurzen/internal/generate"
)

type ShortenService struct {
    repo data.ShortenRepo
}

func NewShortenService(repo data.ShortenRepo) ShortenService {
    return ShortenService{
        repo: repo,
    }
}

func (service ShortenService) GetFullUrl(linkId string) (string, error) {
    return service.repo.GetFullUrl(linkId)
}

func (service ShortenService) StoreUrl(fullUrl string) (string, error) {
    linkId := generate.NewLinkId()

    err := service.repo.StoreLink(linkId, fullUrl)
    if err != nil {
        return "", err
    }

    return linkId, nil
}
