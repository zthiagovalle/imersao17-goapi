package service

import (
	"github.com/zthiagovalle/imersao17/goapi/internal/database"
	"github.com/zthiagovalle/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

func (p *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := p.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := p.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := p.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) CreateProduct(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryID, imageURL, price)
	_, err := p.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
