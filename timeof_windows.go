package timeof

import (
	"os"
	"syscall"
	"time"
)

// Creation returns the time the file or directory was created.
func Creation(f os.FileInfo) time.Time {
	attr := f.Sys().(*syscall.Win32FileAttributeData)
	return time.Unix(0, attr.CreationTime.Nanoseconds())
}

// LastAccess returns the time the file or directory was last written to,
// truncated or overwritten. This time is not updated when file attributes or
// security descriptors are changed.
func LastAccess(f os.FileInfo) time.Time {
	attr := f.Sys().(*syscall.Win32FileAttributeData)
	return time.Unix(0, attr.LastAccessTime.Nanoseconds())
}
