package bitcoin

import (
	b32 "github.com/tyler-smith/go-bip32"
)

// TODO: Build out this interface
type BitcoinWallet interface {
	// Keys
	GetMasterPrivateKey() *b32.Key
	GetMasterPublicKey() *b32.Key
}
