package job

import (
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

// NewsArchiveJob ...
func NewsArchiveJob(conn *utils.Conn) {
	logrus.Println("NewsArchiveJob is running")
	res, err := conn.Mysql.Exec("UPDATE news SET is_live=0, status='ARCHIVED' WHERE status='PUBLISHED' AND end_date > NOW()")
	if err != nil {
		logrus.Error(err)
	}

	// rows affected
	ra, err := res.RowsAffected()

	if err != nil {
		logrus.Error(err)
	}

	logrus.Println("NewsArchiveJob: Rows affected: ", ra)

	return
}
