package repository

// repository adalah perilaku yg akan diberikan pada entity
import (
	"context"
	"go-database/entity"
)

type PaymentRepository interface { // interface = deklarasi dari method
	Insert(ctx context.Context, method entity.Payment) (entity.Payment, error)
	FindById(ctx context.Context, id int32) (entity.Payment, error)
	FindAll(ctx context.Context) ([]entity.Payment, error)
	Update(ctx context.Context, method entity.Payment) (entity.Payment, error)
	Delete(ctx context.Context, method entity.Payment) (entity.Payment, error)
}