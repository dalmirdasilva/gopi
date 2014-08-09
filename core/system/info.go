package system

import (
  "github.com/dalmirdasilva/gorpi/util"
  "strings"
  "strconv"
  "os"
  "errors"
)

type Info struct {
  cpuInfo map[string]string
}

type Board int

const (
  Unknown Board = iota
  ModelARev0
  ModelBRev1
  ModelBRev2
)

func InfoInstance() Info {
  return Info{}
}

func (si Info) CpuInfo() (map[string]string, error) {
  if (si.cpuInfo == nil) {
    si.cpuInfo = make(map[string]string)
    info, err := util.Execute("cat", "/proc/cpuinfo")
    if (err != nil) {
      return nil, err
    }
    for i := range info {
      parts := strings.Split(info[i], ":")
      if (len(parts) >= 2) {
        key := strings.TrimSpace(parts[0])
        val := strings.TrimSpace(parts[1])
        if (key != "" && val != "") {
          si.cpuInfo[key] = val
        }
      }
    }
  }
  return si.cpuInfo, nil
}

func (si Info) CpuInfoEntry(entry string) (string, error) {
  info, err := si.CpuInfo()
  if (err != nil) {
    return "", err
  }
  value, exists := info[entry]
  if (!exists) {
    err := errors.New("Entry not found: " + entry)
    return "", err
  }
  return value, nil
}


func (si Info) Processor() (string, error) {
  return si.CpuInfoEntry("Processor")
}

func (si Info) BogoMIPS() (string, error) {
  return si.CpuInfoEntry("BogoMIPS")
}

func (si Info) CpuFeatures() ([]string, error) {
  features, err := si.CpuInfoEntry("Features")
  if (err != nil) {
    return nil, err
  }
  return strings.Split(features, " "), nil
}

func (si Info) CpuImplementer() (string, error) {
  return si.CpuInfoEntry("CPU implementer")
}

func (si Info) CpuArchitecture() (string, error) {
  return si.CpuInfoEntry("CPU architecture")
}

func (si Info) CpuVariant() (string, error) {
  return si.CpuInfoEntry("CPU variant")
}

func (si Info) CpuPart() (string, error) {
  return si.CpuInfoEntry("CPU part")
}

func (si Info) CpuRevision() (string, error) {
  return si.CpuInfoEntry("CPU revision")
}

func (si Info) Hardware() (string, error) {
  return si.CpuInfoEntry("Hardware")
}

func (si Info) Revision() (string, error) {
  return si.CpuInfoEntry("Revision")
}

func (si Info) Serial() (string, error) {
  return si.CpuInfoEntry("Serial")
}

func (si Info) Memory() (map[string]uint32, error) {
  result := make(map[string]uint32)
  keys := []string{"total", "used", "free", "shared", "buffers", "cached"}
  lines, err := util.Execute("free", "-b")
  if (err != nil) {
    return nil, err
  }
  for i := range lines {
    line := lines[i]
    if (strings.HasPrefix(line, "Mem:")) {
      parts := strings.Split(line, " ")
      for j := range parts {
        if (j > 0) {
          part := strings.TrimSpace(parts[j])
          ui, _ := strconv.ParseUint(part, 10, 32)
          result[keys[j - 1]] = uint32(ui)
        }
      }
    }
  }
  return result, nil
}

func (si Info) memoryOf(of string) uint32 {
  memory, err := si.Memory()
  if (err != nil) {
    return 0
  }
  return memory[of]
}

func (si Info) MemoryTotal() uint32 {
  return si.memoryOf("total")
}

func (si Info) MemoryUsed() uint32 {
  return si.memoryOf("used")
}

func (si Info) MemoryFree() uint32 {
  return si.memoryOf("free")
}

func (si Info) MemoryShared() uint32 {
  return si.memoryOf("shared")
}

func (si Info) MemoryBuffers() uint32 {
  return si.memoryOf("buffers")
}

func (si Info) MemoryCached() uint32 {
  return si.memoryOf("cached")
}

func (si Info) BoardType() Board {
  revision, err := si.Revision()
  if (err == nil) {
    switch revision {
    case "0002", "0003":
      return ModelBRev1
    case "0004", "0005", "0006":
      return ModelBRev2
    case "0007", "0008", "0009":
      return ModelARev0
    case "000d", "000e", "000f":
      return ModelBRev2
    }
  }
  return Unknown
}

func (si Info) Temperature() float32 {
  cmd := "/opt/vc/bin/vcgencmd"
  _, err := os.Stat(cmd)
  if (os.IsNotExist(err)) {
    return 0.0
  }
  lines, err := util.Execute(cmd, "measure_temp")
  if (err != nil) {
    return 0.0
  }
  for i := range lines {
    line := lines[i]
    if (strings.HasPrefix(line, "temp=")) {
      parts := strings.FieldsFunc(line, func(r rune) bool {
        switch r {
        case '=', '\'':
          return true
        }
        return false
      })
      temp, _ := strconv.ParseFloat(parts[1], 32)
      return float32(temp)
    }
  }
  return 0.0
}


func (si Info) ClockFrequency(target string) uint64 {
  cmd := "/opt/vc/bin/vcgencmd"
  _, err := os.Stat(cmd)
  if (os.IsNotExist(err)) {
    return 0
  }
  lines, err := util.Execute(cmd, "measure_clock", strings.TrimSpace(target))
  for i := range lines {
    line := lines[i]
    if (strings.HasPrefix(line, "frequency")) {
      parts := strings.Split(line, "=")
      freq, _ := strconv.ParseUint(parts[1], 10, 64)
      return uint64(freq)
    }
  }
  return 0
}

func (si Info) ClockFrequencyArm() uint64 {
  return si.ClockFrequency("arm")
}

func (si Info) ClockFrequencyCore() uint64 {
  return si.ClockFrequency("core")
}

func (si Info) ClockFrequencyH264() uint64 {
  return si.ClockFrequency("h264")
}

func (si Info) ClockFrequencyISP() uint64 {
  return si.ClockFrequency("isp")
}

func (si Info) ClockFrequencyV3D() uint64 {
  return si.ClockFrequency("v3d")
}

func (si Info) ClockFrequencyUART() uint64 {
  return si.ClockFrequency("uart")
}

func (si Info) ClockFrequencyPWM() uint64 {
  return si.ClockFrequency("pwm")
}

func (si Info) ClockFrequencyEMMC() uint64 {
  return si.ClockFrequency("emmc")
}

func (si Info) ClockFrequencyPixel() uint64 {
  return si.ClockFrequency("pixel")
}

func (si Info) ClockFrequencyVEC() uint64 {
  return si.ClockFrequency("vec")
}

func (si Info) ClockFrequencyHDMI() uint64 {
  return si.ClockFrequency("hdmi")
}

func (si Info) ClockFrequencyDPI() uint64 {
  return si.ClockFrequency("dpi")
}
