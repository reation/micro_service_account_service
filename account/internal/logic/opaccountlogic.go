package logic

import (
	"context"
	"fmt"
	"github.com/reation/micro_service_account_service/config"
	"github.com/reation/micro_service_account_service/model"
	"github.com/reation/micro_service_account_service/tool"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"math/rand"
	"time"

	"github.com/reation/micro_service_account_service/account/internal/svc"
	"github.com/reation/micro_service_account_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpAccountLogic {
	return &OpAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpAccountLogic) OpAccount(in *protoc.OpAccountRequest) (*protoc.OpAccountResponse, error) {

	accountInfo, err := l.svcCtx.AccountModel.GetAccountInfoByUserID(l.ctx, in.GetUserId())
	if err == model.ErrNotFound {
		return &protoc.OpAccountResponse{States: config.RETURN_STATE_USER_NOT_FOUND}, nil
	}

	if accountInfo.State == model.STATE_UNUSUAL {
		return &protoc.OpAccountResponse{States: config.RETURN_STATE_USER_UNUSUAL}, nil
	}

	if accountInfo.State == model.STATE_FORZEN {
		return &protoc.OpAccountResponse{States: config.RETURN_STATE_USER_FORZEN}, nil
	}
	if accountInfo.PayPassword != tool.MD5(in.GetPayPassword()) {
		return &protoc.OpAccountResponse{States: config.RETURN_STATE_PASSWORD_WRONG}, nil
	}

	var states int64
	err = l.svcCtx.Conn.Transact(func(session sqlx.Session) error {
		opBeforeBalance := accountInfo.Balance
		opPrices := in.GetPrices()
		var opAfterBalance float64
		if in.GetType() == model.TYPE_PAY_ORDER || in.GetType() == model.TYPE_PAY_MEMBER {
			if opBeforeBalance < opPrices {
				states = config.RETURN_STATE_MONEY_LESS
				return nil
			}
			opAfterBalance = opBeforeBalance - opPrices
		} else if in.GetType() == model.TYPE_COMPANSATE || in.GetType() == model.TYPE_RETURN_ORDER {
			opAfterBalance = opBeforeBalance + opPrices
		}
		code := rand.Intn(899999999) + 100000000
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		accountSql := fmt.Sprintf("update %s set `balance` = %f, `code`= %d, `update_time`= \"%s\" where `id` = %d and `code` = %d ",
			"user_account", opAfterBalance, code, nowTime, accountInfo.Id, accountInfo.Code)
		_, err = session.Exec(accountSql)
		if err != nil {
			states = config.RETURN_STATE_ERROR
			return err
		}

		logSql := fmt.Sprintf("insert into %s (%s) values (%d, %d, %s, %f, %f, %f, %d)",
			" user_account_log", "`user_id`, `payment_id`, `order_num`, `op_before_balance`, `op_prices`, `op_after_balance`, `type` ",
			in.GetUserId(), in.GetPaymentId(), in.GetOrderNum(), opBeforeBalance, opPrices, opAfterBalance, in.GetType())
		_, err = session.Exec(logSql)
		if err != nil {
			states = config.RETURN_STATE_ERROR
			return err
		}

		states = config.RETURN_STATE_NORMAL
		return nil
	})

	return &protoc.OpAccountResponse{States: states}, nil
}
