package main

import (
	"fmt"

	"os"

	"unsafe"

	_ "image/jpeg"

	"nyctal-dri/drm"
	"nyctal-dri/drm/mode"
	"nyctal/model"

	"golang.org/x/sys/unix"
)

type (
	framebuffer struct {
		id     uint32
		handle uint32
		data   []byte
		fb     *mode.FB
		size   uint64
		stride uint32
	}

	msetData struct {
		mode      *mode.Modeset
		fbs       [2]framebuffer
		frontbuf  uint
		savedCrtc *mode.Crtc
	}
)

func createFramebuffer(file *os.File, dev *mode.Modeset) (framebuffer, error) {
	fb, err := mode.CreateFB(file, dev.Width, dev.Height, 32)
	if err != nil {
		return framebuffer{}, fmt.Errorf("failed to create framebuffer: %s", err.Error())
	}
	stride := fb.Pitch
	size := fb.Size
	handle := fb.Handle

	fbID, err := mode.AddFB(file, dev.Width, dev.Height, 24, 32, stride, handle)
	if err != nil {
		return framebuffer{}, fmt.Errorf("cannot create dumb buffer: %s", err.Error())
	}

	offset, err := mode.MapDumb(file, handle)
	if err != nil {
		return framebuffer{}, err
	}

	mmap, err := unix.Mmap(int(file.Fd()), int64(offset), int(size), int(unix.PROT_READ)|int(unix.PROT_WRITE), int(unix.MAP_SHARED))

	if err != nil {
		return framebuffer{}, fmt.Errorf("failed to mmap framebuffer: %s", err.Error())
	}
	for i := uint64(0); i < size; i++ {
		mmap[i] = 0
	}
	framebuf := framebuffer{
		id:     fbID,
		handle: handle,
		data:   mmap,
		fb:     fb,
		size:   size,
		stride: stride,
	}
	return framebuf, nil
}

func destroyFramebuffer(modeset *mode.SimpleModeset, mset msetData, file *os.File) {
	fbs := mset.fbs

	for _, fb := range fbs {
		handle := fb.handle
		data := fb.data

		err := unix.Munmap(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to munmap memory: %s\n", err.Error())
			continue
		}
		err = mode.RmFB(file, fb.id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to remove frame buffer: %s\n", err.Error())
			continue
		}

		err = mode.DestroyDumb(file, handle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to destroy dumb buffer: %s\n", err.Error())
			continue
		}

		err = modeset.SetCrtc(mset.mode, mset.savedCrtc)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			continue
		}
	}
}

func cleanup(modeset *mode.SimpleModeset, msets []msetData, file *os.File) {
	for _, mset := range msets {
		destroyFramebuffer(modeset, mset, file)
	}
}

type DrmState struct {
	modeset *mode.SimpleModeset
	file    *os.File
	msets   []msetData
}

func (ds *DrmState) Stats() (uint32, uint32) {
	return ds.msets[0].fbs[0].fb.Width, ds.msets[0].fbs[0].fb.Height
}

func (ds *DrmState) RenderBuffer(img *model.BGRA) error {
	var off uint32
	bounds := img.Bounds()
	for j := 0; j < len(ds.msets); j++ {
		mset := ds.msets[j]
		buf := &mset.fbs[mset.frontbuf^1]
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, _ := img.AtRaw(x, y).RGBA()
				off = (buf.stride * uint32(y)) + (uint32(x) * 4)
				val := uint32((uint32(r) << 16) | (uint32(g) << 8) | uint32(b))
				*(*uint32)(unsafe.Pointer(&buf.data[off])) = val
			}
		}
		err := mode.SetCrtc(ds.file, mset.mode.Crtc, buf.id, 0, 0, &mset.mode.Conn, 1, &mset.mode.Mode)
		if err != nil {
			return err
		}

		mset.frontbuf ^= 1
	}
	return nil
}

func (ds *DrmState) Clenup() {
	cleanup(ds.modeset, ds.msets, ds.file)
}

func DrmInit() *DrmState {
	file, err := drm.OpenCard(0)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return nil
	}
	if !drm.HasDumbBuffer(file) {
		fmt.Printf("drm device does not support dumb buffers")
		return nil
	}

	modeset, err := mode.NewSimpleModeset(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}

	var msets []msetData
	for _, mod := range modeset.Modesets {
		framebuf1, err := createFramebuffer(file, &mod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			cleanup(modeset, msets, file)
			return nil
		}

		framebuf2, err := createFramebuffer(file, &mod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			cleanup(modeset, msets, file)
			return nil
		}

		// save current CRTC of this mode to restore at exit
		savedCrtc, err := mode.GetCrtc(file, mod.Crtc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: Cannot get CRTC for connector %d: %s", mod.Conn, err.Error())
			cleanup(modeset, msets, file)
			return nil
		}
		// change the mode using framebuf1 initially
		err = mode.SetCrtc(file, mod.Crtc, framebuf1.id, 0, 0, &mod.Conn, 1, &mod.Mode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot set CRTC for connector %d: %s", mod.Conn, err.Error())
			cleanup(modeset, msets, file)
			return nil
		}
		msets = append(msets, msetData{
			frontbuf: 0,
			mode:     &mod,
			fbs: [2]framebuffer{
				framebuf1, framebuf2,
			},
			savedCrtc: savedCrtc,
		})
	}
	return &DrmState{
		modeset: modeset,
		msets:   msets,
		file:    file,
	}
}
