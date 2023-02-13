package job

import (
	"context"

	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

func PopUpBannerDeactivateJob(conn *utils.Conn, cfg *config.Config) {
	logrus.Println("PopUpBannerDeactivateJob is running...")

	// deactivate banner
	deactivateQuery := `UPDATE pop_up_banners 
	SET status = 'NON-ACTIVE',
	is_live = 0,
	updated_at = now()
	WHERE status = 'ACTIVE'
	AND end_date < NOW()`

	stmt, err := conn.Mysql.PrepareContext(context.TODO(), deactivateQuery)
	if err != nil {
		logrus.Error(err)
	}

	res, err := stmt.ExecContext(context.TODO())
	if err != nil {
		logrus.Error(err)
	}

	rowAffected, _ := res.RowsAffected()

	// just printed out job activities
	if rowAffected != 0 {
		logrus.Println("PopUpBannerDeactivateJob: Deactivate banner row affected: ", rowAffected)
	} else {
		logrus.Println("PopUpBannerDeactivateJob: No-one deactivate pop up banner")
	}
}
