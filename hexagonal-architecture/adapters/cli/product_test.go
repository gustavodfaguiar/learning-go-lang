package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gustavodfaguiar/learning-golang/hexagonal-architecture/adapters/cli"
	"github.com/gustavodfaguiar/learning-golang/hexagonal-architecture/application"
	mock_application "github.com/gustavodfaguiar/learning-golang/hexagonal-architecture/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := application.Product{
		Name:   "Package test",
		Price:  24.99,
		Status: "enabled",
		ID:     "abc",
	}

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(product.ID).AnyTimes()
	productMock.EXPECT().GetStatus().Return(product.Status).AnyTimes()
	productMock.EXPECT().GetName().Return(product.Name).AnyTimes()
	productMock.EXPECT().GetPrice().Return(product.Price).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(product.Name, product.Price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(product.ID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price%f and status %s",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	result, err := cli.Run(service, "create", "", product.Name, product.Price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled.", product.GetName())
	result, err = cli.Run(service, "enable", product.ID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID:  %s", product.ID)
	result, err = cli.Run(service, "get", product.ID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
