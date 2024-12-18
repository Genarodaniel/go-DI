package product

type ProductUseCase struct {
	repository IProductRepository
}

func NewProductUseCase(repository IProductRepository) *ProductUseCase {
	return &ProductUseCase{repository}
}

// This Product entity was not supposed to be returned. We should return a DTO instead
// However, we'll return as entity now to keep it simple.
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
