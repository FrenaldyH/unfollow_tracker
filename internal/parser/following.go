package parser

import (
	"encoding/json"
	"fmt"

	"unfollow_tracker/models"
	"unfollow_tracker/pkg/helper"
	"unfollow_tracker/pkg/logger"
)

const relativeFollowingPath = "connections/followers_and_following/following.json"

func FollowingParser(path string) (models.MediaRelationshipsFollowing, error) {
	followingPath := fmt.Sprintf("%s/%s", path, relativeFollowingPath)

	byteValue, err := helper.ReadFile(followingPath)
	if err != nil {
		return models.MediaRelationshipsFollowing{}, err
	}

	var data models.MediaRelationshipsFollowing
	if err := json.Unmarshal(byteValue, &data); err != nil {
		logger.Log.Warn("Failed to decode following",
			"error", err,
			"file", followingPath,
		)
		return models.MediaRelationshipsFollowing{}, err
	}

	logger.Log.Info("Parsing Followers Successfully",
		"file", followingPath,
	)
	return data, nil
}
