package protocol

type LogPlugin interface {
	PrintMessage(message string)
}
