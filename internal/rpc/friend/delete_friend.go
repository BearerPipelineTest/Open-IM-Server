package friend

import (
	"Open_IM/pkg/common/constant"
	"Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/token_verify"
	pbFriend "Open_IM/pkg/proto/friend"
	"context"
)

func (s *friendServer) DeleteFriend(ctx context.Context, req *pbFriend.DeleteFriendReq) (*pbFriend.CommonResp, error) {
	log.Info(req.Token, req.OperationID, "rpc delete friend is server,args=%s", req.String())
	//Parse token, to find current user information
	claims, err := token_verify.ParseToken(req.Token)
	if err != nil {
		log.Error(req.Token, req.OperationID, "err=%s,parse token failed", err.Error())
		return &pbFriend.CommonResp{ErrorCode: constant.ErrParseToken.ErrCode, ErrorMsg: constant.ErrParseToken.ErrMsg}, nil
	}
	err = im_mysql_model.DeleteSingleFriendInfo(claims.UID, req.Uid)
	if err != nil {
		log.Error(req.Token, req.OperationID, "err=%s,delete friend failed", err.Error())
		return &pbFriend.CommonResp{ErrorCode: constant.ErrMysql.ErrCode, ErrorMsg: constant.ErrMysql.ErrMsg}, nil
	}
	log.Info(req.Token, req.OperationID, "rpc delete friend success return")
	return &pbFriend.CommonResp{}, nil
}
