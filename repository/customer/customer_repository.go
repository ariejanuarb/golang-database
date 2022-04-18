package repository

// repository adalah perilaku yg akan diberikan pada entity
import (
	"context"
	"go-database/entity"
)

type CustomerRepository interface { // interface = deklarasi dari method
	Insert(ctx context.Context, name entity.Customer) (entity.Customer, error)
	FindById(ctx context.Context, id int32) (entity.Customer, error)
	FindAll(ctx context.Context) ([]entity.Customer, error)
	Update(ctx context.Context, name entity.Customer) (entity.Customer, error)
	Delete(ctx context.Context, name entity.Customer) (entity.Customer, error)
}
