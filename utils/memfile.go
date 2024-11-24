package utils

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func Memfile(name string, b []byte) (int, []byte, error) {
	fd, err := unix.MemfdCreate(name, 0)
	if err != nil {
		return 0, nil, fmt.Errorf("memfdcreate: %v", err)
	}

	err = unix.Ftruncate(fd, int64(len(b)))
	if err != nil {
		return 0, nil, fmt.Errorf("ftruncate: %v", err)
	}

	data, err := unix.Mmap(fd, 0, len(b), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		unix.Close(fd)
		return 0, nil, fmt.Errorf("mmap: %v", err)
	}

	copy(data, b)

	err = unix.Munmap(data)
	if err != nil {
		return 0, data, fmt.Errorf("munmap: %v", err)
	}

	return fd, data, nil
}
