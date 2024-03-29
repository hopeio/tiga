package service

import (
	"context"
	gormi "github.com/hopeio/tiga/utils/dao/db/gorm"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/hopeio/tiga/_example/protobuf/user"
	"github.com/hopeio/tiga/_example/user/conf"
	"github.com/hopeio/tiga/context/http_context"
	"github.com/hopeio/tiga/protobuf/errorcode"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (u *UserService) Signup(ctx context.Context, req *user.SignupReq) (*wrapperspb.StringValue, error) {
	ctxi, span := http_context.ContextFromContext(ctx).StartSpan("")
	defer span.End()
	ctx = ctxi.Context
	if req.Mail == "" && req.Phone == "" {
		return nil, errorcode.DBError.Message("请填写邮箱或手机号")
	}

	formatNow := ctxi.TimeString
	var user = &user.User{
		Name: req.Name,

		Mail:   req.Mail,
		Phone:  req.Phone,
		Gender: req.Gender,

		Role:      user.RoleNormal,
		CreatedAt: formatNow,
		Status:    user.UserStatusInActive,
	}

	db := gormi.NewTraceDB(conf.Dao.GORMDB.DB, ctxi.TraceID)
	err := db.Create(&user).Error
	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.DBError.Message("新建出错"), err, "UserService.Creat")
	}
	return &wrapperspb.StringValue{Value: "注册成功"}, nil
}
