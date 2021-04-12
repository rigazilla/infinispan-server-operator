// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Autoscale) DeepCopyInto(out *Autoscale) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Autoscale.
func (in *Autoscale) DeepCopy() *Autoscale {
	if in == nil {
		return nil
	}
	out := new(Autoscale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossSiteExposeSpec) DeepCopyInto(out *CrossSiteExposeSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossSiteExposeSpec.
func (in *CrossSiteExposeSpec) DeepCopy() *CrossSiteExposeSpec {
	if in == nil {
		return nil
	}
	out := new(CrossSiteExposeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointEncryption) DeepCopyInto(out *EndpointEncryption) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointEncryption.
func (in *EndpointEncryption) DeepCopy() *EndpointEncryption {
	if in == nil {
		return nil
	}
	out := new(EndpointEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposeSpec) DeepCopyInto(out *ExposeSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposeSpec.
func (in *ExposeSpec) DeepCopy() *ExposeSpec {
	if in == nil {
		return nil
	}
	out := new(ExposeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infinispan) DeepCopyInto(out *Infinispan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infinispan.
func (in *Infinispan) DeepCopy() *Infinispan {
	if in == nil {
		return nil
	}
	out := new(Infinispan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Infinispan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanCloudEvents) DeepCopyInto(out *InfinispanCloudEvents) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanCloudEvents.
func (in *InfinispanCloudEvents) DeepCopy() *InfinispanCloudEvents {
	if in == nil {
		return nil
	}
	out := new(InfinispanCloudEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanCondition) DeepCopyInto(out *InfinispanCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanCondition.
func (in *InfinispanCondition) DeepCopy() *InfinispanCondition {
	if in == nil {
		return nil
	}
	out := new(InfinispanCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanContainerSpec) DeepCopyInto(out *InfinispanContainerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanContainerSpec.
func (in *InfinispanContainerSpec) DeepCopy() *InfinispanContainerSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanExternalDependencies) DeepCopyInto(out *InfinispanExternalDependencies) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanExternalDependencies.
func (in *InfinispanExternalDependencies) DeepCopy() *InfinispanExternalDependencies {
	if in == nil {
		return nil
	}
	out := new(InfinispanExternalDependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanList) DeepCopyInto(out *InfinispanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Infinispan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanList.
func (in *InfinispanList) DeepCopy() *InfinispanList {
	if in == nil {
		return nil
	}
	out := new(InfinispanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfinispanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanLoggingSpec) DeepCopyInto(out *InfinispanLoggingSpec) {
	*out = *in
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make(map[string]LoggingLevelType, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanLoggingSpec.
func (in *InfinispanLoggingSpec) DeepCopy() *InfinispanLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSecurity) DeepCopyInto(out *InfinispanSecurity) {
	*out = *in
	if in.EndpointAuthentication != nil {
		in, out := &in.EndpointAuthentication, &out.EndpointAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.EndpointEncryption != nil {
		in, out := &in.EndpointEncryption, &out.EndpointEncryption
		*out = new(EndpointEncryption)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSecurity.
func (in *InfinispanSecurity) DeepCopy() *InfinispanSecurity {
	if in == nil {
		return nil
	}
	out := new(InfinispanSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanServiceContainerSpec) DeepCopyInto(out *InfinispanServiceContainerSpec) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanServiceContainerSpec.
func (in *InfinispanServiceContainerSpec) DeepCopy() *InfinispanServiceContainerSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanServiceContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanServiceSpec) DeepCopyInto(out *InfinispanServiceSpec) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(InfinispanServiceContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Sites != nil {
		in, out := &in.Sites, &out.Sites
		*out = new(InfinispanSitesSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanServiceSpec.
func (in *InfinispanServiceSpec) DeepCopy() *InfinispanServiceSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSiteLocationSpec) DeepCopyInto(out *InfinispanSiteLocationSpec) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSiteLocationSpec.
func (in *InfinispanSiteLocationSpec) DeepCopy() *InfinispanSiteLocationSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSiteLocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSitesLocalSpec) DeepCopyInto(out *InfinispanSitesLocalSpec) {
	*out = *in
	in.Expose.DeepCopyInto(&out.Expose)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSitesLocalSpec.
func (in *InfinispanSitesLocalSpec) DeepCopy() *InfinispanSitesLocalSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSitesLocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSitesSpec) DeepCopyInto(out *InfinispanSitesSpec) {
	*out = *in
	in.Local.DeepCopyInto(&out.Local)
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]InfinispanSiteLocationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSitesSpec.
func (in *InfinispanSitesSpec) DeepCopy() *InfinispanSitesSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSitesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSpec) DeepCopyInto(out *InfinispanSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	in.Security.DeepCopyInto(&out.Security)
	out.Container = in.Container
	in.Service.DeepCopyInto(&out.Service)
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(InfinispanLoggingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(ExposeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(Autoscale)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudEvents != nil {
		in, out := &in.CloudEvents, &out.CloudEvents
		*out = new(InfinispanCloudEvents)
		**out = **in
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(InfinispanUpgrade)
		**out = **in
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = new(InfinispanExternalDependencies)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSpec.
func (in *InfinispanSpec) DeepCopy() *InfinispanSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanStatus) DeepCopyInto(out *InfinispanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]InfinispanCondition, len(*in))
		copy(*out, *in)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(InfinispanSecurity)
		(*in).DeepCopyInto(*out)
	}
	in.PodStatus.DeepCopyInto(&out.PodStatus)
	if in.ConsoleUrl != nil {
		in, out := &in.ConsoleUrl, &out.ConsoleUrl
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanStatus.
func (in *InfinispanStatus) DeepCopy() *InfinispanStatus {
	if in == nil {
		return nil
	}
	out := new(InfinispanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanUpgrade) DeepCopyInto(out *InfinispanUpgrade) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanUpgrade.
func (in *InfinispanUpgrade) DeepCopy() *InfinispanUpgrade {
	if in == nil {
		return nil
	}
	out := new(InfinispanUpgrade)
	in.DeepCopyInto(out)
	return out
}
