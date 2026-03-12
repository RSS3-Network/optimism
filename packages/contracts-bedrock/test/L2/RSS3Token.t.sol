// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Test } from "forge-std/Test.sol";
import { RSS3Token } from "src/L2/RSS3Token.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract RSS3Token_Test is Test {
    RSS3Token public token;

    address constant L2_BRIDGE = address(0x1234);
    address constant L1_TOKEN = address(0x5678);
    address constant SETTLEMENT_ADDRESS = 0x0cE3159BF19F3C55B648D04E8f0Ae1Ae118D2A0B;
    uint256 constant INITIAL_MINT_AMOUNT = 30_000_000 * 10 ** 18;
    uint256 settlementBalance;

    function setUp() public {
        // Create and select the fork
        vm.createSelectFork("https://rpc.rss3.io", 31726799);

        // Deploy the token contract
        RSS3Token tokenV2 = new RSS3Token(L2_BRIDGE, L1_TOKEN);

        address payable tokenProxy = payable(0x4200000000000000000000000000000000000042);
        token = RSS3Token(tokenProxy);

        settlementBalance = token.balanceOf(SETTLEMENT_ADDRESS);

        TransparentUpgradeableProxy proxy = TransparentUpgradeableProxy(tokenProxy);

        vm.startPrank(0x4200000000000000000000000000000000000018);
        proxy.upgradeToAndCall(address(tokenV2), abi.encodeWithSelector(tokenV2.initialize.selector));
        vm.stopPrank();
    }

    function test_constructor() public {
        assertEq(token.name(), "RSS3");
        assertEq(token.symbol(), "RSS3");
        assertEq(token.decimals(), 18);
        assertEq(token.l2Bridge(), L2_BRIDGE);
        assertEq(token.l1Token(), L1_TOKEN);
    }

    function test_initialize() public {
        // Check the settlement address balance
        assertEq(
            token.balanceOf(SETTLEMENT_ADDRESS),
            INITIAL_MINT_AMOUNT + settlementBalance,
            "Settlement address should have initial mint amount"
        );
    }

    function test_initialize_cannotBeCalledTwice() public {
        // Second initialization should revert
        vm.expectRevert("Initializable: contract is already initialized");
        token.initialize();
    }
}
