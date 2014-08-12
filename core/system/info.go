package system

import (
  "github.com/dalmirdasilva/gorpi/util"
  "strings"
  "strconv"
  "os"
  "errors"
  "io/ioutil"
)

type Info struct {
  cpuInfoFilePath string
  cpuInfo map[string]string
}

type Board int

const DEFAULT_CPU_INFO_PATH = "/proc/cpuinfo"

const (
  Unknown Board = iota
  ModelARev0
  ModelBRev1
  ModelBRev2
)

func InfoInstance() Info {
  return Info{}
}

func (info *Info) SetCpuInfoFilePath(filePath string) {
  info.cpuInfoFilePath = filePath
}

func (info Info) CpuInfo() (map[string]string, error) {
  if (info.cpuInfo == nil) {
    info.cpuInfo = make(map[string]string)
    lines, err := info.readCpuInfoFile()
    if err != nil {
      return nil, err
    }
    for i := range lines {
      parts := strings.Split(lines[i], ":")
      if len(parts) >= 2 {
        key := strings.ToLower(strings.TrimSpace(parts[0]))
        val := strings.ToLower(strings.TrimSpace(parts[1]))
        if key != "" && val != "" {
          info.cpuInfo[key] = val
        }
      }
    }
  }
  return info.cpuInfo, nil
}

func (info Info) CpuInfoEntry(entry string) (string, error) {
  cpuInfo, err := info.CpuInfo()
  if err != nil {
    return "", err
  }
  entry = strings.ToLower(entry)
  value, exists := cpuInfo[entry]
  if !exists {
    err := errors.New("Entry not found: " + entry)
    return "", err
  }
  return value, nil
}


func (info Info) Processor() (string, error) {
  return info.CpuInfoEntry("processor")
}

func (info Info) BogoMIPS() (string, error) {
  return info.CpuInfoEntry("BogoMIPS")
}

func (info Info) CpuFeatures() ([]string, error) {
  features, err := info.CpuInfoEntry("Features")
  if (err != nil) {
    return nil, err
  }
  return strings.Split(features, " "), nil
}

func (info Info) CpuImplementer() (string, error) {
  return info.CpuInfoEntry("CPU implementer")
}

func (info Info) CpuArchitecture() (string, error) {
  return info.CpuInfoEntry("CPU architecture")
}

func (info Info) CpuVariant() (string, error) {
  return info.CpuInfoEntry("CPU variant")
}

func (info Info) CpuPart() (string, error) {
  return info.CpuInfoEntry("CPU part")
}

func (info Info) CpuRevision() (string, error) {
  return info.CpuInfoEntry("CPU revision")
}

func (info Info) Hardware() (string, error) {
  return info.CpuInfoEntry("Hardware")
}

func (info Info) Revision() (string, error) {
  return info.CpuInfoEntry("Revision")
}

func (info Info) Serial() (string, error) {
  return info.CpuInfoEntry("Serial")
}

func (info Info) Memory() (map[string]uint32, error) {
  result := make(map[string]uint32)
  keys := []string{"total", "used", "free", "shared", "buffers", "cached"}
  lines, err := util.Execute("free", "-b")
  if err != nil {
    return nil, err
  }
  for i := range lines {
    line := lines[i]
    if strings.HasPrefix(line, "Mem:") {
      parts := strings.Split(line, " ")
      for j := range parts {
        if j > 0 {
          part := strings.TrimSpace(parts[j])
          if len(part) > 0 {
            ui, _ := strconv.ParseUint(part, 10, 32)
            result[keys[j-1]] = uint32(ui)
          }
        }
      }
    }
  }
  return result, nil
}

func (info Info) memoryOf(of string) uint32 {
  memory, err := info.Memory()
  if err != nil {
    return 0
  }
  return memory[of]
}

func (info Info) MemoryTotal() uint32 {
  return info.memoryOf("total")
}

func (info Info) MemoryUsed() uint32 {
  return info.memoryOf("used")
}

func (info Info) MemoryFree() uint32 {
  return info.memoryOf("free")
}

func (info Info) MemoryShared() uint32 {
  return info.memoryOf("shared")
}

func (info Info) MemoryBuffers() uint32 {
  return info.memoryOf("buffers")
}

func (info Info) MemoryCached() uint32 {
  return info.memoryOf("cached")
}

func (info Info) BoardModel() Board {
  revision, err := info.Revision()
  if err == nil {
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

func (info Info) Temperature() float32 {
  cmd := "/opt/vc/bin/vcgencmd"
  _, err := os.Stat(cmd)
  if os.IsNotExist(err) {
    return 0.0
  }
  lines, err := util.Execute(cmd, "measure_temp")
  if err != nil {
    return 0.0
  }
  for i := range lines {
    line := lines[i]
    if strings.HasPrefix(line, "temp=") {
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


func (info Info) ClockFrequency(target string) uint64 {
  cmd := "/opt/vc/bin/vcgencmd"
  _, err := os.Stat(cmd)
  if os.IsNotExist(err) {
    return 0
  }
  lines, err := util.Execute(cmd, "measure_clock", strings.TrimSpace(target))
  for i := range lines {
    line := lines[i]
    if strings.HasPrefix(line, "frequency") {
      parts := strings.Split(line, "=")
      freq, _ := strconv.ParseUint(parts[1], 10, 64)
      return uint64(freq)
    }
  }
  return 0
}

func (info Info) ClockFrequencyArm() uint64 {
  return info.ClockFrequency("arm")
}

func (info Info) ClockFrequencyCore() uint64 {
  return info.ClockFrequency("core")
}

func (info Info) ClockFrequencyH264() uint64 {
  return info.ClockFrequency("h264")
}

func (info Info) ClockFrequencyISP() uint64 {
  return info.ClockFrequency("isp")
}

func (info Info) ClockFrequencyV3D() uint64 {
  return info.ClockFrequency("v3d")
}

func (info Info) ClockFrequencyUART() uint64 {
  return info.ClockFrequency("uart")
}

func (info Info) ClockFrequencyPWM() uint64 {
  return info.ClockFrequency("pwm")
}

func (info Info) ClockFrequencyEMMC() uint64 {
  return info.ClockFrequency("emmc")
}

func (info Info) ClockFrequencyPixel() uint64 {
  return info.ClockFrequency("pixel")
}

func (info Info) ClockFrequencyVEC() uint64 {
  return info.ClockFrequency("vec")
}

func (info Info) ClockFrequencyHDMI() uint64 {
  return info.ClockFrequency("hdmi")
}

func (info Info) ClockFrequencyDPI() uint64 {
  return info.ClockFrequency("dpi")
}

func (info Info) readCpuInfoFile() ([]string, error) {
  if info.cpuInfoFilePath == "" {
    info.cpuInfoFilePath = DEFAULT_CPU_INFO_PATH
  }
  content, err := ioutil.ReadFile(info.cpuInfoFilePath)
  if err != nil {
    return nil, err
  }
  return strings.Split(string(content), "\n"), nil
}
