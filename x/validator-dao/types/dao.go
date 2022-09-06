package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DaoBizs []DaoBiz
type AcqBizs []AcqBiz

func NewDaoBiz(name string, fee sdk.Coin) DaoBiz {
	return DaoBiz {
		Name: name,
		Fee: fee,
	}
}

func NewAcqBiz(from, bizName string, amount, price sdk.Coin) AcqBiz {
	return AcqBiz {
		From: from,
		BizName: bizName,
		Amount: amount,
		Price: price,
	}
}

func NewGranteeBizs(granteeAddr sdk.AccAddress, bizs AcqBizs) GranteeBizs {
	return GranteeBizs {
		Grantee: granteeAddr.String(),
		Bizs: bizs,
	}
}
