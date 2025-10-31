package models

type MinioConfig struct {
	InputBucketName string `json:"input_bucket_name"`
	InputKey    string `json:"input_key"`
	OutputBucketName string `json:"ouput_bucket_name"`
	OutputKey string `json:"output_key"`
}