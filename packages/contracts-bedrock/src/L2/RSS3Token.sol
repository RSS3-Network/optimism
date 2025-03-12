// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OptimismMintableERC20 as ERC20} from "src/universal/OptimismMintableERC20.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

/// @custom:proxied
/// @custom:predeploy 0x4200000000000000000000000000000000000042
/// @title RSS3Token
contract RSS3Token is ERC20, Initializable {
    /// @param _l2Bridge    Address of the L2 standard bridge.
    /// @param _l1Token     Address of the L1 RSS3 token.
    constructor(
        address _l2Bridge,
        address _l1Token
    ) ERC20(_l2Bridge, _l1Token, "RSS3", "RSS3", 18) {}

    /// @dev Can only be called once through initializer modifier
    function initialize() external initializer {
        uint256 amount = 30_000_000 * 10 ** 18; // annual issuance rate of 3%
        _mint(0x0cE3159BF19F3C55B648D04E8f0Ae1Ae118D2A0B, amount);
    }
}
