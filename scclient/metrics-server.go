package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func main() {
	// 加载 kubeconfig 配置
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	// 创建 Metrics 客户端
	metricsClient, err := metricsv.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 获取 Pod 指标
	podMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, podMetrics := range podMetricsList.Items {
		fmt.Printf("Pod: %s\n", podMetrics.Name)
		for _, container := range podMetrics.Containers {
			fmt.Printf("  Container: %s\n", container.Name)
			fmt.Printf("    CPU Usage: %s\n", container.Usage.Cpu().String())
			fmt.Printf("    Memory Usage: %s\n", container.Usage.Memory().String())

		}
	}
}
