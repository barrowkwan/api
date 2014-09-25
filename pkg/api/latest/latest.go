/*
Copyright 2014 Google Inc. All rights reserved.

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

package latest

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta1"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta2"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
)

// Version is the string that represents the current external default version.
const Version = "v1beta1"

// OldestVersion is the string that represents the oldest server version supported,
// for client code that wants to hardcode the lowest common denominator.
const OldestVersion = "v1beta1"

// Versions is the list of versions that are recognized in code. The order provided
// may be assumed to be least feature rich to most feature rich, and clients may
// choose to prefer the latter items in the list over the former items when presented
// with a set of versions to choose.
var Versions = []string{"v1beta1", "v1beta2"}

// Codec is the default codec for serializing output that should use
// the latest supported version.  Use this Codec when writing to
// disk, a data store that is not dynamically versioned, or in tests.
// This codec can decode any object that Kubernetes is aware of.
var Codec = v1beta1.Codec

// ResourceVersioner describes a default versioner that can handle all types
// of versioning.
// TODO: when versioning changes, make this part of each API definition.
var ResourceVersioner = runtime.NewJSONBaseResourceVersioner()

// SelfLinker can set or get the SelfLink field of all API types.
// TODO: when versioning changes, make this part of each API definition.
// TODO(lavalamp): Combine SelfLinker & ResourceVersioner interfaces, force all uses
// to go through the InterfacesFor method below.
var SelfLinker = runtime.NewJSONBaseSelfLinker()

// VersionInterfaces contains the interfaces one should use for dealing with types of a particular version.
type VersionInterfaces struct {
	runtime.Codec
	runtime.ResourceVersioner
	runtime.SelfLinker
}

// InterfacesFor returns the default Codec and ResourceVersioner for a given version
// string, or an error if the version is not known.
func InterfacesFor(version string) (*VersionInterfaces, error) {
	switch version {
	case "v1beta1":
		return &VersionInterfaces{
			Codec:             v1beta1.Codec,
			ResourceVersioner: ResourceVersioner,
			SelfLinker:        SelfLinker,
		}, nil
	case "v1beta2":
		return &VersionInterfaces{
			Codec:             v1beta2.Codec,
			ResourceVersioner: ResourceVersioner,
			SelfLinker:        SelfLinker,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported storage version: %s (valid: %s)", version, strings.Join(Versions, ", "))
	}
}
