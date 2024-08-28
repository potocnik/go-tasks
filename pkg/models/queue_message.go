package models

type OperationData struct {
	Index int
	Text  string
}

type QueueMessage struct {
	Operation string
	Data      *OperationData
}

func NewQueueMessage(operation string, text string, index int) QueueMessage {
	result := QueueMessage{}
	result.Operation = operation
	if index > -1 || text != "" {
		data := OperationData{}
		result.Data = &data
		result.Data.Index = index
		result.Data.Text = text
	}
	return result
}

func IsEmptyMessage(message *QueueMessage) bool {
	return message == nil || message == &QueueMessage{}
}
