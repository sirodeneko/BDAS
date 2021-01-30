package serializer

import "singo/model"

// Message  各种需要管理员进行处理的通知
type Message struct {
	ID               uint             `json:"id"`
	CreatedAt        int64            `json:"created_at"`
	MsgType          string           `json:"msg_type"`
	Description      string           `json:"description"`
	StudentAcMsg     StudentAcMsg     `json:"student_ac_msg"`
	EducationalAcMsg EducationalAcMsg `json:"educational_ac_msg"`
}

type StudentAcMsg struct {
	ID           uint   `json:"id"`
	CreatedAt    int64  `json:"created_at"`
	UserId       uint   `json:"user_id"`
	Name         string `json:"name"`
	CardCode     string `json:"card_code"`
	FrontFaceImg string `json:"front_face_img"`
	BackFaceImg  string `json:"back_face_img"`
}

type EducationalAcMsg struct {
	ID                uint   `json:"id"`
	CreatedAt         int64  `json:"created_at"`
	Name              string `json:"name"`
	Sex               uint   `json:"sex"`                // 0 男 1女
	Ethnic            string `json:"ethnic"`             // 民族
	Birthday          int64  `json:"birthday"`           // 生日
	CardCode          string `json:"card_code"`          // 身份证号
	EducationCategory string `json:"education_category"` // 学历类别
	Level             string `json:"level"`              // 层次
	University        string `json:"university"`         // 学校
	Professional      string `json:"professional"`       // 专业
	LearningFormat    string `json:"learning_format"`    // 学习形式
	EducationalSystem string `json:"educational_system"` // 学制
	AdmissionDate     string `json:"admission_date"`     // 入学日期
	GraduationDate    string `json:"graduation_date"`    // 毕业日期
	Status            string `json:"status"`             // 状态（是否结业）
	StudentAvatar     string `json:"student_avatar"`     // 照片
}

// BuildMessage 序列化消息
func BuildMessage(message model.Message) Message {
	return Message{
		ID:          message.ID,
		CreatedAt:   message.CreatedAt.Unix(),
		MsgType:     message.MsgType,
		Description: message.Description,
		StudentAcMsg: StudentAcMsg{
			ID:           message.StudentAcMsg.ID,
			CreatedAt:    message.StudentAcMsg.CreatedAt.Unix(),
			UserId:       message.StudentAcMsg.UserId,
			Name:         message.StudentAcMsg.Name,
			CardCode:     message.StudentAcMsg.CardCode,
			FrontFaceImg: message.StudentAcMsg.FrontFaceImg,
			BackFaceImg:  message.StudentAcMsg.BackFaceImg,
		},
		EducationalAcMsg: EducationalAcMsg{
			ID:                message.EducationalAcMsg.ID,
			CreatedAt:         message.EducationalAcMsg.CreatedAt.Unix(),
			Name:              message.EducationalAcMsg.Name,
			Sex:               message.EducationalAcMsg.Sex,
			Ethnic:            message.EducationalAcMsg.Ethnic,
			Birthday:          message.EducationalAcMsg.Birthday.Unix(),
			CardCode:          message.EducationalAcMsg.CardCode,
			EducationCategory: message.EducationalAcMsg.EducationCategory,
			Level:             message.EducationalAcMsg.Level,
			University:        message.EducationalAcMsg.University,
			Professional:      message.EducationalAcMsg.Professional,
			LearningFormat:    message.EducationalAcMsg.LearningFormat,
			EducationalSystem: message.EducationalAcMsg.EducationalSystem,
			AdmissionDate:     message.EducationalAcMsg.AdmissionDate,
			GraduationDate:    message.EducationalAcMsg.GraduationDate,
			Status:            message.EducationalAcMsg.Status,
			StudentAvatar:     message.EducationalAcMsg.StudentAvatar,
		},
	}
}

// BuildMessages 序列化消息列表
func BuildMessages(items []model.Message) []Message {
	var messages []Message

	for _, item := range items {
		message := BuildMessage(item)
		messages = append(messages, message)
	}
	return messages
}

// BuildMessageResponse 序列化消息响应
func BuildMessageResponse(message model.Message) Response {
	return Response{
		Code: 0,
		Data: BuildMessage(message),
	}
}
