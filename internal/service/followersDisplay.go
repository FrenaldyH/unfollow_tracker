package service

import (
	"fmt"

	"unfollow_tracker/models"
	"unfollow_tracker/pkg/logger"
)

func FollowersDisplay(data []models.Relation) error {
	for i, val := range data {
		fmt.Printf("%d. %s\n", i, val.StringListData[0].Href)
	}

	logger.Log.Info("Display followers successfully")
	return nil
}
