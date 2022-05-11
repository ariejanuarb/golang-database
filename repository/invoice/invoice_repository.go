package repository

// repository adalah perilaku yg akan diberikan pada entity
import (
	"context"
	"go-database/entity"
)

type InvoiceRepository interface { // interface = deklarasi dari method
	Insert(ctx context.Context, no_invoice entity.Invoice) (entity.Invoice, error)
	FindById(ctx context.Context, id int32) (entity.Invoice, error)
	FindAll(ctx context.Context) ([]entity.Invoice, error)
	Update(ctx context.Context, no_invoice entity.Invoice) (entity.Invoice, error)
	Delete(ctx context.Context, id int32) (int32, error)
}
