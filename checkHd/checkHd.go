package checkHd

import (
	"github.com/fatih/color"
	syscall "golang.org/x/sys/unix"
	"math"
)

type DiskStatus struct {
	All uint64 `json:"all"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	return
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func CheckHd() {
	disk := DiskUsage("/")
	diskSize := float64(disk.All) / float64(GB)
	hdCompare(diskSize)
}

func hdCompare(size float64) {
	intSize := math.Round(size)
	if intSize > 100 {
		color.Green("SUCCESS: HD is above 100 GB")
	} else {
		color.Red("ERROR: HD is below 100 GB")
	}
}
