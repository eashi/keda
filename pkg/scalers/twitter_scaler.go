package scalers

import (
	"context"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"k8s.io/api/autoscaling/v2beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

const (
	twitterMetricName = "twitterMentionMetric"
)

type twitterScaler struct {
	metadata *twitterMetadata
}

type twitterMetadata struct {
	targetTwitterStatus int
	accountToMonitor    string
	accessKey           string
	accessSecret        string
}

func NewTwitterScaler(resolvedEnv, metadata map[string]string) (Scaler, error) {
	meta, err := ParseTwitterScalerMetadata(metadata, resolvedEnv)
	if err != nil {
		return nil, fmt.Errorf("error parsing Twitter scaler metadata: %s", err)
	}

	return &twitterScaler{
		metadata: meta,
	}, nil
}

func ParseTwitterScalerMetadata(metadata, resolvedEnv map[string]string) (*twitterMetadata, error) {
	meta := twitterMetadata{}
	meta.targetTwitterStatus = 10

	if val, ok := metadata[twitterMetricName]; ok {
		tempParsedValue, err := strconv.Atoi(val)
		if err != nil {
			log.Errorf("Error parsing twitter scaler metadata %s: %s", twitterMetricName, err.Error())
			return nil, fmt.Errorf("Error parsing twitter scaler metadata %s: %s", twitterMetricName, err.Error())
		}

		meta.targetTwitterStatus = tempParsedValue
	}

	if val, ok := metadata["accountToMonitor"]; ok && val != "" {
		meta.accountToMonitor = val
	} else {
		return nil, fmt.Errorf("no accountToMonitor given")
	}

	return &meta, nil
}

func (s *twitterScaler) IsActive(ctx context.Context) (bool, error) {
	//TODO: give proper implementation to IsActive
	return true, nil
}

func (s *twitterScaler) GetMetricSpecForScaling() []v2beta1.MetricSpec {
	targetTwitterStatus := resource.NewQuantity(int64(s.metadata.targetTwitterStatus), resource.DecimalSI)
	externalMetric := &v2beta1.ExternalMetricSource{MetricName: twitterMetricName, TargetAverageValue: targetTwitterStatus}
	metricSpec := v2beta1.MetricSpec{External: externalMetric, Type: externalMetricType}
	return []v2beta1.MetricSpec{metricSpec}
}

func (s *twitterScaler) GetMetrics(ctx context.Context, metricName string, metricSelector labels.Selector) ([]external_metrics.ExternalMetricValue, error) {

	//TODO: Get value from Twitter

	//TODO: Transform value you got from Twitter and give it the right weight
	xxx := 5
	metric := external_metrics.ExternalMetricValue{
		MetricName: metricName,
		Value:      *resource.NewQuantity(int64(xxx), resource.DecimalSI),
		Timestamp:  metav1.Now(),
	}

	return append([]external_metrics.ExternalMetricValue{}, metric), nil
}

func (s *twitterScaler) Close() error {
	return nil
}
