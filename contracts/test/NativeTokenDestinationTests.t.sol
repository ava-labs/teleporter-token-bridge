// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {NativeTokenBridgeTest} from "./NativeTokenBridgeTests.t.sol";
import {
    NativeTokenDestination,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "../src/NativeTokenDestination.sol";
import {IWrappedNativeToken} from "../src/interfaces/IWrappedNativeToken.sol";
import {ExampleWAVAX} from "../src/mocks/ExampleWAVAX.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";
import {SendTokensInput} from "../src/interfaces/ITeleporterTokenBridge.sol";

contract NativeTokenDestinationTest is NativeTokenBridgeTest, TeleporterTokenDestinationTest {
    NativeTokenDestination public app;
    IWrappedNativeToken public mockWrappedToken;

    event CollateralAdded(uint256 amount, uint256 remaining);
    event NativeTokensMinted(address indexed recipient, uint256 amount);
    event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned);

    function setUp() public override {
        TeleporterTokenDestinationTest.setUp();

        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(INativeMinter.mintNativeCoin.selector),
            ""
        );
        mockWrappedToken = new ExampleWAVAX();
        app = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID_: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress_: TOKEN_SOURCE_ADDRESS,
            feeTokenAddress_: address(mockWrappedToken),
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true
        });
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        feeToken = mockWrappedToken;
        collateralizeBridge();
    }

    function collateralizeBridge() public {
        vm.expectEmit(true, true, true, true, address(app));
        emit CollateralAdded({amount: _DEFAULT_INITIAL_RESERVE_IMBALANCE, remaining: 0});

        bool isCollateralized = app.isCollateralized();
        assertEq(isCollateralized, false);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(
                _DEFAULT_INITIAL_RESERVE_IMBALANCE / _DEFAULT_TOKEN_MULTIPLIER,
                DEFAULT_RECIPIENT_ADDRESS
            )
        );

        isCollateralized = app.isCollateralized();
        assertEq(isCollateralized, true);
    }

    function testTransferToSource() public {
        vm.expectEmit(true, true, true, true, address(app));
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = _DEFAULT_TRANSFER_AMOUNT;
        uint256 scaledAmount = _scaleTokens(amount, false);
        emit TokensSent({
            teleporterMessageID: _MOCK_MESSAGE_ID,
            sender: address(this),
            input: input,
            amount: scaledAmount
        });

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: app.feeTokenAddress(),
                amount: input.primaryFee
            }),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(scaledAmount, input.recipient)
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        app.send{value: amount}(input);
    }

    function _checkExpectedWithdrawal(address addr, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(app));
        emit NativeTokensMinted(addr, _scaleTokens(amount, true));
    }

    function _scaleTokens(
        uint256 amount,
        bool isReceive
    ) internal view override returns (uint256) {
        if (app.multiplyOnReceive() == isReceive) {
            return amount * app.tokenMultiplier();
        }
        return amount / app.tokenMultiplier();
    }
}