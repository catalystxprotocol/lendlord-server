package response

import (
	error1 "github.com/lendlord/lendlord-server/internal/app/error"
	"github.com/lendlord/lendlord-server/internal/app/models/vo"
)

func FailError(err error1.Error) vo.FailResp {
	return vo.FailResp{
		Error: vo.Fail{
			Code:    err.Code(),
			Message: err.Msg(),
		},
	}
}

func Success(data interface{}) vo.Success {
	return vo.Success{
		Data: data,
	}
}

func FailFormBindRequest(err error) vo.FailResp {
	return FailError(error1.WrapErrorDetail(error1.CodeInvalidParam, error1.ParamsIllegal+":"+err.Error()))
}

func FailRequest(msg string) vo.FailResp {
	return FailError(error1.WrapErrorDetail(error1.CodeInvalidParam, error1.ParamsIllegal+":"+msg))
}
