// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { OptimismMintableERC20 as ERC20 } from "src/universal/OptimismMintableERC20.sol";

/// @title RSS3Token
contract RSS3Token is ERC20 {
    /// @param _l2Bridge    Address of the L2 standard bridge.
    /// @param _l1Token     Address of the L1 RSS3 token.
    constructor(address _l2Bridge, address _l1Token) ERC20(_l2Bridge, _l1Token, "RSS3", "RSS3", 18) { }
}
