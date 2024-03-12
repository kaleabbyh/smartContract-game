package config

import (
	// "context"
	// "crypto/ecdsa"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	// "math"
	// "math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var Ks *keystore.KeyStore
var Admin string
var hasRun bool

func GetClient() (*ethclient.Client, error) {
	if Client != nil {
		return Client, nil
	}
	nodeUrl := "https://ethereum-sepolia.publicnode.com"
	var err error
	Client, err = ethclient.Dial(nodeUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to dial client, error: %w", err)
	}

	return Client, nil
}

func CreateKs() error {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {

		return fmt.Errorf("failed to create keystore, error: %w", err)
	}

	fmt.Println(account.Address.Hex())
	return nil
}

func RestoreKs() (*keystore.KeyStore, error) {
	dirPath, err := filepath.Abs("./tmp2")
	if err != nil {
		fmt.Printf("Absolute Path not found: 1 \n")
	}
	fmt.Println(dirPath)
	ks := keystore.NewKeyStore("./tmp2", keystore.StandardScryptN, keystore.StandardScryptP)

	var password string
	PASSWORD_jsonBytes, err := os.ReadFile(os.Getenv("PASSWORD"))
	if err != nil {
		fmt.Printf("failed to read password file: no such file \n")
		password = os.Getenv("PASSWORD")
	} else {
		password = string(PASSWORD_jsonBytes)
	}

	dirPath2, err := filepath.Abs("./tmp")
	if err != nil {
		fmt.Printf("Absolute Path not found: 2 \n")
	}
	fmt.Println(dirPath2)
	files, err := os.ReadDir("./tmp")
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s", err.Error())
	}
	// var jsonBytes []byte
	for _, file := range files {
		jsonBytes, err := os.ReadFile("./tmp" + "/" + file.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %s", err.Error())
		}
		fmt.Println("Creating Keystore...")
		account1, err := ks.Import(jsonBytes, password, password)
		if err != nil {
			return nil, errors.New("failed to restore accounts")
		}
		if ks.Accounts() != nil {
			fmt.Println("Account-1: ", account1.Address.Hex())
			Ks = ks
			// return ks, nil
		}
		// jsonAcc, err := ks.Export(account1, password, password)
		// if err != nil {

		// 	return nil, fmt.Errorf("failed to export keystore, error: %w", err)
		// }
		// if err := os.WriteFile("./env/keystore.txt", jsonAcc, 0644); err != nil {

		// 	return nil, fmt.Errorf("failed to create keystore file, error: %w", err)
		// }
		// os.Remove file in tmp, setDir to ./tmp, and run the above functions again (importing from ./tmp2 to ./tmp)
		if err := os.Remove("./tmp" + "/" + file.Name()); err != nil {
			return nil, fmt.Errorf("failed to remove keystore file: %s", err.Error())
		}
		ks2 := keystore.NewKeyStore("./tmp"+"/", keystore.StandardScryptN, keystore.StandardScryptP)
		account, err := ks2.Import(jsonBytes, password, password)
		if err != nil {
			return nil, errors.New("failed to replace keystore")
		}
		if err := os.RemoveAll("./tmp2"); err != nil {
			fmt.Printf("failed to  remove dirPath: %s \n", err)
		}

		err = ks2.Unlock(account, password)
		if err != nil {
			// fmt.Println("NewKeystoreWallet - ", account.Address.Hex(), " could not be unlock: ", err.Error())
			return nil, fmt.Errorf("NewKeystoreWallet - %s could not be unlock: %s", account.Address.Hex(), err.Error())
		}
		if ks2.Accounts() != nil {
			fmt.Println("Account-2: ", account.Address.Hex())
			Ks = ks2
			return ks2, nil
		}
	}
	return nil, errors.New("keystore has no accounts")
}
func ImportKs() (*keystore.KeyStore, error) {
	if Ks != nil {
		return Ks, nil
	}
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	// file := "./tmp/UTC--2024-02-07T20-05-31.183939189Z--a44a4205019cafb1694debce99e72d0c67a5ec4c"

	var password string
	PASSWORD_jsonBytes, err := os.ReadFile(os.Getenv("PASSWORD"))
	if err != nil {
		fmt.Printf("Failed to read password file: %s \n", err.Error())
		password = os.Getenv("PASSWORD")
	} else {
		password = string(PASSWORD_jsonBytes)
	}

	Admin = os.Getenv("ADMIN")

	a1 := accounts.Account{}
	a1.Address = common.HexToAddress(Admin)
	var account accounts.Account
	// if os.Getenv start with /run import ks
	if strings.HasPrefix(os.Getenv("KEYSTORE"), "/run") && !hasRun {

		jsonBytes, err := os.ReadFile(os.Getenv("KEYSTORE"))
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %s", err.Error())
		}
		account, err = ks.Import(jsonBytes, password, password)
		if err != nil {
			return nil, errors.Join(errors.New("failed to import keystore: "), err)
		}
		hasRun = true
	} else if strings.HasPrefix(os.Getenv("KEYSTORE"), "{") && !hasRun {
		jsonBytes := []byte(os.Getenv("KEYSTORE"))
		_, err := os.Stat("./tmp")
		if err != nil {
			if os.IsNotExist(err) {
				account, err = ks.Import(jsonBytes, password, password)
				if err != nil {
					return nil, fmt.Errorf("failed to import keystore: %s", err)
				}
				hasRun = true
			} else {
				return nil, fmt.Errorf("failed to read file: %s", err.Error())
			}
		} else {
			account, err = ks.Find(a1)
			if err != nil {
				return nil, fmt.Errorf("account not found: %s", err.Error())
			}
		}
	} else {
		account, err = ks.Find(a1)
		if err != nil {
			return nil, fmt.Errorf("account not found: %s", err.Error())
		}
	}

	err = ks.Unlock(account, password)
	if err != nil {
		// fmt.Println("NewKeystoreWallet - ", account.Address.Hex(), " could not be unlock: ", err.Error())
		return nil, fmt.Errorf("newKeystoreWallet - %s could not be unlock: %s", account.Address.Hex(), err.Error())
	}

	// sig, err := ks.SignHashWithPassphrase(account, password, accounts.TextHash([]byte("Msg to sign")))
	// if err != nil {
	// 	fmt.Println("NewKeystoreWallet - ", account.Address.Hex(), " could not be unlock: ", err.Error())
	// }
	// fmt.Println("signature: ", string(sig))
	if ks.Accounts() != nil {
		fmt.Println("Account: ", account.Address.Hex())
		Ks = ks
		hasRun = true
		return ks, nil
	}

	// accounts, err := ks.Import(jsonBytes, password, password)
	// if err != nil {
	// 	return nil, errors.Join(errors.New("Failed to import keystore: "), err)
	// }

	// fmt.Println(accounts.Address.Hex())

	// if err := os.Remove(file); err != nil {
	// 	return nil, errors.Join(errors.New("Failed to overwrite keystore: "), err)
	// }

	return nil, errors.New("keystore has no accounts")
}
