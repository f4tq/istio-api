# Istio APIs and Common Configuration Definitions

This repository defines component-level APIs and common configuration formats for the Istio
platform. These definitions are specified using the [protobuf](https://github.com/google/protobuf)
syntax.

This repository depends only on the [tools](https://github.com/istio/tools) repository for tools used during build. This repository *will not* depend on any other repositories. Except for tools, all other Istio repositories can take a dependency on the api repository. 

## API Guidelines

When making changes to the protos in this repository, your changes **must** comply with the [API guidelines](./GUIDELINES.md).

## Updating

After the [protobuf](https://github.com/google/protobuf) definitions
are updated, the corresponding `*pb.go` and `_pb2.py` files must be
generated by running `scripts/generate-protos.sh` and submitted as
part of the same PR as the updated definitions. Also `make
proto-commit` must be run to update the proto.lock file with new
changes.


If releasing a new tagged version, please update python/istio-api/setup.py version to reflect.

## Backwards Incompatible Changes

If a PR tries to make backwards incompatible changes, it will be
blocked by protolock. To force these changes in, install
[protolock](https://github.com/nilslice/protolock) and run
`protolock commit --force`.

You must include a note in your PR that you had to force the
protolock and why.

## Maintaining Generated Kubernetes Types, Clientsets, Informers, Listers

Annotate Istio types that are associated with Kubernetes Custom Resource Definitions
by adding the following tags to the proto message's comments:

```
// +kubetype-gen
// +genclient
// +k8s:deepcopy-gen=true
```

The `kubetype-gen` is used by the kubetype-gen generator to identify types that should
have a corresponding type generated for use with Kubernetes.  `k8s:deepcopy-gen=true`
is used so deepcopy-gen will generate `DeepCopy()` functions for the type.  Note, that
you will have to create custom `DeepCopyInto()` functions for the type.  Please
see one of the existing generated types for an example.  Place any other Kubernetes
specific generator tags here as well, as they will be copied into the generated
Kubernetes types, e.g. `genclient`.

If adding a new package, make sure you add a `doc.go` file with the following tags:

```
//+kubetype-gen
//+kubetype-gen:package=rbac
//+kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
```

The `kubetype-gen` tag is used as trigger to notify the generator that it should
scan the package for types tagged with `kubetype-gen`.  `kubetype-gen:package` is
used to specify the base target package for the generated Kubernetes types.
`kubetype-gen:groupVersion` is used to specify the Kubernetes group/version to
which the types belong.  The generated Kubernetes types will be located in the
package consisting of the `--output-package`, specified through the command
line options, plus the `kubetype-gen:package`, plus the version, e.g. using the
above along with `--output-package istio.io/api/kube/apis` will result in the
target package `istio.io/api/kube/apis/rbac/v1alpha1`.

To update the generated Kubernetes types, simply run
[./scripts/generate-kube-apis.sh].

Note, you'll need to have `dep`, `gofmt` and `goimports` installed on your
system. (`dep` is used for populating the vendor dirctory, while `gofmt` and
`goimports` are used by `go-to-protobuf`.)
