/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by script/embed_resources.sh. DO NOT EDIT.

package deploy

var Resources map[string]string

func init() {
	Resources = make(map[string]string)

	Resources["operator.yaml"] =
		`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: yaks
  labels:
    yaks.dev/component: operator
spec:
  replicas: 1
  selector:
    matchLabels:
      name: yaks
  template:
    metadata:
      labels:
        name: yaks
    spec:
      serviceAccountName: yaks
      containers:
        - name: yaks
          image: yaks/yaks:0.0.1
          command:
          - yaks
          - operator
          imagePullPolicy: IfNotPresent
          env:
            - name: WATCH_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: OPERATOR_NAME
              value: "yaks"

`
	Resources["role_binding.yaml"] =
		`
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: yaks
subjects:
- kind: ServiceAccount
  name: yaks
roleRef:
  kind: Role
  name: yaks
  apiGroup: rbac.authorization.k8s.io

`
	Resources["role.yaml"] =
		`
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  creationTimestamp: null
  name: yaks
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - services
  - endpoints
  - persistentvolumeclaims
  - configmaps
  - secrets
  - serviceaccounts
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - pods/log
  - pods/status
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - roles
  - rolebindings
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - apps
  resources:
  - deployments
  - replicasets
  - statefulsets
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - get
  - create
- apiGroups:
  - yaks.dev
  resources:
  - '*'
  verbs:
  - '*'

`
	Resources["service_account.yaml"] =
		`
apiVersion: v1
kind: ServiceAccount
metadata:
  name: yaks

`
	Resources["user_cluster_role.yaml"] =
		`
# ---------------------------------------------------------------------------
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ---------------------------------------------------------------------------

kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: yaks:edit
  labels:
    # Add these permissions to the "admin" and "edit" default roles.
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
rules:
- apiGroups: ["yaks.dev"]
  resources: ["*"]
  verbs: ["*"]
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - get
  - create
`
	Resources["viewer_role_binding.yaml"] =
		`
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: yaks-viewer
  labels:
    app: "yaks"
subjects:
- kind: ServiceAccount
  name: yaks-viewer
roleRef:
  kind: Role
  name: yaks-viewer
  apiGroup: rbac.authorization.k8s.io

`
	Resources["viewer_role.yaml"] =
		`
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: yaks-viewer
  labels:
    app: "yaks"
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  - endpoints
  - persistentvolumeclaims
  - pods
  - serviceaccounts
  - services
  - secrets
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - pods/log
  - pods/status
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - apps
  resources:
  - deployments
  - replicasets
  - statefulsets
  verbs:
  - get
  - list
  - watch

`
	Resources["viewer_service_account.yaml"] =
		`
apiVersion: v1
kind: ServiceAccount
metadata:
  name: yaks-viewer
  labels:
    app: "yaks"

`
	Resources["crds/yaks_v1alpha1_test_crd.yaml"] =
		`
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: tests.yaks.dev
spec:
  group: yaks.dev
  names:
    kind: Test
    listKind: TestList
    plural: tests
    singular: test
  scope: Namespaced
  subresources:
    status: {}
  additionalPrinterColumns:
    - name: Phase
      type: string
      description: The test phase
      JSONPath: .status.phase
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          type: object
        status:
          type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true

`
	Resources["crds/yaks_v1alpha1_test_cr.yaml"] =
		`
apiVersion: yaks.dev/v1alpha1
kind: Test
metadata:
  name: example-test
spec:
  # Add fields here
  size: 3

`

}
