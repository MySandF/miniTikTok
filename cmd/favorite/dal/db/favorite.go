package db

import (
	"context"
	"errors"
	"fmt"
	"miniTikTok/pkg/constants"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
	Status  bool  `json:"status"`
}

func (f *Favorite) TableName() string {
	return constants.UsersFavoriteTableName
}

// CreateFavorite create favorite info
func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	if err := DB.WithContext(ctx).Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

// SetFavorite set favorite status to 0
func SetFavorite(ctx context.Context, favorite *Favorite) error {
	if err := DB.Where(Favorite{UserId: favorite.UserId, VideoId: favorite.VideoId}).Assign(Favorite{Status: false}).FirstOrCreate(&Favorite{}).Error; err != nil {
		return err
	}
	return nil
}

// CancelFavorite set favorite status to 1
func CancelFavorite(ctx context.Context, favorite *Favorite) error {
	if err := DB.Where(Favorite{UserId: favorite.UserId, VideoId: favorite.VideoId}).Assign(Favorite{Status: true}).FirstOrCreate(&Favorite{}).Error; err != nil {
		return err
	}
	return nil
}

//QueryFavoriteById query favorite video info by user id
func QueryFavoriteById(ctx context.Context, userId int64) ([]int64, error) {
	favorites := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND status = ?", fmt.Sprint(userId), "0").Find(&favorites).Error; err != nil {
		return []int64{}, err
	}
	ret := make([]int64, len(favorites))
	for i, f := range favorites {
		ret[i] = f.VideoId
	}
	return ret, nil
}
