package types

import ()

// bank module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeReqAuthorization = "req_authorization"
	EventTypeWithdrawAuthorization = "withdraw_authorization"
	EventTypeAddAuthBiz       = "add_auth_biz"
	EventTypeUpdateAuthBiz    = "update_auth_biz"
	EventTypeRemoveAuthBiz    = "remove_auth_biz"

	AttributeKeyAuthorizer = "authorizer"
	AttributeKeyBiz        = "biz"
)
