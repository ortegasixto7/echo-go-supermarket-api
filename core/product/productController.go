package product

import (
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/requests"
	"github.com/ortegasixto7/echo-go-supermarket-api/core/product/validations"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductService ProductService
}

func (this ProductController) AddQuantity(request *requests.AddProductQuantityRequest) (requestError string, errorCode string) {
	// requestError, errorCode = new(validations.UpdateProductRequest).Validate(request)
	// if requestError != "" {
	// 	return requestError, errorCode
	// }
	productResult := this.ProductService.GetById(request.Id)
	productResult.Quantity += request.Quantity
	this.ProductService.Update(productResult)
	return requestError, errorCode
}

func (this ProductController) GetById(id string) Product {
	return this.ProductService.GetById(id)
}

func (this ProductController) GetAll() []Product {
	return this.ProductService.GetAll()
}

func (this ProductController) Update(request *requests.UpdateProductRequest) (requestError string, errorCode string) {
	// requestError, errorCode = new(validations.UpdateProductRequest).Validate(request)
	// if requestError != "" {
	// 	return requestError, errorCode
	// }
	productResult := this.ProductService.GetById(request.Id)
	if request.Name != "" {
		productResult.Name = request.Name
	}
	if request.Description != "" {
		productResult.Description = request.Description
	}
	if request.Price != 0 {
		productResult.Price = request.Price
	}
	this.ProductService.Update(productResult)
	return requestError, errorCode
}

func (this ProductController) Create(request *requests.CreateProductRequest) (requestError string, errorCode string) {
	requestError, errorCode = new(validations.CreateProductRequestValidation).Validate(request)
	if requestError != "" {
		return requestError, errorCode
	}
	product := Product{
		Id:          primitive.NewObjectID().Hex(),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Quantity:    0}
	this.ProductService.Save(product)
	return requestError, errorCode
}
