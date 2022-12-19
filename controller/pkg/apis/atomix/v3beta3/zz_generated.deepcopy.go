//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v3beta3

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Binding) DeepCopyInto(out *Binding) {
	*out = *in
	out.Store = in.Store
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(uint32)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Primitives != nil {
		in, out := &in.Primitives, &out.Primitives
		*out = make([]PrimitiveSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Binding.
func (in *Binding) DeepCopy() *Binding {
	if in == nil {
		return nil
	}
	out := new(Binding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStore) DeepCopyInto(out *DataStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStore.
func (in *DataStore) DeepCopy() *DataStore {
	if in == nil {
		return nil
	}
	out := new(DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStoreList) DeepCopyInto(out *DataStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStoreList.
func (in *DataStoreList) DeepCopy() *DataStoreList {
	if in == nil {
		return nil
	}
	out := new(DataStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStoreSpec) DeepCopyInto(out *DataStoreSpec) {
	*out = *in
	out.Driver = in.Driver
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStoreSpec.
func (in *DataStoreSpec) DeepCopy() *DataStoreSpec {
	if in == nil {
		return nil
	}
	out := new(DataStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Driver) DeepCopyInto(out *Driver) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Driver.
func (in *Driver) DeepCopy() *Driver {
	if in == nil {
		return nil
	}
	out := new(Driver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSinkConfig) DeepCopyInto(out *FileSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSinkConfig.
func (in *FileSinkConfig) DeepCopy() *FileSinkConfig {
	if in == nil {
		return nil
	}
	out := new(FileSinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggerConfig) DeepCopyInto(out *LoggerConfig) {
	*out = *in
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = make(map[string]OutputConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggerConfig.
func (in *LoggerConfig) DeepCopy() *LoggerConfig {
	if in == nil {
		return nil
	}
	out := new(LoggerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfig) DeepCopyInto(out *LoggingConfig) {
	*out = *in
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make(map[string]LoggerConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Sinks != nil {
		in, out := &in.Sinks, &out.Sinks
		*out = make(map[string]SinkConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfig.
func (in *LoggingConfig) DeepCopy() *LoggingConfig {
	if in == nil {
		return nil
	}
	out := new(LoggingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputConfig) DeepCopyInto(out *OutputConfig) {
	*out = *in
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(string)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputConfig.
func (in *OutputConfig) DeepCopy() *OutputConfig {
	if in == nil {
		return nil
	}
	out := new(OutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
	in.Proxy.DeepCopyInto(&out.Proxy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyStatus) DeepCopyInto(out *ProxyStatus) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouteStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyStatus.
func (in *ProxyStatus) DeepCopy() *ProxyStatus {
	if in == nil {
		return nil
	}
	out := new(ProxyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	out.Store = in.Store
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimitiveSpec) DeepCopyInto(out *PrimitiveSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimitiveSpec.
func (in *PrimitiveSpec) DeepCopy() *PrimitiveSpec {
	if in == nil {
		return nil
	}
	out := new(PrimitiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkConfig) DeepCopyInto(out *SinkConfig) {
	*out = *in
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(string)
		**out = **in
	}
	if in.Stdout != nil {
		in, out := &in.Stdout, &out.Stdout
		*out = new(StdoutSinkConfig)
		**out = **in
	}
	if in.Stderr != nil {
		in, out := &in.Stderr, &out.Stderr
		*out = new(StderrSinkConfig)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileSinkConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkConfig.
func (in *SinkConfig) DeepCopy() *SinkConfig {
	if in == nil {
		return nil
	}
	out := new(SinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StderrSinkConfig) DeepCopyInto(out *StderrSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StderrSinkConfig.
func (in *StderrSinkConfig) DeepCopy() *StderrSinkConfig {
	if in == nil {
		return nil
	}
	out := new(StderrSinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StdoutSinkConfig) DeepCopyInto(out *StdoutSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StdoutSinkConfig.
func (in *StdoutSinkConfig) DeepCopy() *StdoutSinkConfig {
	if in == nil {
		return nil
	}
	out := new(StdoutSinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfile) DeepCopyInto(out *StorageProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfile.
func (in *StorageProfile) DeepCopy() *StorageProfile {
	if in == nil {
		return nil
	}
	out := new(StorageProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileList) DeepCopyInto(out *StorageProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileList.
func (in *StorageProfileList) DeepCopy() *StorageProfileList {
	if in == nil {
		return nil
	}
	out := new(StorageProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileSpec) DeepCopyInto(out *StorageProfileSpec) {
	*out = *in
	in.Proxy.DeepCopyInto(&out.Proxy)
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]Binding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileSpec.
func (in *StorageProfileSpec) DeepCopy() *StorageProfileSpec {
	if in == nil {
		return nil
	}
	out := new(StorageProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileStatus) DeepCopyInto(out *StorageProfileStatus) {
	*out = *in
	if in.PodStatuses != nil {
		in, out := &in.PodStatuses, &out.PodStatuses
		*out = make([]PodStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileStatus.
func (in *StorageProfileStatus) DeepCopy() *StorageProfileStatus {
	if in == nil {
		return nil
	}
	out := new(StorageProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProxyConfig) DeepCopyInto(out *StorageProxyConfig) {
	*out = *in
	in.Server.DeepCopyInto(&out.Server)
	in.Logging.DeepCopyInto(&out.Logging)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProxyConfig.
func (in *StorageProxyConfig) DeepCopy() *StorageProxyConfig {
	if in == nil {
		return nil
	}
	out := new(StorageProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProxyServerConfig) DeepCopyInto(out *StorageProxyServerConfig) {
	*out = *in
	if in.ReadBufferSize != nil {
		in, out := &in.ReadBufferSize, &out.ReadBufferSize
		*out = new(int)
		**out = **in
	}
	if in.WriteBufferSize != nil {
		in, out := &in.WriteBufferSize, &out.WriteBufferSize
		*out = new(int)
		**out = **in
	}
	if in.MaxRecvMsgSize != nil {
		in, out := &in.MaxRecvMsgSize, &out.MaxRecvMsgSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.MaxSendMsgSize != nil {
		in, out := &in.MaxSendMsgSize, &out.MaxSendMsgSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.NumStreamWorkers != nil {
		in, out := &in.NumStreamWorkers, &out.NumStreamWorkers
		*out = new(uint32)
		**out = **in
	}
	if in.MaxConcurrentStreams != nil {
		in, out := &in.MaxConcurrentStreams, &out.MaxConcurrentStreams
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProxyServerConfig.
func (in *StorageProxyServerConfig) DeepCopy() *StorageProxyServerConfig {
	if in == nil {
		return nil
	}
	out := new(StorageProxyServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProxySpec) DeepCopyInto(out *StorageProxySpec) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProxySpec.
func (in *StorageProxySpec) DeepCopy() *StorageProxySpec {
	if in == nil {
		return nil
	}
	out := new(StorageProxySpec)
	in.DeepCopyInto(out)
	return out
}
