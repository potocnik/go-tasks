package models

type OperationData struct {
	Index int
	Text  string
}

type QueMessage struct {
	Operation string
	Data      *OperationData
}

func NewQueueMessage(operation string, text string, index int) QueMessage {
	result := QueMessage{}
	result.Operation = operation
	if index > -1 || text != "" {
		data := OperationData{}
		result.Data = &data
		result.Data.Index = index
		result.Data.Text = text
	}
	return result
}

func IsEmptyMessage(message *QueMessage) bool {
	return message == nil || message == &QueMessage{}
}
