package stellar

import (
	"github.com/interstellar/starlight/worizon/xlm"
	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/xdr"
)

// BuildPegInTx builds a slidechain peg-in transaction
func BuildPegInTx(source string, txvmPubkey [32]byte, amount xlm.Amount, destination string, hclient *horizon.Client) (*b.TransactionBuilder, error) {
	root, err := hclient.Root()
	if err != nil {
		return nil, err
	}
	return b.Transaction(
		b.Network{Passphrase: root.NetworkPassphrase},
		b.SourceAccount{AddressOrSeed: source},
		b.AutoSequence{SequenceProvider: hclient},
		b.BaseFee{Amount: 100},
		b.MemoHash{Value: xdr.Hash(txvmPubkey)},
		b.Payment(
			b.Destination{AddressOrSeed: destination},
			b.NativeAmount{Amount: amount.HorizonString()},
		),
	)
}
