<!--
order: 0
title: "Validator-DAO Overview"
parent:
  title: "validator-dao"
-->

# `dao`

## Abstract

This document specifies the internal `x/validator-dao` module of the ICPLAZA.

The `x/validator-dao` module enables ICPLAZA to support that validators govern the application layer of the ICPLAZA network, especially for EVM-based smart contract applications.

Why is application governance important? The ICPLAZA network is a decentralized network. In theory, EVM smart contract applications can be deployed by anyone at any time. This will inevitably lead to the accumulation of junk applications and affect the quality of the network ecology; on the other side, the network protocol governance by the validator cannot satisfy the needs of rapid ecological development. In order to solve this, the ability of governance needs to be placed to the application layer to build a healthy community ecological environment.

With the `x/validator-dao` users on ICPLAZA can

- Deploy EVM smart contracts after authorization

## Contents

1. **[Hooks](01_hooks.md)**
2. **[Events](02_events.md)**
3. **[Parameters](03_parameters.md)**
4. **[Clients](04_clients.md)**

