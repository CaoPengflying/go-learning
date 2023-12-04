package facade

type IUserFacade interface {
	CreateUserAndWallet(name string)
}

type UserFacadeImpl struct {
}

func (u UserFacadeImpl) CreateUserAndWallet(name string) {
	var userService IUserService = UserServiceImpl{}
	userService.Register(name)

	var walletService IWalletService = WalletServiceImpl{}

	walletService.CreateWallet(name)
}
