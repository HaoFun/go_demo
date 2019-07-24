package check

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func HealthCheck(ctx *gin.Context) {
	message := "OK"
	ctx.String(http.StatusOK, "\n"+message)
}

func DiskCheck(ctx *gin.Context) {
	u, _ := disk.Usage("/")
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPrecent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPrecent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPrecent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",
		text,
		usedMB,
		usedGB,
		totalMB,
		totalGB,
		usedPrecent)

	ctx.String(status, "\n"+message)
}

func CPUCheck(ctx *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	load1 := a.Load1
	load5 := a.Load5
	load15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if load5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if load5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %s",
		text,
		load1,
		load5,
		load15,
		cores)

	ctx.String(status, "\n"+message)
}

func RAMCheck(ctx *gin.Context) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",
		text,
		usedMB,
		usedGB,
		totalMB,
		totalGB,
		usedPercent)

	ctx.String(status, message)
}
