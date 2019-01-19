// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha3

import (
	internalinterfaces "istio.io/api/pkg/kube/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DestinationRules returns a DestinationRuleInformer.
	DestinationRules() DestinationRuleInformer
	// EnvoyFilters returns a EnvoyFilterInformer.
	EnvoyFilters() EnvoyFilterInformer
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// ServiceEntries returns a ServiceEntryInformer.
	ServiceEntries() ServiceEntryInformer
	// Sidecars returns a SidecarInformer.
	Sidecars() SidecarInformer
	// VirtualServices returns a VirtualServiceInformer.
	VirtualServices() VirtualServiceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// DestinationRules returns a DestinationRuleInformer.
func (v *version) DestinationRules() DestinationRuleInformer {
	return &destinationRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EnvoyFilters returns a EnvoyFilterInformer.
func (v *version) EnvoyFilters() EnvoyFilterInformer {
	return &envoyFilterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceEntries returns a ServiceEntryInformer.
func (v *version) ServiceEntries() ServiceEntryInformer {
	return &serviceEntryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Sidecars returns a SidecarInformer.
func (v *version) Sidecars() SidecarInformer {
	return &sidecarInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualServices returns a VirtualServiceInformer.
func (v *version) VirtualServices() VirtualServiceInformer {
	return &virtualServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
