package structs

type Product struct {
	id          string
	name        string
	description string
	price       string
	maxCount    int
}

type Shirt struct {
	Product
}

type Shoe struct {
	Product
}

type Phone struct {
	Product
}

type ProductService interface {
	Get(string) (Product, error)
	GetAll() ([]Product, error)
	Create(Product) (Product, error)
	Delete(string) (bool, error)
}

type ProductServiceImpl struct {
}

func (p *ProductServiceImpl) Get(id string) (product Product, err error) {
	product = *new(Product)
	return
}

func (p *ProductServiceImpl) GetAll(id string) (product []Product, err error) {
	product = make([]Product, 0)
	return
}

func (p *ProductServiceImpl) Create(in Product) (out Product, err error) {
	in.id = "100"
	return
}
