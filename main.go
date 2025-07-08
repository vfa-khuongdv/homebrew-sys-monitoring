package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcInfo struct {
	Name string
	CPU  float64
	Mem  float32
}

func getTopProcessesBy(limit int, sortBy string) ([]ProcInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var infos []ProcInfo
	for _, p := range procs {
		name, err := p.Name()
		if err != nil || name == "" {
			continue
		}
		cpuPercent, err := p.CPUPercent()
		if err != nil {
			continue
		}
		memPercent, err := p.MemoryPercent()
		if err != nil {
			continue
		}
		infos = append(infos, ProcInfo{
			Name: name,
			CPU:  cpuPercent,
			Mem:  memPercent,
		})
	}

	switch sortBy {
	case "cpu":
		sort.Slice(infos, func(i, j int) bool {
			return infos[i].CPU > infos[j].CPU
		})
	case "mem":
		sort.Slice(infos, func(i, j int) bool {
			return infos[i].Mem > infos[j].Mem
		})
	}

	if len(infos) > limit {
		return infos[:limit], nil
	}
	return infos, nil
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Widgets
	header := widgets.NewParagraph()
	header.Text = "📊 SYSTEM MONITOR (Press 'q' to quit)"
	header.SetRect(0, 0, 100, 3)
	header.TextStyle.Fg = ui.ColorWhite
	header.TextStyle.Bg = ui.ColorBlue
	header.Border = false

	cpuChart := widgets.NewPlot()
	cpuChart.Title = "🧠 CPU Usage (%)"
	cpuChart.Data = [][]float64{make([]float64, 60)}
	cpuChart.SetRect(0, 3, 100, 15)
	cpuChart.AxesColor = ui.ColorWhite
	cpuChart.LineColors[0] = ui.ColorRed
	cpuChart.Marker = widgets.MarkerBraille
	cpuChart.DrawDirection = widgets.DrawLeft

	ramGauge := widgets.NewGauge()
	ramGauge.Title = "📦 RAM Usage"
	ramGauge.SetRect(0, 15, 49, 20)
	ramGauge.BarColor = ui.ColorGreen
	ramGauge.TitleStyle.Fg = ui.ColorWhite
	ramGauge.LabelStyle.Fg = ui.ColorYellow
	ramGauge.BorderStyle.Fg = ui.ColorCyan

	diskGauge := widgets.NewGauge()
	diskGauge.Title = "💽 Disk Usage"
	diskGauge.SetRect(51, 15, 100, 20)
	diskGauge.BarColor = ui.ColorMagenta
	diskGauge.TitleStyle.Fg = ui.ColorWhite
	diskGauge.LabelStyle.Fg = ui.ColorYellow
	diskGauge.BorderStyle.Fg = ui.ColorCyan

	processCPU := widgets.NewTable()
	processCPU.Title = "🔥 Top 5 Processes (by CPU)"
	processCPU.SetRect(0, 20, 50, 30)
	processCPU.TextStyle = ui.NewStyle(ui.ColorWhite)
	processCPU.RowSeparator = false
	processCPU.FillRow = true
	processCPU.RowStyles[0] = ui.NewStyle(ui.ColorCyan, ui.ColorBlack, ui.ModifierBold)

	processMem := widgets.NewTable()
	processMem.Title = "🧠 Top 5 Processes (by Memory)"
	processMem.SetRect(50, 20, 100, 30)
	processMem.TextStyle = ui.NewStyle(ui.ColorWhite)
	processMem.RowSeparator = false
	processMem.FillRow = true
	processMem.RowStyles[0] = ui.NewStyle(ui.ColorCyan, ui.ColorBlack, ui.ModifierBold)

	footer := widgets.NewParagraph()
	footer.Text = "Made with ❤️ using Go + TermUI"
	footer.SetRect(0, 30, 100, 32)
	footer.Border = false
	footer.TextStyle.Fg = ui.ColorMagenta

	ticker := time.NewTicker(time.Second).C
	uiEvents := ui.PollEvents()

	for {
		select {
		case <-ticker:
			// Update CPU chart
			cpuPercent, _ := cpu.Percent(0, false)
			cpuChart.Data[0] = append(cpuChart.Data[0][1:], cpuPercent[0])

			// Update RAM
			vm, _ := mem.VirtualMemory()
			ramGauge.Percent = int(vm.UsedPercent)
			ramGauge.Label = fmt.Sprintf("%d%% (%.1f GB / %.1f GB)",
				int(vm.UsedPercent),
				float64(vm.Used)/math.Pow(1024, 3),
				float64(vm.Total)/math.Pow(1024, 3),
			)

			// Update Disk
			d, _ := disk.Usage("/")
			diskGauge.Percent = int(d.UsedPercent)
			diskGauge.Label = fmt.Sprintf("%d%% (%.1f GB / %.1f GB)",
				int(d.UsedPercent),
				float64(d.Used)/math.Pow(1024, 3),
				float64(d.Total)/math.Pow(1024, 3),
			)

			// Top CPU
			topCPU, err1 := getTopProcessesBy(5, "cpu")
			if err1 == nil {
				processCPU.Rows = [][]string{{"Name", "CPU %", "RAM %"}}
				for _, p := range topCPU {
					processCPU.Rows = append(processCPU.Rows, []string{
						truncate(p.Name, 24),
						fmt.Sprintf("%.1f", p.CPU),
						fmt.Sprintf("%.1f", p.Mem),
					})
				}
			}

			// Top RAM
			topRAM, err2 := getTopProcessesBy(5, "mem")
			if err2 == nil {
				processMem.Rows = [][]string{{"Name", "RAM %", "CPU %"}}
				for _, p := range topRAM {
					processMem.Rows = append(processMem.Rows, []string{
						truncate(p.Name, 24),
						fmt.Sprintf("%.1f", p.Mem),
						fmt.Sprintf("%.1f", p.CPU),
					})
				}
			}

			ui.Render(header, cpuChart, ramGauge, diskGauge, processCPU, processMem, footer)

		case e := <-uiEvents:
			if e.Type == ui.KeyboardEvent && e.ID == "q" {
				return
			}
		}
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return strings.TrimSpace(s[:max-1]) + "…"
}
