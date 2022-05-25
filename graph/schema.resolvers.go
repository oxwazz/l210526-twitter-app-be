package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"l210526-twitter-app-be/app/entities"
	"l210526-twitter-app-be/app/entities/repositories"
	"l210526-twitter-app-be/graph/generated"
	"l210526-twitter-app-be/graph/model"
)

func (r *mutationResolver) CreateTweet(ctx context.Context, input model.NewTweet) (*entities.Tweet, error) {
	newTweet := &entities.Tweet{
		Content: input.Content,
		UserID:  input.UserID,
	}
	return repositories.StoreTweet(newTweet)
}

func (r *queryResolver) Users(ctx context.Context) ([]*entities.User, error) {
	return repositories.FetchAllUser[[]*entities.User]()
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*entities.Tweet, error) {
	return repositories.FetchAllTweet()
}

func (r *tweetResolver) User(ctx context.Context, obj *entities.Tweet) (*entities.User, error) {
	return repositories.FetchUserByID[*entities.User](obj.UserID)
}

func (r *userResolver) Tweets(ctx context.Context, obj *entities.User) ([]*entities.Tweet, error) {
	return repositories.FetchAllTweet()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tweet returns generated.TweetResolver implementation.
func (r *Resolver) Tweet() generated.TweetResolver { return &tweetResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tweetResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
