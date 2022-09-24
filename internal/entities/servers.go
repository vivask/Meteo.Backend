package entities

type StatusXU4 struct {
	TransmissionRunnig bool
	TransmissionIdle   bool
	SambaRunning       bool
}

type StatusN2 struct {
	KodiRunning  bool
	StorageMount bool
}
