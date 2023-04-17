package service

import "github.com/diyor200/golang-gin-poc/entity"

type VideoService interface {
	FindAll() []entity.Video
	Save(entity.Video) entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}

}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
