package models

type MinioConfig struct {
	BucketName string `json:"bucket_name"`
	Prefix     string `json:"prefix"`
}