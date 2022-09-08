<!--
order: 4
-->

# Clients

## CLI

Find below a list of  `icplazad` commands added with the  `x/validator-dao` module. You can obtain the full list by using the `icplazad -h` command. A CLI command can look like this:

```bash
icplazad query validator-dao params
```

### Queries

| Command                 | Subcommand         | Description                                                              |
| ----------------------- | ------------------ | ------------------------------------------------------------------------ |
| `query` `validator-dao` | `params`           | Get validator-dao params                                                 |
| `query` `validator-dao` | `authorizer-bizs`  | Get the business authorization price list of a validator                 |
| `query` `validator-dao` | `grantee-auth-bizs`| Get the business authorization guarantee list that someone pre-purchases |

### Transactions

| Command              | Subcommand           | Description                                       |
| -------------------- | -------------------- | ------------------------------------------------- |
| `tx` `validator-dao` | `req-authorization`  | Apply for business authorization from a validator |
| `tx` `validator-dao` | `add-auth-biz`       | A validator adds business authorization           |
| `tx` `validator-dao` | `update-auth-biz`    | A validator updates business authorization        |
| `tx` `validator-dao` | `remove-auth-biz`    | A validator removes business authorization        |

### Proposals

The `tx gov submit-proposal` commands allow users to query create a proposal using the governance module CLI:

**`param-change`**

Allows users to submit a `ParameterChangeProposal``.

```bash
icplazad tx gov submit-proposal param-change [proposal-file] [flags]
```

## gRPC

### Queries

| Verb   | Method                                            | Description                                                              |
| ------ | ------------------------------------------------- | ------------------------------------------------------------------------ |
| `gRPC` | `icplaza.validatordao.v1.Query/Params`            | Get validator-dao params                                                 |
| `gRPC` | `icplaza.validatordao.v1.Query/AuthorizerBizs`    | Get the business authorization price list of a validator                 |
| `gRPC` | `icplaza.validatordao.v1.Query/GranteeAuthBizs`   | Get the business authorization guarantee list that someone pre-purchases |
| `GET`  | `/icplaza/validatordao/v1/params`                 | Get validator-dao params                                                 |
| `GET`  | `/icplaza/validatordao/v1/authorizer-bizs`        | Get the business authorization price list of a validator                 |
| `GET`  | `/icplaza/validatordao/v1/grantee-auth-bizs`      | Get the business authorization guarantee list that someone pre-purchases |

### Transactions

| Verb   | Method                                               | Description                                         |
| ------ | ---------------------------------------------------- | --------------------------------------------------- |
| `gRPC` | `icplaza.validatordao.v1.Msg/ReqAuthorization`       | Apply for business authorization from a validator   |
| `gRPC` | `icplaza.validatordao.v1.Msg/WithdrawAuthorization`  | Withdraws an application for business authorization |
| `gRPC` | `icplaza.validatordao.v1.Msg/AddAuthBiz`             | A validator adds business authorization             |
| `gRPC` | `icplaza.validatordao.v1.Msg/UpdateAuthBiz`          | A validator updates business authorization          |
| `gRPC` | `icplaza.validatordao.v1.Msg/RemoveAuthBiz`          | A validator removes business authorization          |
| `GET`  | `/icplaza/validatordao/v1/tx/req-authorization`      | Apply for business authorization from a validator   |
| `GET`  | `/icplaza/validatordao/v1/tx/withdraw-authorization` | Withdraws an application for business authorization |
| `GET`  | `/icplaza/validatordao/v1/tx/add-auth-bizs`          | A validator adds business authorization             |
| `GET`  | `/icplaza/validatordao/v1/tx/update-auth-bizs`       | A validator updates business authorization          |
| `GET`  | `/icplaza/validatordao/v1/tx/remove-auth-bizs`       | A validator removes business authorization          |
