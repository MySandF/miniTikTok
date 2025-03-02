package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mime/multipart"
	"miniTikTok/pkg/errno"
)

type UserParam struct {
	UserName string `json:"username" query:"username"`
	PassWord string `json:"password" query:"password"`
}

type RegisterResponse struct {
	Code    int64  `json:"status_code" `
	Message string `json:"status_msg"`
	ID      int64  `json:"user_id"`
	Token   string `json:"token"`
}

func SendRegisterResponse(c *app.RequestContext, err error, id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, RegisterResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		ID:      id,
		Token:   token,
	})
}

type LoginResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	ID      int64  `json:"user_id"`
	Token   string `json:"token"`
}

func SendLoginResponse(c *app.RequestContext, err error, id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, LoginResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		ID:      id,
		Token:   token,
	})
}

// UserInfoParam get请求 tag要设置query
// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/
type UserInfoParam struct {
	ID    string `json:"user_id" form:"user_id" query:"user_id"`
	Token string `json:"token" form:"token" query:"token"`
}

type User struct {
	ID            int64  `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type InfoResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	User    User   `json:"user"`
}

func SendInfoResponse(c *app.RequestContext, err error, user User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, InfoResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		User:    user,
	})
}

type PublishParam struct {
	//绑定文件
	//https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/
	Data  *multipart.FileHeader `json:"data,omitempty" form:"data" query:"data"`
	Token string                `json:"token,omitempty" form:"token" query:"token"`
	Title string                `json:"title,omitempty" form:"title" query:"title"`
}

type PublishResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
}

func SendPublishResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, PublishResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
	})
}

type ListParam struct {
	ID    string `json:"user_id" form:"user_id" query:"user_id"`
	Token string `json:"token" form:"token" query:"token"`
}

type Video struct {
	ID            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type VideoListResponse struct {
	Code      int64   `json:"status_code"`
	Message   string  `json:"status_msg"`
	VideoList []Video `json:"video_list"`
}

func SendListResponse(c *app.RequestContext, videoList []Video, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, FeedResponse{
		Code:      int64(Err.ErrCode),
		Message:   Err.ErrMsg,
		VideoList: videoList,
	})
}

type FeedParam struct {
	LatestTime string `json:"latest_time" query:"latest_time" form:"latest_time"`
	Token      string `json:"token" query:"token" form:"token"`
}

type FeedResponse struct {
	Code      int64   `json:"status_code"`
	Message   string  `json:"status_msg"`
	VideoList []Video `json:"video_list"`
}

func SendFeedResponse(c *app.RequestContext, videoList []Video, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, FeedResponse{
		Code:      int64(Err.ErrCode),
		Message:   Err.ErrMsg,
		VideoList: videoList,
	})
}

type CommentParam struct {
	Token       string `json:"token" query:"token" form:"token"`
	VideoID     string `json:"video_id" query:"video_id" form:"video_id"`
	ActionType  string `json:"action_type" query:"action_type" form:"action_type"`
	CommentText string `json:"comment_text" query:"comment_text" form:"comment_text"`
	ID          string `json:"comment_id" query:"comment_id" form:"comment_id"`
}

type Comments struct {
	ID         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentResponse struct {
	Code    int64    `json:"status_code"`
	Message string   `json:"status_msg"`
	Comment Comments `json:"comment"`
}

func SendCommentResponse(c *app.RequestContext, cmt Comments, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, CommentResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		Comment: cmt,
	})
}

type CommentListParam struct {
	Token   string `json:"token" query:"token" form:"token"`
	VideoID string `json:"video_id" query:"video_id" form:"video_id"`
}

type CommentListResponse struct {
	Code        int64      `json:"status_code"`
	Message     string     `json:"status_msg"`
	CommentList []Comments `json:"comment_list"`
}

func SendCommentListResponse(c *app.RequestContext, cmtList []Comments, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, CommentListResponse{
		Code:        int64(Err.ErrCode),
		Message:     Err.ErrMsg,
		CommentList: cmtList,
	})
}
