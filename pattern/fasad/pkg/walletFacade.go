package pkg

import (
	"errors"
	"fmt"
)

type walletFacade struct {
	account      account
	securityCode securityCode
	wallet       wallet
}

func CreateWallet(idAccount int, accountName string, phone string, securityCode string, walletId int) *walletFacade {
	return &walletFacade{
		account:      *createAccount(idAccount, accountName, phone),
		securityCode: *newSecurityCode(securityCode),
		wallet:       *newWallet(walletId),
	}
}

func (w *walletFacade) AddMoney(phone, securityCode string, money int) error {
	if !w.account.checkAccountByPhone(phone) {
		return errors.New("неверный номмер телефона")
	}
	if !w.securityCode.checkSecurityCode(securityCode) {
		return errors.New("неверный ключ")
	}
	w.wallet.AddMoney(money)
	fmt.Println(fmt.Sprintf("На ваш акаунт: %v, была добалена сумма %d \n  состояние счета:%d", w.account.nameAccount, money, w.wallet.CheckBalance()))
	return nil
}
