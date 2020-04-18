package model

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestGetDB(t *testing.T) {
	get1st, _ := GetDB()
	tests := []struct {
		name    string
		want    Env
		wantErr bool
	}{
		// TODO: Add test cases.
		{"GetDB_defaut", get1st, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
