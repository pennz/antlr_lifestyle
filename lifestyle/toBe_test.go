package lifestyle

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func TestToBe_AddReminder(t *testing.T) {
	defer func() {
		fmt.Print(&buf)
	}()

	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		in0 time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"basic", fields{logger}, args{time.Now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.AddReminder(tt.args.in0)
		})
	}
}

func TestToBe_Countdown(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			if got := test.Countdown(); got != tt.want {
				t.Errorf("ToBe.Countdown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBe_AddResourceUsage(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.AddResourceUsage()
		})
	}
}

func TestToBe_Serialize(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.Serialize()
		})
	}
}

func TestToBe_GetProcess(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.GetProcess()
		})
	}
}

func TestToBe_Review(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			if got := test.Review(); got != tt.want {
				t.Errorf("ToBe.Review() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBe_Exalate(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.Exalate()
		})
	}
}

func TestToBe_Suspend(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.Suspend()
		})
	}
}

func TestToBe_Continue(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.Continue()
		})
	}
}

func TestToBe_AddMilestone(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.AddMilestone()
		})
	}
}

func TestToBe_CheckMilestone(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.CheckMilestone()
		})
	}
}

func TestToBe_AddTag(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		in0 Tag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.AddTag(tt.args.in0)
		})
	}
}

func TestToBe_RemoveTag(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		in0 Tag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.RemoveTag(tt.args.in0)
		})
	}
}

func TestToBe_NewToDo(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   *ToDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			if got := test.NewToDo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBe.NewToDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBe_Do(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			if got := test.Do(); got != tt.want {
				t.Errorf("ToBe.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBe_AddSteps(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.AddSteps()
		})
	}
}

func TestToBe_SetDeadLine(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.SetDeadLine()
		})
	}
}

func TestToBe_AddRelated(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		in0 Thing
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := &ToBe{
				Logger: tt.fields.Logger,
			}
			test.AddRelated(tt.args.in0)
		})
	}
}
