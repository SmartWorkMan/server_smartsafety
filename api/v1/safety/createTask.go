package safety

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/commval"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/safety"
	"go.uber.org/zap"
	"time"
)

func GenerateTask(period string) {
	curDate := time.Now().Format("2006-01-02")
	global.GVA_LOG.Info(fmt.Sprintf("开始生成 %s 巡检任务, 日期:%s", period, curDate))

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
		task.AdminComment = ""

		if taskService.IsExistActiveTask(task) {
			global.GVA_LOG.Warn(fmt.Sprintf("巡检事项存在未完成任务,此次不生成,factoryName:%s, itemId:%d, inspectorUsername:%s", task.FactoryName, item.ID, item.InspectorUsername))
			continue
		}

		err = taskService.CreateTask(task)
		if err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("生成巡检任务失败!task:%+v", task), zap.Error(err))
		}
	}

	global.GVA_LOG.Info(fmt.Sprintf("成功生成 %s 巡检任务, 日期:%s", period, curDate))
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
		GenerateTask(commval.ItemPeriodDay)
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

			var period safety.ItemNextPeriodDate
			period.Period = commval.ItemPeriodWeek
			period.NextDate  = next.Format("2006-01-02")
			period.Interval = 7
			var err error
			if itemService.IsPeriodExist(period) {
				err = itemService.UpdateNextPeriodDate(period)
			} else {
				err = itemService.CreateNextPeriodDate(period)
			}
			if err != nil {
				global.GVA_LOG.Error("设置下一个重复周开始日期失败!", zap.Error(err))
			}
		}
		// 阻塞到执行时间
		t := time.NewTimer(next.Sub(now))
		<-t.C
		// 执行的任务内容
		GenerateTask(commval.ItemPeriodWeek)
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

			var period safety.ItemNextPeriodDate
			period.Period = commval.ItemPeriodMonth
			period.NextDate  = next.Format("2006-01-02")
			period.Interval = 30
			var err error
			if itemService.IsPeriodExist(period) {
				err = itemService.UpdateNextPeriodDate(period)
			} else {
				err = itemService.CreateNextPeriodDate(period)
			}
			if err != nil {
				global.GVA_LOG.Error("设置下一个重复月开始日期失败!", zap.Error(err))
			}
		}
		// 阻塞到执行时间
		t := time.NewTimer(next.Sub(now))
		<-t.C
		// 执行的任务内容
		GenerateTask(commval.ItemPeriodMonth)
	}
}

func DoCronTask() {
	now := time.Now()
	dateTime := time.Date(now.Year(), now.Month(), now.Day(), commval.CronTaskTime, 00, 0, 0, now.Location())

	go dailyTask(dateTime)
	go weeklyTask(dateTime)
	go monthlyTask(dateTime)
}

