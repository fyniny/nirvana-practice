package handlers

import (
	"context"

	meta "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	"github.com/caicloud/nirvana-practice/pkg/apis/middlewares"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

// CreateProduct returns error if the product's key existed
func CreateProduct(ctx context.Context, prod *api.Product) error {
	if prod == nil {
		return  errors.ErrorInvalidParameter.Error("param \"production\" is nil")
	}

	cache := middlewares.GetCache(ctx)

	cache.Lock()
	defer cache.Unlock()

	if _, ok := cache.Products[prod.Name]; ok {
		return errors.ErrorAlreadyExist.Error()
	}

	cache.Products[prod.Name] = prod
	return nil
}

// GetProduct returns product value whose name is equal to ${name}
func GetProduct (ctx context.Context, name string) (*api.Product, error) {
	var (
		product *api.Product
		productInterface interface{}
		ok bool
	)

	cache := middlewares.GetCache(ctx)

	if name == "" {
		return nil, errors.ErrorInvalidParameter.Error("name")
	}

	productInterface, ok = cache.Products[name]
	if !ok {
		return nil, errors.ErrorNotFound.Error()
	}

	product = productInterface.(*api.Product)
	return product, nil
}


// ListProducts returns product list
func ListProducts (ctx context.Context, options *meta.ListOptions) ([]*api.Product, error) {
	var (
		list []*api.Product
		i = 0
		min = options.Limit
		cache = middlewares.GetCache(ctx)
	)

	if min <= 0 {
		return nil, errors.ErrorInvalidParameter.Error("limit")
	}

	if min > len(cache.Products) {
		min = len(cache.Products)
	}

	list = make([]*api.Product, 0, min)

	for _, val := range cache.Products {
		list = append(list, val)
		i++
		if i > min {
			break
		}
	}

	return list, nil
}

// UpdateProduct will replace the product named ${oldName} in m.products
// returns error when ${oldName} doesn't exist
func UpdateProduct (ctx context.Context, product *api.Product, oldName string) error {
	if oldName == "" || product == nil {
		return errors.ErrorInvalidParameter.Error("oldName or product invalid")
	}

	cache := middlewares.GetCache(ctx)

	_, ok := cache.Products[oldName]
	if !ok {
		return errors.ErrorNotFound.Error()
	}

	cache.Lock()
	defer cache.Unlock()

	if oldName != product.Name {
		_, ok := cache.Products[product.Name]
		if ok {
			return errors.ErrorAlreadyExist.Error()
		}
	}

	delete(cache.Products, oldName)
	cache.Products[product.Name] = product

	return nil
}

// Delete product from m.products
// returns error if name is not in m.products
func DeleteProduct (ctx context.Context, name string) error {
	cache := middlewares.GetCache(ctx)
	_, ok := cache.Products[name]
	if !ok {
		return errors.ErrorNotFound.Error()
	}

	cache.Lock()
	delete(cache.Products, name)
	cache.Unlock()

	return nil
}
