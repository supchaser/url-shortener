package usecase

import (
	"context"
	"fmt"
	"url-shortener/internal/models"
	"url-shortener/internal/pkg/shortener"
)

type ShortenerUsecase struct {
	shortenerRepo shortener.ShortenerRepository
}

func CreateShortenerUsecase(shortenerRepo shortener.ShortenerRepository) *ShortenerUsecase {
	return &ShortenerUsecase{
		shortenerRepo: shortenerRepo,
	}
}

// TODO: добавить проверку на права пользователя
func (u *ShortenerUsecase) SaveURL(ctx context.Context, urlToSave string, alias string) (newURLStruct *models.URLStruct, err error) {
	funcName := "SaveURL"
	newURLStruct, err = u.shortenerRepo.SaveURL(ctx, urlToSave, alias)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", funcName, err)
	}

	return newURLStruct, err
}

func (u *ShortenerUsecase) GetURL(ctx context.Context, alias string) (url *models.URLStruct, err error) {
	funcName := "GetURL"
	url, err = u.shortenerRepo.GetURL(ctx, alias)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", funcName, err)
	}

	return url, err
}

// TODO: добавить проверку на права пользователя
func (u *ShortenerUsecase) DeleteURL(ctx context.Context, alias string) (err error) {
	funcName := "DeleteURL"
	err = u.shortenerRepo.DeleteURL(ctx, alias)
	if err != nil {
		return fmt.Errorf("%s : %w", funcName, err)
	}

	return err
}
