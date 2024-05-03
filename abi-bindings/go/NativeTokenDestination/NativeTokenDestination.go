// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokendestination

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NativeTokenDestinationSettings is an auto generated low-level Go binding around an user-defined struct.
type NativeTokenDestinationSettings struct {
	NativeAssetSymbol                   string
	TeleporterRegistryAddress           common.Address
	TeleporterManager                   common.Address
	SourceBlockchainID                  [32]byte
	TokenSourceAddress                  common.Address
	InitialReserveImbalance             *big.Int
	DecimalsShift                       uint8
	MultiplyOnReceive                   bool
	BurnedFeesReportingRewardPercentage *big.Int
}

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	RecipientContract        common.Address
	RecipientPayload         []byte
	RequiredGasLimit         *big.Int
	RecipientGasLimit        *big.Int
	FallbackRecipient        common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
	RequiredGasLimit         *big.Int
	FallbackRecipient        common.Address
}

// NativeTokenDestinationMetaData contains all meta data concerning the NativeTokenDestination contract.
var NativeTokenDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenSourceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalsShift\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnReceive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structNativeTokenDestinationSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"multiplyOnReceive\",\"type\":\"bool\"}],\"name\":\"DestinationRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_DESTINATION_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnedFeesReportingRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastestBurnedFeesReported\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnReceive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenSourceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b50604051620054ed380380620054ed833981016040819052620000359162000e30565b80602001518160400151826060015183608001518460a001518560c001518660e001518686818a6000015160405160200162000072919062000f25565b60408051601f198184030181529190528b51600362000092838262000fe5565b506004620000a1828262000fe5565b50506001600555506001600160a01b0381166200012b5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa15801562000176573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200019c9190620010b1565b60075550620001ab336200061c565b620001b6816200066e565b505060016009819055507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000213573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002399190620010b1565b60a05284620002a05760405162461bcd60e51b81526020600482015260356024820152600080516020620054cd83398151915260448201527f20736f7572636520626c6f636b636861696e2049440000000000000000000000606482015260840162000122565b60a0518503620003285760405162461bcd60e51b815260206004820152604660248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a2063616e6e60448201527f6f74206465706c6f7920746f2073616d6520626c6f636b636861696e20617320606482015265736f7572636560d01b608482015260a40162000122565b6001600160a01b038416620003955760405162461bcd60e51b81526020600482015260356024820152600080516020620054cd83398151915260448201527f20746f6b656e20736f7572636520616464726573730000000000000000000000606482015260840162000122565b60128260ff161115620004055760405162461bcd60e51b815260206004820152603160248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e76616044820152701b1a5908191958da5b585b1cd4da1a599d607a1b606482015260840162000122565b60c08590526001600160a01b03841660e052610140839052600a805460ff1916841517815562000437908390620011e0565b610100908152811515610120908152604080516060808201835287825293516020808301918252935115828401908152835180850185526000808252855185518189015293518487015291511515838801528451808403880181526080938401865281870152845160c0818101875251815260e0516001600160a01b0316818801528551808701875230815280880184905281870152620186a09781019790975284519182528186018552918601529151909391926200051d929160a083019162000505918691016200121f565b60408051601f198184030181529190529052620006ed565b505050505050505050508060a00151600003620005a35760405162461bcd60e51b815260206004820152603660248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f20696e6960448201527f7469616c207265736572766520696d62616c616e636500000000000000000000606482015260840162000122565b6064816101000151106200060d5760405162461bcd60e51b815260206004820152602a60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20696e76616c69642060448201526970657263656e7461676560b01b606482015260840162000122565b610100015161016052620013bf565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b620006786200082a565b6001600160a01b038116620006df5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840162000122565b620006ea816200061c565b50565b600080620006fa62000888565b60408401516020015190915015620007af576040830151516001600160a01b03166200077f5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b606482015260840162000122565b60408301516020808201519151620007af926001600160a01b03909116918491906200138f62000984821b17901c565b604051630624488560e41b81526001600160a01b03821690636244885090620007dd908690600401620012ad565b6020604051808303816000875af1158015620007fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620008239190620010b1565b9392505050565b6008546001600160a01b03163314620008865760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000122565b565b6000806080516001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620008cc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620008f2919062001338565b905062000917816001600160a01b031660009081526006602052604090205460ff1690565b156200097f5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b606482015260840162000122565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015620009d6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620009fc9190620010b1565b62000a08919062001358565b604080516001600160a01b038616602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663095ea7b360e01b1790915291925062000a649186919062000a6a16565b50505050565b600062000ac6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031662000b4d60201b62001463179092919060201c565b80519091501562000b48578080602001905181019062000ae791906200136e565b62000b485760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000122565b505050565b606062000b5e848460008562000b66565b949350505050565b60608247101562000bc95760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000122565b600080866001600160a01b0316858760405162000be791906200138c565b60006040518083038185875af1925050503d806000811462000c26576040519150601f19603f3d011682016040523d82523d6000602084013e62000c2b565b606091505b50909250905062000c3f8783838762000c4a565b979650505050505050565b6060831562000cbe57825160000362000cb6576001600160a01b0385163b62000cb65760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000122565b508162000b5e565b62000b5e838381511562000cd55781518083602001fd5b8060405162461bcd60e51b8152600401620001229190620013aa565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b038111828210171562000d2d5762000d2d62000cf1565b60405290565b60005b8381101562000d5057818101518382015260200162000d36565b50506000910152565b600082601f83011262000d6b57600080fd5b81516001600160401b038082111562000d885762000d8862000cf1565b604051601f8301601f19908116603f0116810190828211818310171562000db35762000db362000cf1565b8160405283815286602085880101111562000dcd57600080fd5b62000de084602083016020890162000d33565b9695505050505050565b6001600160a01b0381168114620006ea57600080fd5b80516200097f8162000dea565b805160ff811681146200097f57600080fd5b805180151581146200097f57600080fd5b60006020828403121562000e4357600080fd5b81516001600160401b038082111562000e5b57600080fd5b90830190610120828603121562000e7157600080fd5b62000e7b62000d07565b82518281111562000e8b57600080fd5b62000e998782860162000d59565b82525062000eaa6020840162000e00565b602082015262000ebd6040840162000e00565b60408201526060830151606082015262000eda6080840162000e00565b608082015260a083015160a082015262000ef760c0840162000e0d565b60c082015262000f0a60e0840162000e1f565b60e08201526101009283015192810192909252509392505050565b6702bb930b83832b2160c51b81526000825162000f4a81600885016020870162000d33565b9190910160080192915050565b600181811c9082168062000f6c57607f821691505b60208210810362000f8d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000b4857600081815260208120601f850160051c8101602086101562000fbc5750805b601f850160051c820191505b8181101562000fdd5782815560010162000fc8565b505050505050565b81516001600160401b0381111562001001576200100162000cf1565b620010198162001012845462000f57565b8462000f93565b602080601f831160018114620010515760008415620010385750858301515b600019600386901b1c1916600185901b17855562000fdd565b600085815260208120601f198616915b82811015620010825788860151825594840194600190910190840162001061565b5085821015620010a15787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060208284031215620010c457600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600181815b8085111562001122578160001904821115620011065762001106620010cb565b808516156200111457918102915b93841c9390800290620010e6565b509250929050565b6000826200113b57506001620011da565b816200114a57506000620011da565b81600181146200116357600281146200116e576200118e565b6001915050620011da565b60ff841115620011825762001182620010cb565b50506001821b620011da565b5060208310610133831016604e8410600b8410161715620011b3575081810a620011da565b620011bf8383620010e1565b8060001904821115620011d657620011d6620010cb565b0290505b92915050565b60006200082360ff8416836200112a565b600081518084526200120b81602086016020860162000d33565b601f01601f19169290920160200192915050565b6020815260008251600581106200124657634e487b7160e01b600052602160045260246000fd5b80602084015250602083015160408084015262000b5e6060840182620011f1565b600081518084526020808501945080840160005b83811015620012a25781516001600160a01b0316875295820195908201906001016200127b565b509495945050505050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526200131061010084018262001267565b905060a0840151601f198483030160e08501526200132f8282620011f1565b95945050505050565b6000602082840312156200134b57600080fd5b8151620008238162000dea565b80820180821115620011da57620011da620010cb565b6000602082840312156200138157600080fd5b620008238262000e1f565b60008251620013a081846020870162000d33565b9190910192915050565b602081526000620008236020830184620011f1565b60805160a05160c05160e05161010051610120516101405161016051614000620014cd600039600081816104c60152610d6c01526000818161060a015261096f01526000818161039401528181610df30152612e4d01526000818161070001528181610dd20152612e2c01526000818161083e01528181610f1601528181611c0901528181611e9701528181612506015281816127e601526129680152600081816103e801528181610ef001528181611be201528181611e71015281816124d6015281816127c001526128f201526000818161079501528181611c7b01528181611d19015261263601526000818161034801528181611122015281816121890152612ef601526140006000f3fe6080604052600436106102765760003560e01c80635eb995141161014f578063c21e5fa7116100c1578063d2cc7a701161007a578063d2cc7a70146107b7578063dd62ed3e146107cc578063ecd4ed1b14610508578063f2fde38b146107ec578063f3f981d81461080c578063f5ea06031461082c57610285565b8063c21e5fa714610722578063c452165e14610735578063c868efaa1461074d578063d0e30db014610285578063d10a5b8c1461076d578063d127dc9b1461078357610285565b806395d89b411161011357806395d89b411461064a578063973142971461065f578063a2309ff814610698578063a457c2d7146106ae578063a9059cbb146106ce578063ba3f5a12146106ee57610285565b80635eb995141461058d57806370a08231146105ad578063715018a6146105e35780638ac7dd20146105f85780638da5cb5b1461062c57610285565b806330079bff116101e85780634511243e116101ac5780634511243e146104e857806347a9a22c1461050857806349e3284e14610525578063525975e61461053f57806355538c8b1461055657806355c267251461057657610285565b806330079bff1461044a578063313ce5671461045d578063329c3e121461047957806339509351146104945780633a23dfe2146104b457610285565b80631a7f5bec1161023a5780631a7f5bec146103365780631ce220751461038257806323b872dd146103b657806329b7b3fd146103d65780632b0d8f181461040a5780632e1a7d4d1461042a57610285565b806306fdde031461028d578063095ea7b3146102b857806315beb59f146102e857806318160ddd1461030c5780631906529c1461032157610285565b3661028557610283610860565b005b610283610860565b34801561029957600080fd5b506102a26108a1565b6040516102af91906133ac565b60405180910390f35b3480156102c457600080fd5b506102d86102d33660046133df565b610933565b60405190151581526020016102af565b3480156102f457600080fd5b506102fe61213481565b6040519081526020016102af565b34801561031857600080fd5b506002546102fe565b34801561032d57600080fd5b506102fe61094d565b34801561034257600080fd5b5061036a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102af565b34801561038e57600080fd5b506102d87f000000000000000000000000000000000000000000000000000000000000000081565b3480156103c257600080fd5b506102d86103d136600461340b565b6109ae565b3480156103e257600080fd5b506102fe7f000000000000000000000000000000000000000000000000000000000000000081565b34801561041657600080fd5b5061028361042536600461344c565b6109d2565b34801561043657600080fd5b50610283610445366004613469565b610ad4565b610283610458366004613482565b610b44565b34801561046957600080fd5b50604051601281526020016102af565b34801561048557600080fd5b5061036a6001600160991b0181565b3480156104a057600080fd5b506102d86104af3660046133df565b610b7b565b3480156104c057600080fd5b506102fe7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104f457600080fd5b5061028361050336600461344c565b610b9d565b34801561051457600080fd5b5061036a62010203600160981b0181565b34801561053157600080fd5b50600a546102d89060ff1681565b34801561054b57600080fd5b506102fe62035b6081565b34801561056257600080fd5b50610283610571366004613469565b610c9a565b34801561058257600080fd5b506102fe620186a081565b34801561059957600080fd5b506102836105a8366004613469565b611026565b3480156105b957600080fd5b506102fe6105c836600461344c565b6001600160a01b031660009081526020819052604090205490565b3480156105ef57600080fd5b50610283611037565b34801561060457600080fd5b506102fe7f000000000000000000000000000000000000000000000000000000000000000081565b34801561063857600080fd5b506008546001600160a01b031661036a565b34801561065657600080fd5b506102a2611049565b34801561066b57600080fd5b506102d861067a36600461344c565b6001600160a01b031660009081526006602052604090205460ff1690565b3480156106a457600080fd5b506102fe600b5481565b3480156106ba57600080fd5b506102d86106c93660046133df565b611058565b3480156106da57600080fd5b506102d86106e93660046133df565b6110d3565b3480156106fa57600080fd5b506102fe7f000000000000000000000000000000000000000000000000000000000000000081565b6102836107303660046134bd565b6110e1565b34801561074157600080fd5b5061036a600160981b81565b34801561075957600080fd5b506102836107683660046134d5565b61110d565b34801561077957600080fd5b506102fe600c5481565b34801561078f57600080fd5b506102fe7f000000000000000000000000000000000000000000000000000000000000000081565b3480156107c357600080fd5b506007546102fe565b3480156107d857600080fd5b506102fe6107e736600461355d565b6112d7565b3480156107f857600080fd5b5061028361080736600461344c565b611302565b34801561081857600080fd5b506102fe610827366004613469565b611378565b34801561083857600080fd5b5061036a7f000000000000000000000000000000000000000000000000000000000000000081565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a261089f333461147a565b565b6060600380546108b090613596565b80601f01602080910402602001604051908101604052809291908181526020018280546108dc90613596565b80156109295780601f106108fe57610100808354040283529160200191610929565b820191906000526020600020905b81548152906001019060200180831161090c57829003601f168201915b5050505050905090565b600033610941818585611539565b60019150505b92915050565b60008061096962010203600160981b0131600160981b316135e0565b905060007f0000000000000000000000000000000000000000000000000000000000000000600b5461099b91906135e0565b90506109a782826135f3565b9250505090565b6000336109bc85828561165e565b6109c78585856116d2565b506001949350505050565b6109da611876565b6001600160a01b038116610a095760405162461bcd60e51b8152600401610a0090613606565b60405180910390fd5b6001600160a01b03811660009081526006602052604090205460ff1615610a885760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610a00565b6001600160a01b038116600081815260066020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610b13338261187e565b604051339082156108fc029083906000818181858888f19350505050158015610b40573d6000803e3d6000fd5b5050565b600a5460ff16610b665760405162461bcd60e51b8152600401610a0090613654565b610b78610b72826137d8565b346119ad565b50565b600033610941818585610b8e83836112d7565b610b9891906135e0565b611539565b610ba5611876565b6001600160a01b038116610bcb5760405162461bcd60e51b8152600401610a0090613606565b6001600160a01b03811660009081526006602052604090205460ff16610c455760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610a00565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600660205260409020805460ff19169055565b600160095414610cbc5760405162461bcd60e51b8152600401610a0090613885565b6002600955600c54600160981b31908111610d515760405162461bcd60e51b815260206004820152604960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a206275726e2061646460448201527f726573732062616c616e6365206e6f742067726561746572207468616e206c616064820152681cdd081c995c1bdc9d60ba1b608482015260a401610a00565b6000600c5482610d6191906135f3565b905060006064610d917f0000000000000000000000000000000000000000000000000000000000000000846138c9565b610d9b91906138e0565b90506000610da982846135f3565b600c85905590508115610dcb57610dc03083611f9c565b610dc98261201e565b505b6000610e1a7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084600061202e565b11610e8d5760405162461bcd60e51b815260206004820152603960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f2073636160448201527f6c656420616d6f756e7420746f207265706f7274206275726e000000000000006064820152608401610a00565b604080518082018252600181528151808301835262010203600160981b0181526020808201859052925160009380840192610eca92909101613918565b60405160208183030381529060405281525090506000610fdc6040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b0316815260200188815250815260200189815260200160006001600160401b03811115610f8557610f856136a8565b604051908082528060200260200182016040528015610fae578160200160208202803683370190505b50815260200184604051602001610fc59190613938565b60405160208183030381529060405281525061205f565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a58460405161101091815260200190565b60405180910390a2505060016009555050505050565b61102e611876565b610b7881612185565b61103f612325565b61089f600061237f565b6060600480546108b090613596565b6000338161106682866112d7565b9050838110156110c65760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610a00565b6109c78286868403611539565b6000336109418185856116d2565b600a5460ff166111035760405162461bcd60e51b8152600401610a0090613654565b610b7881346123d1565b611115612897565b6007546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561118c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b0919061397d565b10156112175760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610a00565b6112203361067a565b156112865760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610a00565b6112c7848484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506128f092505050565b6112d16001600555565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61130a612325565b6001600160a01b03811661136f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a00565b610b788161237f565b6000600561138783601f6135e0565b901c92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156113e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611404919061397d565b61140e91906135e0565b604080516001600160a01b038616602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790529091506112d1908590612b0d565b60606114728484600085612bdf565b949350505050565b6001600160a01b0382166114d05760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610a00565b80600260008282546114e291906135e0565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b03831661159b5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610a00565b6001600160a01b0382166115fc5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610a00565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b600061166a84846112d7565b905060001981146112d157818110156116c55760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610a00565b6112d18484848403611539565b6001600160a01b0383166117365760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610a00565b6001600160a01b0382166117985760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610a00565b6001600160a01b038316600090815260208190526040902054818110156118105760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610a00565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36112d1565b61089f612325565b6001600160a01b0382166118de5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610a00565b6001600160a01b038216600090815260208190526040902054818110156119525760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610a00565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611651565b505050565b6001600954146119cf5760405162461bcd60e51b8152600401610a0090613885565b600260095560408201516001600160a01b0316611a425760405162461bcd60e51b815260206004820152603b6024820152600080516020613f8b83398151915260448201527f20726563697069656e7420636f6e7472616374206164647265737300000000006064820152608401610a00565b6000826080015111611a665760405162461bcd60e51b8152600401610a0090613996565b60008260a0015111611ac55760405162461bcd60e51b81526020600482015260346024820152600080516020613f8b833981519152604482015273081c9958da5c1a595b9d0819d85cc81b1a5b5a5d60621b6064820152608401610a00565b81608001518260a0015110611b305760405162461bcd60e51b81526020600482015260376024820152600080516020613fab83398151915260448201527f6c696420726563697069656e7420676173206c696d69740000000000000000006064820152608401610a00565b60c08201516001600160a01b0316611b9e5760405162461bcd60e51b815260206004820152603b6024820152600080516020613f8b83398151915260448201527f2066616c6c6261636b20726563697069656e74206164647265737300000000006064820152608401610a00565b611bbc82600001518360200151838560e00151866101000151612cba565b9050611bd960408051808201909152600081526060602082015290565b608083015183517f00000000000000000000000000000000000000000000000000000000000000009003611d15577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031684602001516001600160a01b031614611c5c5760405162461bcd60e51b8152600401610a00906139d7565b6040805180820190915280600281526020016040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000008152602001336001600160a01b0316815260200187604001516001600160a01b03168152602001868152602001876060015181526020018760a0015181526020018760c001516001600160a01b0316815250604051602001611cfc9190613a22565b6040516020818303038152906040528152509150611e5f565b83517f00000000000000000000000000000000000000000000000000000000000000009003611d7357306001600160a01b031684602001516001600160a01b031603611d735760405162461bcd60e51b8152600401610a00906139d7565b604080518082019091528060048152602001604051806101400160405280336001600160a01b031681526020018760000151815260200187602001516001600160a01b0316815260200187604001516001600160a01b03168152602001868152602001876060015181526020018760a0015181526020018760c001516001600160a01b0316815260200187608001518152602001876101000151815250604051602001611e209190613a98565b6040516020818303038152906040528152509150612134611e45856060015151611378565b611e4f91906138c9565b611e5c9062035b606135e0565b90505b6000611f4a6040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b031681526020018960e00151815250815260200184815260200160006001600160401b03811115611f0a57611f0a6136a8565b604051908082528060200260200182016040528015611f33578160200160208202803683370190505b50815260200185604051602001610fc59190613938565b9050336001600160a01b0316817f76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e665628787604051611f88929190613b59565b60405180910390a350506001600955505050565b80600b6000828254611fae91906135e0565b90915550506040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba90604401600060405180830381600087803b15801561200257600080fd5b505af1158015612016573d6000803e3d6000fd5b505050505050565b600061202a308361147a565b5090565b60008115158415150361204c5761204585846138c9565b9050611472565b61205685846138e0565b95945050505050565b60008061206a612ef1565b6040840151602001519091501561210f576040830151516001600160a01b03166120ec5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610a00565b60408301516020810151905161210f916001600160a01b0390911690839061138f565b604051630624488560e41b81526001600160a01b0382169063624488509061213b908690600401613c02565b6020604051808303816000875af115801561215a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061217e919061397d565b9392505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612209919061397d565b600754909150818311156122795760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610a00565b8083116122ee5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610a00565b6007839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6008546001600160a01b0316331461089f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a00565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600954146123f35760405162461bcd60e51b8152600401610a0090613885565b6002600955600061240a606084016040850161344c565b6001600160a01b0316036124695760405162461bcd60e51b81526020600482015260326024820152600080516020613f8b83398151915260448201527120726563697069656e74206164647265737360701b6064820152608401610a00565b60008260a001351161248d5760405162461bcd60e51b8152600401610a0090613996565b6124b282356124a2604085016020860161344c565b8385606001358660800135612cba565b90506124cf60408051808201909152600081526060602082015290565b60a08301357f0000000000000000000000000000000000000000000000000000000000000000843503612634576001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016612536604086016020870161344c565b6001600160a01b03161461255c5760405162461bcd60e51b8152600401610a00906139d7565b6080840135156125c95760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a206e6f6e2d6044820152717a65726f207365636f6e646172792066656560701b6064820152608401610a00565b60408051808201909152806001815260200160405180604001604052808760400160208101906125f9919061344c565b6001600160a01b031681526020018681525060405160200161261b9190613918565b60405160208183030381529060405281525091506127ae565b7f0000000000000000000000000000000000000000000000000000000000000000843503612693573061266d604086016020870161344c565b6001600160a01b0316036126935760405162461bcd60e51b8152600401610a00906139d7565b6040805180820190915280600381526020016040518060e00160405280876000013581526020018760200160208101906126cd919061344c565b6001600160a01b031681526020016126eb6060890160408a0161344c565b6001600160a01b0316815260208101879052608080890135604083015260a089013560608301520161272360e0890160c08a0161344c565b6001600160a01b031690526040516127939190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b604051602081830303815290604052815250915062035b6090505b60006128596040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b031681526020018960600135815250815260200184815260200160006001600160401b03811115611f0a57611f0a6136a8565b9050336001600160a01b0316817fc117b57e5bc66f5f81fcb2b851a3ceee905d99d34038a194c791161d63e0ce718787604051611f88929190613cbb565b6002600554036128e95760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610a00565b6002600555565b7f000000000000000000000000000000000000000000000000000000000000000083146129665760405162461bcd60e51b81526020600482015260306024820152600080516020613fab83398151915260448201526f3634b21039b7bab931b29031b430b4b760811b6064820152608401610a00565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316146129fb5760405162461bcd60e51b81526020600482015260386024820152600080516020613fab83398151915260448201527f6c696420746f6b656e20736f75726365206164647265737300000000000000006064820152608401610a00565b600081806020019051810190612a119190613d80565b600a5490915060ff16612a2c57600a805460ff191660011790555b600181516004811115612a4157612a41613902565b03612a7a5760008160200151806020019051810190612a609190613e0e565b9050612a7481600001518260200151613005565b506112d1565b600281516004811115612a8f57612a8f613902565b03612abe5760008160200151806020019051810190612aae9190613e48565b9050612a74818260600151613052565b60405162461bcd60e51b81526020600482015260306024820152600080516020613fab83398151915260448201526f6c6964206d657373616765207479706560801b6064820152608401610a00565b6000612b62826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166114639092919063ffffffff16565b8051909150156119a85780806020019051810190612b809190613f05565b6119a85760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610a00565b606082471015612c405760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610a00565b600080866001600160a01b03168587604051612c5c9190613f27565b60006040518083038185875af1925050503d8060008114612c99576040519150601f19603f3d011682016040523d82523d6000602084013e612c9e565b606091505b5091509150612caf878383876131a4565b979650505050505050565b600085612d1d5760405162461bcd60e51b815260206004820152603a6024820152600080516020613f8b83398151915260448201527f2064657374696e6174696f6e20626c6f636b636861696e2049440000000000006064820152608401610a00565b6001600160a01b038516612d875760405162461bcd60e51b815260206004820152603b6024820152600080516020613f8b83398151915260448201527f2064657374696e6174696f6e20627269646765206164647265737300000000006064820152608401610a00565b612d908461201e565b9350612d9c82846135e0565b8411612e105760405162461bcd60e51b815260206004820152603d60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e737560448201527f6666696369656e7420616d6f756e7420746f20636f76657220666565730000006064820152608401610a00565b612e1a83856135f3565b9350612e258461321d565b6000612e747f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000087600061202e565b11612ee75760405162461bcd60e51b815260206004820152603b60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e737560448201527f6666696369656e7420746f6b656e7320746f207472616e7366657200000000006064820152608401610a00565b5091949350505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f769190613f43565b9050612f9a816001600160a01b031660009081526006602052604090205460ff1690565b156130005760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610a00565b919050565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b8260405161304091815260200190565b60405180910390a2610b408282611f9c565b61305c3082611f9c565b600082600001518360200151846080015160405160240161307f93929190613f60565b60408051601f198184030181529181526020820180516001600160e01b0316630d57eddf60e21b17905260a0850151908501519192506000916130c5919085908561325d565b905080156131195783604001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161310c91815260200190565b60405180910390a26112d1565b83604001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08460405161315891815260200190565b60405180910390a28360c001516001600160a01b03166108fc849081150290604051600060405180830381858888f1935050505015801561319d573d6000803e3d6000fd5b5050505050565b6060831561321357825160000361320c576001600160a01b0385163b61320c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a00565b5081611472565b6114728383613332565b613227308261187e565b60405162010203600160981b019082156108fc029083906000818181858888f19350505050158015610b40573d6000803e3d6000fd5b6000845a10156132af5760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610a00565b834710156132ff5760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610a00565b826001600160a01b03163b60000361331957506000611472565b600080600084516020860188888bf19695505050505050565b8151156133425781518083602001fd5b8060405162461bcd60e51b8152600401610a0091906133ac565b60005b8381101561337757818101518382015260200161335f565b50506000910152565b6000815180845261339881602086016020860161335c565b601f01601f19169290920160200192915050565b60208152600061217e6020830184613380565b6001600160a01b0381168114610b7857600080fd5b8035613000816133bf565b600080604083850312156133f257600080fd5b82356133fd816133bf565b946020939093013593505050565b60008060006060848603121561342057600080fd5b833561342b816133bf565b9250602084013561343b816133bf565b929592945050506040919091013590565b60006020828403121561345e57600080fd5b813561217e816133bf565b60006020828403121561347b57600080fd5b5035919050565b60006020828403121561349457600080fd5b81356001600160401b038111156134aa57600080fd5b8201610120818503121561217e57600080fd5b600060e082840312156134cf57600080fd5b50919050565b600080600080606085870312156134eb57600080fd5b8435935060208501356134fd816133bf565b925060408501356001600160401b038082111561351957600080fd5b818701915087601f83011261352d57600080fd5b81358181111561353c57600080fd5b88602082850101111561354e57600080fd5b95989497505060200194505050565b6000806040838503121561357057600080fd5b823561357b816133bf565b9150602083013561358b816133bf565b809150509250929050565b600181811c908216806135aa57607f821691505b6020821081036134cf57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610947576109476135ca565b81810381811115610947576109476135ca565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526034908201527f4e6174697665546f6b656e44657374696e6174696f6e3a20636f6e7472616374604082015273081d5b99195c98dbdb1b185d195c985b1a5e995960621b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b03811182821017156136e1576136e16136a8565b60405290565b604080519081016001600160401b03811182821017156136e1576136e16136a8565b60405160e081016001600160401b03811182821017156136e1576136e16136a8565b604051601f8201601f191681016001600160401b0381118282101715613753576137536136a8565b604052919050565b60006001600160401b03821115613774576137746136a8565b50601f01601f191660200190565b600082601f83011261379357600080fd5b81356137a66137a18261375b565b61372b565b8181528460208386010111156137bb57600080fd5b816020850160208301376000918101602001919091529392505050565b600061012082360312156137eb57600080fd5b6137f36136be565b82358152613803602084016133d4565b6020820152613814604084016133d4565b604082015260608301356001600160401b0381111561383257600080fd5b61383e36828601613782565b6060830152506080830135608082015260a083013560a082015261386460c084016133d4565b60c082015260e0838101359082015261010092830135928101929092525090565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b8082028115828204841417610947576109476135ca565b6000826138fd57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b81516001600160a01b031681526020808301519082015260408101610947565b60208152600082516005811061395e57634e487b7160e01b600052602160045260246000fd5b8060208401525060208301516040808401526114726060840182613380565b60006020828403121561398f57600080fd5b5051919050565b6020808252603390820152600080516020613f8b833981519152604082015272081c995c5d5a5c99590819d85cc81b1a5b5a5d606a1b606082015260800190565b6020808252603e90820152600080516020613fab83398151915260408201527f6c69642064657374696e6174696f6e2062726964676520616464726573730000606082015260800190565b60208152815160208201526000602083015160018060a01b038082166040850152806040860151166060850152606085015160808501526080850151915060e060a0850152613a75610100850183613380565b915060a085015160c08501528060c08601511660e0850152508091505092915050565b60208152613ab26020820183516001600160a01b03169052565b6020820151604082015260006040830151613ad860608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101408060c0850152613b14610160850183613380565b915060c085015160e085015260e0850151610100613b3c818701836001600160a01b03169052565b860151610120868101919091529095015193019290925250919050565b604081528251604082015260006020840151613b8060608401826001600160a01b03169052565b5060408401516001600160a01b03166080830152606084015161012060a08401819052613bb1610160850183613380565b9150608086015160c085015260a086015160e085015260c0860151610100613be3818701836001600160a01b03169052565b60e0880151928601929092525090940151610140830152506020015290565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b80841015613c9157845183168252938501936001939093019290850190613c6f565b5060a0880151878203601f190160e08901529450613caf8186613380565b98975050505050505050565b8235815261010081016020840135613cd2816133bf565b6001600160a01b039081166020840152604085013590613cf1826133bf565b8082166040850152606086013560608501526080860135608085015260a086013560a085015260c08601359150613d27826133bf565b1660c083015260e090910191909152919050565b600082601f830112613d4c57600080fd5b8151613d5a6137a18261375b565b818152846020838601011115613d6f57600080fd5b61147282602083016020870161335c565b600060208284031215613d9257600080fd5b81516001600160401b0380821115613da957600080fd5b9083019060408286031215613dbd57600080fd5b613dc56136e7565b825160058110613dd457600080fd5b8152602083015182811115613de857600080fd5b613df487828601613d3b565b60208301525095945050505050565b8051613000816133bf565b600060408284031215613e2057600080fd5b613e286136e7565b8251613e33816133bf565b81526020928301519281019290925250919050565b600060208284031215613e5a57600080fd5b81516001600160401b0380821115613e7157600080fd5b9083019060e08286031215613e8557600080fd5b613e8d613709565b82518152613e9d60208401613e03565b6020820152613eae60408401613e03565b604082015260608301516060820152608083015182811115613ecf57600080fd5b613edb87828601613d3b565b60808301525060a083015160a0820152613ef760c08401613e03565b60c082015295945050505050565b600060208284031215613f1757600080fd5b8151801515811461217e57600080fd5b60008251613f3981846020870161335c565b9190910192915050565b600060208284031215613f5557600080fd5b815161217e816133bf565b8381526001600160a01b03831660208201526060604082018190526000906120569083018461338056fe54656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e7661a26469706673582212208933a9127c56748fcf9f88235b113cee1ac37172ea56a6bbcbbf7fa3bca4293264736f6c6343000812003354656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f",
}

// NativeTokenDestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenDestinationMetaData.ABI instead.
var NativeTokenDestinationABI = NativeTokenDestinationMetaData.ABI

// NativeTokenDestinationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenDestinationMetaData.Bin instead.
var NativeTokenDestinationBin = NativeTokenDestinationMetaData.Bin

// DeployNativeTokenDestination deploys a new Ethereum contract, binding an instance of NativeTokenDestination to it.
func DeployNativeTokenDestination(auth *bind.TransactOpts, backend bind.ContractBackend, settings NativeTokenDestinationSettings) (common.Address, *types.Transaction, *NativeTokenDestination, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenDestinationBin), backend, settings)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NativeTokenDestination is an auto generated Go binding around an Ethereum contract.
type NativeTokenDestination struct {
	NativeTokenDestinationCaller     // Read-only binding to the contract
	NativeTokenDestinationTransactor // Write-only binding to the contract
	NativeTokenDestinationFilterer   // Log filterer for contract events
}

// NativeTokenDestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenDestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenDestinationSession struct {
	Contract     *NativeTokenDestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeTokenDestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenDestinationCallerSession struct {
	Contract *NativeTokenDestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NativeTokenDestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenDestinationTransactorSession struct {
	Contract     *NativeTokenDestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NativeTokenDestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenDestinationRaw struct {
	Contract *NativeTokenDestination // Generic contract binding to access the raw methods on
}

// NativeTokenDestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCallerRaw struct {
	Contract *NativeTokenDestinationCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenDestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactorRaw struct {
	Contract *NativeTokenDestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenDestination creates a new instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestination(address common.Address, backend bind.ContractBackend) (*NativeTokenDestination, error) {
	contract, err := bindNativeTokenDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NewNativeTokenDestinationCaller creates a new read-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenDestinationCaller, error) {
	contract, err := bindNativeTokenDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCaller{contract: contract}, nil
}

// NewNativeTokenDestinationTransactor creates a new write-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenDestinationTransactor, error) {
	contract, err := bindNativeTokenDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransactor{contract: contract}, nil
}

// NewNativeTokenDestinationFilterer creates a new log filterer instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenDestinationFilterer, error) {
	contract, err := bindNativeTokenDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationFilterer{contract: contract}, nil
}

// bindNativeTokenDestination binds a generic wrapper to an already deployed contract.
func bindNativeTokenDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.NativeTokenDestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDFORBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_FOR_BRIDGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MULTIHOPREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "MULTI_HOP_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenDestination.Contract.NATIVEMINTER(&_NativeTokenDestination.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenDestination.Contract.NATIVEMINTER(&_NativeTokenDestination.CallOpts)
}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) REGISTERDESTINATIONREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "REGISTER_DESTINATION_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) REGISTERDESTINATIONREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REGISTERDESTINATIONREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) REGISTERDESTINATIONREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REGISTERDESTINATIONREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SOURCECHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "SOURCE_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) SOURCECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.SOURCECHAINBURNADDRESS(&_NativeTokenDestination.CallOpts)
}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SOURCECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.SOURCECHAINBURNADDRESS(&_NativeTokenDestination.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.Allowance(&_NativeTokenDestination.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.Allowance(&_NativeTokenDestination.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.BalanceOf(&_NativeTokenDestination.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.BalanceOf(&_NativeTokenDestination.CallOpts, account)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.BlockchainID(&_NativeTokenDestination.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.BlockchainID(&_NativeTokenDestination.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BurnedFeesReportingRewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "burnedFeesReportingRewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenDestination.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenDestination.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenDestination.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenDestination.CallOpts)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenDestination.Contract.CalculateNumWords(&_NativeTokenDestination.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenDestination.Contract.CalculateNumWords(&_NativeTokenDestination.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationSession) Decimals() (uint8, error) {
	return _NativeTokenDestination.Contract.Decimals(&_NativeTokenDestination.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Decimals() (uint8, error) {
	return _NativeTokenDestination.Contract.Decimals(&_NativeTokenDestination.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenDestination.Contract.GetMinTeleporterVersion(&_NativeTokenDestination.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenDestination.Contract.GetMinTeleporterVersion(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) InitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "initialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenDestination.Contract.IsTeleporterAddressPaused(&_NativeTokenDestination.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenDestination.Contract.IsTeleporterAddressPaused(&_NativeTokenDestination.CallOpts, teleporterAddress)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) LastestBurnedFeesReported(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "lastestBurnedFeesReported")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenDestination.Contract.LastestBurnedFeesReported(&_NativeTokenDestination.CallOpts)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenDestination.Contract.LastestBurnedFeesReported(&_NativeTokenDestination.CallOpts)
}

// MultiplyOnReceive is a free data retrieval call binding the contract method 0x1ce22075.
//
// Solidity: function multiplyOnReceive() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MultiplyOnReceive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "multiplyOnReceive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplyOnReceive is a free data retrieval call binding the contract method 0x1ce22075.
//
// Solidity: function multiplyOnReceive() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) MultiplyOnReceive() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnReceive(&_NativeTokenDestination.CallOpts)
}

// MultiplyOnReceive is a free data retrieval call binding the contract method 0x1ce22075.
//
// Solidity: function multiplyOnReceive() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MultiplyOnReceive() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnReceive(&_NativeTokenDestination.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationSession) Name() (string, error) {
	return _NativeTokenDestination.Contract.Name(&_NativeTokenDestination.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Name() (string, error) {
	return _NativeTokenDestination.Contract.Name(&_NativeTokenDestination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) Owner() (common.Address, error) {
	return _NativeTokenDestination.Contract.Owner(&_NativeTokenDestination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Owner() (common.Address, error) {
	return _NativeTokenDestination.Contract.Owner(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationSession) Symbol() (string, error) {
	return _NativeTokenDestination.Contract.Symbol(&_NativeTokenDestination.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Symbol() (string, error) {
	return _NativeTokenDestination.Contract.Symbol(&_NativeTokenDestination.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterRegistry(&_NativeTokenDestination.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterRegistry(&_NativeTokenDestination.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "tokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TokenMultiplier(&_NativeTokenDestination.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TokenMultiplier(&_NativeTokenDestination.CallOpts)
}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TokenSourceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "tokenSourceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.TokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.TokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalNativeAssetSupply(&_NativeTokenDestination.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalNativeAssetSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Approve(&_NativeTokenDestination.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Approve(&_NativeTokenDestination.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.DecreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.DecreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Deposit(&_NativeTokenDestination.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Deposit(&_NativeTokenDestination.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.IncreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.IncreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, addedValue)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.PauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.PauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RenounceOwnership(&_NativeTokenDestination.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RenounceOwnership(&_NativeTokenDestination.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportBurnedTxFees(&_NativeTokenDestination.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportBurnedTxFees(&_NativeTokenDestination.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0xc21e5fa7.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xc21e5fa7.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xc21e5fa7.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x30079bff.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x30079bff.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCall(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x30079bff.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCall(&_NativeTokenDestination.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Transfer(&_NativeTokenDestination.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Transfer(&_NativeTokenDestination.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferFrom(&_NativeTokenDestination.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferFrom(&_NativeTokenDestination.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferOwnership(&_NativeTokenDestination.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferOwnership(&_NativeTokenDestination.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UnpauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UnpauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UpdateMinTeleporterVersion(&_NativeTokenDestination.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UpdateMinTeleporterVersion(&_NativeTokenDestination.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Withdraw(&_NativeTokenDestination.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Withdraw(&_NativeTokenDestination.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Fallback(&_NativeTokenDestination.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Fallback(&_NativeTokenDestination.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Receive() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Receive(&_NativeTokenDestination.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Receive(&_NativeTokenDestination.TransactOpts)
}

// NativeTokenDestinationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenDestination contract.
type NativeTokenDestinationApprovalIterator struct {
	Event *NativeTokenDestinationApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationApproval represents a Approval event raised by the NativeTokenDestination contract.
type NativeTokenDestinationApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenDestinationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationApprovalIterator{contract: _NativeTokenDestination.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationApproval)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseApproval(log types.Log) (*NativeTokenDestinationApproval, error) {
	event := new(NativeTokenDestinationApproval)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallFailedIterator struct {
	Event *NativeTokenDestinationCallFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCallFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationCallFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCallFailed represents a CallFailed event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenDestinationCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCallFailedIterator{contract: _NativeTokenDestination.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCallFailed)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CallFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCallFailed(log types.Log) (*NativeTokenDestinationCallFailed, error) {
	event := new(NativeTokenDestinationCallFailed)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallSucceededIterator struct {
	Event *NativeTokenDestinationCallSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCallSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationCallSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCallSucceeded represents a CallSucceeded event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenDestinationCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCallSucceededIterator{contract: _NativeTokenDestination.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCallSucceeded)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenDestinationCallSucceeded, error) {
	event := new(NativeTokenDestinationCallSucceeded)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenDestination contract.
type NativeTokenDestinationDepositIterator struct {
	Event *NativeTokenDestinationDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationDeposit represents a Deposit event raised by the NativeTokenDestination contract.
type NativeTokenDestinationDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenDestinationDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationDepositIterator{contract: _NativeTokenDestination.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationDeposit)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseDeposit(log types.Log) (*NativeTokenDestinationDeposit, error) {
	event := new(NativeTokenDestinationDeposit)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationDestinationRegisteredIterator is returned from FilterDestinationRegistered and is used to iterate over the raw logs and unpacked data for DestinationRegistered events raised by the NativeTokenDestination contract.
type NativeTokenDestinationDestinationRegisteredIterator struct {
	Event *NativeTokenDestinationDestinationRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationDestinationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationDestinationRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationDestinationRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationDestinationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationDestinationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationDestinationRegistered represents a DestinationRegistered event raised by the NativeTokenDestination contract.
type NativeTokenDestinationDestinationRegistered struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	InitialReserveImbalance  *big.Int
	TokenMultiplier          *big.Int
	MultiplyOnReceive        bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDestinationRegistered is a free log retrieval operation binding the contract event 0x34e0f289c6952b50ea58fa31d8c3732fb85091cfc3d5ded7767d7562ce0bfb22.
//
// Solidity: event DestinationRegistered(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 initialReserveImbalance, uint256 tokenMultiplier, bool multiplyOnReceive)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterDestinationRegistered(opts *bind.FilterOpts, destinationBlockchainID [][32]byte, destinationBridgeAddress []common.Address) (*NativeTokenDestinationDestinationRegisteredIterator, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "DestinationRegistered", destinationBlockchainIDRule, destinationBridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationDestinationRegisteredIterator{contract: _NativeTokenDestination.contract, event: "DestinationRegistered", logs: logs, sub: sub}, nil
}

// WatchDestinationRegistered is a free log subscription operation binding the contract event 0x34e0f289c6952b50ea58fa31d8c3732fb85091cfc3d5ded7767d7562ce0bfb22.
//
// Solidity: event DestinationRegistered(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 initialReserveImbalance, uint256 tokenMultiplier, bool multiplyOnReceive)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchDestinationRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationDestinationRegistered, destinationBlockchainID [][32]byte, destinationBridgeAddress []common.Address) (event.Subscription, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "DestinationRegistered", destinationBlockchainIDRule, destinationBridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationDestinationRegistered)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "DestinationRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDestinationRegistered is a log parse operation binding the contract event 0x34e0f289c6952b50ea58fa31d8c3732fb85091cfc3d5ded7767d7562ce0bfb22.
//
// Solidity: event DestinationRegistered(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 initialReserveImbalance, uint256 tokenMultiplier, bool multiplyOnReceive)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseDestinationRegistered(log types.Log) (*NativeTokenDestinationDestinationRegistered, error) {
	event := new(NativeTokenDestinationDestinationRegistered)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "DestinationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenDestination contract.
type NativeTokenDestinationMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenDestinationMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationMinTeleporterVersionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationMinTeleporterVersionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenDestination contract.
type NativeTokenDestinationMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenDestinationMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationMinTeleporterVersionUpdatedIterator{contract: _NativeTokenDestination.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationMinTeleporterVersionUpdated)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenDestinationMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenDestinationMinTeleporterVersionUpdated)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenDestination contract.
type NativeTokenDestinationOwnershipTransferredIterator struct {
	Event *NativeTokenDestinationOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenDestination contract.
type NativeTokenDestinationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenDestinationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationOwnershipTransferredIterator{contract: _NativeTokenDestination.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationOwnershipTransferred)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenDestinationOwnershipTransferred, error) {
	event := new(NativeTokenDestinationOwnershipTransferred)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportBurnedTxFeesIterator struct {
	Event *NativeTokenDestinationReportBurnedTxFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationReportBurnedTxFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationReportBurnedTxFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenDestinationReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationReportBurnedTxFeesIterator{contract: _NativeTokenDestination.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationReportBurnedTxFees)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportBurnedTxFees is a log parse operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenDestinationReportBurnedTxFees, error) {
	event := new(NativeTokenDestinationReportBurnedTxFees)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressPausedIterator struct {
	Event *NativeTokenDestinationTeleporterAddressPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTeleporterAddressPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTeleporterAddressPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenDestinationTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTeleporterAddressPausedIterator{contract: _NativeTokenDestination.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTeleporterAddressPaused)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenDestinationTeleporterAddressPaused, error) {
	event := new(NativeTokenDestinationTeleporterAddressPaused)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenDestinationTeleporterAddressUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTeleporterAddressUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTeleporterAddressUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenDestinationTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTeleporterAddressUnpausedIterator{contract: _NativeTokenDestination.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTeleporterAddressUnpaused)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenDestinationTeleporterAddressUnpaused, error) {
	event := new(NativeTokenDestinationTeleporterAddressUnpaused)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallSentIterator struct {
	Event *NativeTokenDestinationTokensAndCallSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensAndCallSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTokensAndCallSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e66562.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenDestinationTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensAndCallSentIterator{contract: _NativeTokenDestination.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e66562.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensAndCallSent)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e66562.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenDestinationTokensAndCallSent, error) {
	event := new(NativeTokenDestinationTokensAndCallSent)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensSentIterator struct {
	Event *NativeTokenDestinationTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensSent represents a TokensSent event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0xc117b57e5bc66f5f81fcb2b851a3ceee905d99d34038a194c791161d63e0ce71.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenDestinationTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensSentIterator{contract: _NativeTokenDestination.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0xc117b57e5bc66f5f81fcb2b851a3ceee905d99d34038a194c791161d63e0ce71.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensSent)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensSent is a log parse operation binding the contract event 0xc117b57e5bc66f5f81fcb2b851a3ceee905d99d34038a194c791161d63e0ce71.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensSent(log types.Log) (*NativeTokenDestinationTokensSent, error) {
	event := new(NativeTokenDestinationTokensSent)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensWithdrawnIterator struct {
	Event *NativeTokenDestinationTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenDestinationTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensWithdrawnIterator{contract: _NativeTokenDestination.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensWithdrawn)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenDestinationTokensWithdrawn, error) {
	event := new(NativeTokenDestinationTokensWithdrawn)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransferIterator struct {
	Event *NativeTokenDestinationTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTransfer represents a Transfer event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenDestinationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransferIterator{contract: _NativeTokenDestination.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTransfer)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTransfer(log types.Log) (*NativeTokenDestinationTransfer, error) {
	event := new(NativeTokenDestinationTransfer)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenDestination contract.
type NativeTokenDestinationWithdrawalIterator struct {
	Event *NativeTokenDestinationWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationWithdrawal represents a Withdrawal event raised by the NativeTokenDestination contract.
type NativeTokenDestinationWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenDestinationWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationWithdrawalIterator{contract: _NativeTokenDestination.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationWithdrawal)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseWithdrawal(log types.Log) (*NativeTokenDestinationWithdrawal, error) {
	event := new(NativeTokenDestinationWithdrawal)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
