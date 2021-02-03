package usecase

import "github.com/wwwillian/codepix-go/domain/model"

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pixUseCase *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := pixUseCase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	pixUseCase.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (pixUseCase *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pixUseCase.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

