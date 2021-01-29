package iolimit

type Request struct {
	DeviceName string
	PodUid     string
	ContainerRuntime string
	IOLimit    *IOMax
}

type ValidRequest struct {
	FilePath string
	DeviceNumber *DeviceNumber
	IOMax    *IOMax
}

type IOMax struct {
	Riops uint64
	Wiops uint64
	Rbps  uint64
	Wbps  uint64
}

type DeviceNumber struct {
	Major uint64
	Minor uint64
}