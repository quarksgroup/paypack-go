package paypack

// Driver identifies the paypack platform driver
type Driver uint

// Supported drivers
const (
	DriverUnknown Driver = iota
	DriverPaypack
)

func (d Driver) String() string {
	switch d {
	case DriverPaypack:
		return "paypack"
	default:
		return "unknown"
	}
}
