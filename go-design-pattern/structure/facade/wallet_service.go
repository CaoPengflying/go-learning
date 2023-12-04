package facade

import "log"

type IWalletService interface {
	CreateWallet(name string)
}

type WalletServiceImpl struct {
}

func (w WalletServiceImpl) CreateWallet(name string) {
	log.Print("create wallet")
}
