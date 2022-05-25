package repositories

import (
	"github.com/oxwazz/l210526-twitter-app-be/app/entities"
	"github.com/oxwazz/l210526-twitter-app-be/app/entities/databases"
)

//func FetchTweetByID(ID string) (*entities.Tweet, error) {
//	db := databases.CreateConnection()
//	var tweet *entities.Tweet
//	result := db.Where("user_id = ?", ID).Find(&tweet)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return tweet, nil
//}
func FetchAllTweet() ([]*entities.Tweet, error) {
	db := databases.CreateConnection()
	var tweets []*entities.Tweet
	result := db.Order("created_at desc").Find(&tweets)
	if result.Error != nil {
		return nil, result.Error
	}
	return tweets, nil
}

func StoreTweet(newTweet *entities.Tweet) (*entities.Tweet, error) {
	db := databases.CreateConnection()
	result := db.Create(&newTweet)
	if result.Error != nil {
		return nil, result.Error
	}
	return newTweet, nil
}
