package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestMsgTransferNft(t *testing.T) {
	//var sender = "unique15hng4pll3upv5gj842d22lw83x38yuj6n9z9sx"
	//var recipient = "unique15hng4pll3upv5gj842d22lw83x38yuj6n9z9sx"
	//var cateId = "cate-id"
	//var tokenId = "token-id"
	//msg := NewMsgTransferNft(sender, recipient, cateId, tokenId)
	//fmt.Println(msg)
	//data := msg.GetSignBytes()
	//fmt.Println(string(data))

	v,err := sdk.NewDecFromStr("5")
	if err!=nil {
		panic(err)
	}
	if v.GT(sdk.NewDec(1)) {
		fmt.Println(v)
	}
}
