/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package encoding

import (
	admissionv1 "k8s.io/api/admission/v1"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"

	apiserverinternal "k8s.io/api/apiserverinternal/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	authenticationv1 "k8s.io/api/authentication/v1"
	authenticationv1beta1 "k8s.io/api/authentication/v1beta1"
	authorizationv1 "k8s.io/api/authorization/v1"
	authorizationv1beta1 "k8s.io/api/authorization/v1beta1"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"

	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"

	certificatesv1 "k8s.io/api/certificates/v1"
	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"

	coordinationv1 "k8s.io/api/coordination/v1"
	coordinationv1beta1 "k8s.io/api/coordination/v1beta1"

	discoveryv1 "k8s.io/api/discovery/v1"
	discoveryv1beta1 "k8s.io/api/discovery/v1beta1"

	corev1 "k8s.io/api/core/v1"

	eventsv1 "k8s.io/api/events/v1"
	eventsv1beta1 "k8s.io/api/events/v1beta1"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

	flowcontrolv1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	flowcontrolv1beta1 "k8s.io/api/flowcontrol/v1beta1"
	flowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"

	imagepolicyv1alpha1 "k8s.io/api/imagepolicy/v1alpha1"

	networkingv1 "k8s.io/api/networking/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"

	nodev1 "k8s.io/api/node/v1"
	nodev1alpha1 "k8s.io/api/node/v1alpha1"
	nodev1beta1 "k8s.io/api/node/v1beta1"

	policyv1 "k8s.io/api/policy/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"

	rbacv1 "k8s.io/api/rbac/v1"
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"

	schedulingv1 "k8s.io/api/scheduling/v1"
	schedulingv1alpha1 "k8s.io/api/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/api/scheduling/v1beta1"

	storagev1 "k8s.io/api/storage/v1"
	storagev1alpha1 "k8s.io/api/storage/v1alpha1"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	AddToScheme(Scheme)
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kuberentes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
//nolint:errcheck
func AddToScheme(scheme *runtime.Scheme) {
	admissionv1beta1.AddToScheme(scheme)
	admissionv1.AddToScheme(scheme)

	admissionregistrationv1.AddToScheme(scheme)
	admissionregistrationv1beta1.AddToScheme(scheme)

	apiserverinternal.AddToScheme(scheme)

	appsv1.AddToScheme(scheme)
	appsv1beta1.AddToScheme(scheme)
	appsv1beta2.AddToScheme(scheme)
	authenticationv1.AddToScheme(scheme)
	authenticationv1beta1.AddToScheme(scheme)
	authorizationv1.AddToScheme(scheme)
	authorizationv1beta1.AddToScheme(scheme)

	autoscalingv1.AddToScheme(scheme)
	autoscalingv2beta1.AddToScheme(scheme)
	autoscalingv2beta2.AddToScheme(scheme)
	autoscalingv2.AddToScheme(scheme)

	batchv1.AddToScheme(scheme)
	batchv1beta1.AddToScheme(scheme)

	certificatesv1beta1.AddToScheme(scheme)
	certificatesv1.AddToScheme(scheme)

	coordinationv1beta1.AddToScheme(scheme)
	coordinationv1.AddToScheme(scheme)

	discoveryv1beta1.AddToScheme(scheme)
	discoveryv1.AddToScheme(scheme)

	corev1.AddToScheme(scheme)

	eventsv1beta1.AddToScheme(scheme)
	eventsv1.AddToScheme(scheme)

	extensionsv1beta1.AddToScheme(scheme)

	flowcontrolv1alpha1.AddToScheme(scheme)
	flowcontrolv1beta1.AddToScheme(scheme)
	flowcontrolv1beta2.AddToScheme(scheme)

	imagepolicyv1alpha1.AddToScheme(scheme)

	networkingv1.AddToScheme(scheme)
	networkingv1beta1.AddToScheme(scheme)

	nodev1alpha1.AddToScheme(scheme)
	nodev1beta1.AddToScheme(scheme)
	nodev1.AddToScheme(scheme)

	policyv1beta1.AddToScheme(scheme)
	policyv1.AddToScheme(scheme)

	rbacv1.AddToScheme(scheme)
	rbacv1alpha1.AddToScheme(scheme)
	rbacv1beta1.AddToScheme(scheme)

	schedulingv1alpha1.AddToScheme(scheme)
	schedulingv1beta1.AddToScheme(scheme)
	schedulingv1.AddToScheme(scheme)

	storagev1.AddToScheme(scheme)
	storagev1alpha1.AddToScheme(scheme)
	storagev1beta1.AddToScheme(scheme)
}
