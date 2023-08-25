package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/hopeio/lemon/_example/protobuf/user"
	"github.com/hopeio/lemon/_example/user/conf"
	"github.com/hopeio/lemon/context/http_context"
	"github.com/hopeio/lemon/protobuf/errorcode"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (u *UserService) Signup(ctx context.Context, req *user.SignupReq) (*wrappers.StringValue, error) {
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

	db := ctxi.NewDB(conf.Dao.GORMDB.DB)
	err := db.Create(&user).Error
	if err != nil {
		return nil, ctxi.ErrorLog(errorcode.DBError.Message("新建出错"), err, "UserService.Creat")
	}
	return &wrappers.StringValue{Value: "注册成功"}, nil
}
