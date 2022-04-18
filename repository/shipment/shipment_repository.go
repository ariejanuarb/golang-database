package repository

// repository adalah perilaku yg akan diberikan pada entity
import (
	"context"
	"go-database/entity"
)

type ShipmentRepository interface { // interface = deklarasi dari method
	Insert(ctx context.Context, adress entity.Shipment) (entity.Shipment, error)
	FindById(ctx context.Context, id int32) (entity.Shipment, error)
	FindAll(ctx context.Context) ([]entity.Shipment, error)
	Update(ctx context.Context, adress entity.Shipment) (entity.Shipment, error)
	Delete(ctx context.Context, adress entity.Shipment) (entity.Shipment, error)
}
