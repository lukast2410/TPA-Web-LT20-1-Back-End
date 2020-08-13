package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ElTube_BackEnd/graph/generated"
	"ElTube_BackEnd/graph/model"
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
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
	channelLink := strings.ToLower(input.Name)
	channelLink = strings.ReplaceAll(channelLink, " ", "-")

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
		ChannelLink: channelLink,
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
	channelLink := strings.ToLower(input.Name)
	channelLink = strings.ReplaceAll(channelLink, " ", "-")

	channel.Name = input.Name
	channel.ChannelLink = channelLink

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

func (r *mutationResolver) UpdateChannel(ctx context.Context, id string, input *model.NewChannel) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("use not found")
	}

	channel.Description = input.Description
	channel.Stats = input.Stats
	channel.Description = input.Description
	channel.IconPath = input.IconPath
	channel.ArtPath = input.ArtPath
	channel.Subscriber = input.Subscriber

	_, updateErr := r.DB.Model(&channel).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update channel failed")
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
	if table == "like comment" {
		var comment model.Comment

		err := r.DB.Model(&comment).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("comment not found")
		}

		comment.Like = comment.Like + do

		_, updateErr := r.DB.Model(&comment).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update like comment failed")
		}

		return true, nil
	} else if table == "dislike comment" {
		var comment model.Comment

		err := r.DB.Model(&comment).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("comment not found")
		}

		comment.Dislike = comment.Dislike + do

		_, updateErr := r.DB.Model(&comment).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update like comment failed")
		}

		return true, nil
	} else if table == "like reverse comment" {
		var comment model.Comment

		err := r.DB.Model(&comment).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("comment not found")
		}

		comment.Like = comment.Like + do
		comment.Dislike = comment.Dislike - do

		_, updateErr := r.DB.Model(&comment).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update like comment failed")
		}

		return true, nil
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
	} else if table == "like post" {
		var post model.Post

		err := r.DB.Model(&post).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("post not found")
		}

		post.Like = post.Like + do

		_, updateErr := r.DB.Model(&post).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update like post failed")
		}

		return true, nil
	} else if table == "dislike post" {
		var post model.Post

		err := r.DB.Model(&post).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("post not found")
		}

		post.Dislike = post.Dislike + do

		_, updateErr := r.DB.Model(&post).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update dislike post failed")
		}

		return true, nil
	} else if table == "like post reverse" {
		var post model.Post

		err := r.DB.Model(&post).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("post not found")
		}

		post.Like = post.Like + do
		post.Dislike = post.Dislike - do

		_, updateErr := r.DB.Model(&post).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update like post reverse failed")
		}

		return true, nil
	} else if table == "channel" {
		var channel model.Channel

		err := r.DB.Model(&channel).Where("id = ?", id).First()

		if err != nil {
			return false, errors.New("channel not found")
		}

		channel.Subscriber = channel.Subscriber + do

		_, updateErr := r.DB.Model(&channel).Where("id = ?", id).Update()

		if updateErr != nil {
			return false, errors.New("update channel subs failed")
		}

		return true, nil
	}

	return false, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*model.Post, error) {
	post := model.Post{
		Content:     input.Content,
		PicturePath: input.PicturePath,
		ChannelID:   input.ChannelID,
		Day:         input.Day,
		Month:       input.Month,
		Year:        input.Year,
		Hours:       input.Hours,
		Minutes:     input.Minutes,
		Seconds:     input.Seconds,
		Like:        input.Like,
		Dislike:     input.Dislike,
	}

	_, err := r.DB.Model(&post).Insert()

	if err != nil {
		return nil, errors.New("failed to insert post")
	}

	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input *model.NewPost) (*model.Post, error) {
	var post model.Post

	err := r.DB.Model(&post).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("post not found")
	}

	post.PicturePath = input.PicturePath
	post.Content = input.Content
	post.Like = input.Like
	post.Dislike = input.Dislike

	_, updateErr := r.DB.Model(&post).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update post failed")
	}

	return &post, nil
}

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	playlist := model.Playlist{
		Name:        input.Name,
		Visibility:  input.Visibility,
		UserID:      input.UserID,
		Views:       input.Views,
		Description: input.Description,
		VideoList:   input.VideoList,
		Day:         input.Day,
		Month:       input.Month,
		Year:        input.Year,
	}

	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		return nil, errors.New("failed to insert playlist")
	}

	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, id string, input *model.NewPlaylist) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("playlist not found")
	}

	playlist.VideoList = input.VideoList
	playlist.Views = input.Views
	playlist.Description = input.Description
	playlist.Day = input.Day
	playlist.Month = input.Month
	playlist.Year = input.Year
	playlist.UserID = input.UserID
	playlist.Name = input.Name
	playlist.Visibility = input.Visibility

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update post failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, id string) (bool, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("playlist not found")
	}

	_, delErr := r.DB.Model(&playlist).Where("id = ?", id).Delete()

	if delErr != nil {
		return false, errors.New("delete playlist failed")
	}

	return true, nil
}

func (r *mutationResolver) ViewPlaylist(ctx context.Context, id string) (bool, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("playlist not found")
	}

	playlist.Views = playlist.Views + 1

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		return false, errors.New("update playlist video failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		UserID:    input.UserID,
		VideoID:   input.VideoID,
		Content:   input.Content,
		Day:       input.Day,
		Month:     input.Month,
		Year:      input.Year,
		Hours:     input.Hours,
		Minutes:   input.Minutes,
		Seconds:   input.Seconds,
		CommentID: input.CommentID,
		Like:      input.Like,
		Dislike:   input.Dislike,
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, errors.New("failed to insert comment")
	}

	return &comment, nil
}

func (r *mutationResolver) CreateQueue(ctx context.Context, input *model.NewQueue) (*model.Queue, error) {
	queue := model.Queue{
		UserID:  input.UserID,
		VideoID: input.VideoID,
	}

	_, err := r.DB.Model(&queue).Insert()

	if err != nil {
		return nil, errors.New("failed to insert queue")
	}

	return &queue, nil
}

func (r *mutationResolver) DeleteQueue(ctx context.Context, id string) (bool, error) {
	var queue model.Queue

	err := r.DB.Model(&queue).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("queue not found")
	}

	_, delErr := r.DB.Model(&queue).Where("id = ?", id).Delete()

	if delErr != nil {
		return false, errors.New("delete queue failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateBilling(ctx context.Context, input *model.NewBilling) (*model.Billing, error) {
	billing := model.Billing{
		UserID: input.UserID,
		Day:    input.Day,
		Month:  input.Month,
		Year:   input.Year,
		Tipe:   input.Tipe,
	}

	_, err := r.DB.Model(&billing).Insert()

	if err != nil {
		return nil, errors.New("failed to insert billing")
	}

	return &billing, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *queryResolver) GetAllVideos(ctx context.Context) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) VideosForHome(ctx context.Context, location string, visibility string, premium string, restrict string) ([]*model.Video, error) {
	if premium == "true" {
		premium = "%%"
	} else {
		premium = "%false%"
	}
	if restrict == "true" {
		restrict = "%false%"
	} else {
		restrict = "%%"
	}

	var videoL []*model.Video
	var video []*model.Video

	err := r.DB.Model(&video).Where("location != ? and visibility = ? and premium like ? and restrict like ?", location, visibility, premium, restrict).Select()
	errL := r.DB.Model(&videoL).Where("location = ? and visibility = ? and premium like ? and restrict like ?", location, visibility, premium, restrict).Select()

	if err != nil || errL != nil {
		return nil, errors.New("failed to query videos")
	}
	v := random(videoL)
	vs := random(video)

	for _, va := range vs {
		v = append(v, va)
	}
	return v, nil
}

func (r *queryResolver) RelatedVideos(ctx context.Context, id string, location string, category string, visibility string, premium string) ([]*model.Video, error) {
	if premium == "true" {
		premium = "%%"
	} else if premium == "false" {
		premium = "%false%"
	}

	var video []*model.Video

	err := r.DB.Model(&video).Where("visibility = ? and premium like ? and id != ?", visibility, premium, id).
		OrderExpr("case when location = ? and category = ? then '1' "+
			"when location = ? or category = ? then '2' "+
			"else location end asc", location, category, location, category).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) GetActivity(ctx context.Context) ([]*model.Activity, error) {
	var act []*model.Activity

	err := r.DB.Model(&act).Select()

	if err != nil {
		return nil, errors.New("failed to query activity")
	}

	return act, nil
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
	} else if cond == "channel" {
		typeOne = "Subscribed"
		typeTwo = "Subscribed"
	} else if cond == "post" {
		typeOne = "Like Post"
		typeTwo = "Dislike Post"
	}

	var activity model.Activity

	_, err := r.DB.Query(&activity, `select * from activities where "to" = ? and "from" = ? and 
		(tipe = ? or tipe = ?)`, to, from, typeOne, typeTwo)

	if err != nil {
		return nil, errors.New("query activity failed")
	}

	return &activity, nil
}

func (r *queryResolver) GetMyActivity(ctx context.Context, from string, cond string) ([]*model.Activity, error) {
	var act []*model.Activity

	_, err := r.DB.Query(&act, `select * from activities where "from" = ? and 
		tipe = ?`, from, cond)

	if err != nil {
		return nil, errors.New("failed to get my activity")
	}

	return act, nil
}

func (r *queryResolver) GetVideo(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("channel not found")
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

	err := r.DB.Model(&video).Where("category = ?", category).Order("views desc").Limit(20).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) CategoryThisWeek(ctx context.Context, category string) ([]*model.Video, error) {
	var video []*model.Video

	//err := r.DB.Model(&video).Where("category = ?", category).Order("views desc").Limit(5);
	_, err := r.DB.Query(&video, `select * from videos where category = ? and
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
	_, err := r.DB.Query(&video, `select * from videos where category = ? and
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

	err := r.DB.Model(&video).Where("category = ?", category).Order("id desc").Limit(20).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) GetTrendingVideos(ctx context.Context) ([]*model.Video, error) {
	var visibility = "public"
	var video []*model.Video

	//err := r.DB.Model(&video).Where("category = ?", category).Order("views desc").Limit(5);
	_, err := r.DB.Query(&video, `select * from videos where visibility = ? and 
cast(cast("year" as text) || '-' || cast("month" as text) || '-' || cast("day" as text) as date)
<= CURRENT_DATE and
cast(cast("year" as text) || '-' || cast("month" as text) || '-' || cast("day" as text) as date)
>= (CURRENT_DATE - interval '7 days') order by "views" asc limit 20`, visibility)

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) GetOneChannelByLink(ctx context.Context, link string) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("channel_link = ?", link).First()

	if err != nil {
		return nil, errors.New("channel not found")
	}

	return &channel, nil
}

func (r *queryResolver) GetOneChannelByUser(ctx context.Context, userID int) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("user_id = ?", userID).First()

	if err != nil {
		return nil, errors.New("channel not found")
	}

	return &channel, nil
}

func (r *queryResolver) GetOneChannelByID(ctx context.Context, id string) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("channel not found")
	}

	return &channel, nil
}

func (r *queryResolver) GetChannelVideos(ctx context.Context, channelID int) ([]*model.Video, error) {
	var path = ""
	var video []*model.Video

	err := r.DB.Model(&video).Where("channel_id = ? and video_path != ?", channelID, path).Order("id desc").Select()

	if err != nil {
		return nil, errors.New("failed to get channel videos")
	}

	return video, nil
}

func (r *queryResolver) GetAllPost(ctx context.Context, channelID int) ([]*model.Post, error) {
	var post []*model.Post

	err := r.DB.Model(&post).Where("channel_id = ?", channelID).
		Order("id desc").Select()

	if err != nil {
		return nil, errors.New("failed to get post")
	}

	return post, nil
}

func (r *queryResolver) GetOnePost(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post

	err := r.DB.Model(&post).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("failed to get post")
	}

	return &post, nil
}

func (r *queryResolver) GetTotalViews(ctx context.Context, channelID int) (*model.Video, error) {
	var video model.Video

	_, err := r.DB.Query(&video, `select sum(coalesce(views, 0)) as views from videos where channel_id = ?`, channelID)

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return &video, nil
}

func (r *queryResolver) GetMyPlaylist(ctx context.Context, userID int) ([]*model.Playlist, error) {
	var playlist []*model.Playlist

	err := r.DB.Model(&playlist).Where("user_id = ?", userID).Select()

	if err != nil {
		return nil, errors.New("failed to query playlist")
	}

	return playlist, nil
}

func (r *queryResolver) GetOnePlaylist(ctx context.Context, id string) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("failed to get playlist")
	}

	return &playlist, nil
}

func (r *queryResolver) GetComment(ctx context.Context, videoID int) ([]*model.Comment, error) {
	var comment []*model.Comment

	err := r.DB.Model(&comment).Where("video_id = ?", videoID).Select()

	if err != nil {
		return nil, errors.New("failed to query comment")
	}

	return comment, nil
}

func (r *queryResolver) GetReply(ctx context.Context, commentID int) ([]*model.Comment, error) {
	var comment []*model.Comment

	err := r.DB.Model(&comment).Where("comment_id = ?", commentID).Select()

	if err != nil {
		return nil, errors.New("failed to query comment")
	}

	return comment, nil
}

func (r *queryResolver) SearchVideo(ctx context.Context, word string) ([]*model.Video, error) {
	word = "%" + word + "%"
	var video []*model.Video

	err := r.DB.Model(&video).Where(`title ilike ? or description ilike ?`, word, word).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	return video, nil
}

func (r *queryResolver) SearchPlaylist(ctx context.Context, word string) ([]*model.Playlist, error) {
	word = "%" + word + "%"
	var playlist []*model.Playlist

	err := r.DB.Model(&playlist).Where(`name ilike ? or description ilike ?`, word, word).Select()

	if err != nil {
		return nil, errors.New("failed to query playlist")
	}

	return playlist, nil
}

func (r *queryResolver) SearchChannel(ctx context.Context, word string) ([]*model.Channel, error) {
	word = "%" + word + "%"
	var channel []*model.Channel

	err := r.DB.Model(&channel).Where(`name ilike ? or description ilike ?`, word, word).Select()

	if err != nil {
		return nil, errors.New("failed to query playlist")
	}

	return channel, nil
}

func (r *queryResolver) GetMyQueue(ctx context.Context, userID int) ([]*model.Queue, error) {
	var queue []*model.Queue

	err := r.DB.Model(&queue).Where(`user_id = ?`, userID).Select()

	if err != nil {
		return nil, errors.New("failed to query playlist")
	}

	return queue, nil
}

func (r *queryResolver) Autocomplete(ctx context.Context, word string) ([]string, error) {
	word = "%" + word + "%"
	var video []*model.Video

	err := r.DB.Model(&video).Where(`title ilike ?`, word).Select()

	if err != nil {
		return nil, errors.New("failed to query videos")
	}

	var result []string
	for i := 0; i < len(video); i++ {
		result = append(result, strings.ToLower(video[i].Title))
	}

	var playlist []*model.Playlist

	errPl := r.DB.Model(&playlist).Where(`name ilike ?`, word).Select()

	if errPl != nil {
		return nil, errors.New("failed to query playlist")
	}

	for i := 0; i < len(playlist); i++ {
		result = append(result, strings.ToLower(playlist[i].Name))
	}

	var channel []*model.Channel

	errCh := r.DB.Model(&channel).Where(`name ilike ?`, word).Select()

	if errCh != nil {
		return nil, errors.New("failed to query channel")
	}

	for i := 0; i < len(channel); i++ {
		result = append(result, strings.ToLower(channel[i].Name))
	}

	return result, nil
}

func (r *queryResolver) GetMyBilling(ctx context.Context, userID int) ([]*model.Billing, error) {
	var billing []*model.Billing

	err := r.DB.Model(&billing).Where(`user_id = ?`, userID).Order("id desc").Select()

	if err != nil {
		return nil, errors.New("failed to query billing")
	}

	return billing, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func random(v []*model.Video) []*model.Video {
	for i := len(v) - 1; i > 0; i-- {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		j := r1.Intn(i + 1)
		v[i], v[j] = v[j], v[i]
	}
	return v
}
