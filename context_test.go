package tracr

import (
	"context"
	"github.com/gofrs/uuid"
	"math"
	"reflect"
	"testing"
)

func TestContextWithCID_with_string(t *testing.T) {
	type args struct {
		ctx           context.Context
		correlationID string
	}
	tests := []struct {
		name                   string
		args                   args
		want                   context.Context
		wantValue              string
		wantErr                bool
		correlationIDGenerator func() string // override the cid generator for testing blank values
	}{
		{
			name: "provided cid is returned",
			args: args{
				ctx:           context.TODO(),
				correlationID: "1",
			},
			want:                   context.WithValue(context.TODO(), contextKeyCorrelationID, 1),
			wantValue:              "1",
			wantErr:                false,
			correlationIDGenerator: uuid.Must(uuid.NewV4()).String,
		},
		{
			name: "blank cid generates a new cid",
			args: args{
				ctx:           context.TODO(),
				correlationID: "",
			},
			want:                   context.WithValue(context.TODO(), contextKeyCorrelationID, "hello-world"),
			wantValue:              "hello-world",
			wantErr:                false,
			correlationIDGenerator: func() string { return "hello-world" },
		},
		{
			name: "blank cid generates a new cid",
			args: args{
				ctx:           context.TODO(),
				correlationID: "",
			},
			want:                   context.WithValue(context.TODO(), contextKeyCorrelationID, nil),
			wantValue:              "hello-world",
			wantErr:                false,
			correlationIDGenerator: func() string { return "hello-world" },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			correlationIDGenerator = tt.correlationIDGenerator
			got := ContextWithCID(tt.args.ctx, tt.args.correlationID)
			cid, err := FromContext[string](contextKeyCorrelationID, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContextWithCID error = %v, wantErr %v", err, tt.wantErr)
			}
			if cid != tt.wantValue {
				t.Errorf("ContextWithCID() = %v, wantContext %v", cid, tt.wantValue)
			}
		})
	}
}

func TestFromContextUsingInt(t *testing.T) {
	type args struct {
		key contextKey
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "little int test",
			args: args{
				key: contextKeyCorrelationID,
				ctx: context.WithValue(context.TODO(), contextKeyCorrelationID, 1),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "big int test",
			args: args{
				key: contextKeyCorrelationID,
				ctx: context.WithValue(context.TODO(), contextKeyCorrelationID, math.MaxInt),
			},
			want:    math.MaxInt,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromContext[int](tt.args.key, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContext() got = %v, wantContext %v", got, tt.want)
			}
		})
	}
}

func TestFromContextUsingString(t *testing.T) {
	type args struct {
		key contextKey
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple string test",
			args: args{
				key: contextKeyCorrelationID,
				ctx: context.WithValue(context.TODO(), contextKeyCorrelationID, "hello world"),
			},
			want:    "hello world",
			wantErr: false,
		},
		{
			name: "uuid string test",
			args: args{
				key: contextKeyCorrelationID,
				ctx: context.WithValue(context.TODO(), contextKeyCorrelationID, "c5342e32-831b-4fc7-87f3-45049c9f50f3"),
			},
			want:    "c5342e32-831b-4fc7-87f3-45049c9f50f3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromContext[string](tt.args.key, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContext() got = %v, wantContext %v", got, tt.want)
			}
		})
	}
}
