package job

import (
	"context"

	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

// PopUpBannerActivateJob ...
func PopUpBannerActivateJob(conn *utils.Conn, cfg *config.Config) {
	logrus.Println("PopUpBannerActivateJob is running")

	// Get pop_up_banners ids from pop_up_banners will be archived
	var ID string
	query := `SELECT id FROM pop_up_banners
		WHERE status='ACTIVE' 
		AND is_live=0 
		AND start_date <= NOW()`
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
		ID = id
	}

	if ID != "" {
		// deactivate banner
		deactivateQuery := `UPDATE pop_up_banners 
		SET status = 'NON-ACTIVE', is_live = 0 
		WHERE status = 'ACTIVE'`

		stmt, err := conn.Mysql.PrepareContext(context.TODO(), deactivateQuery)
		if err != nil {
			logrus.Error(err)
		}

		_, err = stmt.ExecContext(context.TODO())
		if err != nil {
			logrus.Error(err)
		}
	}

	// activate banner
	activateQuery := `UPDATE pop_up_banners SET is_live=1, status='ACTIVE', updated_at=now() WHERE id=?`
	stmt, err := conn.Mysql.PrepareContext(context.TODO(), activateQuery)
	if err != nil {
		logrus.Error(err)
	}

	res, err := stmt.ExecContext(context.TODO(), ID)
	if err != nil {
		logrus.Error(err)
	}

	// rows affected
	if ra, err := res.RowsAffected(); err != nil {
		logrus.Error("ErrPopUpBannerActivateJob: ", err)
	} else {
		logrus.Println("PopUpBannerActivateJob: Rows affected: ", ra)
	}
}
