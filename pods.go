package main

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodsInterface interface {
	List() ([]corev1.Pod, error)
}

type Pods struct {
	client    *kubernetes.Clientset
	namespace string
}

func (p *Pods) List() ([]corev1.Pod, error) {
	pods, err := p.client.CoreV1().Pods(p.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods.Items, nil
}
