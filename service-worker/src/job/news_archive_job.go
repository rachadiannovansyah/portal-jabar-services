package job

import (
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

// NewsArchiveJob ...
func NewsArchiveJob(conn *utils.Conn) {
	logrus.Println("NewsArchiveJob is running")
	_, err := conn.Mysql.Exec("update news set is_live=0, status='ARCHIVED' where end_date < now() and status='PUBLISHED'")
	if err != nil {
		logrus.Error(err)
	}
	return
}
