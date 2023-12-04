package facade

import "testing"

func TestUserFacadeImpl_CreateUserAndWallet(t *testing.T) {
	var facadeService IUserFacade = UserFacadeImpl{}
	facadeService.CreateUserAndWallet("cpf")
}
