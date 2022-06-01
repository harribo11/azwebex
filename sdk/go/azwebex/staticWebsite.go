// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azwebex

import (
	"context"
	"reflect"

	resources "github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	storage "github.com/pulumi/pulumi-azure-native/sdk/go/azure/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"storage"
)

type StaticWebsite struct {
	pulumi.ResourceState

	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewStaticWebsite registers a new resource with the given unique name, arguments, and options.
func NewStaticWebsite(ctx *pulumi.Context,
	name string, args *StaticWebsiteArgs, opts ...pulumi.ResourceOption) (*StaticWebsite, error) {
	if args == nil {
		args = &StaticWebsiteArgs{}
	}

	if isZero(args.AppSkuName) {
		args.AppSkuName = pulumi.StringPtr("B1")
	}
	if isZero(args.AppSkuTier) {
		args.AppSkuTier = pulumi.StringPtr("Basic")
	}
	var resource StaticWebsite
	err := ctx.RegisterRemoteComponentResource("azwebex:index:staticWebsite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type staticWebsiteArgs struct {
	// The name of the compute instance running the server. Also see appSkuTier
	AppSkuName *string `pulumi:"appSkuName"`
	// The tier of the compute instance running the server. Also see appSkuName
	AppSkuTier *string `pulumi:"appSkuTier"`
	// The public access level of the BlobContainer containg the website data.
	ContainerPublicAccess *storage.PublicAccess `pulumi:"containerPublicAccess"`
	// The relative file path to the folder containing web files.
	FilePath *string `pulumi:"filePath"`
	// The name prefix given to child resources of this component. Should not contain dashes.
	NamePrefix *string `pulumi:"namePrefix"`
	// The resource group to use. One will be created if not provided.
	ResourceGroup *resources.ResourceGroup `pulumi:"resourceGroup"`
	// The storage account to use. One will be created if not provided.
	StorageAccount *storage.StorageAccount `pulumi:"storageAccount"`
	// The name of the SKU of the storage account created, if storageAccount is not provided
	StorageSkuName *storage.SkuName `pulumi:"storageSkuName"`
}

// The set of arguments for constructing a StaticWebsite resource.
type StaticWebsiteArgs struct {
	// The name of the compute instance running the server. Also see appSkuTier
	AppSkuName pulumi.StringPtrInput
	// The tier of the compute instance running the server. Also see appSkuName
	AppSkuTier pulumi.StringPtrInput
	// The public access level of the BlobContainer containg the website data.
	ContainerPublicAccess storage.PublicAccessPtrInput
	// The relative file path to the folder containing web files.
	FilePath pulumi.StringPtrInput
	// The name prefix given to child resources of this component. Should not contain dashes.
	NamePrefix pulumi.StringPtrInput
	// The resource group to use. One will be created if not provided.
	ResourceGroup resources.ResourceGroupInput
	// The storage account to use. One will be created if not provided.
	StorageAccount storage.StorageAccountInput
	// The name of the SKU of the storage account created, if storageAccount is not provided
	StorageSkuName storage.SkuNamePtrInput
}

func (StaticWebsiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticWebsiteArgs)(nil)).Elem()
}

type StaticWebsiteInput interface {
	pulumi.Input

	ToStaticWebsiteOutput() StaticWebsiteOutput
	ToStaticWebsiteOutputWithContext(ctx context.Context) StaticWebsiteOutput
}

func (*StaticWebsite) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticWebsite)(nil)).Elem()
}

func (i *StaticWebsite) ToStaticWebsiteOutput() StaticWebsiteOutput {
	return i.ToStaticWebsiteOutputWithContext(context.Background())
}

func (i *StaticWebsite) ToStaticWebsiteOutputWithContext(ctx context.Context) StaticWebsiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebsiteOutput)
}

type StaticWebsiteOutput struct{ *pulumi.OutputState }

func (StaticWebsiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticWebsite)(nil)).Elem()
}

func (o StaticWebsiteOutput) ToStaticWebsiteOutput() StaticWebsiteOutput {
	return o
}

func (o StaticWebsiteOutput) ToStaticWebsiteOutputWithContext(ctx context.Context) StaticWebsiteOutput {
	return o
}

func (o StaticWebsiteOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticWebsite) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticWebsiteInput)(nil)).Elem(), &StaticWebsite{})
	pulumi.RegisterOutputType(StaticWebsiteOutput{})
}
