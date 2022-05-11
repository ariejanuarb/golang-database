package repository

// repository adalah perilaku yg akan diberikan pada entity
import (
	"context"
	"go-database/entity"
)

type ProductRepository interface { // interface = deklarasi dari method
	Insert(ctx context.Context, name entity.Product) (entity.Product, error)
	FindById(ctx context.Context, id int32) (entity.Product, error)
	FindAll(ctx context.Context) ([]entity.Product, error)
	Update(ctx context.Context, name entity.Product) (entity.Product, error)
	Delete(ctx context.Context, id int32) (int32, error)
}
