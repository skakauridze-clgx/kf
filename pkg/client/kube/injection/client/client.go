// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by injection-gen. DO NOT EDIT.

package client

import (
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"

	versioned "github.com/google/kf/v2/pkg/client/kube/clientset/versioned"
	typednetworkingv1 "github.com/google/kf/v2/pkg/client/kube/clientset/versioned/typed/networking/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1 "k8s.io/client-go/applyconfigurations/networking/v1"
	discovery "k8s.io/client-go/discovery"
	dynamic "k8s.io/client-go/dynamic"
	rest "k8s.io/client-go/rest"
	injection "knative.dev/pkg/injection"
	dynamicclient "knative.dev/pkg/injection/clients/dynamicclient"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterClient(withClientFromConfig)
	injection.Default.RegisterClientFetcher(func(ctx context.Context) interface{} {
		return Get(ctx)
	})
	injection.Dynamic.RegisterDynamicClient(withClientFromDynamic)
}

// Key is used as the key for associating information with a context.Context.
type Key struct{}

func withClientFromConfig(ctx context.Context, cfg *rest.Config) context.Context {
	return context.WithValue(ctx, Key{}, versioned.NewForConfigOrDie(cfg))
}

func withClientFromDynamic(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key{}, &wrapClient{dyn: dynamicclient.Get(ctx)})
}

// Get extracts the versioned.Interface client from the context.
func Get(ctx context.Context) versioned.Interface {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		if injection.GetConfig(ctx) == nil {
			logging.FromContext(ctx).Panic(
				"Unable to fetch github.com/google/kf/v2/pkg/client/kube/clientset/versioned.Interface from context. This context is not the application context (which is typically given to constructors via sharedmain).")
		} else {
			logging.FromContext(ctx).Panic(
				"Unable to fetch github.com/google/kf/v2/pkg/client/kube/clientset/versioned.Interface from context.")
		}
	}
	return untyped.(versioned.Interface)
}

type wrapClient struct {
	dyn dynamic.Interface
}

var _ versioned.Interface = (*wrapClient)(nil)

func (w *wrapClient) Discovery() discovery.DiscoveryInterface {
	panic("Discovery called on dynamic client!")
}

func convert(from interface{}, to runtime.Object) error {
	bs, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("Marshal() = %w", err)
	}
	if err := json.Unmarshal(bs, to); err != nil {
		return fmt.Errorf("Unmarshal() = %w", err)
	}
	return nil
}

// NetworkingV1 retrieves the NetworkingV1Client
func (w *wrapClient) NetworkingV1() typednetworkingv1.NetworkingV1Interface {
	return &wrapNetworkingV1{
		dyn: w.dyn,
	}
}

type wrapNetworkingV1 struct {
	dyn dynamic.Interface
}

func (w *wrapNetworkingV1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapNetworkingV1) Ingresses(namespace string) typednetworkingv1.IngressInterface {
	return &wrapNetworkingV1IngressImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "networking.k8s.io",
			Version:  "v1",
			Resource: "ingresses",
		}),

		namespace: namespace,
	}
}

type wrapNetworkingV1IngressImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typednetworkingv1.IngressInterface = (*wrapNetworkingV1IngressImpl)(nil)

func (w *wrapNetworkingV1IngressImpl) Apply(ctx context.Context, in *v1.IngressApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.Ingress, err error) {
	panic("NYI")
}

func (w *wrapNetworkingV1IngressImpl) ApplyStatus(ctx context.Context, in *v1.IngressApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.Ingress, err error) {
	panic("NYI")
}

func (w *wrapNetworkingV1IngressImpl) Create(ctx context.Context, in *networkingv1.Ingress, opts metav1.CreateOptions) (*networkingv1.Ingress, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "Ingress",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.Ingress{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressImpl) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapNetworkingV1IngressImpl) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapNetworkingV1IngressImpl) Get(ctx context.Context, name string, opts metav1.GetOptions) (*networkingv1.Ingress, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.Ingress{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressImpl) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.IngressList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.IngressList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *networkingv1.Ingress, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.Ingress{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressImpl) Update(ctx context.Context, in *networkingv1.Ingress, opts metav1.UpdateOptions) (*networkingv1.Ingress, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "Ingress",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.Ingress{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressImpl) UpdateStatus(ctx context.Context, in *networkingv1.Ingress, opts metav1.UpdateOptions) (*networkingv1.Ingress, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "Ingress",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.Ingress{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressImpl) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapNetworkingV1) IngressClasses() typednetworkingv1.IngressClassInterface {
	return &wrapNetworkingV1IngressClassImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "networking.k8s.io",
			Version:  "v1",
			Resource: "ingressclasses",
		}),
	}
}

type wrapNetworkingV1IngressClassImpl struct {
	dyn dynamic.NamespaceableResourceInterface
}

var _ typednetworkingv1.IngressClassInterface = (*wrapNetworkingV1IngressClassImpl)(nil)

func (w *wrapNetworkingV1IngressClassImpl) Apply(ctx context.Context, in *v1.IngressClassApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.IngressClass, err error) {
	panic("NYI")
}

func (w *wrapNetworkingV1IngressClassImpl) ApplyStatus(ctx context.Context, in *v1.IngressClassApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.IngressClass, err error) {
	panic("NYI")
}

func (w *wrapNetworkingV1IngressClassImpl) Create(ctx context.Context, in *networkingv1.IngressClass, opts metav1.CreateOptions) (*networkingv1.IngressClass, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "IngressClass",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.IngressClass{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressClassImpl) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return w.dyn.Delete(ctx, name, opts)
}

func (w *wrapNetworkingV1IngressClassImpl) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return w.dyn.DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapNetworkingV1IngressClassImpl) Get(ctx context.Context, name string, opts metav1.GetOptions) (*networkingv1.IngressClass, error) {
	uo, err := w.dyn.Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.IngressClass{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressClassImpl) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.IngressClassList, error) {
	uo, err := w.dyn.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.IngressClassList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressClassImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *networkingv1.IngressClass, err error) {
	uo, err := w.dyn.Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.IngressClass{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressClassImpl) Update(ctx context.Context, in *networkingv1.IngressClass, opts metav1.UpdateOptions) (*networkingv1.IngressClass, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "IngressClass",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.IngressClass{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressClassImpl) UpdateStatus(ctx context.Context, in *networkingv1.IngressClass, opts metav1.UpdateOptions) (*networkingv1.IngressClass, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "IngressClass",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.IngressClass{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1IngressClassImpl) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapNetworkingV1) NetworkPolicies(namespace string) typednetworkingv1.NetworkPolicyInterface {
	return &wrapNetworkingV1NetworkPolicyImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "networking.k8s.io",
			Version:  "v1",
			Resource: "networkpolicies",
		}),

		namespace: namespace,
	}
}

type wrapNetworkingV1NetworkPolicyImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typednetworkingv1.NetworkPolicyInterface = (*wrapNetworkingV1NetworkPolicyImpl)(nil)

func (w *wrapNetworkingV1NetworkPolicyImpl) Apply(ctx context.Context, in *v1.NetworkPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.NetworkPolicy, err error) {
	panic("NYI")
}

func (w *wrapNetworkingV1NetworkPolicyImpl) ApplyStatus(ctx context.Context, in *v1.NetworkPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *networkingv1.NetworkPolicy, err error) {
	panic("NYI")
}

func (w *wrapNetworkingV1NetworkPolicyImpl) Create(ctx context.Context, in *networkingv1.NetworkPolicy, opts metav1.CreateOptions) (*networkingv1.NetworkPolicy, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "NetworkPolicy",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.NetworkPolicy{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1NetworkPolicyImpl) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapNetworkingV1NetworkPolicyImpl) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapNetworkingV1NetworkPolicyImpl) Get(ctx context.Context, name string, opts metav1.GetOptions) (*networkingv1.NetworkPolicy, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.NetworkPolicy{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1NetworkPolicyImpl) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.NetworkPolicyList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.NetworkPolicyList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1NetworkPolicyImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *networkingv1.NetworkPolicy, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.NetworkPolicy{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1NetworkPolicyImpl) Update(ctx context.Context, in *networkingv1.NetworkPolicy, opts metav1.UpdateOptions) (*networkingv1.NetworkPolicy, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "NetworkPolicy",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.NetworkPolicy{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1NetworkPolicyImpl) UpdateStatus(ctx context.Context, in *networkingv1.NetworkPolicy, opts metav1.UpdateOptions) (*networkingv1.NetworkPolicy, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "NetworkPolicy",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &networkingv1.NetworkPolicy{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapNetworkingV1NetworkPolicyImpl) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}
