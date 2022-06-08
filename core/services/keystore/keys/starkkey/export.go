package starkkey

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	keys_export "github.com/smartcontractkit/chainlink/core/services/keystore/keys"
	"github.com/smartcontractkit/chainlink/core/utils"
)

const keyTypeIdentifier = "StarkNet"

// FromEncryptedJSON gets key from json and password
func FromEncryptedJSON(keyJSON []byte, password string) (Key, error) {
	return keys_export.FromEncryptedJSON[keys_export.EncryptedKeyExport](
		keyTypeIdentifier,
		keyJSON,
		password,
		adulteratedPassword,
		func(_ keys_export.EncryptedKeyExport, rawPrivKey []byte) (Key, error) {
			return Raw(rawPrivKey).Key(), nil
		},
	)
}

// ToEncryptedJSON returns encrypted JSON representing key
func (key Key) ToEncryptedJSON(password string, scryptParams utils.ScryptParams) (export []byte, err error) {
	return keys_export.ToEncryptedJSON[keys_export.EncryptedKeyExport](
		keyTypeIdentifier,
		key.Raw(),
		key,
		password,
		scryptParams,
		adulteratedPassword,
		func(id string, key Key, cryptoJSON keystore.CryptoJSON) (keys_export.EncryptedKeyExport, error) {
			return keys_export.EncryptedKeyExport{
				KeyType:   id,
				PublicKey: key.PublicKeyStr(),
				Crypto:    cryptoJSON,
			}, nil
		},
	)
}

func adulteratedPassword(password string) string {
	return "starkkey" + password
}
