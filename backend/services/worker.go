package services 

import " github.com/kasyap1234/form-reader/pkg"
type ImageWorker struct {
	redisClient *pkg.AsynqRedisClient 
	// image processing libraries required 
}


func NewImageWorker(redisClient *pkg.AsynqRedisClient)*ImageWorker{
	return &ImageWorker{
		redisClient: redisClient,
	}
}

func(iw *ImageWorker)