package handlers

import (
	"github.com/hibiken/asynq"
	"github.com/kasyap1234/form-reader/models"
	"github.com/kasyap1234/form-reader/pkg"
)

// ingestion handler will ingest tasks to the task queue
func PushToQueue(task *models.TaskImagePayload)