package job

import (
	"context"

	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

func PopUpBannerDeactivateJob(conn *utils.Conn, cfg *config.Config) {
	logrus.Println("PopUpBannerDeactivateJob is running...")

	// get banners id will be deactivated
	var IDs string

	query := `SELECT id FROM pop_up_banners
		WHERE status = 'ACTIVE'
		AND end_date < NOW()`

	rows, err := conn.Mysql.Query(query)
	if err != nil {
		logrus.Error(err)
		return
	}

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			logrus.Error(err)
			return
		}
		IDs = id
	}

	// deactivate banner
	deactivateQuery := `UPDATE pop_up_banners SET status = ? WHERE id = ?`
	stmt, err := conn.Mysql.PrepareContext(context.TODO(), deactivateQuery)
	if err != nil {
		logrus.Error(err)
	}

	_, err = stmt.ExecContext(context.TODO(),
		"NON-ACTIVE",
		IDs,
	)
	if err != nil {
		logrus.Error(err)
	}

	// just printed out job activities
	if IDs != "" {
		logrus.Println("PopUpBannerDeactivateJob: Deactivate banner with id: ", IDs)
	} else {
		logrus.Println("PopUpBannerDeactivateJob: No-one active pop up banner")
	}
}
