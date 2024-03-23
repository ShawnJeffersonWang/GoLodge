// Code generated by goctl. DO NOT EDIT.
package types

type AddCommentReq struct {
	HomestayComment HomestayComment `json:"homestayId"`
}

type AddCommentResp struct {
	Success bool `json:"success"`
}

type AddHomestayReq struct {
	Homestay Homestay `json:"homestay"`
}

type AddHomestayResp struct {
	Success bool `json:"success"`
}

type BusinessListReq struct {
	LastId             int64 `json:"lastId"`
	PageSize           int64 `json:"pageSize"`
	HomestayBusinessId int64 `json:"homestayBusinessId"`
}

type BusinessListResp struct {
	List []Homestay `json:"list"`
}

type CommentListReq struct {
	HomestayId int64 `json:"homestayId"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"pageSize"`
}

type CommentListResp struct {
	List []HomestayComment `json:"list"`
}

type GoodBossReq struct {
}

type GoodBossResp struct {
	List []HomestayBusinessBoss `json:"list"`
}

type GuessListReq struct {
}

type GuessListResp struct {
	List []Homestay `json:"list"`
}

type Homestay struct {
	Id                 int64   `json:"id"`
	Title              string  `json:"title"`
	Cover              string  `json:"cover"`
	Intro              string  `json:"intro"`
	Location           string  `json:"location"`
	HomestayBusinessId int64   `json:"homestayBusinessId"` //店铺id
	UserId             int64   `json:"userId"`             //房东id
	RowState           int64   `json:"rowState"`           //0:下架 1:上架
	RatingStars        float64 `json:"ratingStars"`
	PriceBefore        int64   `json:"priceBefore"` //民宿价格
	PriceAfter         int64   `json:"priceAfter"`
}

type HomestayBusiness struct {
	Id        int64   `json:"id"`
	Title     string  `json:"title"` //店铺名称
	Info      string  `json:"info"`  //店铺介绍
	Tags      string  `json:"tags"`  //标签，多个用“,”分割
	Cover     string  `json:"cover"` //
	Star      float64 `json:"star"`
	IsFav     int64   `json:"isFav"`
	HeaderImg string  `json:"headerImg"` //店招门头图片
}

type HomestayBusinessBoss struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"userId"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"` //房东介绍
	Rank     int64  `json:"rank"` //排名
}

type HomestayBusinessListInfo struct {
	HomestayBusiness
	SellMonth     int64 `json:"sellMonth"`     //月销售
	PersonConsume int64 `json:"personConsume"` //个人消费
}

type HomestayBussinessDetailReq struct {
	Id int64 `json:"id"`
}

type HomestayBussinessDetailResp struct {
	Boss HomestayBusinessBoss `json:"boss"`
}

type HomestayBussinessListReq struct {
	LastId   int64 `json:"lastId"`
	PageSize int64 `json:"pageSize"`
}

type HomestayBussinessListResp struct {
	List []HomestayBusinessListInfo `json:"list"`
}

type HomestayComment struct {
	Id         int64  `json:"id"`
	HomestayId int64  `json:"homestayId"`
	Content    string `json:"content"`
	Star       string `json:"star"`
	UserId     int64  `json:"userId"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
}

type HomestayDetailReq struct {
	HomestayId int64 `json:"homestayId"`
}

type HomestayDetailResp struct {
	Homestay Homestay `json:"homestay"`
}

type HomestayListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type HomestayListResp struct {
	List []Homestay `json:"list"`
}
