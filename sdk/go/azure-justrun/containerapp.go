// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurejustrun

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	containerregistry "github.com/pulumi/pulumi-azure-native/sdk/go/azure/containerregistry"
	resources "github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Containerapp struct {
	pulumi.ResourceState

	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewContainerapp registers a new resource with the given unique name, arguments, and options.
func NewContainerapp(ctx *pulumi.Context,
	name string, args *ContainerappArgs, opts ...pulumi.ResourceOption) (*Containerapp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DockerImageName == nil {
		return nil, errors.New("invalid value for required argument 'DockerImageName'")
	}
	var resource Containerapp
	err := ctx.RegisterRemoteComponentResource("azure-justrun:index:containerapp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type containerappArgs struct {
	// The name of the docker image to use. Required.
	DockerImageName string `pulumi:"dockerImageName"`
	// The name prefix given to child resources of this component. Should not contain dashes.
	NamePrefix *string `pulumi:"namePrefix"`
	// The container registry to use. One will be created if not provided.
	Registry *containerregistry.Registry `pulumi:"registry"`
	// The resource group to use. One will be created if not provided.
	ResourceGroup *resources.ResourceGroup `pulumi:"resourceGroup"`
}

// The set of arguments for constructing a Containerapp resource.
type ContainerappArgs struct {
	// The name of the docker image to use. Required.
	DockerImageName pulumi.StringInput
	// The name prefix given to child resources of this component. Should not contain dashes.
	NamePrefix pulumi.StringPtrInput
	// The container registry to use. One will be created if not provided.
	Registry containerregistry.RegistryInput
	// The resource group to use. One will be created if not provided.
	ResourceGroup resources.ResourceGroupInput
}

func (ContainerappArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerappArgs)(nil)).Elem()
}

type ContainerappInput interface {
	pulumi.Input

	ToContainerappOutput() ContainerappOutput
	ToContainerappOutputWithContext(ctx context.Context) ContainerappOutput
}

func (*Containerapp) ElementType() reflect.Type {
	return reflect.TypeOf((**Containerapp)(nil)).Elem()
}

func (i *Containerapp) ToContainerappOutput() ContainerappOutput {
	return i.ToContainerappOutputWithContext(context.Background())
}

func (i *Containerapp) ToContainerappOutputWithContext(ctx context.Context) ContainerappOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerappOutput)
}

// ContainerappArrayInput is an input type that accepts ContainerappArray and ContainerappArrayOutput values.
// You can construct a concrete instance of `ContainerappArrayInput` via:
//
//          ContainerappArray{ ContainerappArgs{...} }
type ContainerappArrayInput interface {
	pulumi.Input

	ToContainerappArrayOutput() ContainerappArrayOutput
	ToContainerappArrayOutputWithContext(context.Context) ContainerappArrayOutput
}

type ContainerappArray []ContainerappInput

func (ContainerappArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Containerapp)(nil)).Elem()
}

func (i ContainerappArray) ToContainerappArrayOutput() ContainerappArrayOutput {
	return i.ToContainerappArrayOutputWithContext(context.Background())
}

func (i ContainerappArray) ToContainerappArrayOutputWithContext(ctx context.Context) ContainerappArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerappArrayOutput)
}

// ContainerappMapInput is an input type that accepts ContainerappMap and ContainerappMapOutput values.
// You can construct a concrete instance of `ContainerappMapInput` via:
//
//          ContainerappMap{ "key": ContainerappArgs{...} }
type ContainerappMapInput interface {
	pulumi.Input

	ToContainerappMapOutput() ContainerappMapOutput
	ToContainerappMapOutputWithContext(context.Context) ContainerappMapOutput
}

type ContainerappMap map[string]ContainerappInput

func (ContainerappMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Containerapp)(nil)).Elem()
}

func (i ContainerappMap) ToContainerappMapOutput() ContainerappMapOutput {
	return i.ToContainerappMapOutputWithContext(context.Background())
}

func (i ContainerappMap) ToContainerappMapOutputWithContext(ctx context.Context) ContainerappMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerappMapOutput)
}

type ContainerappOutput struct{ *pulumi.OutputState }

func (ContainerappOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Containerapp)(nil)).Elem()
}

func (o ContainerappOutput) ToContainerappOutput() ContainerappOutput {
	return o
}

func (o ContainerappOutput) ToContainerappOutputWithContext(ctx context.Context) ContainerappOutput {
	return o
}

func (o ContainerappOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Containerapp) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

type ContainerappArrayOutput struct{ *pulumi.OutputState }

func (ContainerappArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Containerapp)(nil)).Elem()
}

func (o ContainerappArrayOutput) ToContainerappArrayOutput() ContainerappArrayOutput {
	return o
}

func (o ContainerappArrayOutput) ToContainerappArrayOutputWithContext(ctx context.Context) ContainerappArrayOutput {
	return o
}

func (o ContainerappArrayOutput) Index(i pulumi.IntInput) ContainerappOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Containerapp {
		return vs[0].([]*Containerapp)[vs[1].(int)]
	}).(ContainerappOutput)
}

type ContainerappMapOutput struct{ *pulumi.OutputState }

func (ContainerappMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Containerapp)(nil)).Elem()
}

func (o ContainerappMapOutput) ToContainerappMapOutput() ContainerappMapOutput {
	return o
}

func (o ContainerappMapOutput) ToContainerappMapOutputWithContext(ctx context.Context) ContainerappMapOutput {
	return o
}

func (o ContainerappMapOutput) MapIndex(k pulumi.StringInput) ContainerappOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Containerapp {
		return vs[0].(map[string]*Containerapp)[vs[1].(string)]
	}).(ContainerappOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerappInput)(nil)).Elem(), &Containerapp{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerappArrayInput)(nil)).Elem(), ContainerappArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerappMapInput)(nil)).Elem(), ContainerappMap{})
	pulumi.RegisterOutputType(ContainerappOutput{})
	pulumi.RegisterOutputType(ContainerappArrayOutput{})
	pulumi.RegisterOutputType(ContainerappMapOutput{})
}
