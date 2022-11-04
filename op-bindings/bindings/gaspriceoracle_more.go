// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":29883,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":2206,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"spacer_1_0_32\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":2209,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"spacer_2_0_32\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":2212,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"spacer_3_0_32\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint256\"},{\"astId\":2215,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"spacer_4_0_32\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100d45760003560e01c8063715018a611610081578063f2fde38b1161005b578063f2fde38b14610177578063f45e65d81461018a578063fe173b971461012c57600080fd5b8063715018a6146101325780638da5cb5b1461013c578063de26c4a11461016457600080fd5b8063519b4bd3116100b2578063519b4bd31461010f57806354fd4d50146101175780636ef25c3a1461012c57600080fd5b80630c18c162146100d9578063313ce567146100f457806349948e0e146100fc575b600080fd5b6100e1610192565b6040519081526020015b60405180910390f35b6100e1600681565b6100e161010a3660046107c3565b61021c565b6100e161027d565b61011f6102de565b6040516100eb91906108c2565b486100e1565b61013a610381565b005b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100eb565b6100e16101723660046107c3565b610395565b61013a610185366004610913565b610444565b6100e1610500565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102179190610950565b905090565b60008061022883610395565b9050600061023461027d565b61023e9083610998565b9050600061024e6006600a610af7565b9050600061025a610500565b6102649084610998565b905060006102728383610b32565b979650505050505050565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101f3573d6000803e3d6000fd5b60606103097f0000000000000000000000000000000000000000000000000000000000000000610561565b6103327f0000000000000000000000000000000000000000000000000000000000000000610561565b61035b7f0000000000000000000000000000000000000000000000000000000000000000610561565b60405160200161036d93929190610b46565b604051602081830303815290604052905090565b61038961069e565b610393600061071f565b565b80516000908190815b81811015610418578481815181106103b8576103b8610bbc565b01602001517fff00000000000000000000000000000000000000000000000000000000000000166000036103f8576103f1600484610beb565b9250610406565b610403601084610beb565b92505b8061041081610c03565b91505061039e565b506000610423610192565b61042d9084610beb565b905061043b81610440610beb565b95945050505050565b61044c61069e565b73ffffffffffffffffffffffffffffffffffffffff81166104f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6104fd8161071f565b50565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101f3573d6000803e3d6000fd5b6060816000036105a457505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156105ce57806105b881610c03565b91506105c79050600a83610b32565b91506105a8565b60008167ffffffffffffffff8111156105e9576105e9610794565b6040519080825280601f01601f191660200182016040528015610613576020820181803683370190505b5090505b841561069657610628600183610c3b565b9150610635600a86610c52565b610640906030610beb565b60f81b81838151811061065557610655610bbc565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061068f600a86610b32565b9450610617565b949350505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104eb565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156107d557600080fd5b813567ffffffffffffffff808211156107ed57600080fd5b818401915084601f83011261080157600080fd5b81358181111561081357610813610794565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561085957610859610794565b8160405282815287602084870101111561087257600080fd5b826020860160208301376000928101602001929092525095945050505050565b60005b838110156108ad578181015183820152602001610895565b838111156108bc576000848401525b50505050565b60208152600082518060208401526108e1816040850160208701610892565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561092557600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461094957600080fd5b9392505050565b60006020828403121561096257600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156109d0576109d0610969565b500290565b600181815b80851115610a2e57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610a1457610a14610969565b80851615610a2157918102915b93841c93908002906109da565b509250929050565b600082610a4557506001610af1565b81610a5257506000610af1565b8160018114610a685760028114610a7257610a8e565b6001915050610af1565b60ff841115610a8357610a83610969565b50506001821b610af1565b5060208310610133831016604e8410600b8410161715610ab1575081810a610af1565b610abb83836109d5565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610aed57610aed610969565b0290505b92915050565b60006109498383610a36565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610b4157610b41610b03565b500490565b60008451610b58818460208901610892565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610b94816001850160208a01610892565b60019201918201528351610baf816002840160208801610892565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115610bfe57610bfe610969565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c3457610c34610969565b5060010190565b600082821015610c4d57610c4d610969565b500390565b600082610c6157610c61610b03565b50069056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
}
