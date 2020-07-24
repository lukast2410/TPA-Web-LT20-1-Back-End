package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ElTube_BackEnd/graph/generated"
	"ElTube_BackEnd/graph/model"
	"context"
	"errors"
	"strconv"
	"time"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	video := model.Video{
		Title:         input.Title,
		ChannelID:     input.ChannelID,
		Description:   input.Description,
		Day:           input.Day,
		Month:         input.Month,
		Year:          input.Year,
		Hours:         input.Hours,
		Minutes:       input.Minutes,
		Seconds:       input.Seconds,
		Views:         input.Views,
		Like:          input.Like,
		Dislike:       input.Dislike,
		Comment:       input.Comment,
		Location:      input.Location,
		Category:      input.Category,
		Visibility:    input.Visibility,
		Restrict:      input.Restrict,
		VideoPath:     input.VideoPath,
		ThumbnailPath: input.ThumbnailPath,
		Premium:       input.Premium,
		Duration:      input.Duration,
	}

	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		return nil, errors.New("failed to insert video")
	}

	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input *model.NewVideo) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("video not found")
	}

	video.Title = input.Title
	video.ChannelID = input.ChannelID
	video.Description = input.Description
	video.Day = input.Day
	video.Month = input.Month
	video.Year = input.Year
	video.Hours = input.Hours
	video.Minutes = input.Minutes
	video.Seconds = input.Seconds
	video.Views = input.Views
	video.Like = input.Like
	video.Dislike = input.Dislike
	video.Comment = input.Comment
	video.Location = input.Location
	video.Category = input.Category
	video.Visibility = input.Visibility
	video.Restrict = input.Restrict
	video.VideoPath = input.VideoPath
	video.ThumbnailPath = input.ThumbnailPath
	video.Premium = input.Premium
	video.Duration = input.Duration

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update failed")
	}

	return &video, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("video not found")
	}

	_, error := r.DB.Model(&video).Where("id = ?", id).Delete()

	if error != nil {
		return false, errors.New("failed to delete")
	}

	return true, nil
}

func (r *mutationResolver) GetOneVideo(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("channel not found")
	}

	video.Views = video.Views + 1

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update views failed")
	}

	return &video, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	user := model.User{
		Restricted: input.Restricted,
		Premium:    input.Premium,
		Name:       input.Name,
		Photourl:   input.Photourl,
		Location:   input.Location,
		Email:      input.Email,
	}

	_, err := r.DB.Model(&user).Insert()

	if err != nil {
		return nil, errors.New("failed to insert user")
	}

	i64, _ := strconv.ParseInt(user.ID, 10, 64)
	currentTime := time.Now()

	channel := model.Channel{
		Name:        input.Name,
		UserID:      int(i64),
		Description: "",
		Day:         currentTime.Day(),
		Month:       int(currentTime.Month()),
		Year:        currentTime.Year(),
		Stats:       "",
		Subscriber:  1,
		ArtPath:     "",
		IconPath:    input.Photourl,
		ChannelLink: "",
	}

	_, chanErr := r.DB.Model(&channel).Insert()

	if chanErr != nil {
		return nil, errors.New("failed to insert channel")
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, email string, input *model.NewUser) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("email = ?", email).First()

	if err != nil {
		return nil, errors.New("use not found")
	}

	user.Email = input.Email
	user.Location = input.Location
	user.Photourl = input.Photourl
	user.Name = input.Name
	user.Premium = input.Premium
	user.Restricted = input.Restricted

	_, updateErr := r.DB.Model(&user).Where("email = ?", email).Update()

	if updateErr != nil {
		return nil, errors.New("Update failed")
	}

	i64, _ := strconv.ParseInt(user.ID, 10, 64)
	var channel model.Channel

	chanErr := r.DB.Model(&channel).Where("user_id = ?", int(i64)).First()

	if chanErr != nil {
		return nil, errors.New("channel not found")
	}

	channel.IconPath = input.Photourl
	channel.Name = input.Name

	_, updateChanErr := r.DB.Model(&channel).Where("user_id = ?", int(i64)).Update()

	if updateChanErr != nil {
		return nil, errors.New("Update channel failed")
	}

	return &user, nil
}

func (r *mutationResolver) GetOneUser(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("email = ?", email).First()

	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *mutationResolver) CreateChannel(ctx context.Context, input *model.NewChannel) (*model.Channel, error) {
	channel := model.Channel{
		Name:        input.Name,
		UserID:      input.UserID,
		Description: input.Description,
		Day:         input.Day,
		Month:       input.Month,
		Year:        input.Year,
		Stats:       input.Stats,
		Subscriber:  input.Subscriber,
		ArtPath:     input.ArtPath,
		IconPath:    input.IconPath,
		ChannelLink: input.ChannelLink,
	}

	_, err := r.DB.Model(&channel).Insert()

	if err != nil {
		return nil, errors.New("failed to insert channel")
	}

	return &channel, nil
}

func (r *mutationResolver) GetOneChannelByUser(ctx context.Context, userID int) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("user_id = ?", userID).First()

	if err != nil {
		return nil, errors.New("channel not found")
	}

	return &channel, nil
}

func (r *mutationResolver) GetOneChannelByID(ctx context.Context, id string) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("channel not found")
	}

	return &channel, nil
}

func (r *mutationResolver) CreateActivity(ctx context.Context, input *model.NewActivity) (*model.Activity, error) {
	activity := model.Activity{
		To:   input.To,
		From: input.From,
		Tipe: input.Tipe,
	}

	_, err := r.DB.Model(&activity).Insert()

	if err != nil {
		return nil, errors.New("failed to insert video")
	}

	return &activity, nil
}

func (r *mutationResolver) DeleteActivity(ctx context.Context, input *model.NewActivity) (bool, error) {
	var activity model.Activity

	_, err := r.DB.Query(&activity, `delete from activities where "to" = ? and "from" = ? 
		and tipe = ?`, input.To, input.From, input.Tipe)

	if err != nil {
		return false, errors.New("delete activity failed")
	}

	return true, nil
}

func (r *mutationResolver) DoActivity(ctx context.Context, table string, id string, do int) (bool, error) {
	if table == "comment" {

	} else if table == "like video" {
		var video model.Video

		err := r.DB.Model(&video).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("video not found")
		}

		video.Like = video.Like + do

		_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update like video failed")
		}

		return true, nil
	} else if table == "dislike video" {
		var video model.Video

		err := r.DB.Model(&video).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("video not found")
		}

		video.Dislike = video.Dislike + do

		_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update dislike video failed")
		}

		return true, nil
	} else if table == "like video reverse" {
		var video model.Video

		err := r.DB.Model(&video).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("video not found")
		}

		video.Like = video.Like + do
		video.Dislike = video.Dislike - do

		_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update dislike video failed")
		}

		return true, nil
	} else if table == "channel" {

	}

	return false, nil
}

func (r *queryResolver) VideosForHome(ctx context.Context, location string, visibility string, premium string) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("visibility = ? and premium = ?", visibility, premium).
		OrderExpr("case when location = ? then '1' else location end asc", location).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) RelatedVideos(ctx context.Context, id string, location string, category string, visibility string, premium string) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("visibility = ? and premium = ? and id != ?", visibility, premium, id).
		OrderExpr("case when location = ? and category = ? then '1' "+
			"when location = ? or category = ? then '2' "+
			"else location end asc", location, category, location, category).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) CheckActivity(ctx context.Context, cond string, to string, from string) (*model.Activity, error) {
	var typeOne = ""
	var typeTwo = ""
	if cond == "video" {
		typeOne = "Like Video"
		typeTwo = "Dislike Video"
	} else if cond == "comment" {
		typeOne = "Like Comment"
		typeTwo = "Dislike Comment"
	} else if cond == "subscribe" {
		typeOne = "Subscribed"
		typeTwo = "Subscribed"
	}

	var activity model.Activity

	_, err := r.DB.Query(&activity, `select * from activities where "to" = ? and "from" = ? and 
		(tipe = ? or tipe = ?)`, to, from, typeOne, typeTwo)

	if err != nil {
		return nil, errors.New("query activity failed")
	}

	return &activity, nil
}

func (r *queryResolver) GetVideo(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("channel not found")
	}

	video.Views = video.Views + 1

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update views failed")
	}

	return &video, nil
}

func (r *queryResolver) GetOneCategory(ctx context.Context, category string) (*model.Category, error) {
	var cat model.Category

	err := r.DB.Model(&cat).Where("name = ?", category).First()

	if err != nil {
		return nil, errors.New("category not found")
	}

	return &cat, nil
}

func (r *queryResolver) CategoryAllTime(ctx context.Context, category string) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("category = ?", category).Order("views desc").Limit(20).Select();

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) CategoryThisWeek(ctx context.Context, category string) ([]*model.Video, error) {
	var video []*model.Video

	//err := r.DB.Model(&video).Where("category = ?", category).Order("views desc").Limit(5);
	_,err := r.DB.Query(&video, `select * from videos where category = ? and
cast(cast("year" as text) || '-' || cast("month" as text) || '-' || cast("day" as text) as date)
<= CURRENT_DATE and
cast(cast("year" as text) || '-' || cast("month" as text) || '-' || cast("day" as text) as date)
>= (CURRENT_DATE - interval '7 days') order by "views" desc limit 20`, category)

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) CategoryThisMonth(ctx context.Context, category string) ([]*model.Video, error) {
	var video []*model.Video

	//err := r.DB.Model(&video).Where("category = ?", category).Order("id desc").Limit(5);
	_,err := r.DB.Query(&video, `select * from videos where category = ? and
cast(cast("year" as text) || '-' || cast("month" as text) || '-' || cast("day" as text) as date)
<= CURRENT_DATE and
cast(cast("year" as text) || '-' || cast("month" as text) || '-' || cast("day" as text) as date)
>= (CURRENT_DATE - interval '31 days') order by "views" desc limit 20`, category)

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) CategoryRecently(ctx context.Context, category string) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("category = ?", category).Order("id desc").Limit(20).Select();

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
