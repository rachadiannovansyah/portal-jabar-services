package job

import (
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

// NewsPublishingJob ...
func NewsPublishingJob(conn *utils.Conn) {
	logrus.Println("NewsPublishingJob is running")
	_, err := conn.Mysql.Exec("update news set is_live=1 where start_date <= now() and end_date >= now() and status='PUBLISHED' and is_live=0")
	if err != nil {
		logrus.Error(err)
	}
	return
}
