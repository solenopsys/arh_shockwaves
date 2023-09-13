package wrappers

import (
	"context"
	helmapiv1 "github.com/k3s-io/helm-controller/pkg/apis/helm.cattle.io/v1"
	helmcln "github.com/k3s-io/helm-controller/pkg/generated/clientset/versioned"
	"xs/pkg/io"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
)

type HelmApi struct {
	HelmClientSet *helmcln.Clientset
	Namespace     string
}

func NewAPI(rc *rest.Config) HelmApi {
	helmcln, err := helmcln.NewForConfig(rc)
	if err != nil {
		io.Panic("Error create helmcln: %v\n", err)
	}
	return HelmApi{HelmClientSet: helmcln, Namespace: "installers"}
}

func (f *HelmApi) CreateHelmChartSimple(name string, repo string, version string, targetNamespace string) (*helmapiv1.HelmChart, error) {
	chart := &helmapiv1.HelmChart{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name, // название в кубере
			Namespace: f.Namespace,
			Labels: map[string]string{
				"helm-test": "true",
			},
		},
		Spec: helmapiv1.HelmChartSpec{
			Chart:           name, // исходные данные
			Version:         version,
			Repo:            repo,
			TargetNamespace: targetNamespace,
			Set: map[string]intstr.IntOrString{
				"rbac.enabled": {
					Type:   intstr.String,
					StrVal: "true",
				},
				"ssl.enabled": {
					Type:   intstr.String,
					StrVal: "true",
				},
			},
			HelmVersion: "v3",
		},
	}
	return f.CreateHelmChart(chart, f.Namespace)
}

func (f *HelmApi) DeleteHelmChart(name string) error {
	grace := int64(10)
	background := metav1.DeletePropagationBackground
	options := metav1.DeleteOptions{GracePeriodSeconds: &grace, PropagationPolicy: &background}
	return f.HelmClientSet.HelmV1().HelmCharts(f.Namespace).Delete(context.TODO(), name, options)
}

func (f *HelmApi) GetHelmChart(name string) (*helmapiv1.HelmChart, error) {
	return f.HelmClientSet.HelmV1().HelmCharts(f.Namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (f *HelmApi) ListHelmCharts(labelSelector string) (*helmapiv1.HelmChartList, error) {
	return f.HelmClientSet.HelmV1().HelmCharts(f.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
}

func (f *HelmApi) CreateHelmChart(chart *helmapiv1.HelmChart, namespace string) (*helmapiv1.HelmChart, error) {
	return f.HelmClientSet.HelmV1().HelmCharts(namespace).Create(context.TODO(), chart, metav1.CreateOptions{})
}

func (f *HelmApi) UpdateHelmChart(chart *helmapiv1.HelmChart, namespace string) (updated *helmapiv1.HelmChart, err error) {
	hcs := f.HelmClientSet.HelmV1()
	if err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		updated, err = hcs.HelmCharts(namespace).Get(context.TODO(), chart.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		updated.Spec = chart.Spec
		_, err = hcs.HelmCharts(namespace).Update(context.TODO(), updated, metav1.UpdateOptions{})
		return err
	}); err != nil {
		updated = nil
	}
	return
}
