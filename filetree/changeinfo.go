package filetree

import (
	"bytes"
	"fmt"
)

type FileChangeInfo struct {
	Path     string
	Typeflag byte
	MD5sum   [16]byte
	DiffType DiffType
}

type DiffType int

// enum to show whether a file has changed
const (
	Unchanged DiffType = iota
	Changed
	Added
	Removed
)

func (d DiffType) String() string {
	switch d {
	case Unchanged:
		return "Unchanged"
	case Changed:
		return "Changed"
	case Added:
		return "Added"
	case Removed:
		return "Removed"
	default:
		return fmt.Sprintf("%d", int(d))
	}
}

func (a DiffType) merge(b DiffType) DiffType {
	if a == b {
		return a
	}
	return Changed
}

func (a *FileChangeInfo) getDiffType(b *FileChangeInfo) DiffType {
	if a == nil && b == nil {
		return Unchanged
	}
	if a == nil || b == nil {
		return Changed
	}
	if a.Typeflag == b.Typeflag {
		if bytes.Compare(a.MD5sum[:], b.MD5sum[:]) == 0 {
			return Unchanged
		}
	}
	return Changed
}