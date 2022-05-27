package safety

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"go.uber.org/zap"
	"time"
)

func generateTask(period string) {
	curDate := time.Now().Format("2006-01-02")
	global.GVA_LOG.Info(fmt.Sprintf("开始生成巡检任务, 日期:%s", curDate))

	err, itemList := itemService.GetAllValidItemList(period)
	if err != nil {
		global.GVA_LOG.Error("生成巡检任务失败!", zap.Error(err))
		return
	}

	for _, item := range itemList {
		var task safety.Task
		task.AreaId = item.AreaId
		task.AreaName = item.AreaName
		task.FactoryName = item.FactoryName
		task.InspectorName = item.InspectorName
		task.InspectorUsername = item.InspectorUsername
		task.ItemName = item.ItemName
		task.ItemSn = item.ItemSn
		task.Period = item.Period
		task.PlanInspectionDate = curDate
		task.Standard = item.Standard
		task.ItemId = item.ID
		task.TaskStatus = commval.TaskStatusNotStart
		task.TaskStatusStr = commval.TaskStatus[commval.TaskStatusNotStart]

		err = taskService.CreateTask(task)
		if err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("生成巡检任务失败!task:%+v", task), zap.Error(err))
		}
	}

	global.GVA_LOG.Info(fmt.Sprintf("成功生成巡检任务, 日期:%s", curDate))
	return
}

func dailyTask(dateTime time.Time) {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		// 检查是否超过当日的时间
		if next.Sub(now) < 0 {
			global.GVA_LOG.Info(fmt.Sprintf("执行时间未到, 日期:%s", now.String()))
			next = now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		}
		// 阻塞到执行时间
		t := time.NewTimer(next.Sub(now))
		<-t.C
		// 执行的任务内容
		generateTask(commval.ItemPeriodDay)
	}
}

func weeklyTask(dateTime time.Time) {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		// 检查是否超过当日的时间
		if next.Sub(now) < 0 {
			global.GVA_LOG.Info(fmt.Sprintf("执行时间未到, 日期:%s", now.String()))
			next = now.Add(time.Hour * 24 * 7)
			next = time.Date(next.Year(), next.Month(), next.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		}
		// 阻塞到执行时间
		t := time.NewTimer(next.Sub(now))
		<-t.C
		// 执行的任务内容
		generateTask(commval.ItemPeriodWeek)
	}
}

func monthlyTask(dateTime time.Time) {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		// 检查是否超过当日的时间
		if next.Sub(now) < 0 {
			global.GVA_LOG.Info(fmt.Sprintf("执行时间未到, 日期:%s", now.String()))
			next = now.Add(time.Hour * 24 * 30)
			next = time.Date(next.Year(), next.Month(), next.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), dateTime.Location())
		}
		// 阻塞到执行时间
		t := time.NewTimer(next.Sub(now))
		<-t.C
		// 执行的任务内容
		generateTask(commval.ItemPeriodMonth)
	}
}

func DoCronTask() {
	now := time.Now()
	dateTime := time.Date(now.Year(), now.Month(), now.Day(), commval.CronTaskTime, 00, 0, 0, now.Location())

	go dailyTask(dateTime)
	go weeklyTask(dateTime)
	go monthlyTask(dateTime)
}

