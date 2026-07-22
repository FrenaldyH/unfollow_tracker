package parser

import (
	"encoding/json"
	"fmt"

	"unfollow_tracker/models"
	"unfollow_tracker/pkg/helper"
	"unfollow_tracker/pkg/logger"
)

const relativeFollowersPath = "connections/followers_and_following/followers_1.json"

func FollowersParser(path string) ([]models.Relation, error) {
	followerPath := fmt.Sprintf("%s/%s", path, relativeFollowersPath)

	byteValue, err := helper.ReadFile(followerPath)
	if err != nil {
		return nil, err
	}

	var data []models.Relation
	if err := json.Unmarshal(byteValue, &data); err != nil {
		logger.Log.Warn("Failed to decode followers",
			"error", err,
			"file", followerPath,
		)
		return nil, err
	}

	logger.Log.Info("Parsing Followers Successfully",
		"file", followerPath,
	)
	return data, nil
}
