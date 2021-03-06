// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package transaction

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/kowala-tech/equilibrium/common/hexutil"
	"github.com/kowala-tech/equilibrium/crypto"
	"github.com/kowala-tech/equilibrium/state/accounts"
)

var _ = (*txDataMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (t txData) MarshalJSON() ([]byte, error) {
	type txData struct {
		AccountNonce hexutil.Uint64    `json:"accountNonce" gencodec:"required"`
		ComputeLimit hexutil.Uint64    `json:"computeLimit" gencodec:"required"`
		Recipient    *accounts.Address `json:"recipient"           rlp:"nil"`
		Amount       *hexutil.Big      `json:"amount"       gencodec:"required"`
		Payload      hexutil.Bytes     `json:"payload"      gencodec:"required"`
		V            *hexutil.Big      `json:"v" gencodec:"required"`
		R            *hexutil.Big      `json:"r" gencodec:"required"`
		S            *hexutil.Big      `json:"s" gencodec:"required"`
		Hash         *crypto.Hash      `json:"hash" rlp:"-"`
	}
	var enc txData
	enc.AccountNonce = hexutil.Uint64(t.AccountNonce)
	enc.ComputeLimit = hexutil.Uint64(t.ComputeLimit)
	enc.Recipient = t.Recipient
	enc.Amount = (*hexutil.Big)(t.Amount)
	enc.Payload = t.Payload
	enc.V = (*hexutil.Big)(t.V)
	enc.R = (*hexutil.Big)(t.R)
	enc.S = (*hexutil.Big)(t.S)
	enc.Hash = t.Hash
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *txData) UnmarshalJSON(input []byte) error {
	type txData struct {
		AccountNonce *hexutil.Uint64   `json:"accountNonce" gencodec:"required"`
		ComputeLimit *hexutil.Uint64   `json:"computeLimit" gencodec:"required"`
		Recipient    *accounts.Address `json:"recipient"           rlp:"nil"`
		Amount       *hexutil.Big      `json:"amount"       gencodec:"required"`
		Payload      *hexutil.Bytes    `json:"payload"      gencodec:"required"`
		V            *hexutil.Big      `json:"v" gencodec:"required"`
		R            *hexutil.Big      `json:"r" gencodec:"required"`
		S            *hexutil.Big      `json:"s" gencodec:"required"`
		Hash         *crypto.Hash      `json:"hash" rlp:"-"`
	}
	var dec txData
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.AccountNonce == nil {
		return errors.New("missing required field 'accountNonce' for txData")
	}
	t.AccountNonce = uint64(*dec.AccountNonce)
	if dec.ComputeLimit == nil {
		return errors.New("missing required field 'computeLimit' for txData")
	}
	t.ComputeLimit = uint64(*dec.ComputeLimit)
	if dec.Recipient != nil {
		t.Recipient = dec.Recipient
	}
	if dec.Amount == nil {
		return errors.New("missing required field 'amount' for txData")
	}
	t.Amount = (*big.Int)(dec.Amount)
	if dec.Payload == nil {
		return errors.New("missing required field 'payload' for txData")
	}
	t.Payload = *dec.Payload
	if dec.V == nil {
		return errors.New("missing required field 'v' for txData")
	}
	t.V = (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for txData")
	}
	t.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for txData")
	}
	t.S = (*big.Int)(dec.S)
	if dec.Hash != nil {
		t.Hash = dec.Hash
	}
	return nil
}
