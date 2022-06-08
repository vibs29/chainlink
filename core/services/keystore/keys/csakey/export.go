package csakey

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	keys_export "github.com/smartcontractkit/chainlink/core/services/keystore/keys"
	"github.com/smartcontractkit/chainlink/core/utils"
)

const keyTypeIdentifier = "CSA"

func FromEncryptedJSON(keyJSON []byte, password string) (KeyV2, error) {
	return keys_export.FromEncryptedJSON[keys_export.EncryptedKeyExport](
		keyTypeIdentifier,
		keyJSON,
		password,
		adulteratedPassword,
		func(_ keys_export.EncryptedKeyExport, rawPrivKey []byte) (KeyV2, error) {
			return Raw(rawPrivKey).Key(), nil
		},
	)
}

func (key KeyV2) ToEncryptedJSON(password string, scryptParams utils.ScryptParams) (export []byte, err error) {
	return keys_export.ToEncryptedJSON[keys_export.EncryptedKeyExport](
		keyTypeIdentifier,
		key.Raw(),
		key,
		password,
		scryptParams,
		adulteratedPassword,
		func(id string, key KeyV2, cryptoJSON keystore.CryptoJSON) (keys_export.EncryptedKeyExport, error) {
			return keys_export.EncryptedKeyExport{
				KeyType:   id,
				PublicKey: key.PublicKeyString(),
				Crypto:    cryptoJSON,
			}, nil
		},
	)
}

func adulteratedPassword(password string) string {
	return "csakey" + password
}
