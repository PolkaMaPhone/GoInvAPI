package itemDomain

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgtype"
	"reflect"
	"testing"
	"time"
)

func TestMapDBItemToDomainItem(t *testing.T) {
	tests := []struct {
		name string
		args *db.Item
		want *Item
	}{
		{
			name: "AllFieldsFilled",
			args: &db.Item{
				ItemID: 1,
				Name:   "Item1",
				Description: pgtype.Text{
					String: "Description1",
					Valid:  true,
				},
				CategoryID: pgtype.Int4{
					Int32: 1,
					Valid: true,
				},
				GroupID: pgtype.Int4{
					Int32: 1,
					Valid: true,
				},
				LocationID: pgtype.Int4{
					Int32: 1,
					Valid: true,
				},
				IsStored: pgtype.Bool{
					Bool:  true,
					Valid: true,
				},
				CreatedAt: pgtype.Timestamptz{Time: time.Now()},
				UpdatedAt: pgtype.Timestamptz{Time: time.Now()},
			},
			want: &Item{
				ItemID: 1,
				Name:   "Item1",
				Description: pgtype.Text{
					String: "Description1",
					Valid:  true,
				},
				CategoryID: pgtype.Int4{
					Int32: 1,
					Valid: true,
				},
				GroupID: pgtype.Int4{
					Int32: 1,
					Valid: true,
				},
				LocationID: pgtype.Int4{
					Int32: 1,
					Valid: true,
				},
				IsStored:  pgtype.Bool{Bool: true, Valid: true},
				CreatedAt: pgtype.Timestamptz{Time: time.Now()},
				UpdatedAt: pgtype.Timestamptz{Time: time.Now()},
			},
		},
		{
			name: "EmptyFields",
			args: &db.Item{},
			want: &Item{},
		},
		{
			name: "EmptyDBItem",
			args: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := MapDBItemToDomainItem(tt.args)
			if got != nil && tt.want != nil {
				// Check the type of the CreatedAt and UpdatedAt fields
				if reflect.TypeOf(got.CreatedAt) != reflect.TypeOf(tt.want.CreatedAt) {
					t.Errorf("CreatedAt field type = %v, want %v", reflect.TypeOf(got.CreatedAt), reflect.TypeOf(tt.want.CreatedAt))
				}
				if reflect.TypeOf(got.UpdatedAt) != reflect.TypeOf(tt.want.UpdatedAt) {
					t.Errorf("UpdatedAt field type = %v, want %v", reflect.TypeOf(got.UpdatedAt), reflect.TypeOf(tt.want.UpdatedAt))
				}

				// Set the CreatedAt and UpdatedAt fields to nil before comparing the rest of the fields
				got.CreatedAt = pgtype.Timestamptz{}
				got.UpdatedAt = pgtype.Timestamptz{}
				tt.want.CreatedAt = pgtype.Timestamptz{}
				tt.want.UpdatedAt = pgtype.Timestamptz{}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapDBItemToDomainItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
