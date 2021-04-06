/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:scheduler.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package serializer

import "singo/model"

// Scheduler 调度消息记录
type Scheduler struct {
	ID               uint   `json:"id"`
	CreatedAt        int64  `json:"created_at"`
	UniversityName   string `json:"university_name"`
	UniversityUserID uint   `json:"university_user_id"`
	MessageID        uint   `json:"message_id"`
	//CertificationID  uint   `json:"certification_id"`
	//Err              string `json:"-"`
	States      int    `json:"states"`
	StudentName string `json:"student_name"`
}

// BuildScheduler 序列化消息
func BuildScheduler(scheduler model.Scheduler) Scheduler {
	return Scheduler{
		ID:               scheduler.ID,
		CreatedAt:        scheduler.CreatedAt.Unix(),
		UniversityName:   scheduler.UniversityName,
		UniversityUserID: scheduler.UniversityUserID,
		MessageID:        scheduler.MessageID,
		StudentName:      scheduler.StudentName,
		States:           scheduler.Status,
	}
}

// BuildSchedulers 序列化消息列表
func BuildSchedulers(items []model.Scheduler) []Scheduler {
	var schedulers []Scheduler

	for _, item := range items {
		scheduler := BuildScheduler(item)
		schedulers = append(schedulers, scheduler)
	}
	return schedulers
}

// BuildSchedulerResponse 序列化消息响应
func BuildSchedulerResponse(scheduler model.Scheduler) Response {
	return Response{
		Code: 0,
		Data: BuildScheduler(scheduler),
	}
}
