package service

import (
	"fmt"

	"unfollow_tracker/models"
)

func UnfollowedDisplay(followersData []models.Relation, followingData models.MediaRelationshipsFollowing) []string {
	followersSet := make(map[string]bool)

	for _, val := range followersData {
		followersSet[val.StringListData[0].Href] = true
	}

	// count := 1
	// for _, val := range followingData.RelationshipsFollowing {
	// 	if !followersSet[fmt.Sprintf("https://www.instagram.com/%s", val.Title)] {
	// 		fmt.Printf("%d. https://www.instagram.com/%s\n", count, val.Title)
	// 		count += 1
	// 	}
	// }

	for _, val := range followersData {
		if len(val.StringListData) == 0 {
			continue
		}
		followersSet[val.StringListData[0].Href] = true
	}

	var resultData []string
	for _, val := range followingData.RelationshipsFollowing {
		if !followersSet[fmt.Sprintf("https://www.instagram.com/%s", val.Title)] {
			resultData = append(resultData, fmt.Sprintf("https://www.instagram.com/%s", val.Title))
		}
	}

	return resultData
}
