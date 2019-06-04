package cmdutil

import (
	"fmt"
	"github.com/luxas/ignite/pkg/filter"
	"github.com/luxas/ignite/pkg/metadata"
	"github.com/luxas/ignite/pkg/metadata/imgmd"
	"github.com/luxas/ignite/pkg/metadata/kernmd"
	"github.com/luxas/ignite/pkg/metadata/vmmd"
)

// TODO: Make the filter match only strings?
// Then load the actual metadata later, used together with a custom Args validator

func MatchSingleVM(match string) (*vmmd.VMMetadata, error) {
	var md *vmmd.VMMetadata

	// Match a single VM using the VMFilter
	if matches, err := filter.NewFilterer(vmmd.NewVMFilter(match), metadata.VM.Path(), vmmd.LoadVMMetadataFilterable); err == nil {
		if filterable, err := matches.Single(); err == nil {
			if md, err = vmmd.ToVMMetadata(filterable); err != nil {
				return nil, err
			}
		} else {
			return nil, filterError(err, "VM", match)
		}
	} else {
		return nil, err
	}

	return md, nil
}

func MatchSingleVMs(matches []string) ([]*vmmd.VMMetadata, error) {
	mds := make([]*vmmd.VMMetadata, 0, len(matches))

	for _, match := range matches {
		if md, err := MatchSingleVM(match); err == nil {
			mds = append(mds, md)
		} else {
			return nil, err
		}
	}

	return mds, nil
}

func MatchSingleImage(match string) (*imgmd.ImageMetadata, error) {
	var md *imgmd.ImageMetadata

	// Match a single Image using the ImageFilter
	if matches, err := filter.NewFilterer(imgmd.NewImageFilter(match), metadata.Image.Path(), imgmd.LoadImageMetadataFilterable); err == nil {
		if filterable, err := matches.Single(); err == nil {
			if md, err = imgmd.ToImageMetadata(filterable); err != nil {
				return nil, err
			}
		} else {
			return nil, filterError(err, "image", match)
		}
	} else {
		return nil, err
	}

	return md, nil
}

func MatchSingleImages(matches []string) ([]*imgmd.ImageMetadata, error) {
	mds := make([]*imgmd.ImageMetadata, 0, len(matches))

	for _, match := range matches {
		if md, err := MatchSingleImage(match); err == nil {
			mds = append(mds, md)
		} else {
			return nil, err
		}
	}

	return mds, nil
}

func MatchSingleKernel(match string) (*kernmd.KernelMetadata, error) {
	var md *kernmd.KernelMetadata

	// Match a single Kernel using the KernelFilter
	if matches, err := filter.NewFilterer(kernmd.NewKernelFilter(match), metadata.Kernel.Path(), kernmd.LoadKernelMetadataFilterable); err == nil {
		if filterable, err := matches.Single(); err == nil {
			if md, err = kernmd.ToKernelMetadata(filterable); err != nil {
				return nil, err
			}
		} else {
			return nil, filterError(err, "kernel", match)
		}
	} else {
		return nil, err
	}

	return md, nil
}

func MatchSingleKernels(matches []string) ([]*kernmd.KernelMetadata, error) {
	mds := make([]*kernmd.KernelMetadata, 0, len(matches))

	for _, match := range matches {
		if md, err := MatchSingleKernel(match); err == nil {
			mds = append(mds, md)
		} else {
			return nil, err
		}
	}

	return mds, nil
}

func MatchAllVMs(all bool) ([]*vmmd.VMMetadata, error) {
	var mds []*vmmd.VMMetadata

	// Match all VMs using the VMFilter with state checking
	if matches, err := filter.NewFilterer(vmmd.NewVMFilterAll("", all), metadata.VM.Path(), vmmd.LoadVMMetadataFilterable); err == nil {
		if all, err := matches.All(); err == nil {
			if mds, err = vmmd.ToVMMetadataAll(all); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	return mds, nil
}

func MatchAllImages() ([]*imgmd.ImageMetadata, error) {
	var mds []*imgmd.ImageMetadata

	// Match all Images using the ImageFilter
	if matches, err := filter.NewFilterer(imgmd.NewImageFilter(""), metadata.Image.Path(), imgmd.LoadImageMetadataFilterable); err == nil {
		if all, err := matches.All(); err == nil {
			if mds, err = imgmd.ToImageMetadataAll(all); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	return mds, nil
}

func MatchAllKernels() ([]*kernmd.KernelMetadata, error) {
	var mds []*kernmd.KernelMetadata

	// Match all Kernels using the KernelFilter
	if matches, err := filter.NewFilterer(kernmd.NewKernelFilter(""), metadata.Kernel.Path(), kernmd.LoadKernelMetadataFilterable); err == nil {
		if all, err := matches.All(); err == nil {
			if mds, err = kernmd.ToKernelMetadataAll(all); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	return mds, nil
}

func filterError(err error, object, match string) error {
	switch err.(type) {
	case *filter.ErrNonexistent, *filter.ErrAmbiguous:
		return fmt.Errorf(err.Error(), object, match)
	}

	return err
}

func MatchAllVMNames() ([]*metadata.Name, error) {
	mds, err := MatchAllVMs(true)
	if err != nil {
		return nil, err
	}

	names := make([]*metadata.Name, 0, len(mds))

	for _, md := range mds {
		names = append(names, &md.Name)
	}

	return names, nil
}

func MatchAllImageNames() ([]*metadata.Name, error) {
	mds, err := MatchAllImages()
	if err != nil {
		return nil, err
	}

	names := make([]*metadata.Name, 0, len(mds))

	for _, md := range mds {
		names = append(names, &md.Name)
	}

	return names, nil
}

func MatchAllKernelNames() ([]*metadata.Name, error) {
	mds, err := MatchAllKernels()
	if err != nil {
		return nil, err
	}

	names := make([]*metadata.Name, 0, len(mds))

	for _, md := range mds {
		names = append(names, &md.Name)
	}

	return names, nil
}
