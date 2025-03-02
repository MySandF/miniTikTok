package service

import (
	"context"
	"miniTikTok/cmd/feed/dal/db"
	"miniTikTok/kitex_gen/feed"
	"miniTikTok/pkg/constants"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) FeedVideo(req *feed.DouyinFeedRequest) ([]*feed.Video, *int64, error) {
	t := time.Unix(*req.LatestTime, 0)

	var resp []*feed.Video
	videos, err := db.QueryVideoBeforeTime(s.ctx, t, constants.LimitQuery)
	if err != nil || len(videos) == 0 {
		return nil, req.LatestTime, err
	}
	nextTime := videos[0].CreatedAt.Unix()
	for _, v := range videos {
		authorID := v.AuthorId
		users, err := db.QueryUserByID(s.ctx, authorID)
		if err != nil {
			return nil, req.LatestTime, err
		}
		var u feed.User
		u.Id = int64(users[0].ID)
		u.Name = users[0].UserName
		u.FollowCount = &users[0].FollowCount
		u.FollowerCount = &users[0].FollowerCount
		vd := feed.Video{
			Id:            int64(v.ID),
			Author:        &u,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		}
		resp = append(resp, &vd)
	}

	return resp, &nextTime, nil
}
