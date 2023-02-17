<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [evmos/claims/v1/genesis.proto](#evmos/claims/v1/genesis.proto)
    - [GenesisState](#evmos.claims.v1.GenesisState)
    - [Params](#evmos.claims.v1.Params)
  
- [evmos/claims/v1/query.proto](#evmos/claims/v1/query.proto)
    - [QueryClaimsRecordRequest](#evmos.claims.v1.QueryClaimsRecordRequest)
    - [QueryClaimsRecordResponse](#evmos.claims.v1.QueryClaimsRecordResponse)
    - [QueryClaimsRecordsRequest](#evmos.claims.v1.QueryClaimsRecordsRequest)
    - [QueryClaimsRecordsResponse](#evmos.claims.v1.QueryClaimsRecordsResponse)
    - [QueryParamsRequest](#evmos.claims.v1.QueryParamsRequest)
    - [QueryParamsResponse](#evmos.claims.v1.QueryParamsResponse)
    - [QueryTotalUnclaimedRequest](#evmos.claims.v1.QueryTotalUnclaimedRequest)
    - [QueryTotalUnclaimedResponse](#evmos.claims.v1.QueryTotalUnclaimedResponse)
  
    - [Query](#evmos.claims.v1.Query)
  
- [evmos/claims/v1/claims.proto](#evmos/claims/v1/claims.proto)
    - [Claim](#evmos.claims.v1.Claim)
    - [ClaimsRecord](#evmos.claims.v1.ClaimsRecord)
    - [ClaimsRecordAddress](#evmos.claims.v1.ClaimsRecordAddress)
  
    - [Action](#evmos.claims.v1.Action)
  
- [evmos/epochs/v1/query.proto](#evmos/epochs/v1/query.proto)
    - [QueryCurrentEpochRequest](#evmos.epochs.v1.QueryCurrentEpochRequest)
    - [QueryCurrentEpochResponse](#evmos.epochs.v1.QueryCurrentEpochResponse)
    - [QueryEpochsInfoRequest](#evmos.epochs.v1.QueryEpochsInfoRequest)
    - [QueryEpochsInfoResponse](#evmos.epochs.v1.QueryEpochsInfoResponse)
  
    - [Query](#evmos.epochs.v1.Query)
  
- [evmos/epochs/v1/genesis.proto](#evmos/epochs/v1/genesis.proto)
    - [EpochInfo](#evmos.epochs.v1.EpochInfo)
    - [GenesisState](#evmos.epochs.v1.GenesisState)
  
- [evmos/erc20/v1/genesis.proto](#evmos/erc20/v1/genesis.proto)
    - [GenesisState](#evmos.erc20.v1.GenesisState)
    - [Params](#evmos.erc20.v1.Params)
  
- [evmos/erc20/v1/query.proto](#evmos/erc20/v1/query.proto)
    - [QueryParamsRequest](#evmos.erc20.v1.QueryParamsRequest)
    - [QueryParamsResponse](#evmos.erc20.v1.QueryParamsResponse)
    - [QueryTokenPairRequest](#evmos.erc20.v1.QueryTokenPairRequest)
    - [QueryTokenPairResponse](#evmos.erc20.v1.QueryTokenPairResponse)
    - [QueryTokenPairsRequest](#evmos.erc20.v1.QueryTokenPairsRequest)
    - [QueryTokenPairsResponse](#evmos.erc20.v1.QueryTokenPairsResponse)
  
    - [Query](#evmos.erc20.v1.Query)
  
- [evmos/erc20/v1/tx.proto](#evmos/erc20/v1/tx.proto)
    - [MsgConvertCoin](#evmos.erc20.v1.MsgConvertCoin)
    - [MsgConvertCoinResponse](#evmos.erc20.v1.MsgConvertCoinResponse)
    - [MsgConvertERC20](#evmos.erc20.v1.MsgConvertERC20)
    - [MsgConvertERC20Response](#evmos.erc20.v1.MsgConvertERC20Response)
  
    - [Msg](#evmos.erc20.v1.Msg)
  
- [evmos/erc20/v1/erc20.proto](#evmos/erc20/v1/erc20.proto)
    - [RegisterCoinProposal](#evmos.erc20.v1.RegisterCoinProposal)
    - [RegisterERC20Proposal](#evmos.erc20.v1.RegisterERC20Proposal)
    - [ToggleTokenConversionProposal](#evmos.erc20.v1.ToggleTokenConversionProposal)
    - [TokenPair](#evmos.erc20.v1.TokenPair)
  
    - [Owner](#evmos.erc20.v1.Owner)
  
- [evmos/incentives/v1/incentives.proto](#evmos/incentives/v1/incentives.proto)
    - [CancelIncentiveProposal](#evmos.incentives.v1.CancelIncentiveProposal)
    - [GasMeter](#evmos.incentives.v1.GasMeter)
    - [Incentive](#evmos.incentives.v1.Incentive)
    - [RegisterIncentiveProposal](#evmos.incentives.v1.RegisterIncentiveProposal)
  
- [evmos/incentives/v1/query.proto](#evmos/incentives/v1/query.proto)
    - [QueryAllocationMeterRequest](#evmos.incentives.v1.QueryAllocationMeterRequest)
    - [QueryAllocationMeterResponse](#evmos.incentives.v1.QueryAllocationMeterResponse)
    - [QueryAllocationMetersRequest](#evmos.incentives.v1.QueryAllocationMetersRequest)
    - [QueryAllocationMetersResponse](#evmos.incentives.v1.QueryAllocationMetersResponse)
    - [QueryGasMeterRequest](#evmos.incentives.v1.QueryGasMeterRequest)
    - [QueryGasMeterResponse](#evmos.incentives.v1.QueryGasMeterResponse)
    - [QueryGasMetersRequest](#evmos.incentives.v1.QueryGasMetersRequest)
    - [QueryGasMetersResponse](#evmos.incentives.v1.QueryGasMetersResponse)
    - [QueryIncentiveRequest](#evmos.incentives.v1.QueryIncentiveRequest)
    - [QueryIncentiveResponse](#evmos.incentives.v1.QueryIncentiveResponse)
    - [QueryIncentivesRequest](#evmos.incentives.v1.QueryIncentivesRequest)
    - [QueryIncentivesResponse](#evmos.incentives.v1.QueryIncentivesResponse)
    - [QueryParamsRequest](#evmos.incentives.v1.QueryParamsRequest)
    - [QueryParamsResponse](#evmos.incentives.v1.QueryParamsResponse)
  
    - [Query](#evmos.incentives.v1.Query)
  
- [evmos/incentives/v1/genesis.proto](#evmos/incentives/v1/genesis.proto)
    - [GenesisState](#evmos.incentives.v1.GenesisState)
    - [Params](#evmos.incentives.v1.Params)
  
- [evmos/inflation/v1/inflation.proto](#evmos/inflation/v1/inflation.proto)
    - [ExponentialCalculation](#evmos.inflation.v1.ExponentialCalculation)
    - [InflationDistribution](#evmos.inflation.v1.InflationDistribution)
  
- [evmos/inflation/v1/query.proto](#evmos/inflation/v1/query.proto)
    - [QueryCirculatingSupplyRequest](#evmos.inflation.v1.QueryCirculatingSupplyRequest)
    - [QueryCirculatingSupplyResponse](#evmos.inflation.v1.QueryCirculatingSupplyResponse)
    - [QueryEpochMintProvisionRequest](#evmos.inflation.v1.QueryEpochMintProvisionRequest)
    - [QueryEpochMintProvisionResponse](#evmos.inflation.v1.QueryEpochMintProvisionResponse)
    - [QueryInflationRateRequest](#evmos.inflation.v1.QueryInflationRateRequest)
    - [QueryInflationRateResponse](#evmos.inflation.v1.QueryInflationRateResponse)
    - [QueryParamsRequest](#evmos.inflation.v1.QueryParamsRequest)
    - [QueryParamsResponse](#evmos.inflation.v1.QueryParamsResponse)
    - [QueryPeriodRequest](#evmos.inflation.v1.QueryPeriodRequest)
    - [QueryPeriodResponse](#evmos.inflation.v1.QueryPeriodResponse)
    - [QuerySkippedEpochsRequest](#evmos.inflation.v1.QuerySkippedEpochsRequest)
    - [QuerySkippedEpochsResponse](#evmos.inflation.v1.QuerySkippedEpochsResponse)
  
    - [Query](#evmos.inflation.v1.Query)
  
- [evmos/inflation/v1/genesis.proto](#evmos/inflation/v1/genesis.proto)
    - [GenesisState](#evmos.inflation.v1.GenesisState)
    - [Params](#evmos.inflation.v1.Params)
  
- [evmos/recovery/v1/query.proto](#evmos/recovery/v1/query.proto)
    - [QueryParamsRequest](#evmos.recovery.v1.QueryParamsRequest)
    - [QueryParamsResponse](#evmos.recovery.v1.QueryParamsResponse)
  
    - [Query](#evmos.recovery.v1.Query)
  
- [evmos/recovery/v1/genesis.proto](#evmos/recovery/v1/genesis.proto)
    - [GenesisState](#evmos.recovery.v1.GenesisState)
    - [Params](#evmos.recovery.v1.Params)
  
- [evmos/vesting/v1/tx.proto](#evmos/vesting/v1/tx.proto)
    - [MsgClawback](#evmos.vesting.v1.MsgClawback)
    - [MsgClawbackResponse](#evmos.vesting.v1.MsgClawbackResponse)
    - [MsgCreateClawbackVestingAccount](#evmos.vesting.v1.MsgCreateClawbackVestingAccount)
    - [MsgCreateClawbackVestingAccountResponse](#evmos.vesting.v1.MsgCreateClawbackVestingAccountResponse)
  
    - [Msg](#evmos.vesting.v1.Msg)
  
- [evmos/vesting/v1/vesting.proto](#evmos/vesting/v1/vesting.proto)
    - [ClawbackVestingAccount](#evmos.vesting.v1.ClawbackVestingAccount)
  
- [evmos/vesting/v1/query.proto](#evmos/vesting/v1/query.proto)
    - [QueryBalancesRequest](#evmos.vesting.v1.QueryBalancesRequest)
    - [QueryBalancesResponse](#evmos.vesting.v1.QueryBalancesResponse)
  
    - [Query](#evmos.vesting.v1.Query)
  
- [gauss/auction/v1/params.proto](#gauss/auction/v1/params.proto)
    - [Params](#gauss.auction.v1.Params)
  
- [gauss/auction/v1/tx.proto](#gauss/auction/v1/tx.proto)
    - [MsgBidOrder](#gauss.auction.v1.MsgBidOrder)
    - [MsgBidOrderResponse](#gauss.auction.v1.MsgBidOrderResponse)
    - [MsgCreateOrder](#gauss.auction.v1.MsgCreateOrder)
    - [MsgCreateOrderResponse](#gauss.auction.v1.MsgCreateOrderResponse)
    - [MsgDeleteOrder](#gauss.auction.v1.MsgDeleteOrder)
    - [MsgDeleteOrderResponse](#gauss.auction.v1.MsgDeleteOrderResponse)
    - [Order](#gauss.auction.v1.Order)
    - [PoolAddress](#gauss.auction.v1.PoolAddress)
  
    - [Msg](#gauss.auction.v1.Msg)
  
- [gauss/auction/v1/genesis.proto](#gauss/auction/v1/genesis.proto)
    - [GenesisState](#gauss.auction.v1.GenesisState)
  
- [gauss/auction/v1/query.proto](#gauss/auction/v1/query.proto)
    - [QueryOrderRequest](#gauss.auction.v1.QueryOrderRequest)
    - [QueryOrderResponse](#gauss.auction.v1.QueryOrderResponse)
    - [QueryOrdersRequest](#gauss.auction.v1.QueryOrdersRequest)
    - [QueryOrdersResponse](#gauss.auction.v1.QueryOrdersResponse)
  
    - [Query](#gauss.auction.v1.Query)
  
- [gauss/blindbox/v1/query.proto](#gauss/blindbox/v1/query.proto)
    - [QueryGetBoxRequest](#gauss.blindbox.v1.QueryGetBoxRequest)
    - [QueryGetBoxResponse](#gauss.blindbox.v1.QueryGetBoxResponse)
    - [QueryGetGroupRequest](#gauss.blindbox.v1.QueryGetGroupRequest)
    - [QueryGetGroupResponse](#gauss.blindbox.v1.QueryGetGroupResponse)
    - [QueryGetGroupsRequest](#gauss.blindbox.v1.QueryGetGroupsRequest)
    - [QueryGetGroupsResponse](#gauss.blindbox.v1.QueryGetGroupsResponse)
    - [QueryGetPoolRequest](#gauss.blindbox.v1.QueryGetPoolRequest)
    - [QueryGetPoolResponse](#gauss.blindbox.v1.QueryGetPoolResponse)
    - [QueryGetPoolsRequest](#gauss.blindbox.v1.QueryGetPoolsRequest)
    - [QueryGetPoolsResponse](#gauss.blindbox.v1.QueryGetPoolsResponse)
    - [QueryParamsRequest](#gauss.blindbox.v1.QueryParamsRequest)
    - [QueryParamsResponse](#gauss.blindbox.v1.QueryParamsResponse)
  
    - [Query](#gauss.blindbox.v1.Query)
  
- [gauss/blindbox/v1/group.proto](#gauss/blindbox/v1/group.proto)
    - [Group](#gauss.blindbox.v1.Group)
    - [PoolIdToGroupId](#gauss.blindbox.v1.PoolIdToGroupId)
  
- [gauss/blindbox/v1/params.proto](#gauss/blindbox/v1/params.proto)
    - [Params](#gauss.blindbox.v1.Params)
  
- [gauss/blindbox/v1/pool.proto](#gauss/blindbox/v1/pool.proto)
    - [CreatorToPool](#gauss.blindbox.v1.CreatorToPool)
    - [Pool](#gauss.blindbox.v1.Pool)
  
- [gauss/blindbox/v1/tx.proto](#gauss/blindbox/v1/tx.proto)
    - [MsgCreateBox](#gauss.blindbox.v1.MsgCreateBox)
    - [MsgCreateBoxResponse](#gauss.blindbox.v1.MsgCreateBoxResponse)
    - [MsgCreatePool](#gauss.blindbox.v1.MsgCreatePool)
    - [MsgCreatePoolResponse](#gauss.blindbox.v1.MsgCreatePoolResponse)
    - [MsgOpenBox](#gauss.blindbox.v1.MsgOpenBox)
    - [MsgOpenBoxResponse](#gauss.blindbox.v1.MsgOpenBoxResponse)
    - [MsgRemovePool](#gauss.blindbox.v1.MsgRemovePool)
    - [MsgRemovePoolResponse](#gauss.blindbox.v1.MsgRemovePoolResponse)
    - [MsgRevokeBoxGroup](#gauss.blindbox.v1.MsgRevokeBoxGroup)
    - [MsgRevokeBoxGroupResponse](#gauss.blindbox.v1.MsgRevokeBoxGroupResponse)
    - [MsgUpdatePool](#gauss.blindbox.v1.MsgUpdatePool)
    - [MsgUpdatePoolResponse](#gauss.blindbox.v1.MsgUpdatePoolResponse)
  
    - [Msg](#gauss.blindbox.v1.Msg)
  
- [gauss/blindbox/v1/box.proto](#gauss/blindbox/v1/box.proto)
    - [Box](#gauss.blindbox.v1.Box)
  
- [gauss/blindbox/v1/genesis.proto](#gauss/blindbox/v1/genesis.proto)
    - [GenesisState](#gauss.blindbox.v1.GenesisState)
    - [GroupBox](#gauss.blindbox.v1.GroupBox)
  
- [gauss/farm/v1/query.proto](#gauss/farm/v1/query.proto)
    - [FarmPoolEntry](#gauss.farm.v1.FarmPoolEntry)
    - [LockedList](#gauss.farm.v1.LockedList)
    - [QueryFarmPoolRequest](#gauss.farm.v1.QueryFarmPoolRequest)
    - [QueryFarmPoolResponse](#gauss.farm.v1.QueryFarmPoolResponse)
    - [QueryFarmPoolsRequest](#gauss.farm.v1.QueryFarmPoolsRequest)
    - [QueryFarmPoolsResponse](#gauss.farm.v1.QueryFarmPoolsResponse)
    - [QueryFarmerRequest](#gauss.farm.v1.QueryFarmerRequest)
    - [QueryFarmerResponse](#gauss.farm.v1.QueryFarmerResponse)
    - [QueryParamsRequest](#gauss.farm.v1.QueryParamsRequest)
    - [QueryParamsResponse](#gauss.farm.v1.QueryParamsResponse)
  
    - [Query](#gauss.farm.v1.Query)
  
- [gauss/farm/v1/genesis.proto](#gauss/farm/v1/genesis.proto)
    - [GenesisState](#gauss.farm.v1.GenesisState)
  
- [gauss/farm/v1/tx.proto](#gauss/farm/v1/tx.proto)
    - [MsgCreatePool](#gauss.farm.v1.MsgCreatePool)
    - [MsgCreatePoolResponse](#gauss.farm.v1.MsgCreatePoolResponse)
    - [MsgDestroyPool](#gauss.farm.v1.MsgDestroyPool)
    - [MsgDestroyPoolResponse](#gauss.farm.v1.MsgDestroyPoolResponse)
    - [MsgStake](#gauss.farm.v1.MsgStake)
    - [MsgStakeResponse](#gauss.farm.v1.MsgStakeResponse)
    - [MsgUnbond](#gauss.farm.v1.MsgUnbond)
    - [MsgUnbondResponse](#gauss.farm.v1.MsgUnbondResponse)
    - [MsgWithdraw](#gauss.farm.v1.MsgWithdraw)
    - [MsgWithdrawResponse](#gauss.farm.v1.MsgWithdrawResponse)
  
    - [Msg](#gauss.farm.v1.Msg)
  
- [gauss/farm/v1/farm.proto](#gauss/farm/v1/farm.proto)
    - [FarmPool](#gauss.farm.v1.FarmPool)
    - [FarmRewardRule](#gauss.farm.v1.FarmRewardRule)
    - [Farmer](#gauss.farm.v1.Farmer)
    - [FarmerRewards](#gauss.farm.v1.FarmerRewards)
    - [FarmersRewardEntity](#gauss.farm.v1.FarmersRewardEntity)
    - [FarmersRewards](#gauss.farm.v1.FarmersRewards)
    - [Params](#gauss.farm.v1.Params)
  
- [gauss/fixedprice/v1/tx.proto](#gauss/fixedprice/v1/tx.proto)
    - [MsgBidOrder](#gauss.fixedprice.v1.MsgBidOrder)
    - [MsgBidOrderResponse](#gauss.fixedprice.v1.MsgBidOrderResponse)
    - [MsgCreateOrder](#gauss.fixedprice.v1.MsgCreateOrder)
    - [MsgCreateOrderResponse](#gauss.fixedprice.v1.MsgCreateOrderResponse)
    - [MsgDeleteOrder](#gauss.fixedprice.v1.MsgDeleteOrder)
    - [MsgDeleteOrderResponse](#gauss.fixedprice.v1.MsgDeleteOrderResponse)
    - [Order](#gauss.fixedprice.v1.Order)
    - [OrderClose](#gauss.fixedprice.v1.OrderClose)
    - [OrderPriceAdjust](#gauss.fixedprice.v1.OrderPriceAdjust)
    - [PoolAddress](#gauss.fixedprice.v1.PoolAddress)
  
    - [Msg](#gauss.fixedprice.v1.Msg)
  
- [gauss/fixedprice/v1/genesis.proto](#gauss/fixedprice/v1/genesis.proto)
    - [GenesisState](#gauss.fixedprice.v1.GenesisState)
    - [Params](#gauss.fixedprice.v1.Params)
  
- [gauss/fixedprice/v1/query.proto](#gauss/fixedprice/v1/query.proto)
    - [QueryOrderRequest](#gauss.fixedprice.v1.QueryOrderRequest)
    - [QueryOrderResponse](#gauss.fixedprice.v1.QueryOrderResponse)
    - [QueryOrdersRequest](#gauss.fixedprice.v1.QueryOrdersRequest)
    - [QueryOrdersResponse](#gauss.fixedprice.v1.QueryOrdersResponse)
  
    - [Query](#gauss.fixedprice.v1.Query)
  
- [gauss/nftexpool/v1/params.proto](#gauss/nftexpool/v1/params.proto)
    - [Params](#gauss.pool.v1.Params)
  
- [gauss/nftexpool/v1/tx.proto](#gauss/nftexpool/v1/tx.proto)
    - [Delegation](#gauss.pool.v1.Delegation)
    - [MsgCreatePool](#gauss.pool.v1.MsgCreatePool)
    - [MsgCreatePoolResponse](#gauss.pool.v1.MsgCreatePoolResponse)
    - [MsgDelegate](#gauss.pool.v1.MsgDelegate)
    - [MsgDelegateResponse](#gauss.pool.v1.MsgDelegateResponse)
    - [MsgUndelegate](#gauss.pool.v1.MsgUndelegate)
    - [MsgUndelegateResponse](#gauss.pool.v1.MsgUndelegateResponse)
    - [MsgUpdatePool](#gauss.pool.v1.MsgUpdatePool)
    - [MsgUpdatePoolResponse](#gauss.pool.v1.MsgUpdatePoolResponse)
    - [Pool](#gauss.pool.v1.Pool)
  
    - [Msg](#gauss.pool.v1.Msg)
  
- [gauss/nftexpool/v1/genesis.proto](#gauss/nftexpool/v1/genesis.proto)
    - [GenesisState](#gauss.pool.v1.GenesisState)
  
- [gauss/nftexpool/v1/query.proto](#gauss/nftexpool/v1/query.proto)
    - [QueryDelegateRequest](#gauss.pool.v1.QueryDelegateRequest)
    - [QueryDelegateResponse](#gauss.pool.v1.QueryDelegateResponse)
    - [QueryPoolsRequest](#gauss.pool.v1.QueryPoolsRequest)
    - [QueryPoolsResponse](#gauss.pool.v1.QueryPoolsResponse)
  
    - [Query](#gauss.pool.v1.Query)
  
- [gauss/token/v1/token.proto](#gauss/token/v1/token.proto)
    - [Params](#gauss.token.v1.Params)
    - [Token](#gauss.token.v1.Token)
  
- [gauss/token/v1/tx.proto](#gauss/token/v1/tx.proto)
    - [MsgBurnToken](#gauss.token.v1.MsgBurnToken)
    - [MsgBurnTokenResponse](#gauss.token.v1.MsgBurnTokenResponse)
    - [MsgEditToken](#gauss.token.v1.MsgEditToken)
    - [MsgEditTokenResponse](#gauss.token.v1.MsgEditTokenResponse)
    - [MsgIssueToken](#gauss.token.v1.MsgIssueToken)
    - [MsgIssueTokenResponse](#gauss.token.v1.MsgIssueTokenResponse)
    - [MsgMintToken](#gauss.token.v1.MsgMintToken)
    - [MsgMintTokenResponse](#gauss.token.v1.MsgMintTokenResponse)
    - [MsgTransferTokenOwner](#gauss.token.v1.MsgTransferTokenOwner)
    - [MsgTransferTokenOwnerResponse](#gauss.token.v1.MsgTransferTokenOwnerResponse)
    - [MsgUnlockToken](#gauss.token.v1.MsgUnlockToken)
    - [MsgUnlockTokenResponse](#gauss.token.v1.MsgUnlockTokenResponse)
  
    - [Msg](#gauss.token.v1.Msg)
  
- [gauss/token/v1/genesis.proto](#gauss/token/v1/genesis.proto)
    - [GenesisState](#gauss.token.v1.GenesisState)
  
- [gauss/token/v1/query.proto](#gauss/token/v1/query.proto)
    - [QueryBurntokenRequest](#gauss.token.v1.QueryBurntokenRequest)
    - [QueryBurntokenResponse](#gauss.token.v1.QueryBurntokenResponse)
    - [QueryFeesRequest](#gauss.token.v1.QueryFeesRequest)
    - [QueryFeesResponse](#gauss.token.v1.QueryFeesResponse)
    - [QueryParamsRequest](#gauss.token.v1.QueryParamsRequest)
    - [QueryParamsResponse](#gauss.token.v1.QueryParamsResponse)
    - [QueryTokenRequest](#gauss.token.v1.QueryTokenRequest)
    - [QueryTokenResponse](#gauss.token.v1.QueryTokenResponse)
    - [QueryTokensRequest](#gauss.token.v1.QueryTokensRequest)
    - [QueryTokensResponse](#gauss.token.v1.QueryTokensResponse)
  
    - [Query](#gauss.token.v1.Query)
  
- [gauss/validator-dao/v1/params.proto](#gauss/validator-dao/v1/params.proto)
    - [Params](#gauss.validatordao.v1.Params)
  
- [gauss/validator-dao/v1/query.proto](#gauss/validator-dao/v1/query.proto)
    - [QueryAuthorizerBizsRequest](#gauss.validatordao.v1.QueryAuthorizerBizsRequest)
    - [QueryAuthorizerBizsResponse](#gauss.validatordao.v1.QueryAuthorizerBizsResponse)
    - [QueryGranteeAuthBizsRequest](#gauss.validatordao.v1.QueryGranteeAuthBizsRequest)
    - [QueryGranteeAuthBizsResponse](#gauss.validatordao.v1.QueryGranteeAuthBizsResponse)
    - [QueryParamsRequest](#gauss.validatordao.v1.QueryParamsRequest)
    - [QueryParamsResponse](#gauss.validatordao.v1.QueryParamsResponse)
  
    - [Query](#gauss.validatordao.v1.Query)
  
- [gauss/validator-dao/v1/dao.proto](#gauss/validator-dao/v1/dao.proto)
    - [AcqBiz](#gauss.validatordao.v1.AcqBiz)
    - [AuthorizerBizs](#gauss.validatordao.v1.AuthorizerBizs)
    - [DaoBiz](#gauss.validatordao.v1.DaoBiz)
    - [GranteeBizs](#gauss.validatordao.v1.GranteeBizs)
  
- [gauss/validator-dao/v1/tx.proto](#gauss/validator-dao/v1/tx.proto)
    - [MsgAddAuthBiz](#gauss.validatordao.v1.MsgAddAuthBiz)
    - [MsgAddAuthBizResponse](#gauss.validatordao.v1.MsgAddAuthBizResponse)
    - [MsgRemoveAuthBiz](#gauss.validatordao.v1.MsgRemoveAuthBiz)
    - [MsgRemoveAuthBizResponse](#gauss.validatordao.v1.MsgRemoveAuthBizResponse)
    - [MsgReqAuthorization](#gauss.validatordao.v1.MsgReqAuthorization)
    - [MsgReqAuthorizationResponse](#gauss.validatordao.v1.MsgReqAuthorizationResponse)
    - [MsgUpdateAuthBiz](#gauss.validatordao.v1.MsgUpdateAuthBiz)
    - [MsgUpdateAuthBizResponse](#gauss.validatordao.v1.MsgUpdateAuthBizResponse)
    - [MsgWithdrawAuthorization](#gauss.validatordao.v1.MsgWithdrawAuthorization)
    - [MsgWithdrawAuthorizationResponse](#gauss.validatordao.v1.MsgWithdrawAuthorizationResponse)
  
    - [Msg](#gauss.validatordao.v1.Msg)
  
- [gauss/validator-dao/v1/genesis.proto](#gauss/validator-dao/v1/genesis.proto)
    - [GenesisState](#gauss.validatordao.v1.GenesisState)
  
- [gauss/nft/v1/genesis.proto](#gauss/nft/v1/genesis.proto)
    - [GenesisState](#gauss.nft.v1.GenesisState)
  
- [gauss/nft/v1/nft.proto](#gauss/nft/v1/nft.proto)
    - [LastPrice](#gauss.nft.v1.LastPrice)
    - [Params](#gauss.nft.v1.Params)
  
- [gauss/nft/v1/tx.proto](#gauss/nft/v1/tx.proto)
    - [Cate](#gauss.nft.v1.Cate)
    - [Collection](#gauss.nft.v1.Collection)
    - [Component](#gauss.nft.v1.Component)
    - [IDCollection](#gauss.nft.v1.IDCollection)
    - [MintBatchItem](#gauss.nft.v1.MintBatchItem)
    - [MsgFrozenNft](#gauss.nft.v1.MsgFrozenNft)
    - [MsgFrozenNftResponse](#gauss.nft.v1.MsgFrozenNftResponse)
    - [MsgMintBatch](#gauss.nft.v1.MsgMintBatch)
    - [MsgMintBatchResponse](#gauss.nft.v1.MsgMintBatchResponse)
    - [MsgMintNft](#gauss.nft.v1.MsgMintNft)
    - [MsgMintNftResponse](#gauss.nft.v1.MsgMintNftResponse)
    - [MsgTransferNft](#gauss.nft.v1.MsgTransferNft)
    - [MsgTransferNftResponse](#gauss.nft.v1.MsgTransferNftResponse)
    - [NfToken](#gauss.nft.v1.NfToken)
    - [Nft](#gauss.nft.v1.Nft)
    - [Owner](#gauss.nft.v1.Owner)
    - [TransferedCNFT](#gauss.nft.v1.TransferedCNFT)
  
    - [Msg](#gauss.nft.v1.Msg)
  
- [gauss/nft/v1/query.proto](#gauss/nft/v1/query.proto)
    - [QueryCollectionRequest](#gauss.nft.v1.QueryCollectionRequest)
    - [QueryCollectionResponse](#gauss.nft.v1.QueryCollectionResponse)
    - [QueryCollectionsRequest](#gauss.nft.v1.QueryCollectionsRequest)
    - [QueryCollectionsResponse](#gauss.nft.v1.QueryCollectionsResponse)
    - [QueryCreatorRequest](#gauss.nft.v1.QueryCreatorRequest)
    - [QueryCreatorResponse](#gauss.nft.v1.QueryCreatorResponse)
    - [QueryNftokenRequest](#gauss.nft.v1.QueryNftokenRequest)
    - [QueryNftokenResponse](#gauss.nft.v1.QueryNftokenResponse)
    - [QueryNftsByNameOrTokenRequest](#gauss.nft.v1.QueryNftsByNameOrTokenRequest)
    - [QueryNftsByNameOrTokenResponse](#gauss.nft.v1.QueryNftsByNameOrTokenResponse)
    - [QueryOwnerRequest](#gauss.nft.v1.QueryOwnerRequest)
    - [QueryOwnerResponse](#gauss.nft.v1.QueryOwnerResponse)
  
    - [Query](#gauss.nft.v1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="evmos/claims/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/claims/v1/genesis.proto



<a name="evmos.claims.v1.GenesisState"></a>

### GenesisState
GenesisState define the claims module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.claims.v1.Params) |  | params defines all the parameters of the module. |
| `claims_records` | [ClaimsRecordAddress](#evmos.claims.v1.ClaimsRecordAddress) | repeated | list of claim records with the corresponding airdrop recipient |






<a name="evmos.claims.v1.Params"></a>

### Params
Params defines the claims module's parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enable_claims` | [bool](#bool) |  | enable claiming process |
| `airdrop_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp of the airdrop start |
| `duration_until_decay` | [google.protobuf.Duration](#google.protobuf.Duration) |  | duration until decay of claimable tokens begin |
| `duration_of_decay` | [google.protobuf.Duration](#google.protobuf.Duration) |  | duration of the token claim decay period |
| `claims_denom` | [string](#string) |  | denom of claimable coin |
| `authorized_channels` | [string](#string) | repeated | list of authorized channel identifiers that can perform address attestations via IBC. |
| `evm_channels` | [string](#string) | repeated | list of channel identifiers from EVM compatible chains |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/claims/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/claims/v1/query.proto



<a name="evmos.claims.v1.QueryClaimsRecordRequest"></a>

### QueryClaimsRecordRequest
QueryClaimsRecordRequest is the request type for the Query/ClaimsRecord RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address defines the user to query claims record for |






<a name="evmos.claims.v1.QueryClaimsRecordResponse"></a>

### QueryClaimsRecordResponse
QueryClaimsRecordResponse is the response type for the Query/ClaimsRecord RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `initial_claimable_amount` | [string](#string) |  | total initial claimable amount for the user |
| `claims` | [Claim](#evmos.claims.v1.Claim) | repeated | the claims of the user |






<a name="evmos.claims.v1.QueryClaimsRecordsRequest"></a>

### QueryClaimsRecordsRequest
QueryClaimsRecordsRequest is the request type for the Query/ClaimsRecords RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="evmos.claims.v1.QueryClaimsRecordsResponse"></a>

### QueryClaimsRecordsResponse
QueryClaimsRecordsResponse is the response type for the Query/ClaimsRecords
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `claims` | [ClaimsRecordAddress](#evmos.claims.v1.ClaimsRecordAddress) | repeated | claims defines all claims records |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="evmos.claims.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="evmos.claims.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.claims.v1.Params) |  | params defines the parameters of the module. |






<a name="evmos.claims.v1.QueryTotalUnclaimedRequest"></a>

### QueryTotalUnclaimedRequest
QueryTotalUnclaimedRequest is the request type for the Query/TotalUnclaimed
RPC method.






<a name="evmos.claims.v1.QueryTotalUnclaimedResponse"></a>

### QueryTotalUnclaimedResponse
QueryTotalUnclaimedResponse is the response type for the Query/TotalUnclaimed
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | coins defines the unclaimed coins |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.claims.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `TotalUnclaimed` | [QueryTotalUnclaimedRequest](#evmos.claims.v1.QueryTotalUnclaimedRequest) | [QueryTotalUnclaimedResponse](#evmos.claims.v1.QueryTotalUnclaimedResponse) | TotalUnclaimed queries the total unclaimed tokens from the airdrop | GET|/evmos/claims/v1/total_unclaimed|
| `Params` | [QueryParamsRequest](#evmos.claims.v1.QueryParamsRequest) | [QueryParamsResponse](#evmos.claims.v1.QueryParamsResponse) | Params returns the claims module parameters | GET|/evmos/claims/v1/params|
| `ClaimsRecords` | [QueryClaimsRecordsRequest](#evmos.claims.v1.QueryClaimsRecordsRequest) | [QueryClaimsRecordsResponse](#evmos.claims.v1.QueryClaimsRecordsResponse) | ClaimsRecords returns all claims records | GET|/evmos/claims/v1/claims_records|
| `ClaimsRecord` | [QueryClaimsRecordRequest](#evmos.claims.v1.QueryClaimsRecordRequest) | [QueryClaimsRecordResponse](#evmos.claims.v1.QueryClaimsRecordResponse) | ClaimsRecord returns the claims record for a given address | GET|/evmos/claims/v1/claims_records/{address}|

 <!-- end services -->



<a name="evmos/claims/v1/claims.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/claims/v1/claims.proto



<a name="evmos.claims.v1.Claim"></a>

### Claim
Claim defines the action, completed flag and the remaining claimable amount
for a given user. This is only used during client queries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `action` | [Action](#evmos.claims.v1.Action) |  | action enum |
| `completed` | [bool](#bool) |  | true if the action has been completed |
| `claimable_amount` | [string](#string) |  | claimable token amount for the action. Zero if completed |






<a name="evmos.claims.v1.ClaimsRecord"></a>

### ClaimsRecord
ClaimsRecord defines the initial claimable airdrop amount and the list of
completed actions to claim the tokens.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `initial_claimable_amount` | [string](#string) |  | total initial claimable amount for the user |
| `actions_completed` | [bool](#bool) | repeated | slice of the available actions completed |






<a name="evmos.claims.v1.ClaimsRecordAddress"></a>

### ClaimsRecordAddress
ClaimsRecordAddress is the claims metadata per address that is used at
Genesis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | bech32 or hex address of claim user |
| `initial_claimable_amount` | [string](#string) |  | total initial claimable amount for the user |
| `actions_completed` | [bool](#bool) | repeated | slice of the available actions completed |





 <!-- end messages -->


<a name="evmos.claims.v1.Action"></a>

### Action
Action defines the list of available actions to claim the airdrop tokens.

| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_UNSPECIFIED | 0 | UNSPECIFIED defines an invalid action. |
| ACTION_VOTE | 1 | VOTE defines a proposal vote. |
| ACTION_DELEGATE | 2 | DELEGATE defines an staking delegation. |
| ACTION_EVM | 3 | EVM defines an EVM transaction. |
| ACTION_IBC_TRANSFER | 4 | IBC Transfer defines a fungible token transfer transaction via IBC. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/epochs/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/epochs/v1/query.proto



<a name="evmos.epochs.v1.QueryCurrentEpochRequest"></a>

### QueryCurrentEpochRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `identifier` | [string](#string) |  |  |






<a name="evmos.epochs.v1.QueryCurrentEpochResponse"></a>

### QueryCurrentEpochResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `current_epoch` | [int64](#int64) |  |  |






<a name="evmos.epochs.v1.QueryEpochsInfoRequest"></a>

### QueryEpochsInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="evmos.epochs.v1.QueryEpochsInfoResponse"></a>

### QueryEpochsInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `epochs` | [EpochInfo](#evmos.epochs.v1.EpochInfo) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.epochs.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `EpochInfos` | [QueryEpochsInfoRequest](#evmos.epochs.v1.QueryEpochsInfoRequest) | [QueryEpochsInfoResponse](#evmos.epochs.v1.QueryEpochsInfoResponse) | EpochInfos provide running epochInfos | GET|/evmos/epochs/v1/epochs|
| `CurrentEpoch` | [QueryCurrentEpochRequest](#evmos.epochs.v1.QueryCurrentEpochRequest) | [QueryCurrentEpochResponse](#evmos.epochs.v1.QueryCurrentEpochResponse) | CurrentEpoch provide current epoch of specified identifier | GET|/evmos/epochs/v1/current_epoch|

 <!-- end services -->



<a name="evmos/epochs/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/epochs/v1/genesis.proto



<a name="evmos.epochs.v1.EpochInfo"></a>

### EpochInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `identifier` | [string](#string) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `current_epoch` | [int64](#int64) |  |  |
| `current_epoch_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `epoch_counting_started` | [bool](#bool) |  |  |
| `current_epoch_start_height` | [int64](#int64) |  |  |






<a name="evmos.epochs.v1.GenesisState"></a>

### GenesisState
GenesisState defines the epochs module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `epochs` | [EpochInfo](#evmos.epochs.v1.EpochInfo) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/erc20/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/erc20/v1/genesis.proto



<a name="evmos.erc20.v1.GenesisState"></a>

### GenesisState
GenesisState defines the module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.erc20.v1.Params) |  | module parameters |
| `token_pairs` | [TokenPair](#evmos.erc20.v1.TokenPair) | repeated | registered token pairs |






<a name="evmos.erc20.v1.Params"></a>

### Params
Params defines the erc20 module params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enable_erc20` | [bool](#bool) |  | parameter to enable the conversion of Cosmos coins <--> ERC20 tokens. |
| `enable_evm_hook` | [bool](#bool) |  | parameter to enable the EVM hook that converts an ERC20 token to a Cosmos Coin by transferring the Tokens through a MsgEthereumTx to the ModuleAddress Ethereum address. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/erc20/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/erc20/v1/query.proto



<a name="evmos.erc20.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="evmos.erc20.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.erc20.v1.Params) |  |  |






<a name="evmos.erc20.v1.QueryTokenPairRequest"></a>

### QueryTokenPairRequest
QueryTokenPairRequest is the request type for the Query/TokenPair RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token` | [string](#string) |  | token identifier can be either the hex contract address of the ERC20 or the Cosmos base denomination |






<a name="evmos.erc20.v1.QueryTokenPairResponse"></a>

### QueryTokenPairResponse
QueryTokenPairResponse is the response type for the Query/TokenPair RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pair` | [TokenPair](#evmos.erc20.v1.TokenPair) |  |  |






<a name="evmos.erc20.v1.QueryTokenPairsRequest"></a>

### QueryTokenPairsRequest
QueryTokenPairsRequest is the request type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="evmos.erc20.v1.QueryTokenPairsResponse"></a>

### QueryTokenPairsResponse
QueryTokenPairsResponse is the response type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pairs` | [TokenPair](#evmos.erc20.v1.TokenPair) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.erc20.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `TokenPairs` | [QueryTokenPairsRequest](#evmos.erc20.v1.QueryTokenPairsRequest) | [QueryTokenPairsResponse](#evmos.erc20.v1.QueryTokenPairsResponse) | TokenPairs retrieves registered token pairs | GET|/evmos/erc20/v1/token_pairs|
| `TokenPair` | [QueryTokenPairRequest](#evmos.erc20.v1.QueryTokenPairRequest) | [QueryTokenPairResponse](#evmos.erc20.v1.QueryTokenPairResponse) | TokenPair retrieves a registered token pair | GET|/evmos/erc20/v1/token_pairs/{token}|
| `Params` | [QueryParamsRequest](#evmos.erc20.v1.QueryParamsRequest) | [QueryParamsResponse](#evmos.erc20.v1.QueryParamsResponse) | Params retrieves the erc20 module params | GET|/evmos/erc20/v1/params|

 <!-- end services -->



<a name="evmos/erc20/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/erc20/v1/tx.proto



<a name="evmos.erc20.v1.MsgConvertCoin"></a>

### MsgConvertCoin
MsgConvertCoin defines a Msg to convert a native Cosmos coin to a ERC20 token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Cosmos coin which denomination is registered in a token pair. The coin amount defines the amount of coins to convert. |
| `receiver` | [string](#string) |  | recipient hex address to receive ERC20 token |
| `sender` | [string](#string) |  | cosmos bech32 address from the owner of the given Cosmos coins |






<a name="evmos.erc20.v1.MsgConvertCoinResponse"></a>

### MsgConvertCoinResponse
MsgConvertCoinResponse returns no fields






<a name="evmos.erc20.v1.MsgConvertERC20"></a>

### MsgConvertERC20
MsgConvertERC20 defines a Msg to convert a ERC20 token to a native Cosmos
coin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | ERC20 token contract address registered in a token pair |
| `amount` | [string](#string) |  | amount of ERC20 tokens to convert |
| `receiver` | [string](#string) |  | bech32 address to receive native Cosmos coins |
| `sender` | [string](#string) |  | sender hex address from the owner of the given ERC20 tokens |






<a name="evmos.erc20.v1.MsgConvertERC20Response"></a>

### MsgConvertERC20Response
MsgConvertERC20Response returns no fields





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.erc20.v1.Msg"></a>

### Msg
Msg defines the erc20 Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ConvertCoin` | [MsgConvertCoin](#evmos.erc20.v1.MsgConvertCoin) | [MsgConvertCoinResponse](#evmos.erc20.v1.MsgConvertCoinResponse) | ConvertCoin mints a ERC20 representation of the native Cosmos coin denom that is registered on the token mapping. | GET|/evmos/erc20/v1/tx/convert_coin|
| `ConvertERC20` | [MsgConvertERC20](#evmos.erc20.v1.MsgConvertERC20) | [MsgConvertERC20Response](#evmos.erc20.v1.MsgConvertERC20Response) | ConvertERC20 mints a native Cosmos coin representation of the ERC20 token contract that is registered on the token mapping. | GET|/evmos/erc20/v1/tx/convert_erc20|

 <!-- end services -->



<a name="evmos/erc20/v1/erc20.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/erc20/v1/erc20.proto



<a name="evmos.erc20.v1.RegisterCoinProposal"></a>

### RegisterCoinProposal
RegisterCoinProposal is a gov Content type to register a token pair for a
native Cosmos coin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `metadata` | [cosmos.bank.v1beta1.Metadata](#cosmos.bank.v1beta1.Metadata) |  | metadata of the native Cosmos coin |






<a name="evmos.erc20.v1.RegisterERC20Proposal"></a>

### RegisterERC20Proposal
RegisterERC20Proposal is a gov Content type to register a token pair for an
ERC20 token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `erc20address` | [string](#string) |  | contract address of ERC20 token |






<a name="evmos.erc20.v1.ToggleTokenConversionProposal"></a>

### ToggleTokenConversionProposal
ToggleTokenConversionProposal is a gov Content type to toggle the conversion
of a token pair.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `token` | [string](#string) |  | token identifier can be either the hex contract address of the ERC20 or the Cosmos base denomination |






<a name="evmos.erc20.v1.TokenPair"></a>

### TokenPair
TokenPair defines an instance that records a pairing consisting of a native
 Cosmos Coin and an ERC20 token address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc20_address` | [string](#string) |  | address of ERC20 contract token |
| `denom` | [string](#string) |  | cosmos base denomination to be mapped to |
| `enabled` | [bool](#bool) |  | shows token mapping enable status |
| `contract_owner` | [Owner](#evmos.erc20.v1.Owner) |  | ERC20 owner address ENUM (0 invalid, 1 ModuleAccount, 2 external address) |





 <!-- end messages -->


<a name="evmos.erc20.v1.Owner"></a>

### Owner
Owner enumerates the ownership of a ERC20 contract.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OWNER_UNSPECIFIED | 0 | OWNER_UNSPECIFIED defines an invalid/undefined owner. |
| OWNER_MODULE | 1 | OWNER_MODULE erc20 is owned by the erc20 module account. |
| OWNER_EXTERNAL | 2 | EXTERNAL erc20 is owned by an external account. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/incentives/v1/incentives.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/incentives/v1/incentives.proto



<a name="evmos.incentives.v1.CancelIncentiveProposal"></a>

### CancelIncentiveProposal
CancelIncentiveProposal is a gov Content type to cancel an incentive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `contract` | [string](#string) |  | contract address |






<a name="evmos.incentives.v1.GasMeter"></a>

### GasMeter
GasMeter tracks the cumulative gas spent per participant in one epoch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | hex address of the incentivized contract |
| `participant` | [string](#string) |  | participant address that interacts with the incentive |
| `cumulative_gas` | [uint64](#uint64) |  | cumulative gas spent during the epoch |






<a name="evmos.incentives.v1.Incentive"></a>

### Incentive
Incentive defines an instance that organizes distribution conditions for a
given smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract address |
| `allocations` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | denoms and percentage of rewards to be allocated |
| `epochs` | [uint32](#uint32) |  | number of remaining epochs |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | distribution start time |
| `total_gas` | [uint64](#uint64) |  | cumulative gas spent by all gasmeters of the incentive during the epoch |






<a name="evmos.incentives.v1.RegisterIncentiveProposal"></a>

### RegisterIncentiveProposal
RegisterIncentiveProposal is a gov Content type to register an incentive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `contract` | [string](#string) |  | contract address |
| `allocations` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | denoms and percentage of rewards to be allocated |
| `epochs` | [uint32](#uint32) |  | number of remaining epochs |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/incentives/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/incentives/v1/query.proto



<a name="evmos.incentives.v1.QueryAllocationMeterRequest"></a>

### QueryAllocationMeterRequest
QueryAllocationMeterRequest is the request type for the Query/AllocationMeter
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom is the coin denom to query an allocation meter for. |






<a name="evmos.incentives.v1.QueryAllocationMeterResponse"></a>

### QueryAllocationMeterResponse
QueryAllocationMeterResponse is the response type for the
Query/AllocationMeter RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allocation_meter` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  |  |






<a name="evmos.incentives.v1.QueryAllocationMetersRequest"></a>

### QueryAllocationMetersRequest
QueryAllocationMetersRequest is the request type for the
Query/AllocationMeters RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="evmos.incentives.v1.QueryAllocationMetersResponse"></a>

### QueryAllocationMetersResponse
QueryAllocationMetersResponse is the response type for the
Query/AllocationMeters RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allocation_meters` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="evmos.incentives.v1.QueryGasMeterRequest"></a>

### QueryGasMeterRequest
QueryGasMeterRequest is the request type for the Query/Incentive RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract identifier is the hex contract address of a contract |
| `participant` | [string](#string) |  | participant identifier is the hex address of a user |






<a name="evmos.incentives.v1.QueryGasMeterResponse"></a>

### QueryGasMeterResponse
QueryGasMeterResponse is the response type for the Query/Incentive RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_meter` | [uint64](#uint64) |  |  |






<a name="evmos.incentives.v1.QueryGasMetersRequest"></a>

### QueryGasMetersRequest
QueryGasMetersRequest is the request type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract is the hex contract address of a incentivized smart contract |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="evmos.incentives.v1.QueryGasMetersResponse"></a>

### QueryGasMetersResponse
QueryGasMetersResponse is the response type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_meters` | [GasMeter](#evmos.incentives.v1.GasMeter) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="evmos.incentives.v1.QueryIncentiveRequest"></a>

### QueryIncentiveRequest
QueryIncentiveRequest is the request type for the Query/Incentive RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract identifier is the hex contract address of a contract |






<a name="evmos.incentives.v1.QueryIncentiveResponse"></a>

### QueryIncentiveResponse
QueryIncentiveResponse is the response type for the Query/Incentive RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive` | [Incentive](#evmos.incentives.v1.Incentive) |  |  |






<a name="evmos.incentives.v1.QueryIncentivesRequest"></a>

### QueryIncentivesRequest
QueryIncentivesRequest is the request type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="evmos.incentives.v1.QueryIncentivesResponse"></a>

### QueryIncentivesResponse
QueryIncentivesResponse is the response type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentives` | [Incentive](#evmos.incentives.v1.Incentive) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="evmos.incentives.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="evmos.incentives.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.incentives.v1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.incentives.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Incentives` | [QueryIncentivesRequest](#evmos.incentives.v1.QueryIncentivesRequest) | [QueryIncentivesResponse](#evmos.incentives.v1.QueryIncentivesResponse) | Incentives retrieves registered incentives | GET|/evmos/incentives/v1/incentives|
| `Incentive` | [QueryIncentiveRequest](#evmos.incentives.v1.QueryIncentiveRequest) | [QueryIncentiveResponse](#evmos.incentives.v1.QueryIncentiveResponse) | Incentive retrieves a registered incentive | GET|/evmos/incentives/v1/incentives/{contract}|
| `GasMeters` | [QueryGasMetersRequest](#evmos.incentives.v1.QueryGasMetersRequest) | [QueryGasMetersResponse](#evmos.incentives.v1.QueryGasMetersResponse) | GasMeters retrieves active gas meters for a given contract | GET|/evmos/incentives/v1/gas_meters/{contract}|
| `GasMeter` | [QueryGasMeterRequest](#evmos.incentives.v1.QueryGasMeterRequest) | [QueryGasMeterResponse](#evmos.incentives.v1.QueryGasMeterResponse) | GasMeter Retrieves a active gas meter | GET|/evmos/incentives/v1/gas_meters/{contract}/{participant}|
| `AllocationMeters` | [QueryAllocationMetersRequest](#evmos.incentives.v1.QueryAllocationMetersRequest) | [QueryAllocationMetersResponse](#evmos.incentives.v1.QueryAllocationMetersResponse) | AllocationMeters retrieves active allocation meters for a given denomination | GET|/evmos/incentives/v1/allocation_meters|
| `AllocationMeter` | [QueryAllocationMeterRequest](#evmos.incentives.v1.QueryAllocationMeterRequest) | [QueryAllocationMeterResponse](#evmos.incentives.v1.QueryAllocationMeterResponse) | AllocationMeter Retrieves a active gas meter | GET|/evmos/incentives/v1/allocation_meters/{denom}|
| `Params` | [QueryParamsRequest](#evmos.incentives.v1.QueryParamsRequest) | [QueryParamsResponse](#evmos.incentives.v1.QueryParamsResponse) | Params retrieves the incentives module params | GET|/evmos/incentives/v1/params|

 <!-- end services -->



<a name="evmos/incentives/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/incentives/v1/genesis.proto



<a name="evmos.incentives.v1.GenesisState"></a>

### GenesisState
GenesisState defines the module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.incentives.v1.Params) |  | module parameters |
| `incentives` | [Incentive](#evmos.incentives.v1.Incentive) | repeated | active incentives |
| `gas_meters` | [GasMeter](#evmos.incentives.v1.GasMeter) | repeated | active Gasmeters |






<a name="evmos.incentives.v1.Params"></a>

### Params
Params defines the incentives module params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enable_incentives` | [bool](#bool) |  | parameter to enable incentives |
| `allocation_limit` | [string](#string) |  | maximum percentage an incentive can allocate per denomination |
| `incentives_epoch_identifier` | [string](#string) |  | identifier for the epochs module hooks |
| `reward_scaler` | [string](#string) |  | scaling factor for capping rewards |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/inflation/v1/inflation.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/inflation/v1/inflation.proto



<a name="evmos.inflation.v1.ExponentialCalculation"></a>

### ExponentialCalculation
ExponentialCalculation holds factors to calculate exponential inflation on
each period. Calculation reference:
periodProvision = exponentialDecay       *  bondingIncentive
f(x)            = (a * (1 - r) ^ x + c)  *  (1 + max_variance - bondedRatio *
(max_variance / bonding_target))


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `a` | [string](#string) |  | initial value |
| `r` | [string](#string) |  | reduction factor |
| `c` | [string](#string) |  | long term inflation |
| `bonding_target` | [string](#string) |  | bonding target |
| `max_variance` | [string](#string) |  | max variance |






<a name="evmos.inflation.v1.InflationDistribution"></a>

### InflationDistribution
InflationDistribution defines the distribution in which inflation is
allocated through minting on each epoch (staking, incentives, community). It
excludes the team vesting distribution, as this is minted once at genesis.
The initial InflationDistribution can be calculated from the Evmos Token
Model like this:
mintDistribution1 = distribution1 / (1 - teamVestingDistribution)
0.5333333         = 40%           / (1 - 25%)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `staking_rewards` | [string](#string) |  | staking_rewards defines the proportion of the minted minted_denom that is to be allocated as staking rewards |
| `usage_incentives` | [string](#string) |  | usage_incentives defines the proportion of the minted minted_denom that is to be allocated to the incentives module address |
| `community_pool` | [string](#string) |  | community_pool defines the proportion of the minted minted_denom that is to be allocated to the community pool |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/inflation/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/inflation/v1/query.proto



<a name="evmos.inflation.v1.QueryCirculatingSupplyRequest"></a>

### QueryCirculatingSupplyRequest
QueryCirculatingSupplyRequest is the request type for the
Query/CirculatingSupply RPC method.






<a name="evmos.inflation.v1.QueryCirculatingSupplyResponse"></a>

### QueryCirculatingSupplyResponse
QueryCirculatingSupplyResponse is the response type for the
Query/CirculatingSupply RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `circulating_supply` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | total amount of coins in circulation |






<a name="evmos.inflation.v1.QueryEpochMintProvisionRequest"></a>

### QueryEpochMintProvisionRequest
QueryEpochMintProvisionRequest is the request type for the
Query/EpochMintProvision RPC method.






<a name="evmos.inflation.v1.QueryEpochMintProvisionResponse"></a>

### QueryEpochMintProvisionResponse
QueryEpochMintProvisionResponse is the response type for the
Query/EpochMintProvision RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `epoch_mint_provision` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | epoch_mint_provision is the current minting per epoch provision value. |






<a name="evmos.inflation.v1.QueryInflationRateRequest"></a>

### QueryInflationRateRequest
QueryInflationRateRequest is the request type for the Query/InflationRate RPC
method.






<a name="evmos.inflation.v1.QueryInflationRateResponse"></a>

### QueryInflationRateResponse
QueryInflationRateResponse is the response type for the Query/InflationRate
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation_rate` | [string](#string) |  | rate by which the total supply increases within one period |






<a name="evmos.inflation.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="evmos.inflation.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.inflation.v1.Params) |  | params defines the parameters of the module. |






<a name="evmos.inflation.v1.QueryPeriodRequest"></a>

### QueryPeriodRequest
QueryPeriodRequest is the request type for the Query/Period RPC method.






<a name="evmos.inflation.v1.QueryPeriodResponse"></a>

### QueryPeriodResponse
QueryPeriodResponse is the response type for the Query/Period RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `period` | [uint64](#uint64) |  | period is the current minting per epoch provision value. |






<a name="evmos.inflation.v1.QuerySkippedEpochsRequest"></a>

### QuerySkippedEpochsRequest
QuerySkippedEpochsRequest is the request type for the Query/SkippedEpochs RPC
method.






<a name="evmos.inflation.v1.QuerySkippedEpochsResponse"></a>

### QuerySkippedEpochsResponse
QuerySkippedEpochsResponse is the response type for the Query/SkippedEpochs
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `skipped_epochs` | [uint64](#uint64) |  | number of epochs that the inflation module has been disabled. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.inflation.v1.Query"></a>

### Query
Query provides defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Period` | [QueryPeriodRequest](#evmos.inflation.v1.QueryPeriodRequest) | [QueryPeriodResponse](#evmos.inflation.v1.QueryPeriodResponse) | Period retrieves current period. | GET|/evmos/inflation/v1/period|
| `EpochMintProvision` | [QueryEpochMintProvisionRequest](#evmos.inflation.v1.QueryEpochMintProvisionRequest) | [QueryEpochMintProvisionResponse](#evmos.inflation.v1.QueryEpochMintProvisionResponse) | EpochMintProvision retrieves current minting epoch provision value. | GET|/evmos/inflation/v1/epoch_mint_provision|
| `SkippedEpochs` | [QuerySkippedEpochsRequest](#evmos.inflation.v1.QuerySkippedEpochsRequest) | [QuerySkippedEpochsResponse](#evmos.inflation.v1.QuerySkippedEpochsResponse) | SkippedEpochs retrieves the total number of skipped epochs. | GET|/evmos/inflation/v1/skipped_epochs|
| `CirculatingSupply` | [QueryCirculatingSupplyRequest](#evmos.inflation.v1.QueryCirculatingSupplyRequest) | [QueryCirculatingSupplyResponse](#evmos.inflation.v1.QueryCirculatingSupplyResponse) | CirculatingSupply retrieves the total number of tokens that are in circulation (i.e. excluding unvested tokens). | GET|/evmos/inflation/v1/circulating_supply|
| `InflationRate` | [QueryInflationRateRequest](#evmos.inflation.v1.QueryInflationRateRequest) | [QueryInflationRateResponse](#evmos.inflation.v1.QueryInflationRateResponse) | InflationRate retrieves the inflation rate of the current period. | GET|/evmos/inflation/v1/inflation_rate|
| `Params` | [QueryParamsRequest](#evmos.inflation.v1.QueryParamsRequest) | [QueryParamsResponse](#evmos.inflation.v1.QueryParamsResponse) | Params retrieves the total set of minting parameters. | GET|/evmos/inflation/v1/params|

 <!-- end services -->



<a name="evmos/inflation/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/inflation/v1/genesis.proto



<a name="evmos.inflation.v1.GenesisState"></a>

### GenesisState
GenesisState defines the inflation module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.inflation.v1.Params) |  | params defines all the paramaters of the module. |
| `period` | [uint64](#uint64) |  | amount of past periods, based on the epochs per period param |
| `epoch_identifier` | [string](#string) |  | inflation epoch identifier |
| `epochs_per_period` | [int64](#int64) |  | number of epochs after which inflation is recalculated |
| `skipped_epochs` | [uint64](#uint64) |  | number of epochs that have passed while inflation is disabled |






<a name="evmos.inflation.v1.Params"></a>

### Params
Params holds parameters for the inflation module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_denom` | [string](#string) |  | type of coin to mint |
| `exponential_calculation` | [ExponentialCalculation](#evmos.inflation.v1.ExponentialCalculation) |  | variables to calculate exponential inflation |
| `inflation_distribution` | [InflationDistribution](#evmos.inflation.v1.InflationDistribution) |  | inflation distribution of the minted denom |
| `enable_inflation` | [bool](#bool) |  | parameter to enable inflation and halt increasing the skipped_epochs |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/recovery/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/recovery/v1/query.proto



<a name="evmos.recovery.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="evmos.recovery.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.recovery.v1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.recovery.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#evmos.recovery.v1.QueryParamsRequest) | [QueryParamsResponse](#evmos.recovery.v1.QueryParamsResponse) | Params retrieves the total set of recovery parameters. | GET|/evmos/recovery/v1/params|

 <!-- end services -->



<a name="evmos/recovery/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/recovery/v1/genesis.proto



<a name="evmos.recovery.v1.GenesisState"></a>

### GenesisState
GenesisState defines the recovery module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#evmos.recovery.v1.Params) |  | params defines all the paramaters of the module. |






<a name="evmos.recovery.v1.Params"></a>

### Params
Params holds parameters for the recovery module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enable_recovery` | [bool](#bool) |  | enable recovery IBC middleware |
| `packet_timeout_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | duration added to timeout timestamp for balances recovered via IBC packets |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/vesting/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/vesting/v1/tx.proto



<a name="evmos.vesting.v1.MsgClawback"></a>

### MsgClawback
MsgClawback defines a message that removes unvested tokens from a
ClawbackVestingAccount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `funder_address` | [string](#string) |  | funder_address is the address which funded the account |
| `account_address` | [string](#string) |  | account_address is the address of the ClawbackVestingAccount to claw back from. |
| `dest_address` | [string](#string) |  | dest_address specifies where the clawed-back tokens should be transferred to. If empty, the tokens will be transferred back to the original funder of the account. |






<a name="evmos.vesting.v1.MsgClawbackResponse"></a>

### MsgClawbackResponse
MsgClawbackResponse defines the MsgClawback response type.






<a name="evmos.vesting.v1.MsgCreateClawbackVestingAccount"></a>

### MsgCreateClawbackVestingAccount
MsgCreateClawbackVestingAccount defines a message that enables creating a
ClawbackVestingAccount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  | from_address specifies the account to provide the funds and sign the clawback request |
| `to_address` | [string](#string) |  | to_address specifies the account to receive the funds |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_time defines the time at which the vesting period begins |
| `lockup_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | lockup_periods defines the unlocking schedule relative to the start_time |
| `vesting_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | vesting_periods defines thevesting schedule relative to the start_time |
| `merge` | [bool](#bool) |  | merge specifies a the creation mechanism for existing ClawbackVestingAccounts. If true, merge this new grant into an existing ClawbackVestingAccount, or create it if it does not exist. If false, creates a new account. New grants to an existing account must be from the same from_address. |






<a name="evmos.vesting.v1.MsgCreateClawbackVestingAccountResponse"></a>

### MsgCreateClawbackVestingAccountResponse
MsgCreateClawbackVestingAccountResponse defines the
MsgCreateClawbackVestingAccount response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.vesting.v1.Msg"></a>

### Msg
Msg defines the vesting Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateClawbackVestingAccount` | [MsgCreateClawbackVestingAccount](#evmos.vesting.v1.MsgCreateClawbackVestingAccount) | [MsgCreateClawbackVestingAccountResponse](#evmos.vesting.v1.MsgCreateClawbackVestingAccountResponse) | CreateClawbackVestingAccount creats a vesting account that is subject to clawback and the configuration of vesting and lockup schedules. | GET|/evmos/vesting/v1/tx/create_clawback_vesting_account|
| `Clawback` | [MsgClawback](#evmos.vesting.v1.MsgClawback) | [MsgClawbackResponse](#evmos.vesting.v1.MsgClawbackResponse) | Clawback removes the unvested tokens from a ClawbackVestingAccount. | GET|/evmos/vesting/v1/tx/clawback|

 <!-- end services -->



<a name="evmos/vesting/v1/vesting.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/vesting/v1/vesting.proto



<a name="evmos.vesting.v1.ClawbackVestingAccount"></a>

### ClawbackVestingAccount
ClawbackVestingAccount implements the VestingAccount interface. It provides
an account that can hold contributions subject to "lockup" (like a
PeriodicVestingAccount), or vesting which is subject to clawback
of unvested tokens, or a combination (tokens vest, but are still locked).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_vesting_account` | [cosmos.vesting.v1beta1.BaseVestingAccount](#cosmos.vesting.v1beta1.BaseVestingAccount) |  | base_vesting_account implements the VestingAccount interface. It contains all the necessary fields needed for any vesting account implementation |
| `funder_address` | [string](#string) |  | funder_address specifies the account which can perform clawback |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_time defines the time at which the vesting period begins |
| `lockup_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | lockup_periods defines the unlocking schedule relative to the start_time |
| `vesting_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | vesting_periods defines the vesting schedule relative to the start_time |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="evmos/vesting/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## evmos/vesting/v1/query.proto



<a name="evmos.vesting.v1.QueryBalancesRequest"></a>

### QueryBalancesRequest
QueryBalancesRequest is the request type for the Query/Balances RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address of the clawback vesting account |






<a name="evmos.vesting.v1.QueryBalancesResponse"></a>

### QueryBalancesResponse
QueryBalancesResponse is the response type for the Query/Balances RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `locked` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | current amount of locked tokens |
| `unvested` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | current amount of unvested tokens |
| `vested` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | current amount of vested tokens |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="evmos.vesting.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Balances` | [QueryBalancesRequest](#evmos.vesting.v1.QueryBalancesRequest) | [QueryBalancesResponse](#evmos.vesting.v1.QueryBalancesResponse) | Retrieves the unvested, vested and locked tokens for a vesting account | GET|/evmos/vesting/v1/balances/{address}|

 <!-- end services -->



<a name="gauss/auction/v1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/auction/v1/params.proto



<a name="gauss.auction.v1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auto_agree_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | 自动成交间隔 |
| `orders` | [Order](#gauss.auction.v1.Order) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/auction/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/auction/v1/tx.proto



<a name="gauss.auction.v1.MsgBidOrder"></a>

### MsgBidOrder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `pool_address` | [string](#string) |  |  |






<a name="gauss.auction.v1.MsgBidOrderResponse"></a>

### MsgBidOrderResponse







<a name="gauss.auction.v1.MsgCreateOrder"></a>

### MsgCreateOrder
create order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `start_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `min_step` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `auto_agree_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="gauss.auction.v1.MsgCreateOrderResponse"></a>

### MsgCreateOrderResponse







<a name="gauss.auction.v1.MsgDeleteOrder"></a>

### MsgDeleteOrder
delete order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |






<a name="gauss.auction.v1.MsgDeleteOrderResponse"></a>

### MsgDeleteOrderResponse







<a name="gauss.auction.v1.Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bid_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `nft` | [gauss.nft.v1.Nft](#gauss.nft.v1.Nft) |  |  |
| `bidder` | [string](#string) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `min_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `min_step` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `pool_address` | [string](#string) |  |  |
| `auto_agree_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="gauss.auction.v1.PoolAddress"></a>

### PoolAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.auction.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateOrder` | [MsgCreateOrder](#gauss.auction.v1.MsgCreateOrder) | [MsgCreateOrderResponse](#gauss.auction.v1.MsgCreateOrderResponse) |  | |
| `DeleteOrder` | [MsgDeleteOrder](#gauss.auction.v1.MsgDeleteOrder) | [MsgDeleteOrderResponse](#gauss.auction.v1.MsgDeleteOrderResponse) |  | |
| `BidOrder` | [MsgBidOrder](#gauss.auction.v1.MsgBidOrder) | [MsgBidOrderResponse](#gauss.auction.v1.MsgBidOrderResponse) |  | |

 <!-- end services -->



<a name="gauss/auction/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/auction/v1/genesis.proto



<a name="gauss.auction.v1.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.auction.v1.Params) |  | params defines all the parameters for the nft module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/auction/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/auction/v1/query.proto



<a name="gauss.auction.v1.QueryOrderRequest"></a>

### QueryOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |






<a name="gauss.auction.v1.QueryOrderResponse"></a>

### QueryOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `order` | [Order](#gauss.auction.v1.Order) |  |  |






<a name="gauss.auction.v1.QueryOrdersRequest"></a>

### QueryOrdersRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_address` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.auction.v1.QueryOrdersResponse"></a>

### QueryOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `orders` | [Order](#gauss.auction.v1.Order) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.auction.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Orders` | [QueryOrdersRequest](#gauss.auction.v1.QueryOrdersRequest) | [QueryOrdersResponse](#gauss.auction.v1.QueryOrdersResponse) |  | GET|/icplaza/auction/orders|
| `Order` | [QueryOrderRequest](#gauss.auction.v1.QueryOrderRequest) | [QueryOrderResponse](#gauss.auction.v1.QueryOrderResponse) |  | GET|/icplaza/auction/order/{token_id}|

 <!-- end services -->



<a name="gauss/blindbox/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/blindbox/v1/query.proto



<a name="gauss.blindbox.v1.QueryGetBoxRequest"></a>

### QueryGetBoxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [string](#string) |  |  |
| `box_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.QueryGetBoxResponse"></a>

### QueryGetBoxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `box` | [Box](#gauss.blindbox.v1.Box) |  |  |
| `group` | [Group](#gauss.blindbox.v1.Group) |  |  |






<a name="gauss.blindbox.v1.QueryGetGroupRequest"></a>

### QueryGetGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.QueryGetGroupResponse"></a>

### QueryGetGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group` | [Group](#gauss.blindbox.v1.Group) |  |  |
| `boxes` | [Box](#gauss.blindbox.v1.Box) | repeated |  |






<a name="gauss.blindbox.v1.QueryGetGroupsRequest"></a>

### QueryGetGroupsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.blindbox.v1.QueryGetGroupsResponse"></a>

### QueryGetGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `groups` | [QueryGetGroupResponse](#gauss.blindbox.v1.QueryGetGroupResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gauss.blindbox.v1.QueryGetPoolRequest"></a>

### QueryGetPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.QueryGetPoolResponse"></a>

### QueryGetPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [Pool](#gauss.blindbox.v1.Pool) |  |  |






<a name="gauss.blindbox.v1.QueryGetPoolsRequest"></a>

### QueryGetPoolsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.blindbox.v1.QueryGetPoolsResponse"></a>

### QueryGetPoolsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pools` | [Pool](#gauss.blindbox.v1.Pool) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gauss.blindbox.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="gauss.blindbox.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.blindbox.v1.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.blindbox.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gauss.blindbox.v1.QueryParamsRequest) | [QueryParamsResponse](#gauss.blindbox.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/icplaza/blindbox/params|
| `GetPools` | [QueryGetPoolsRequest](#gauss.blindbox.v1.QueryGetPoolsRequest) | [QueryGetPoolsResponse](#gauss.blindbox.v1.QueryGetPoolsResponse) | Queries a list of getpool_ids items. | GET|/icplaza/blindbox/getPools|
| `GetPool` | [QueryGetPoolRequest](#gauss.blindbox.v1.QueryGetPoolRequest) | [QueryGetPoolResponse](#gauss.blindbox.v1.QueryGetPoolResponse) | Queries a list of GetPool items. | GET|/icplaza/blindbox/getPool/{pool_id}|
| `GetGroups` | [QueryGetGroupsRequest](#gauss.blindbox.v1.QueryGetGroupsRequest) | [QueryGetGroupsResponse](#gauss.blindbox.v1.QueryGetGroupsResponse) | Queries a list of getgroup_ids items. | GET|/icplaza/blindbox/getGroups/{pool_id}|
| `GetGroup` | [QueryGetGroupRequest](#gauss.blindbox.v1.QueryGetGroupRequest) | [QueryGetGroupResponse](#gauss.blindbox.v1.QueryGetGroupResponse) | Queries a list of GetGroup items. | GET|/icplaza/blindbox/getGroup/{group_id}|
| `GetBox` | [QueryGetBoxRequest](#gauss.blindbox.v1.QueryGetBoxRequest) | [QueryGetBoxResponse](#gauss.blindbox.v1.QueryGetBoxResponse) | Queries a list of GetBox items. | GET|/icplaza/blindbox/getBox/{group_id}/{box_id}|

 <!-- end services -->



<a name="gauss/blindbox/v1/group.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/blindbox/v1/group.proto



<a name="gauss.blindbox.v1.Group"></a>

### Group



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [string](#string) |  |  |
| `box_size` | [uint64](#uint64) |  |  |
| `box_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `open_size` | [uint64](#uint64) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `creator` | [string](#string) |  |  |
| `pool_id` | [string](#string) |  |  |
| `random_min` | [uint64](#uint64) |  |  |
| `random_max` | [uint64](#uint64) |  |  |
| `random_nfts` | [string](#string) | repeated |  |
| `fixed_nfts` | [string](#string) | repeated |  |
| `left_random_nfts` | [string](#string) | repeated |  |
| `left_fixed_nfts` | [string](#string) | repeated |  |






<a name="gauss.blindbox.v1.PoolIdToGroupId"></a>

### PoolIdToGroupId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/blindbox/v1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/blindbox/v1/params.proto



<a name="gauss.blindbox.v1.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/blindbox/v1/pool.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/blindbox/v1/pool.proto



<a name="gauss.blindbox.v1.CreatorToPool"></a>

### CreatorToPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_ids` | [string](#string) | repeated |  |






<a name="gauss.blindbox.v1.Pool"></a>

### Pool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [string](#string) |  |  |
| `fee_rate` | [string](#string) |  |  |
| `fee_address` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/blindbox/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/blindbox/v1/tx.proto



<a name="gauss.blindbox.v1.MsgCreateBox"></a>

### MsgCreateBox



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `group_id` | [string](#string) |  |  |
| `box_size` | [uint64](#uint64) |  |  |
| `box_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `random_min` | [uint64](#uint64) |  |  |
| `random_max` | [uint64](#uint64) |  |  |
| `random_nfts` | [string](#string) | repeated |  |
| `fixed_nfts` | [string](#string) | repeated |  |
| `pool_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.MsgCreateBoxResponse"></a>

### MsgCreateBoxResponse







<a name="gauss.blindbox.v1.MsgCreatePool"></a>

### MsgCreatePool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `fee_rate` | [string](#string) |  |  |
| `fee_address` | [string](#string) |  |  |
| `pool_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.MsgCreatePoolResponse"></a>

### MsgCreatePoolResponse







<a name="gauss.blindbox.v1.MsgOpenBox"></a>

### MsgOpenBox



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `group_id` | [string](#string) |  |  |
| `box_ids` | [string](#string) | repeated |  |






<a name="gauss.blindbox.v1.MsgOpenBoxResponse"></a>

### MsgOpenBoxResponse







<a name="gauss.blindbox.v1.MsgRemovePool"></a>

### MsgRemovePool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `pool_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.MsgRemovePoolResponse"></a>

### MsgRemovePoolResponse







<a name="gauss.blindbox.v1.MsgRevokeBoxGroup"></a>

### MsgRevokeBoxGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `group_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.MsgRevokeBoxGroupResponse"></a>

### MsgRevokeBoxGroupResponse







<a name="gauss.blindbox.v1.MsgUpdatePool"></a>

### MsgUpdatePool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `fee_rate` | [string](#string) |  |  |
| `fee_address` | [string](#string) |  |  |
| `pool_id` | [string](#string) |  |  |






<a name="gauss.blindbox.v1.MsgUpdatePoolResponse"></a>

### MsgUpdatePoolResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.blindbox.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreatePool` | [MsgCreatePool](#gauss.blindbox.v1.MsgCreatePool) | [MsgCreatePoolResponse](#gauss.blindbox.v1.MsgCreatePoolResponse) |  | |
| `UpdatePool` | [MsgUpdatePool](#gauss.blindbox.v1.MsgUpdatePool) | [MsgUpdatePoolResponse](#gauss.blindbox.v1.MsgUpdatePoolResponse) |  | |
| `RemovePool` | [MsgRemovePool](#gauss.blindbox.v1.MsgRemovePool) | [MsgRemovePoolResponse](#gauss.blindbox.v1.MsgRemovePoolResponse) |  | |
| `CreateBox` | [MsgCreateBox](#gauss.blindbox.v1.MsgCreateBox) | [MsgCreateBoxResponse](#gauss.blindbox.v1.MsgCreateBoxResponse) |  | |
| `RevokeBoxGroup` | [MsgRevokeBoxGroup](#gauss.blindbox.v1.MsgRevokeBoxGroup) | [MsgRevokeBoxGroupResponse](#gauss.blindbox.v1.MsgRevokeBoxGroupResponse) |  | |
| `OpenBox` | [MsgOpenBox](#gauss.blindbox.v1.MsgOpenBox) | [MsgOpenBoxResponse](#gauss.blindbox.v1.MsgOpenBoxResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="gauss/blindbox/v1/box.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/blindbox/v1/box.proto



<a name="gauss.blindbox.v1.Box"></a>

### Box



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `box_id` | [string](#string) |  |  |
| `opened` | [bool](#bool) |  |  |
| `group_id` | [string](#string) |  |  |
| `nfts` | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/blindbox/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/blindbox/v1/genesis.proto



<a name="gauss.blindbox.v1.GenesisState"></a>

### GenesisState
GenesisState defines the blindbox module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.blindbox.v1.Params) |  |  |
| `pools` | [Pool](#gauss.blindbox.v1.Pool) | repeated |  |
| `group_boxes` | [GroupBox](#gauss.blindbox.v1.GroupBox) | repeated | this line is used by starport scaffolding # genesis/proto/state |






<a name="gauss.blindbox.v1.GroupBox"></a>

### GroupBox



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group` | [Group](#gauss.blindbox.v1.Group) |  |  |
| `boxes` | [Box](#gauss.blindbox.v1.Box) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/farm/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/farm/v1/query.proto



<a name="gauss.farm.v1.FarmPoolEntry"></a>

### FarmPoolEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [FarmPool](#gauss.farm.v1.FarmPool) |  |  |
| `remaining_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gauss.farm.v1.LockedList"></a>

### LockedList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `farmer` | [Farmer](#gauss.farm.v1.Farmer) |  |  |
| `pending_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gauss.farm.v1.QueryFarmPoolRequest"></a>

### QueryFarmPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |






<a name="gauss.farm.v1.QueryFarmPoolResponse"></a>

### QueryFarmPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [FarmPoolEntry](#gauss.farm.v1.FarmPoolEntry) |  |  |






<a name="gauss.farm.v1.QueryFarmPoolsRequest"></a>

### QueryFarmPoolsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.farm.v1.QueryFarmPoolsResponse"></a>

### QueryFarmPoolsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pools` | [FarmPoolEntry](#gauss.farm.v1.FarmPoolEntry) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gauss.farm.v1.QueryFarmerRequest"></a>

### QueryFarmerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `farmer` | [string](#string) |  |  |
| `pool_name` | [string](#string) |  |  |






<a name="gauss.farm.v1.QueryFarmerResponse"></a>

### QueryFarmerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `list` | [LockedList](#gauss.farm.v1.LockedList) | repeated |  |
| `height` | [int64](#int64) |  |  |






<a name="gauss.farm.v1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="gauss.farm.v1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.farm.v1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.farm.v1.Query"></a>

### Query


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `FarmPools` | [QueryFarmPoolsRequest](#gauss.farm.v1.QueryFarmPoolsRequest) | [QueryFarmPoolsResponse](#gauss.farm.v1.QueryFarmPoolsResponse) |  | GET|/icplaza/farm/pools|
| `FarmPool` | [QueryFarmPoolRequest](#gauss.farm.v1.QueryFarmPoolRequest) | [QueryFarmPoolResponse](#gauss.farm.v1.QueryFarmPoolResponse) |  | GET|/icplaza/farm/pool/{name}|
| `Farmer` | [QueryFarmerRequest](#gauss.farm.v1.QueryFarmerRequest) | [QueryFarmerResponse](#gauss.farm.v1.QueryFarmerResponse) |  | GET|/icplaza/farm/farmer/{farmer}|
| `Params` | [QueryParamsRequest](#gauss.farm.v1.QueryParamsRequest) | [QueryParamsResponse](#gauss.farm.v1.QueryParamsResponse) | Params queries the htlc parameters | GET|/icplaza/farm/params|

 <!-- end services -->



<a name="gauss/farm/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/farm/v1/genesis.proto



<a name="gauss.farm.v1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.farm.v1.Params) |  |  |
| `farm_pools` | [FarmPool](#gauss.farm.v1.FarmPool) | repeated |  |
| `farmers` | [Farmer](#gauss.farm.v1.Farmer) | repeated |  |
| `farmers_rewards` | [FarmersRewards](#gauss.farm.v1.FarmersRewards) | repeated |  |
| `farmer_rewards` | [FarmerRewards](#gauss.farm.v1.FarmerRewards) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/farm/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/farm/v1/tx.proto



<a name="gauss.farm.v1.MsgCreatePool"></a>

### MsgCreatePool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `start_height` | [int64](#int64) |  |  |
| `min_staking` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `rewards_per_block` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `total_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `creator` | [string](#string) |  |  |






<a name="gauss.farm.v1.MsgCreatePoolResponse"></a>

### MsgCreatePoolResponse







<a name="gauss.farm.v1.MsgDestroyPool"></a>

### MsgDestroyPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_name` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |






<a name="gauss.farm.v1.MsgDestroyPoolResponse"></a>

### MsgDestroyPoolResponse







<a name="gauss.farm.v1.MsgStake"></a>

### MsgStake



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_name` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `sender` | [string](#string) |  |  |






<a name="gauss.farm.v1.MsgStakeResponse"></a>

### MsgStakeResponse







<a name="gauss.farm.v1.MsgUnbond"></a>

### MsgUnbond



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_name` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `sender` | [string](#string) |  |  |






<a name="gauss.farm.v1.MsgUnbondResponse"></a>

### MsgUnbondResponse







<a name="gauss.farm.v1.MsgWithdraw"></a>

### MsgWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_name` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |






<a name="gauss.farm.v1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.farm.v1.Msg"></a>

### Msg
Msg defines the farm Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreatePool` | [MsgCreatePool](#gauss.farm.v1.MsgCreatePool) | [MsgCreatePoolResponse](#gauss.farm.v1.MsgCreatePoolResponse) | CreatePool defines a method for creating a new farm pool | |
| `DestroyPool` | [MsgDestroyPool](#gauss.farm.v1.MsgDestroyPool) | [MsgDestroyPoolResponse](#gauss.farm.v1.MsgDestroyPoolResponse) | DestroyPool defines a method for destroying a existed farm pool | |
| `Stake` | [MsgStake](#gauss.farm.v1.MsgStake) | [MsgStakeResponse](#gauss.farm.v1.MsgStakeResponse) | Stake defines a method for staking some lp token to a farm pool | |
| `Unbond` | [MsgUnbond](#gauss.farm.v1.MsgUnbond) | [MsgUnbondResponse](#gauss.farm.v1.MsgUnbondResponse) | Unbond defines a method for unstaking some lp token from a farm pool and withdraw some reward | |
| `Withdraw` | [MsgWithdraw](#gauss.farm.v1.MsgWithdraw) | [MsgWithdrawResponse](#gauss.farm.v1.MsgWithdrawResponse) | Withdraw defines a method withdraw some reward from a farm pool | |

 <!-- end services -->



<a name="gauss/farm/v1/farm.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/farm/v1/farm.proto



<a name="gauss.farm.v1.FarmPool"></a>

### FarmPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `start_height` | [int64](#int64) |  |  |
| `end_height` | [int64](#int64) |  |  |
| `min_staking` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `tokens` | [string](#string) |  |  |
| `farm_reward_rules` | [FarmRewardRule](#gauss.farm.v1.FarmRewardRule) | repeated |  |
| `farmer_count` | [uint64](#uint64) |  |  |






<a name="gauss.farm.v1.FarmRewardRule"></a>

### FarmRewardRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `rewards_per_block` | [string](#string) |  |  |






<a name="gauss.farm.v1.Farmer"></a>

### Farmer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `farmer_address` | [string](#string) |  |  |
| `tokens` | [string](#string) |  |  |
| `pool_name` | [string](#string) |  |  |






<a name="gauss.farm.v1.FarmerRewards"></a>

### FarmerRewards



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `debts` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `farmer_address` | [string](#string) |  |  |
| `pool_name` | [string](#string) |  |  |






<a name="gauss.farm.v1.FarmersRewardEntity"></a>

### FarmersRewardEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `remaining_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `rewards_per_share` | [string](#string) |  |  |






<a name="gauss.farm.v1.FarmersRewards"></a>

### FarmersRewards



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `previous_height` | [int64](#int64) |  |  |
| `farmers_rewards` | [FarmersRewardEntity](#gauss.farm.v1.FarmersRewardEntity) | repeated |  |
| `pool_name` | [string](#string) |  |  |






<a name="gauss.farm.v1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creating_pool_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_reward_categories` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/fixedprice/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/fixedprice/v1/tx.proto



<a name="gauss.fixedprice.v1.MsgBidOrder"></a>

### MsgBidOrder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `pool_address` | [string](#string) |  |  |






<a name="gauss.fixedprice.v1.MsgBidOrderResponse"></a>

### MsgBidOrderResponse







<a name="gauss.fixedprice.v1.MsgCreateOrder"></a>

### MsgCreateOrder
create order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `start_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `end_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="gauss.fixedprice.v1.MsgCreateOrderResponse"></a>

### MsgCreateOrderResponse







<a name="gauss.fixedprice.v1.MsgDeleteOrder"></a>

### MsgDeleteOrder
delete order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |






<a name="gauss.fixedprice.v1.MsgDeleteOrderResponse"></a>

### MsgDeleteOrderResponse







<a name="gauss.fixedprice.v1.Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `start_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `end_price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `nft` | [gauss.nft.v1.Nft](#gauss.nft.v1.Nft) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `next_time_adjust_price` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `pool_address` | [string](#string) |  |  |






<a name="gauss.fixedprice.v1.OrderClose"></a>

### OrderClose



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |






<a name="gauss.fixedprice.v1.OrderPriceAdjust"></a>

### OrderPriceAdjust



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.fixedprice.v1.PoolAddress"></a>

### PoolAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.fixedprice.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateOrder` | [MsgCreateOrder](#gauss.fixedprice.v1.MsgCreateOrder) | [MsgCreateOrderResponse](#gauss.fixedprice.v1.MsgCreateOrderResponse) |  | |
| `DeleteOrder` | [MsgDeleteOrder](#gauss.fixedprice.v1.MsgDeleteOrder) | [MsgDeleteOrderResponse](#gauss.fixedprice.v1.MsgDeleteOrderResponse) |  | |
| `BidOrder` | [MsgBidOrder](#gauss.fixedprice.v1.MsgBidOrder) | [MsgBidOrderResponse](#gauss.fixedprice.v1.MsgBidOrderResponse) |  | |

 <!-- end services -->



<a name="gauss/fixedprice/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/fixedprice/v1/genesis.proto



<a name="gauss.fixedprice.v1.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.fixedprice.v1.Params) |  | params defines all the parameters for the nft module. |






<a name="gauss.fixedprice.v1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `adjust_price_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | 自动成交间隔 |
| `orders` | [Order](#gauss.fixedprice.v1.Order) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/fixedprice/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/fixedprice/v1/query.proto



<a name="gauss.fixedprice.v1.QueryOrderRequest"></a>

### QueryOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |






<a name="gauss.fixedprice.v1.QueryOrderResponse"></a>

### QueryOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `order` | [Order](#gauss.fixedprice.v1.Order) |  |  |






<a name="gauss.fixedprice.v1.QueryOrdersRequest"></a>

### QueryOrdersRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_address` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.fixedprice.v1.QueryOrdersResponse"></a>

### QueryOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `orders` | [Order](#gauss.fixedprice.v1.Order) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.fixedprice.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Orders` | [QueryOrdersRequest](#gauss.fixedprice.v1.QueryOrdersRequest) | [QueryOrdersResponse](#gauss.fixedprice.v1.QueryOrdersResponse) |  | GET|/icplaza/fixedprice/orders|
| `Order` | [QueryOrderRequest](#gauss.fixedprice.v1.QueryOrderRequest) | [QueryOrderResponse](#gauss.fixedprice.v1.QueryOrderResponse) |  | GET|/icplaza/fixedprice/order/{token_id}|

 <!-- end services -->



<a name="gauss/nftexpool/v1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nftexpool/v1/params.proto



<a name="gauss.pool.v1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_creation_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | create fee |
| `burn_percent` | [string](#string) |  | fee allot |
| `community_percent` | [string](#string) |  |  |
| `validators_percent` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/nftexpool/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nftexpool/v1/tx.proto



<a name="gauss.pool.v1.Delegation"></a>

### Delegation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gauss.pool.v1.MsgCreatePool"></a>

### MsgCreatePool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `commission_rate` | [string](#string) |  |  |
| `commission_address` | [string](#string) |  |  |
| `value_added_tax_address` | [string](#string) |  |  |






<a name="gauss.pool.v1.MsgCreatePoolResponse"></a>

### MsgCreatePoolResponse







<a name="gauss.pool.v1.MsgDelegate"></a>

### MsgDelegate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.pool.v1.MsgDelegateResponse"></a>

### MsgDelegateResponse







<a name="gauss.pool.v1.MsgUndelegate"></a>

### MsgUndelegate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `delegator` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.pool.v1.MsgUndelegateResponse"></a>

### MsgUndelegateResponse







<a name="gauss.pool.v1.MsgUpdatePool"></a>

### MsgUpdatePool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `commission_rate` | [string](#string) |  |  |
| `commission_address` | [string](#string) |  |  |
| `value_added_tax_address` | [string](#string) |  |  |






<a name="gauss.pool.v1.MsgUpdatePoolResponse"></a>

### MsgUpdatePoolResponse







<a name="gauss.pool.v1.Pool"></a>

### Pool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |
| `commission_rate` | [string](#string) |  |  |
| `commission_address` | [string](#string) |  |  |
| `value_added_tax_address` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.pool.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Undelegate` | [MsgUndelegate](#gauss.pool.v1.MsgUndelegate) | [MsgUndelegateResponse](#gauss.pool.v1.MsgUndelegateResponse) |  | |
| `Delegate` | [MsgDelegate](#gauss.pool.v1.MsgDelegate) | [MsgDelegateResponse](#gauss.pool.v1.MsgDelegateResponse) |  | |
| `UpdatePool` | [MsgUpdatePool](#gauss.pool.v1.MsgUpdatePool) | [MsgUpdatePoolResponse](#gauss.pool.v1.MsgUpdatePoolResponse) |  | |
| `CreatePool` | [MsgCreatePool](#gauss.pool.v1.MsgCreatePool) | [MsgCreatePoolResponse](#gauss.pool.v1.MsgCreatePoolResponse) |  | |

 <!-- end services -->



<a name="gauss/nftexpool/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nftexpool/v1/genesis.proto



<a name="gauss.pool.v1.GenesisState"></a>

### GenesisState
GenesisState defines the pool module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.pool.v1.Params) |  | params defines all the parameters for the nft module. |
| `pools` | [Pool](#gauss.pool.v1.Pool) | repeated |  |
| `delegations` | [Delegation](#gauss.pool.v1.Delegation) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/nftexpool/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nftexpool/v1/query.proto



<a name="gauss.pool.v1.QueryDelegateRequest"></a>

### QueryDelegateRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator` | [string](#string) |  |  |
| `pool_address` | [string](#string) |  |  |






<a name="gauss.pool.v1.QueryDelegateResponse"></a>

### QueryDelegateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="gauss.pool.v1.QueryPoolsRequest"></a>

### QueryPoolsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.pool.v1.QueryPoolsResponse"></a>

### QueryPoolsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pools` | [Pool](#gauss.pool.v1.Pool) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.pool.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `QDelegate` | [QueryDelegateRequest](#gauss.pool.v1.QueryDelegateRequest) | [QueryDelegateResponse](#gauss.pool.v1.QueryDelegateResponse) |  | GET|/icplaza/pool/delegate/{delegator}/{pool_address}|
| `QPools` | [QueryPoolsRequest](#gauss.pool.v1.QueryPoolsRequest) | [QueryPoolsResponse](#gauss.pool.v1.QueryPoolsResponse) |  | GET|/icplaza/pool/pools|

 <!-- end services -->



<a name="gauss/token/v1/token.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/token/v1/token.proto



<a name="gauss.token.v1.Params"></a>

### Params
Params defines token module's parameters


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_tax` | [string](#string) |  |  |
| `issue_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_fee_ratio` | [string](#string) |  |  |






<a name="gauss.token.v1.Token"></a>

### Token
Token defines a standard for the fungible token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `smallest_unit` | [string](#string) |  |  |
| `decimals` | [uint32](#uint32) |  |  |
| `initial_supply` | [uint64](#uint64) |  |  |
| `total_supply` | [uint64](#uint64) |  |  |
| `mintable` | [bool](#bool) |  |  |
| `owner` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/token/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/token/v1/tx.proto



<a name="gauss.token.v1.MsgBurnToken"></a>

### MsgBurnToken
MsgBurnToken defines an SDK message for burning some tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |






<a name="gauss.token.v1.MsgBurnTokenResponse"></a>

### MsgBurnTokenResponse
MsgBurnTokenResponse defines the Msg/BurnToken response type






<a name="gauss.token.v1.MsgEditToken"></a>

### MsgEditToken
MsgEditToken defines an SDK message for editing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `mintable` | [bool](#bool) |  |  |
| `owner` | [string](#string) |  |  |






<a name="gauss.token.v1.MsgEditTokenResponse"></a>

### MsgEditTokenResponse
MsgEditTokenResponse defines the Msg/EditToken response type






<a name="gauss.token.v1.MsgIssueToken"></a>

### MsgIssueToken
MsgIssueToken defines an SDK message for issuing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `smallest_unit` | [string](#string) |  |  |
| `decimals` | [uint32](#uint32) |  |  |
| `initial_supply` | [uint64](#uint64) |  |  |
| `total_supply` | [uint64](#uint64) |  |  |
| `mintable` | [bool](#bool) |  |  |
| `unlocked` | [bool](#bool) |  |  |
| `owner` | [string](#string) |  |  |






<a name="gauss.token.v1.MsgIssueTokenResponse"></a>

### MsgIssueTokenResponse
MsgIssueTokenResponse defines the Msg/IssueToken response type






<a name="gauss.token.v1.MsgMintToken"></a>

### MsgMintToken
MsgMintToken defines an SDK message for minting a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |
| `to` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="gauss.token.v1.MsgMintTokenResponse"></a>

### MsgMintTokenResponse
MsgMintTokenResponse defines the Msg/MintToken response type






<a name="gauss.token.v1.MsgTransferTokenOwner"></a>

### MsgTransferTokenOwner
MsgTransferTokenOwner defines an SDK message for transferring the token owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `old_owner` | [string](#string) |  |  |
| `new_owner` | [string](#string) |  |  |






<a name="gauss.token.v1.MsgTransferTokenOwnerResponse"></a>

### MsgTransferTokenOwnerResponse
MsgTransferTokenOwnerResponse defines the Msg/TransferTokenOwner response type






<a name="gauss.token.v1.MsgUnlockToken"></a>

### MsgUnlockToken
MsgUnlockToken defines an SDK message for locking the token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="gauss.token.v1.MsgUnlockTokenResponse"></a>

### MsgUnlockTokenResponse
MsgUnlockTokenResponse defines the Msg/UnlockToken response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.token.v1.Msg"></a>

### Msg
Msg defines the oracle Msg service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `IssueToken` | [MsgIssueToken](#gauss.token.v1.MsgIssueToken) | [MsgIssueTokenResponse](#gauss.token.v1.MsgIssueTokenResponse) | IssueToken defines a method for issuing a new token | |
| `EditToken` | [MsgEditToken](#gauss.token.v1.MsgEditToken) | [MsgEditTokenResponse](#gauss.token.v1.MsgEditTokenResponse) | EditToken defines a method for editing a token | |
| `MintToken` | [MsgMintToken](#gauss.token.v1.MsgMintToken) | [MsgMintTokenResponse](#gauss.token.v1.MsgMintTokenResponse) | MintToken defines a method for minting some tokens | |
| `BurnToken` | [MsgBurnToken](#gauss.token.v1.MsgBurnToken) | [MsgBurnTokenResponse](#gauss.token.v1.MsgBurnTokenResponse) | BurnToken defines a method for burning some tokens | |
| `UnlockToken` | [MsgUnlockToken](#gauss.token.v1.MsgUnlockToken) | [MsgUnlockTokenResponse](#gauss.token.v1.MsgUnlockTokenResponse) | UnLockToken defines a method for burning some tokens | |
| `TransferTokenOwner` | [MsgTransferTokenOwner](#gauss.token.v1.MsgTransferTokenOwner) | [MsgTransferTokenOwnerResponse](#gauss.token.v1.MsgTransferTokenOwnerResponse) | TransferTokenOwner defines a method for minting some tokens | |

 <!-- end services -->



<a name="gauss/token/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/token/v1/genesis.proto



<a name="gauss.token.v1.GenesisState"></a>

### GenesisState
GenesisState defines the token module's genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.token.v1.Params) |  |  |
| `tokens` | [Token](#gauss.token.v1.Token) | repeated |  |
| `burned_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/token/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/token/v1/query.proto



<a name="gauss.token.v1.QueryBurntokenRequest"></a>

### QueryBurntokenRequest
QueryBurntokenRequest is request type for the Query/Token RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |






<a name="gauss.token.v1.QueryBurntokenResponse"></a>

### QueryBurntokenResponse
QueryBurntokenResponse is response type for the Query/Tokens RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exist` | [bool](#bool) |  |  |
| `burned_coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.token.v1.QueryFeesRequest"></a>

### QueryFeesRequest
QueryFeesRequest is request type for the Query/Token RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |






<a name="gauss.token.v1.QueryFeesResponse"></a>

### QueryFeesResponse
QueryFeesResponse is response type for the Query/Fees RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exist` | [bool](#bool) |  |  |
| `issue_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.token.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParametersRequest is request type for the Query/Parameters RPC method






<a name="gauss.token.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParametersResponse is response type for the Query/Parameters RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.token.v1.Params) |  |  |






<a name="gauss.token.v1.QueryTokenRequest"></a>

### QueryTokenRequest
QueryTokenRequest is request type for the Query/Token RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |






<a name="gauss.token.v1.QueryTokenResponse"></a>

### QueryTokenResponse
QueryTokenResponse is response type for the Query/Token RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Token` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `unlocked` | [bool](#bool) |  |  |






<a name="gauss.token.v1.QueryTokensRequest"></a>

### QueryTokensRequest
QueryTokensRequest is request type for the Query/Tokens RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="gauss.token.v1.QueryTokensResponse"></a>

### QueryTokensResponse
QueryTokensResponse is response type for the Query/Tokens RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Tokens` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.token.v1.Query"></a>

### Query
Query creates service with token as RPC

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gauss.token.v1.QueryParamsRequest) | [QueryParamsResponse](#gauss.token.v1.QueryParamsResponse) | Params queries the token parameters | GET|/icplaza/token/params|
| `Tokens` | [QueryTokensRequest](#gauss.token.v1.QueryTokensRequest) | [QueryTokensResponse](#gauss.token.v1.QueryTokensResponse) | Tokens returns the token list | GET|/icplaza/token/tokens|
| `Token` | [QueryTokenRequest](#gauss.token.v1.QueryTokenRequest) | [QueryTokenResponse](#gauss.token.v1.QueryTokenResponse) | Token returns token with token name | GET|/icplaza/token/tokens/{symbol}|
| `Fees` | [QueryFeesRequest](#gauss.token.v1.QueryFeesRequest) | [QueryFeesResponse](#gauss.token.v1.QueryFeesResponse) | Fees returns the fees to issue or mint a token | GET|/icplaza/token/tokens/{symbol}/fees|
| `Burntoken` | [QueryBurntokenRequest](#gauss.token.v1.QueryBurntokenRequest) | [QueryBurntokenResponse](#gauss.token.v1.QueryBurntokenResponse) | BurntToken queries the burnt coins | GET|/icplaza/token/{symbol}/burnt|

 <!-- end services -->



<a name="gauss/validator-dao/v1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/validator-dao/v1/params.proto



<a name="gauss.validatordao.v1.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auth_bizs` | [DaoBiz](#gauss.validatordao.v1.DaoBiz) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/validator-dao/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/validator-dao/v1/query.proto



<a name="gauss.validatordao.v1.QueryAuthorizerBizsRequest"></a>

### QueryAuthorizerBizsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authorizer` | [string](#string) |  |  |






<a name="gauss.validatordao.v1.QueryAuthorizerBizsResponse"></a>

### QueryAuthorizerBizsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bizs` | [DaoBiz](#gauss.validatordao.v1.DaoBiz) | repeated |  |






<a name="gauss.validatordao.v1.QueryGranteeAuthBizsRequest"></a>

### QueryGranteeAuthBizsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grantee` | [string](#string) |  |  |






<a name="gauss.validatordao.v1.QueryGranteeAuthBizsResponse"></a>

### QueryGranteeAuthBizsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bizs` | [AcqBiz](#gauss.validatordao.v1.AcqBiz) | repeated |  |






<a name="gauss.validatordao.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="gauss.validatordao.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.validatordao.v1.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.validatordao.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gauss.validatordao.v1.QueryParamsRequest) | [QueryParamsResponse](#gauss.validatordao.v1.QueryParamsResponse) |  | GET|/icplaza/validatordao/params|
| `AuthorizerBizs` | [QueryAuthorizerBizsRequest](#gauss.validatordao.v1.QueryAuthorizerBizsRequest) | [QueryAuthorizerBizsResponse](#gauss.validatordao.v1.QueryAuthorizerBizsResponse) | Queries a list of BizList items. 按验证人查询业务 | GET|/icplaza/validatordao/auth_bizs/authorizer/{authorizer}|
| `GranteeAuthBizs` | [QueryGranteeAuthBizsRequest](#gauss.validatordao.v1.QueryGranteeAuthBizsRequest) | [QueryGranteeAuthBizsResponse](#gauss.validatordao.v1.QueryGranteeAuthBizsResponse) |  | GET|/icplaza/validatordao/auth_bizs/grantee/{grantee}|

 <!-- end services -->



<a name="gauss/validator-dao/v1/dao.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/validator-dao/v1/dao.proto



<a name="gauss.validatordao.v1.AcqBiz"></a>

### AcqBiz



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `biz_name` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `price` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.validatordao.v1.AuthorizerBizs"></a>

### AuthorizerBizs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authorizer` | [string](#string) |  |  |
| `bizs` | [DaoBiz](#gauss.validatordao.v1.DaoBiz) | repeated |  |






<a name="gauss.validatordao.v1.DaoBiz"></a>

### DaoBiz



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.validatordao.v1.GranteeBizs"></a>

### GranteeBizs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grantee` | [string](#string) |  |  |
| `bizs` | [AcqBiz](#gauss.validatordao.v1.AcqBiz) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/validator-dao/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/validator-dao/v1/tx.proto



<a name="gauss.validatordao.v1.MsgAddAuthBiz"></a>

### MsgAddAuthBiz



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `biz_name` | [string](#string) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.validatordao.v1.MsgAddAuthBizResponse"></a>

### MsgAddAuthBizResponse







<a name="gauss.validatordao.v1.MsgRemoveAuthBiz"></a>

### MsgRemoveAuthBiz



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `biz_name` | [string](#string) |  |  |






<a name="gauss.validatordao.v1.MsgRemoveAuthBizResponse"></a>

### MsgRemoveAuthBizResponse







<a name="gauss.validatordao.v1.MsgReqAuthorization"></a>

### MsgReqAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `authorizer` | [string](#string) |  |  |
| `biz_name` | [string](#string) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.validatordao.v1.MsgReqAuthorizationResponse"></a>

### MsgReqAuthorizationResponse







<a name="gauss.validatordao.v1.MsgUpdateAuthBiz"></a>

### MsgUpdateAuthBiz



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `biz_name` | [string](#string) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.validatordao.v1.MsgUpdateAuthBizResponse"></a>

### MsgUpdateAuthBizResponse







<a name="gauss.validatordao.v1.MsgWithdrawAuthorization"></a>

### MsgWithdrawAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `authorizer` | [string](#string) |  |  |
| `biz_name` | [string](#string) |  |  |






<a name="gauss.validatordao.v1.MsgWithdrawAuthorizationResponse"></a>

### MsgWithdrawAuthorizationResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.validatordao.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ReqAuthorization` | [MsgReqAuthorization](#gauss.validatordao.v1.MsgReqAuthorization) | [MsgReqAuthorizationResponse](#gauss.validatordao.v1.MsgReqAuthorizationResponse) |  | |
| `WithdrawAuthorization` | [MsgWithdrawAuthorization](#gauss.validatordao.v1.MsgWithdrawAuthorization) | [MsgWithdrawAuthorizationResponse](#gauss.validatordao.v1.MsgWithdrawAuthorizationResponse) |  | |
| `AddAuthBiz` | [MsgAddAuthBiz](#gauss.validatordao.v1.MsgAddAuthBiz) | [MsgAddAuthBizResponse](#gauss.validatordao.v1.MsgAddAuthBizResponse) |  | |
| `UpdateAuthBiz` | [MsgUpdateAuthBiz](#gauss.validatordao.v1.MsgUpdateAuthBiz) | [MsgUpdateAuthBizResponse](#gauss.validatordao.v1.MsgUpdateAuthBizResponse) |  | |
| `RemoveAuthBiz` | [MsgRemoveAuthBiz](#gauss.validatordao.v1.MsgRemoveAuthBiz) | [MsgRemoveAuthBizResponse](#gauss.validatordao.v1.MsgRemoveAuthBizResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="gauss/validator-dao/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/validator-dao/v1/genesis.proto



<a name="gauss.validatordao.v1.GenesisState"></a>

### GenesisState
GenesisState defines the validatordao module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.validatordao.v1.Params) |  |  |
| `authorizer_bizs` | [AuthorizerBizs](#gauss.validatordao.v1.AuthorizerBizs) | repeated |  |
| `grantee_auth_bizs` | [GranteeBizs](#gauss.validatordao.v1.GranteeBizs) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/nft/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nft/v1/genesis.proto



<a name="gauss.nft.v1.GenesisState"></a>

### GenesisState
GenesisState defines the nft module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gauss.nft.v1.Params) |  | params defines all the parameters for the nft module. |
| `last_prices` | [LastPrice](#gauss.nft.v1.LastPrice) | repeated |  |
| `nfts` | [Nft](#gauss.nft.v1.Nft) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/nft/v1/nft.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nft/v1/nft.proto



<a name="gauss.nft.v1.LastPrice"></a>

### LastPrice
最后成交价格
message LastPrice {
   string token_id = 1;
   repeated cosmos.base.v1beta1.Coin coins = 2;
}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |
| `coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gauss.nft.v1.Params"></a>

### Params
Params defines the parameters for the liquidity module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_creation_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Fee paid to create a Liquidity Pool. Set a fee to prevent spamming. |
| `burn_percent` | [string](#string) |  | fee allot |
| `community_percent` | [string](#string) |  |  |
| `validators_percent` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gauss/nft/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nft/v1/tx.proto



<a name="gauss.nft.v1.Cate"></a>

### Cate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cate_id` | [string](#string) |  |  |






<a name="gauss.nft.v1.Collection"></a>

### Collection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cate_id` | [string](#string) |  |  |
| `nfts` | [Nft](#gauss.nft.v1.Nft) | repeated |  |






<a name="gauss.nft.v1.Component"></a>

### Component



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `min_amount` | [uint64](#uint64) |  |  |
| `max_amount` | [uint64](#uint64) |  |  |
| `type` | [uint64](#uint64) |  |  |






<a name="gauss.nft.v1.IDCollection"></a>

### IDCollection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cate_id` | [string](#string) |  |  |
| `ids` | [string](#string) | repeated |  |






<a name="gauss.nft.v1.MintBatchItem"></a>

### MintBatchItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |
| `cate_id` | [string](#string) |  |  |
| `token_uri` | [string](#string) |  |  |
| `company_uri` | [string](#string) |  |  |
| `value_added_tax` | [string](#string) |  |  |
| `copyright_tax` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `components` | [Component](#gauss.nft.v1.Component) | repeated |  |






<a name="gauss.nft.v1.MsgFrozenNft"></a>

### MsgFrozenNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `cate_id` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |






<a name="gauss.nft.v1.MsgFrozenNftResponse"></a>

### MsgFrozenNftResponse







<a name="gauss.nft.v1.MsgMintBatch"></a>

### MsgMintBatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |
| `items` | [MintBatchItem](#gauss.nft.v1.MintBatchItem) | repeated |  |






<a name="gauss.nft.v1.MsgMintBatchResponse"></a>

### MsgMintBatchResponse







<a name="gauss.nft.v1.MsgMintNft"></a>

### MsgMintNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `cate_id` | [string](#string) |  |  |
| `token_uri` | [string](#string) |  |  |
| `company_uri` | [string](#string) |  |  |
| `value_added_tax` | [string](#string) |  |  |
| `copyright_tax` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `components` | [Component](#gauss.nft.v1.Component) | repeated |  |






<a name="gauss.nft.v1.MsgMintNftResponse"></a>

### MsgMintNftResponse







<a name="gauss.nft.v1.MsgTransferNft"></a>

### MsgTransferNft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |
| `cate_id` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |






<a name="gauss.nft.v1.MsgTransferNftResponse"></a>

### MsgTransferNftResponse







<a name="gauss.nft.v1.NfToken"></a>

### NfToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |






<a name="gauss.nft.v1.Nft"></a>

### Nft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |
| `cate_id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |
| `company_uri` | [string](#string) |  |  |
| `token_uri` | [string](#string) |  |  |
| `value_added_tax` | [string](#string) |  |  |
| `copyright_tax` | [string](#string) |  |  |
| `token_status` | [uint32](#uint32) |  |  |
| `name` | [string](#string) |  |  |
| `components` | [Component](#gauss.nft.v1.Component) | repeated |  |






<a name="gauss.nft.v1.Owner"></a>

### Owner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `id_collections` | [IDCollection](#gauss.nft.v1.IDCollection) | repeated |  |






<a name="gauss.nft.v1.TransferedCNFT"></a>

### TransferedCNFT
TransferedCNFT


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |
| `type` | [uint64](#uint64) |  | 类型 0:普通转移;1:拍卖或者固态价格 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.nft.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MintBatch` | [MsgMintBatch](#gauss.nft.v1.MsgMintBatch) | [MsgMintBatchResponse](#gauss.nft.v1.MsgMintBatchResponse) | this line is used by starport scaffolding # proto/tx/rpc | |
| `FrozenNft` | [MsgFrozenNft](#gauss.nft.v1.MsgFrozenNft) | [MsgFrozenNftResponse](#gauss.nft.v1.MsgFrozenNftResponse) |  | |
| `TransferNft` | [MsgTransferNft](#gauss.nft.v1.MsgTransferNft) | [MsgTransferNftResponse](#gauss.nft.v1.MsgTransferNftResponse) |  | |
| `MintNft` | [MsgMintNft](#gauss.nft.v1.MsgMintNft) | [MsgMintNftResponse](#gauss.nft.v1.MsgMintNftResponse) |  | |

 <!-- end services -->



<a name="gauss/nft/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gauss/nft/v1/query.proto



<a name="gauss.nft.v1.QueryCollectionRequest"></a>

### QueryCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cate_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.nft.v1.QueryCollectionResponse"></a>

### QueryCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nfts` | [Nft](#gauss.nft.v1.Nft) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gauss.nft.v1.QueryCollectionsRequest"></a>

### QueryCollectionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.nft.v1.QueryCollectionsResponse"></a>

### QueryCollectionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nfts` | [Nft](#gauss.nft.v1.Nft) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gauss.nft.v1.QueryCreatorRequest"></a>

### QueryCreatorRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cate_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.nft.v1.QueryCreatorResponse"></a>

### QueryCreatorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nfts` | [Nft](#gauss.nft.v1.Nft) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gauss.nft.v1.QueryNftokenRequest"></a>

### QueryNftokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_id` | [string](#string) |  |  |






<a name="gauss.nft.v1.QueryNftokenResponse"></a>

### QueryNftokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft` | [Nft](#gauss.nft.v1.Nft) |  |  |






<a name="gauss.nft.v1.QueryNftsByNameOrTokenRequest"></a>

### QueryNftsByNameOrTokenRequest
search by name or tokenId


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.nft.v1.QueryNftsByNameOrTokenResponse"></a>

### QueryNftsByNameOrTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nfts` | [Nft](#gauss.nft.v1.Nft) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="gauss.nft.v1.QueryOwnerRequest"></a>

### QueryOwnerRequest
QueryOwnerRequest is the request type for the Query/Owner RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `cate_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="gauss.nft.v1.QueryOwnerResponse"></a>

### QueryOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nfts` | [Nft](#gauss.nft.v1.Nft) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gauss.nft.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Creator` | [QueryCreatorRequest](#gauss.nft.v1.QueryCreatorRequest) | [QueryCreatorResponse](#gauss.nft.v1.QueryCreatorResponse) | this line is used by starport scaffolding # 2 | GET|/icplaza/nft/creator/{creator}|
| `Nftoken` | [QueryNftokenRequest](#gauss.nft.v1.QueryNftokenRequest) | [QueryNftokenResponse](#gauss.nft.v1.QueryNftokenResponse) | Query NFT info | GET|/icplaza/nft/token/{token_id}|
| `Collections` | [QueryCollectionsRequest](#gauss.nft.v1.QueryCollectionsRequest) | [QueryCollectionsResponse](#gauss.nft.v1.QueryCollectionsResponse) | Queries a list of collections items. | GET|/icplaza/nft/collections|
| `Collection` | [QueryCollectionRequest](#gauss.nft.v1.QueryCollectionRequest) | [QueryCollectionResponse](#gauss.nft.v1.QueryCollectionResponse) | Queries a list of collections items. | GET|/icplaza/nft/collection/{cate_id}|
| `Owner` | [QueryOwnerRequest](#gauss.nft.v1.QueryOwnerRequest) | [QueryOwnerResponse](#gauss.nft.v1.QueryOwnerResponse) | Owner queries the NFTs of the specified owner | GET|/icplaza/nft/owners/{owner}|
| `OwnerNfts` | [QueryNftsByNameOrTokenRequest](#gauss.nft.v1.QueryNftsByNameOrTokenRequest) | [QueryNftsByNameOrTokenResponse](#gauss.nft.v1.QueryNftsByNameOrTokenResponse) | search by name or token_id | GET|/icplaza/nft/nfts/{owner}|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

