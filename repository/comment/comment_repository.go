package comment

// repository adalah perilaku yg akan diberikan pada entity
import (
	"context"
	"go-database/entity"
)

type CommentRepository interface { // interface = deklarasi dari method
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	Update(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	Delete(ctx context.Context, id int32) (int32, error)
}
