import winston from 'winston'
import { Config } from '../../config/config.interface'
import Http from '../../transport/http/http'
import { VerifyAuth } from '../../transport/http/middleware/verifyAuth'
import { VerifySettingByParams } from '../../transport/http/middleware/verifySetting'
import Handler from './delivery/http/handler'
import Repository from './repository/mongo/repository'
import Usecase from './usecase/usecase'

class Pages {
    constructor(
        private http: Http,
        private logger: winston.Logger,
        private config: Config
    ) {
        const repository = new Repository(logger)
        const usecase = new Usecase(repository, logger)

        this.loadHttp(usecase)
    }

    private loadHttp(usecase: Usecase) {
        const handler = new Handler(usecase, this.logger, this.config.db.name)
        const verifyAuth = VerifyAuth(this.config.jwt.access_key)
        const verifySetting = VerifySettingByParams(this.config.db.name)
        this.http.app.post(
            '/v1/pages/:idSetting',
            verifyAuth,
            verifySetting,
            handler.Store()
        )
        this.http.app.get(
            '/v1/pages/:idSetting',
            verifyAuth,
            verifySetting,
            handler.FindAll()
        )
        this.http.app.get(
            '/v1/pages/:idSetting/:idPage',
            verifyAuth,
            verifySetting,
            handler.Show()
        )
    }
}

export default Pages
