package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golodge/app/travel/cmd/rpc/internal/svc"
	"golodge/app/travel/cmd/rpc/pb"
	"golodge/app/travel/cmd/rpc/travel"
	"golodge/app/travel/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayLogic {
	return &AddHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddHomestayLogic) AddHomestay(in *pb.AddHomestayReq) (*pb.AddHomestayResp, error) {
	// todo: add your logic here and delete this line

	//_, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Homestay.Id)
	//if err == nil {
	//	return nil, errors.Wrapf(ErrHomestayAlreadyAdded,
	//		"homestay has been added in homestayList homestayId:%d,err:%v", in.Homestay.Id, err)
	//}

	if err := l.svcCtx.HomestayModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		homestay := model.Homestay{
			Title:       in.Homestay.Title,
			Cover:       in.Homestay.Cover,
			CleanVideo:  in.Homestay.CleanVideo,
			ImageUrls:   in.Homestay.ImageUrls,
			Intro:       in.Homestay.Intro,
			Location:    in.Homestay.Location,
			UserId:      in.Homestay.UserId,
			PriceBefore: in.Homestay.PriceBefore,
			RowState:    1,
		}

		res, err := l.svcCtx.HomestayModel.Insert(ctx, session, &homestay)
		if err != nil {
			return err
		}

		// bug: 这里可以直接通过homestay获取新加入的Id, 而不用再去查询, 多此一举, 因为是通过&homestay获取的
		// 就像C语言的传入传出参数一样, 更正: 虽然类似传出参数, 但是不能获取自增的id
		// 之前加入的无法在列表中拿到是因为之加入了Homestay表, 而没有加入HomestayActivity表
		// 困扰了n天的bug: 加入数据库后自增的id使用sql.Result接口中的LastInsertId()方法获取
		dataId, _ := res.LastInsertId()
		homestayActivity := model.HomestayActivity{
			DelState:  0,
			RowType:   "preferredHomestay",
			DataId:    dataId,
			RowStatus: 1,
			Version:   0,
		}

		_, err = l.svcCtx.HomestayActivityModel.Insert(ctx, session, &homestayActivity)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &travel.AddHomestayResp{}, nil
}
