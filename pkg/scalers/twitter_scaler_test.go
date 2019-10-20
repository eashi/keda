package scalers

import (
	"context"
	"reflect"
	"testing"

	"k8s.io/api/autoscaling/v2beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

func TestParseTwitterScalerMetadata(t *testing.T) {
	type args struct {
		metadata    map[string]string
		resolvedEnv map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *twitterMetadata
		wantErr bool
	}{
		{
			"No Access Key",
			args{map[string]string{"targetTwitterStatus": "10", "accountToMonitor": "jeffhallon", "accessKey": "", "accessSecret": ""}, nil},
			nil,
			true,
		},
		{
			"No accountToMonitor supplied",
			args{map[string]string{"targetTwitterStatus": "10", "accountToMonitor": "", "accessKey": "", "accessSecret": ""}, nil},
			nil,
			true,
		},
		{
			"No accessSecret supplied",
			args{map[string]string{"targetTwitterStatus": "10", "accountToMonitor": "", "accessKey": "", "accessSecret": ""}, nil},
			nil,
			true,
		},
		{
			"All configs are good",
			args{map[string]string{"targetTwitterStatus": "10", "accountToMonitor": "jeffhollan", "accessKey": "mykey", "accessSecret": "mysecret"}, nil},
			&twitterMetadata{targetTwitterStatus: 10, accountToMonitor: "jeffhollan", accessKey: "mykey", accessSecret: "mysecret"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTwitterScalerMetadata(tt.args.metadata, tt.args.resolvedEnv)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTwitterScalerMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTwitterScalerMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twitterScaler_IsActive(t *testing.T) {
	type fields struct {
		metadata *twitterMetadata
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &twitterScaler{
				metadata: tt.fields.metadata,
			}
			got, err := s.IsActive(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("twitterScaler.IsActive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("twitterScaler.IsActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twitterScaler_GetMetricSpecForScaling(t *testing.T) {
	type fields struct {
		metadata *twitterMetadata
	}
	tests := []struct {
		name   string
		fields fields
		want   []v2beta1.MetricSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &twitterScaler{
				metadata: tt.fields.metadata,
			}
			if got := s.GetMetricSpecForScaling(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twitterScaler.GetMetricSpecForScaling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twitterScaler_GetMetrics(t *testing.T) {
	type fields struct {
		metadata *twitterMetadata
	}
	type args struct {
		ctx            context.Context
		metricName     string
		metricSelector labels.Selector
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []external_metrics.ExternalMetricValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &twitterScaler{
				metadata: tt.fields.metadata,
			}
			got, err := s.GetMetrics(tt.args.ctx, tt.args.metricName, tt.args.metricSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("twitterScaler.GetMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twitterScaler.GetMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twitterScaler_Close(t *testing.T) {
	type fields struct {
		metadata *twitterMetadata
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &twitterScaler{
				metadata: tt.fields.metadata,
			}
			if err := s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("twitterScaler.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
