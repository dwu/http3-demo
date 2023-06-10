package messages

type MessageSize struct {
	Size uint32
}

type MessageData struct {
	DataLen uint32
	Data    []byte
}

type MessageFilename struct {
	Filename string
}
