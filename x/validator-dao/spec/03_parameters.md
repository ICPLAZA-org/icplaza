<!--
order: 3
-->

# Parameters

The validator-dao module contains the following parameters:

| Key                     | Type          | Default Value                                                     |
| ----------------------- | ------------- | ----------------------------------------------------------------- |
| `AuthBizs`              | DaoBizs       | `{issue-erc20, sdk.NewCoin(sdk.DefaultBondDenom, sdk.ZeroInt())}` |

## Authorization Business

The `AuthBizs` parameter is the global application authorization list that the ICPLAZA network needs to govern.
