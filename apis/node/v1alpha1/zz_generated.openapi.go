/*
Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>

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

package v1alpha1

// Code generated by openapi-gen script. DO NOT EDIT.

const (

	// NodeJoinRequestOpenAPISchema represents the OpenAPI schema for kind NodeJoinRequest
	NodeJoinRequestOpenAPISchema = `description: NodeJoinRequest is the Schema for the nodejoinrequests API
properties:
  apiVersion:
    description: 'APIVersion defines the versioned schema of this representation of
      an object. Servers should convert recognized schemas to the latest internal
      value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
    type: string
  kind:
    description: 'Kind is a string value representing the REST resource this object
      represents. Servers may infer this from the endpoint the client submits requests
      to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
    type: string
  metadata:
    type: object
  spec:
    description: NodeJoinRequestSpec defines the desired state of NodeJoinRequest
    properties:
      apiServerEndpoint:
        type: string
      containerRuntimeEndpoint:
        type: string
      imageServiceEndpoint:
        type: string
      symmetricKey:
        type: string
    type: object
  status:
    description: NodeJoinRequestStatus defines the observed state of NodeJoinRequest
    properties:
      conditions:
        items:
          description: Condition represents the node join request conditions
          type: string
        type: array
      kubeConfig:
        type: string
      kubeletConfig:
        type: string
      kubernetesVersion:
        type: string
      vpnAddress:
        type: string
      vpnPeer:
        type: string
    type: object
type: object`
)
