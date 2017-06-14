package layer

import (
	"io"

	"github.com/docker/distribution"
	"github.com/docker/docker/daemon/graphdriver"
)

func (ls *layerStore) RegisterWithDescriptor(ts io.Reader, parent ChainID, platform Platform, descriptor distribution.Descriptor, opts *RegisterOpts) (Layer, error) {
	return ls.registerWithDescriptor(ts, parent, platform, descriptor, opts)
}

func getApplyDiffOpts(opts *RegisterOpts) *graphdriver.ApplyDiffOpts {
	ado := &graphdriver.ApplyDiffOpts{}
	if opts != nil {
		ado.Uvm = opts.Uvm
	}
	return ado
}
