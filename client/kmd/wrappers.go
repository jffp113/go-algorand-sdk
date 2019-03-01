package kmd

import (
	"golang.org/x/crypto/ed25519"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
)

// Version returns a VersionResponse containing a list of kmd API versions
// supported by this running kmd instance.
func (kcl Client) Version() (resp VersionsResponse, err error) {
	req := VersionsRequest{}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ListWallets returns a ListWalletsResponse containing the list of wallets
// known to kmd. Using a wallet ID returned from this endpoint, you can
// initialize a wallet handle with client.InitWalletHandle
func (kcl Client) ListWallets() (resp ListWalletsResponse, err error) {
	req := ListWalletsRequest{}
	err = kcl.DoV1Request(req, &resp)
	return
}

// CreateWallet creates a wallet with the specified name, password, driver,
// and master derivation key. If the master derivation key is blank, one is
// generated internally to kmd. CreateWallet returns a CreateWalletResponse
// containing information about the new wallet.
func (kcl Client) CreateWallet(walletName, walletPassword, walletDriverName string, walletMDK types.MasterDerivationKey) (resp CreateWalletResponse, err error) {
	req := CreateWalletRequest{
		WalletName:          walletName,
		WalletDriverName:    walletDriverName,
		WalletPassword:      walletPassword,
		MasterDerivationKey: walletMDK,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// InitWalletHandle accepts a wallet ID and a wallet password, and returns an
// InitWalletHandleResponse containing a wallet handle token. This wallet
// handle token can be used for subsequent operations on this wallet, like key
// generation, transaction signing, etc.. WalletHandleTokens expire after a
// configurable number of seconds, and must be renewed periodically with
// RenewWalletHandle. It is good practice to call ReleaseWalletHandle when
// you're done interacting with this wallet.
func (kcl Client) InitWalletHandle(walletID, walletPassword string) (resp InitWalletHandleResponse, err error) {
	req := InitWalletHandleRequest{
		WalletID:       walletID,
		WalletPassword: walletPassword,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ReleaseWalletHandle invalidates the passed wallet handle token, making
// it unusable for subsequent wallet operations.
func (kcl Client) ReleaseWalletHandle(walletHandle string) (resp ReleaseWalletHandleResponse, err error) {
	req := ReleaseWalletHandleRequest{
		WalletHandleToken: walletHandle,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// RenewWalletHandle accepts a wallet handle and attempts to renew it, moving
// the expiration time to some number of seconds in the future. It returns a
// RenewWalletHandleResponse containing the walletHandle and the number of
// seconds until expiration
func (kcl Client) RenewWalletHandle(walletHandle string) (resp RenewWalletHandleResponse, err error) {
	req := RenewWalletHandleRequest{
		WalletHandleToken: walletHandle,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// RenameWallet accepts a wallet ID, wallet password, and a new wallet name,
// and renames the underlying wallet.
func (kcl Client) RenameWallet(walletID, walletPassword, newWalletName string) (resp RenameWalletResponse, err error) {
	req := RenameWalletRequest{
		WalletID: walletID,
		WalletPassword: walletPassword,
		NewWalletName: newWalletName,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// GetWallet accepts a wallet handle and returns high level information about
// this wallet in a GetWalletResponse.
func (kcl Client) GetWallet(walletHandle string) (resp GetWalletResponse, err error) {
	req := GetWalletRequest{
		WalletHandleToken: walletHandle,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ExportMasterDerivationKey accepts a wallet handle and a wallet password, and
// returns an ExportMasterDerivationKeyResponse containing the master
// derivation key. This key can be used as an argument to CreateWallet in
// order to recover the keys generated by this wallet. The master derivation
// key can be encoded as a sequence of words using the mnemonic library, and
// displayed to the user as a backup phrase.
func (kcl Client) ExportMasterDerivationKey(walletHandle, walletPassword string) (resp ExportMasterDerivationKeyResponse, err error) {
	req := ExportMasterDerivationKeyRequest{
		WalletHandleToken: walletHandle,
		WalletPassword:    walletPassword,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ImportKey accepts a wallet handle and an ed25519 private key, and imports
// the key into the wallet. It returns an ImportKeyResponse containing the
// address corresponding to this private key.
func (kcl Client) ImportKey(walletHandle string, secretKey ed25519.PrivateKey) (resp ImportKeyResponse, err error) {
	req := ImportKeyRequest{
		WalletHandleToken: walletHandle,
		PrivateKey:        secretKey,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ExportKey accepts a wallet handle, wallet password, and address, and returns
// an ExportKeyResponse containing the ed25519 private key corresponding to the
// address stored in the wallet.
func (kcl Client) ExportKey(walletHandle, walletPassword, addr string) (resp ExportKeyResponse, err error) {
	req := ExportKeyRequest{
		WalletHandleToken: walletHandle,
		WalletPassword: walletPassword,
		Address: addr,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// GenerateKey accepts a wallet handle, and then generates the next key in the
// wallet using its internal master derivation key. Two wallets with the same
// master derivation key will generate the same sequence of keys.
func (kcl Client) GenerateKey(walletHandle string) (resp GenerateKeyResponse, err error) {
	req := GenerateKeyRequest{
		WalletHandleToken: walletHandle,
		DisplayMnemonic:   false,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// DeleteKey accepts a wallet handle, wallet password, and address, and deletes
// the information about this address from the wallet (including address and
// secret key). If DeleteKey is called on a key generated using GenerateKey,
// the same key will not be generated again. However, if a wallet is recovered
// using the master derivation key, a key generated in this way can be
// recovered.
func (kcl Client) DeleteKey(walletHandle, walletPassword, addr string) (resp DeleteKeyResponse, err error) {
	req := DeleteKeyRequest{
		WalletHandleToken: walletHandle,
		WalletPassword: walletPassword,
		Address: addr,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ListKeys accepts a wallet handle and returns a ListKeysResponse containing
// all of the addresses for which this wallet contains secret keys.
func (kcl Client) ListKeys(walletHandle string) (resp ListKeysResponse, err error) {
	req := ListKeysRequest{
		WalletHandleToken: walletHandle,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// SignTransaction accepts a wallet handle, wallet password, and transaction,
// and returns and SignTransactionResponse containing an encoded, signed
// transaction. The transaction is signed using the key corresponding to the
// Sender field.
func (kcl Client) SignTransaction(walletHandle, walletPassword string, tx types.Transaction) (resp SignTransactionResponse, err error) {
	txBytes := msgpack.Encode(tx)
	req := SignTransactionRequest{
		WalletHandleToken: walletHandle,
		WalletPassword:    walletPassword,
		Transaction:       txBytes,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ListMultisig accepts a wallet handle and returns a ListMultisigResponse
// containing the multisig addresses whose preimages are stored in this wallet.
// A preimage is the information needed to reconstruct this multisig address,
// including multisig version information, threshold information, and a list
// of public keys.
func (kcl Client) ListMultisig(walletHandle string) (resp ListMultisigResponse, err error) {
	req := ListMultisigRequest{
		WalletHandleToken: walletHandle,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ImportMultisig accepts a wallet handle and the information required to
// generate a multisig address. It derives this address, and stores all of the
// information within the wallet. It returns a ImportMultisigResponse with the
// derived address.
func (kcl Client) ImportMultisig(walletHandle string, version, threshold uint8, pks []ed25519.PublicKey) (resp ImportMultisigResponse, err error) {
	req := ImportMultisigRequest{
		WalletHandleToken: walletHandle,
		Version:           version,
		Threshold:         threshold,
		PKs:               pks,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// ExportMultisig accepts a wallet handle, wallet password, and multisig
// address, and returns an ExportMultisigResponse containing the stored
// multisig preimage. The preimage contains all of the information necessary
// to derive the multisig address, including version, threshold, and a list of
// public keys.
func (kcl Client) ExportMultisig(walletHandle, walletPassword, addr string) (resp ExportMultisigResponse, err error) {
	req := ExportMultisigRequest{
		WalletHandleToken: walletHandle,
		WalletPassword: walletPassword,
		Address: addr,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}

// MultisigSignTransaction accepts a wallet handle, wallet password,
// transaction, public key (*not* an address), and an optional partial
// MultisigSig. It looks up the secret key corresponding to the public key, and
// returns a SignMultisigTransactionResponse containing a MultisigSig with a
// signature by the secret key included.
func (kcl Client) MultisigSignTransaction(walletHandle, walletPassword string, tx types.Transaction, pk ed25519.PublicKey, partial types.MultisigSig) (resp SignMultisigTransactionResponse, err error) {
	txBytes := msgpack.Encode(tx)
	req := SignMultisigTransactionRequest{
		WalletHandleToken: walletHandle,
		WalletPassword:    walletPassword,
		Transaction:       txBytes,
		PublicKey:         pk,
		PartialMsig:       partial,
	}
	err = kcl.DoV1Request(req, &resp)
	return
}
