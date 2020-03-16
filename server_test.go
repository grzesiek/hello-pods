package main

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestHandlerError(t *testing.T) {
	handler := Handler{}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	handler.ServeError(errors.New("test-error"))(rec, req)

	require.Equal(t, `Error occured: "test-error"`, rec.Body.String())
}

func TestHandlerPods(t *testing.T) {
	handler := Handler{index: "index.html.tpl"}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	handler.ServePods(pods())(rec, req)

	require.Contains(t, rec.Body.String(), `<html lang="en">`)
	require.Contains(t, rec.Body.String(), "my-pod-1")
	require.Contains(t, rec.Body.String(), "my-pod-1")
}

func pods() (pods []v1.Pod) {
	for i := 0; i < 2; i++ {
		pod := v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("my-pod-%d", i),
				Namespace: "default",
			},
			Status: v1.PodStatus{Message: "Running"},
		}

		pods = append(pods, pod)
	}

	return pods
}
