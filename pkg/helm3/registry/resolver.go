/*
<<<<<<< HEAD
Copyright The Helm Authors.
=======
<<<<<<< HEAD:vendor/k8s.io/api/certificates/v1/doc.go
Copyright 2020 The Kubernetes Authors.
=======
Copyright The Helm Authors.
>>>>>>> origin/1.9.1-Integrate-managed-tools:pkg/registry/resolver.go
>>>>>>> origin/1.9.1-Integrate-managed-tools

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

package registry // import "helm.sh/helm/v3/internal/experimental/registry"

import (
	"github.com/containerd/containerd/remotes"
)

type (
	// Resolver provides remotes based on a locator
	Resolver struct {
		remotes.Resolver
	}
)
