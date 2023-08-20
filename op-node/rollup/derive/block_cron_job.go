package derive

import (
	"math/big"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ExecuteFunctionSignature = "executeCron()"
	ExecuteFunctionBytes4 = crypto.Keccak256([]byte(ExecuteFunctionSignature))[:4]
	ExecuteDepositerAddress = common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001")
	BlockCronJobsAddress = predeploys.BlockCronJobsAddr
)

func MarshalBinary() ([]byte, error) {
	data := make([]byte, 4)
	offset := 0
	copy(data[offset:4], ExecuteFunctionBytes4)
	return data, nil
}

func BlockCronJobDeposit(seqNumber uint64, gasLimit uint64, block eth.BlockInfo, regolith bool) (*types.DepositTx, error) {
	data, err := MarshalBinary()
	if err != nil {
		return nil, err
	}
	source := L1InfoDepositSource{
		L1BlockHash: block.Hash(),
		SeqNumber: seqNumber,
	}
	out := &types.DepositTx{
		SourceHash: source.SourceHash(),
		From: ExecuteDepositerAddress,
		To: &BlockCronJobsAddress,
		Mint: nil,
		Value: big.NewInt(0),
		Gas: gasLimit,
		IsSystemTransaction: true,
		Data: data,
	}

	if regolith {
		out.IsSystemTransaction = false
		out.Gas = RegolithSystemTxGas
	}
	return out, nil
}

func BlockCronJobDepositBytes(seqNumber uint64, gasLimit uint64, block eth.BlockInfo, regolith bool) ([]byte, error) {
	tx, err := BlockCronJobDeposit(seqNumber, gasLimit, block, regolith)
	if err != nil {
		return nil, err
	}
	l1Tx := types.NewTx(tx)
	return l1Tx.MarshalBinary()
}
