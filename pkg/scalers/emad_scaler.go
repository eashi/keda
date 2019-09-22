package scalers

import (
	"strconv"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/external_metrics"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/api/autoscaling/v2beta1"
	"fmt"
	"context"
)

type emadScaler struct {
	metadata *emadMetaData
}

type emadMetaData struct {
	target int
}

func NewEmadScaler(resolvedEnv, metadata map[string]string) (Scaler, error){
	meta := &emadMetaData{}
	val := metadata["target"]
	meta.target, _ = strconv.Atoi(val)
	return &emadScaler{
		metadata: meta,
	},nil
}

func (s *emadScaler) GetMetricSpecForScaling() []v2beta1.MetricSpec {
	targetQty := resource.NewQuantity(int64(s.metadata.target), resource.DecimalSI)
	externalMetric := &v2beta1.ExternalMetricSource{MetricName: "emadQuantity", TargetAverageValue: targetQty}
	metricSpec := v2beta1.MetricSpec{External: externalMetric, Type: "External"}
	return []v2beta1.MetricSpec{metricSpec}
}

func (s *emadScaler) GetMetrics(ctx context.Context, metricname string, metricSelector labels.Selector) ([]external_metrics.ExternalMetricValue, error) {
	fmt.Println("GetMetrics called, target value is: %v, and current value is:%v", s.metadata.target, 100)	

	metric := external_metrics.ExternalMetricValue{
		MetricName: metricname,
		Value: *resource.NewQuantity(int64(100), resource.DecimalSI),
		Timestamp: v1.Now(),
	}

	return append([]external_metrics.ExternalMetricValue{}, metric), nil
}

func (s *emadScaler) IsActive(ctx context.Context) (bool, error) {
	 return true, nil
}

func (s *emadScaler) Close() error {
	return nil
}