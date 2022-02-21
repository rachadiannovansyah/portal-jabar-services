package job

import (
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

// NewsPublishingJob ...
func NewsPublishingJob(conn *utils.Conn) {
	logrus.Println("NewsPublishingJob is running")
	res, err := conn.Mysql.Exec("update news set is_live=1, published_at = now() where status='PUBLISHED' and is_live=0 and start_date <= now()")
	if err != nil {
		logrus.Error(err)
	}

	// rows affected
	ra, err := res.RowsAffected()

	if err != nil {
		logrus.Error(err)
	}

	logrus.Println("NewsPublishingJob: Rows affected: ", ra)

	return
}
