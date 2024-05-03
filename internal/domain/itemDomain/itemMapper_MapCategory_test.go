package itemDomain

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgtype"
	"reflect"
	"testing"
)

func TestMapDBItemWithCategoryToDTO(t *testing.T) {
	tests := []struct {
		name  string
		input *db.GetItemWithCategoryRow
		want  *dto.ItemWithCategory
	}{
		{
			name:  "NilInput",
			input: nil,
			want:  nil,
		},
		{
			name: "ValidInputWithCategory",
			input: &db.GetItemWithCategoryRow{
				CategoryID: pgtype.Int4{Int32: 1, Valid: true},
			},
			want: &dto.ItemWithCategory{
				CategoryID: pgtype.Int4{Int32: 1, Valid: true},
			},
		},
		{
			name: "ValidInputWithoutCategory",
			input: &db.GetItemWithCategoryRow{
				CategoryID: pgtype.Int4{Valid: false},
			},
			want: &dto.ItemWithCategory{
				CategoryID:          pgtype.Int4{Valid: false},
				CategoryName:        pgtype.Text{String: "Uncategorized", Valid: true},
				CategoryDescription: pgtype.Text{String: "An Uncategorized Item", Valid: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapDBItemWithCategoryToDTO(tt.input)

			if !equal(tt.want, got) {
				t.Errorf("MapDBItemWithCategoryToDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

// CPU profiled reflection vs custom function, reflection seems acceptable for this use case
// results approximately 47% slower than custom function but still under 1ms
func equal(a, b *dto.ItemWithCategory) bool {
	return reflect.DeepEqual(a, b)
}
