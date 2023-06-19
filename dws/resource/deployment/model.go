package deployment

import (
	"errors"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type vmResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Image    types.String `tfsdk:"image"`
	Region   types.String `tfsdk:"region"`
	CPU      types.Int64  `tfsdk:"cpu"`
	RAM      types.Int64  `tfsdk:"ram"`
	Disk     types.Int64  `tfsdk:"disk_size"`
	DiskType types.String `tfsdk:"disk_type"`
	IPv4     types.Bool   `tfsdk:"assign_public_ipv4"`
	IPv6     types.Bool   `tfsdk:"assign_public_ipv6"`
	Ygg      types.Bool   `tfsdk:"assign_ygg_ip"`
	SSHKey   types.String `tfsdk:"ssh_key"`
	HostName types.String `tfsdk:"host_name"`
	VPCID    types.String `tfsdk:"vpc_id"`

	// Computed
	PublicIPv4 types.String `tfsdk:"public_ipv4"`
	PublicIPv6 types.String `tfsdk:"public_ipv6"`
	YggIP      types.String `tfsdk:"ygg_ip"`
}

func (v *vmResourceModel) ToClientRequest() (*client.VMConfig, error) {
	r := &client.VMConfig{
		Ipv4:        v.IPv4.ValueBool(),
		Ipv6:        v.IPv6.ValueBool(),
		NetworkUUID: v.VPCID.ValueString(),
	}

	if v.Image.IsUnknown() || v.Image.IsNull() {
		return nil, errors.New("image property is required and cannot be empty")
	}

	r.ImageVersion = v.Image.ValueString()

	if v.Region.IsUnknown() || v.CPU.IsNull() {
		return nil, errors.New("region property is required and cannot be empty")
	}

	r.Region = v.Region.ValueString()

	if v.CPU.IsUnknown() || v.CPU.IsNull() {
		return nil, errors.New("cpu property is required and cannot be empty")
	}

	r.CPU = int(v.CPU.ValueInt64())

	if v.RAM.IsUnknown() || v.RAM.IsNull() {
		return nil, errors.New("ram property is required and cannot be empty")
	}

	r.RAM = int(v.RAM.ValueInt64())

	if v.Disk.IsUnknown() || v.Disk.IsNull() {
		return nil, errors.New("disk property is required and cannot be empty")
	}

	r.Hdd = int(v.Disk.ValueInt64())

	if v.DiskType.IsUnknown() || v.DiskType.IsNull() {
		return nil, errors.New("disk_type property is required and cannot be empty")
	}

	r.HddType = v.DiskType.ValueString()

	if v.SSHKey.IsUnknown() || v.SSHKey.IsNull() {
		return nil, errors.New("ssh_key property is required and cannot be empty")
	}

	r.SSHKey = v.SSHKey.ValueString()

	if v.HostName.IsUnknown() || v.HostName.IsNull() {
		return nil, errors.New("host_name property is required and cannot be empty")
	}

	r.HostName = v.HostName.ValueString()

	return r, nil
}

func (v *vmResourceModel) FromClientResponse(c *client.VMResponse) {
	v.PublicIPv4 = types.StringValue(c.Data.IP)
	v.PublicIPv6 = types.StringValue(c.Data.IPv6)
	v.YggIP = types.StringValue(c.Data.Ygg)

	return
}
