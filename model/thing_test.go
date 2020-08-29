package model

import (
	"reflect"
	"testing"

	"gitlab.com/MrCue/antlr_lifestyle/lifestyle"
)

func TestDB_AllThings(t *testing.T) {

	things2 := make([]*lifestyle.Thing, 0)
	things2 = append(things2, &lifestyle.Thing{Name: "No"})
	things2 = append(things2, &lifestyle.Thing{Name: "Yes"})

	type fields struct {
		DB DataStore
	}

	fdb := fakeDB{}

	tests := []struct {
		name    string
		fields  fields
		want    []*lifestyle.Thing
		wantErr bool
	}{
		{"TestFakeOne",
			fields{&fdb},
			things2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.DB
			got, err := d.AllThings()
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.AllThings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.AllThings() = %v, want %v", got, tt.want)
			}
		})
	}
}
