// Code generated by mockery v1.0.0. DO NOT EDIT.
package accounts

import big "math/big"
import kowala "github.com/kowala-tech/kcoin/client"
import mock "github.com/stretchr/testify/mock"
import types "github.com/kowala-tech/kcoin/client/core/types"

// MockWallet is an autogenerated mock type for the Wallet type
type MockWallet struct {
	mock.Mock
}

// Accounts provides a mock function with given fields:
func (_m *MockWallet) Accounts() []Account {
	ret := _m.Called()

	var r0 []Account
	if rf, ok := ret.Get(0).(func() []Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Account)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockWallet) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Contains provides a mock function with given fields: account
func (_m *MockWallet) Contains(account Account) bool {
	ret := _m.Called(account)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Account) bool); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Derive provides a mock function with given fields: path, pin
func (_m *MockWallet) Derive(path DerivationPath, pin bool) (Account, error) {
	ret := _m.Called(path, pin)

	var r0 Account
	if rf, ok := ret.Get(0).(func(DerivationPath, bool) Account); ok {
		r0 = rf(path, pin)
	} else {
		r0 = ret.Get(0).(Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(DerivationPath, bool) error); ok {
		r1 = rf(path, pin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKeyedTransactor provides a mock function with given fields: account, auth
func (_m *MockWallet) NewKeyedTransactor(account Account, auth string) (*TransactOpts, error) {
	ret := _m.Called(account, auth)

	var r0 *TransactOpts
	if rf, ok := ret.Get(0).(func(Account, string) *TransactOpts); ok {
		r0 = rf(account, auth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TransactOpts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Account, string) error); ok {
		r1 = rf(account, auth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Open provides a mock function with given fields: passphrase
func (_m *MockWallet) Open(passphrase string) error {
	ret := _m.Called(passphrase)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(passphrase)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelfDerive provides a mock function with given fields: base, chain
func (_m *MockWallet) SelfDerive(base DerivationPath, chain kowala.ChainStateReader) {
	_m.Called(base, chain)
}

// SignHash provides a mock function with given fields: account, hash
func (_m *MockWallet) SignHash(account Account, hash []byte) ([]byte, error) {
	ret := _m.Called(account, hash)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(Account, []byte) []byte); ok {
		r0 = rf(account, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Account, []byte) error); ok {
		r1 = rf(account, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignHashWithPassphrase provides a mock function with given fields: account, passphrase, hash
func (_m *MockWallet) SignHashWithPassphrase(account Account, passphrase string, hash []byte) ([]byte, error) {
	ret := _m.Called(account, passphrase, hash)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(Account, string, []byte) []byte); ok {
		r0 = rf(account, passphrase, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Account, string, []byte) error); ok {
		r1 = rf(account, passphrase, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignProposal provides a mock function with given fields: account, proposal, chainID
func (_m *MockWallet) SignProposal(account Account, proposal *types.Proposal, chainID *big.Int) (*types.Proposal, error) {
	ret := _m.Called(account, proposal, chainID)

	var r0 *types.Proposal
	if rf, ok := ret.Get(0).(func(Account, *types.Proposal, *big.Int) *types.Proposal); ok {
		r0 = rf(account, proposal, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Proposal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Account, *types.Proposal, *big.Int) error); ok {
		r1 = rf(account, proposal, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignTx provides a mock function with given fields: account, tx, chainID
func (_m *MockWallet) SignTx(account Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(account, tx, chainID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(Account, *types.Transaction, *big.Int) *types.Transaction); ok {
		r0 = rf(account, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Account, *types.Transaction, *big.Int) error); ok {
		r1 = rf(account, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignTxWithPassphrase provides a mock function with given fields: account, passphrase, tx, chainID
func (_m *MockWallet) SignTxWithPassphrase(account Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(account, passphrase, tx, chainID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(Account, string, *types.Transaction, *big.Int) *types.Transaction); ok {
		r0 = rf(account, passphrase, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Account, string, *types.Transaction, *big.Int) error); ok {
		r1 = rf(account, passphrase, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignVote provides a mock function with given fields: account, vote, chainID
func (_m *MockWallet) SignVote(account Account, vote *types.Vote, chainID *big.Int) (*types.Vote, error) {
	ret := _m.Called(account, vote, chainID)

	var r0 *types.Vote
	if rf, ok := ret.Get(0).(func(Account, *types.Vote, *big.Int) *types.Vote); ok {
		r0 = rf(account, vote, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Vote)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Account, *types.Vote, *big.Int) error); ok {
		r1 = rf(account, vote, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields:
func (_m *MockWallet) Status() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URL provides a mock function with given fields:
func (_m *MockWallet) URL() URL {
	ret := _m.Called()

	var r0 URL
	if rf, ok := ret.Get(0).(func() URL); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(URL)
	}

	return r0
}