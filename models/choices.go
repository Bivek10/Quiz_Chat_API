package models

type Choices struct {
	Base
	Choice      string `json:"choice"`
	IsCorrect   int64  `json:"is_correct"`
	Question_ID int64  `json:"question_id"`
}

func (m Choices) TableName() string {
	return "choices"
}

// to map convert questions to map
func (m Choices) ToMap() map[string]interface{} {
	return map[string]interface{}{

		"choice":     m.Choice,
		"is_correct": m.IsCorrect,
		"q_id":       m.Question_ID,
	}
}
