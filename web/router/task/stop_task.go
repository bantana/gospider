package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nange/gospider/common"
	"github.com/nange/gospider/spider"
	"github.com/nange/gospider/web/core"
	"github.com/nange/gospider/web/model"
	"github.com/nange/gospider/web/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func StopTask(c *gin.Context) {
	taskIDStr := c.Param("id")
	if taskIDStr == "" {
		logrus.Warnf("StopTaskReq taskID is empty")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		logrus.Warnf("StopTaskReq taskID format is invalid, taskID:%v", taskIDStr)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	logrus.Infof("StopTaskReq:%+v", taskID)

	// stop cron task
	if ct := service.GetCronTask(taskID); ct != nil {
		ct.Stop()
	}

	// cancel spider task
	if err := spider.CancelTask(taskID); err != nil {
		logrus.Warnf("spider.CancelTask err:%+v", err)
	}

	// set task status to TaskStatusStopped
	err = model.NewTaskQuerySet(core.GetDB()).IDEq(taskID).GetUpdater().SetStatus(common.TaskStatusStopped).Update()
	if err != nil {
		logrus.Errorf("update task status err:%+v", errors.WithStack(err))
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.String(http.StatusOK, "success")
}
