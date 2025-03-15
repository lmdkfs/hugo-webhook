package service

import "fmt"

type HugoWebHookService interface {
	UpdateWebSite() error
}

type hugoWebHookService struct {
	*Service
}

func NewHugoWebHookService(service *Service) HugoWebHookService {
	return &hugoWebHookService{
		Service: service,
	}
}

func (s *hugoWebHookService) UpdateWebSite() error {
	fmt.Printf("----->\n")
	fmt.Printf("hugoWebHookService UpdateWebSite\n")
	return nil
}
