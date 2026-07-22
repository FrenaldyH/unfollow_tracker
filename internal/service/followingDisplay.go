package service

import (
	"fmt"

	"unfollow_tracker/models"
	"unfollow_tracker/pkg/logger"
)

func FollowingDisplay(data models.MediaRelationshipsFollowing) error {
	for i, val := range data.RelationshipsFollowing {
		fmt.Printf("%d. https://www.instagram.com/%s\n", i, val.Title)
	}

	logger.Log.Info("Display following successfully")
	return nil
}
