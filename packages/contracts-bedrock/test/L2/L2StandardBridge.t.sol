// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
// Target contract is imported by the `Bridge_Initializer`
import { Bridge_Initializer } from "test/setup/Bridge_Initializer.sol";
import { stdStorage, StdStorage } from "forge-std/Test.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// Libraries
import { Hashing } from "src/libraries/Hashing.sol";
import { Types } from "src/libraries/Types.sol";

// Target contract dependencies
import { Predeploys } from "src/libraries/Predeploys.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { OptimismMintableERC20 } from "src/universal/OptimismMintableERC20.sol";

contract L2StandardBridge_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    /// @dev Tests that the bridge is initialized correctly.
    function test_initialize_succeeds() external {
        assertEq(address(l2StandardBridge.messenger()), address(l2CrossDomainMessenger));
        assertEq(l1StandardBridge.l2TokenBridge(), address(l2StandardBridge));
        assertEq(address(l2StandardBridge.OTHER_BRIDGE()), address(l1StandardBridge));
    }

    /// @dev Ensures that the L2StandardBridge is always not paused. The pausability
    ///      happens on L1 and not L2.
    function test_paused_succeeds() external {
        assertFalse(l2StandardBridge.paused());
    }

    /// @dev Tests that `withdraw` reverts.
    function test_withdraw_legacy_eth_reverts() external {
        assertEq(address(l2ToL1MessagePasser).balance, 0);

        vm.expectRevert();
        vm.prank(alice, alice);
        l2StandardBridge.withdraw(address(Predeploys.LEGACY_ERC20_ETH), 100, 1000, hex"");
    }
}

contract PreBridgeERC20 is Bridge_Initializer {
    /// @dev Sets up expected calls and emits for a successful ERC20 withdrawal.
    function _preBridgeERC20(bool _isLegacy, address _l2Token) internal {
        // Alice has 100 L2Token
        deal(_l2Token, alice, 100, true);
        assertEq(ERC20(_l2Token).balanceOf(alice), 100);
        uint256 nonce = l2CrossDomainMessenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L1Token), _l2Token, alice, alice, 100, hex""
        );
        uint64 baseGas = l2CrossDomainMessenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l2StandardBridge),
            address(l1StandardBridge),
            0,
            1000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(l2CrossDomainMessenger),
                target: address(l1CrossDomainMessenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        if (_isLegacy) {
            vm.expectCall(
                address(l2StandardBridge),
                abi.encodeWithSelector(l2StandardBridge.withdraw.selector, _l2Token, 100, 1000, hex"")
            );
        } else {
            vm.expectCall(
                address(l2StandardBridge),
                abi.encodeWithSelector(
                    l2StandardBridge.bridgeERC20.selector, _l2Token, address(L1Token), 100, 1000, hex""
                )
            );
        }

        vm.expectCall(
            address(l2CrossDomainMessenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l1StandardBridge), message, 1000)
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(l1CrossDomainMessenger),
                baseGas,
                withdrawalData
            )
        );

        // The l2StandardBridge should burn the tokens
        vm.expectCall(_l2Token, abi.encodeWithSelector(OptimismMintableERC20.burn.selector, alice, 100));

        vm.expectEmit(true, true, true, true);
        emit WithdrawalInitiated(address(L1Token), _l2Token, alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true);
        emit ERC20BridgeInitiated(_l2Token, address(L1Token), alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true);
        emit MessagePassed(
            nonce,
            address(l2CrossDomainMessenger),
            address(l1CrossDomainMessenger),
            0,
            baseGas,
            withdrawalData,
            withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true);
        emit SentMessage(address(l1StandardBridge), address(l2StandardBridge), message, nonce, 1000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true);
        emit SentMessageExtension1(address(l2StandardBridge), 0);

        vm.prank(alice, alice);
    }
}

contract L2StandardBridge_BridgeERC20_Test is PreBridgeERC20 {
    // withdraw
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_withdraw_withdrawingERC20_succeeds() external {
        _preBridgeERC20({ _isLegacy: true, _l2Token: address(L2Token) });
        l2StandardBridge.withdraw(address(L2Token), 100, 1000, hex"");

        assertEq(L2Token.balanceOf(alice), 0);
    }

    // BridgeERC20
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeERC20_succeeds() external {
        _preBridgeERC20({ _isLegacy: false, _l2Token: address(L2Token) });
        l2StandardBridge.bridgeERC20(address(L2Token), address(L1Token), 100, 1000, hex"");

        assertEq(L2Token.balanceOf(alice), 0);
    }

    function test_withdrawLegacyERC20_succeeds() external {
        _preBridgeERC20({ _isLegacy: true, _l2Token: address(LegacyL2Token) });
        l2StandardBridge.withdraw(address(LegacyL2Token), 100, 1000, hex"");

        assertEq(L2Token.balanceOf(alice), 0);
    }

    function test_bridgeLegacyERC20_succeeds() external {
        _preBridgeERC20({ _isLegacy: false, _l2Token: address(LegacyL2Token) });
        l2StandardBridge.bridgeERC20(address(LegacyL2Token), address(L1Token), 100, 1000, hex"");

        assertEq(L2Token.balanceOf(alice), 0);
    }

    function test_withdraw_notEOA_reverts() external {
        // This contract has 100 L2Token
        deal(address(L2Token), address(this), 100, true);

        vm.expectRevert("StandardBridge: function can only be called from an EOA");
        l2StandardBridge.withdraw(address(L2Token), 100, 1000, hex"");
    }
}

contract PreBridgeERC20To is Bridge_Initializer {
    // withdrawTo and BridgeERC20To should behave the same when transferring ERC20 tokens
    // so they should share the same setup and expectEmit calls
    function _preBridgeERC20To(bool _isLegacy, address _l2Token) internal {
        deal(_l2Token, alice, 100, true);
        assertEq(ERC20(L2Token).balanceOf(alice), 100);
        uint256 nonce = l2CrossDomainMessenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L1Token), _l2Token, alice, bob, 100, hex""
        );
        uint64 baseGas = l2CrossDomainMessenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l2StandardBridge),
            address(l1StandardBridge),
            0,
            1000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(l2CrossDomainMessenger),
                target: address(l1CrossDomainMessenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectEmit(true, true, true, true, address(l2StandardBridge));
        emit WithdrawalInitiated(address(L1Token), _l2Token, alice, bob, 100, hex"");

        vm.expectEmit(true, true, true, true, address(l2StandardBridge));
        emit ERC20BridgeInitiated(_l2Token, address(L1Token), alice, bob, 100, hex"");

        vm.expectEmit(true, true, true, true, address(l2ToL1MessagePasser));
        emit MessagePassed(
            nonce,
            address(l2CrossDomainMessenger),
            address(l1CrossDomainMessenger),
            0,
            baseGas,
            withdrawalData,
            withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(l2CrossDomainMessenger));
        emit SentMessage(address(l1StandardBridge), address(l2StandardBridge), message, nonce, 1000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(l2CrossDomainMessenger));
        emit SentMessageExtension1(address(l2StandardBridge), 0);

        if (_isLegacy) {
            vm.expectCall(
                address(l2StandardBridge),
                abi.encodeWithSelector(l2StandardBridge.withdrawTo.selector, _l2Token, bob, 100, 1000, hex"")
            );
        } else {
            vm.expectCall(
                address(l2StandardBridge),
                abi.encodeWithSelector(
                    l2StandardBridge.bridgeERC20To.selector, _l2Token, address(L1Token), bob, 100, 1000, hex""
                )
            );
        }

        vm.expectCall(
            address(l2CrossDomainMessenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l1StandardBridge), message, 1000)
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(l1CrossDomainMessenger),
                baseGas,
                withdrawalData
            )
        );

        // The l2StandardBridge should burn the tokens
        vm.expectCall(address(L2Token), abi.encodeWithSelector(OptimismMintableERC20.burn.selector, alice, 100));

        vm.prank(alice, alice);
    }
}

contract L2StandardBridge_BridgeERC20To_Test is PreBridgeERC20To {
    /// @dev Tests that `withdrawTo` burns the tokens, emits `WithdrawalInitiated`,
    ///      and initiates a withdrawal with `Withdrawer.initiateWithdrawal`.
    function test_withdrawTo_withdrawingERC20_succeeds() external {
        _preBridgeERC20To({ _isLegacy: true, _l2Token: address(L2Token) });
        l2StandardBridge.withdrawTo(address(L2Token), bob, 100, 1000, hex"");

        assertEq(L2Token.balanceOf(alice), 0);
    }

    /// @dev Tests that `bridgeERC20To` burns the tokens, emits `WithdrawalInitiated`,
    ///      and initiates a withdrawal with `Withdrawer.initiateWithdrawal`.
    function test_bridgeERC20To_succeeds() external {
        _preBridgeERC20To({ _isLegacy: false, _l2Token: address(L2Token) });
        l2StandardBridge.bridgeERC20To(address(L2Token), address(L1Token), bob, 100, 1000, hex"");
        assertEq(L2Token.balanceOf(alice), 0);
    }
}

contract L2StandardBridge_Bridge_Test is Bridge_Initializer {
    /// @dev Tests that `finalizeDeposit` succeeds. It should:
    ///      - only be callable by the l1TokenBridge
    ///      - emit `DepositFinalized` if the token pair is supported
    ///      - call `Withdrawer.initiateWithdrawal` if the token pair is not supported
    function test_finalizeDeposit_depositingERC20_succeeds() external {
        vm.mockCall(
            address(l2StandardBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2StandardBridge.OTHER_BRIDGE()))
        );

        vm.expectCall(address(L2Token), abi.encodeWithSelector(OptimismMintableERC20.mint.selector, alice, 100));

        // Should emit both the bedrock and legacy events
        vm.expectEmit(true, true, true, true, address(l2StandardBridge));
        emit DepositFinalized(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true, address(l2StandardBridge));
        emit ERC20BridgeFinalized(address(L2Token), address(L1Token), alice, alice, 100, hex"");

        vm.prank(address(l2CrossDomainMessenger));
        l2StandardBridge.finalizeDeposit(address(L1Token), address(L2Token), alice, alice, 100, hex"");
    }

    /// @dev Tests that `deposit` reverts.
    function test_finalizeDeposit_depositingLegacyETH_reverts() external {
        vm.expectRevert();
        vm.prank(address(l2CrossDomainMessenger));
        l2StandardBridge.finalizeDeposit(address(0), address(Predeploys.LEGACY_ERC20_ETH), alice, alice, 100, hex"");
    }
}
