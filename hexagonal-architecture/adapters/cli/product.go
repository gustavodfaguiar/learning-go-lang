package cli

import (
	"fmt"

	"github.com/gustavodfaguiar/learning-golang/hexagonal-architecture/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price%f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID:  %s", res.GetID())
	}
	return result, nil
}
