<!--
order: 2
-->

# Events

The `x/validator-dao` module emits the following events:

## Request Business Authorization

| Type                | Attribute Key   | Attribute Value       |
| ------------------- | --------------- | --------------------- |
| `req_authorization` | `"authorizer"`  | `{msg.Authorizer}`    |
| `req_authorization` | `"biz"`         | `{msg.BizName}`       |
| `req_authorization` | `"amount"`      | `{msg.Fee.String()}`  |

## Withdraw Business Authorization Requestion

| Type                     | Attribute Key   | Attribute Value   |
| ------------------------ | --------------- | ----------------- |
| `withdraw_authorization` | `"authorizer"`  | `{msg.Authorizer}`|
| `withdraw_authorization` | `"biz"`         | `{msg.BizName}`   |
| `withdraw_authorization` | `"amount"`      | `{fees.String()}` |

## Add Business To Be Authorized

| Type           | Attribute Key   | Attribute Value   |
| -------------- | --------------- | ----------------- |
| `add_auth_biz` | `"sender"`      | `{msg.Sender}`    |
| `add_auth_biz` | `"biz"`         | `{msg.BizName}`   |
| `add_auth_biz` | `"amount"`      | `{msg.Fee.String()}`       |

## Update Business To Be Authorized

| Type              | Attribute Key   | Attribute Value              |
| ----------------- | --------------- | ---------------------------- |
| `update_auth_biz` | `"sender"`      | `{msg.Sender}`               |
| `update_auth_biz` | `"amount"`      | `{msg.Fee.String()}` |

## Remove Business To Be Authorized

| Type              | Attribute Key   | Attribute Value         |
| ----------------- | --------------- | ----------------------- |
| `remove_auth_biz` | `"sender"`      | `{msg.Sender}`          |
| `remove_auth_biz` | `"biz"`         | `{msg.BizName}`         |
