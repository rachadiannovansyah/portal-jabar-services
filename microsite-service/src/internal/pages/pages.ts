import winston from 'winston'
import { Config } from '../../config/config.interface'
import Http from '../../transport/http/http'
import { VerifyAuth } from '../../transport/http/middleware/verifyAuth'
import {
    VerifySettingByDomain,
    VerifySettingById,
} from '../../transport/http/middleware/verifySetting'
import HttpHandler from './delivery/http/handler'
import Repository from './repository/mongo/repository'
import Usecase from './usecase/usecase'

class Pages {
    private httpHandler: HttpHandler
    constructor(
        private http: Http,
        private logger: winston.Logger,
        private config: Config
    ) {
        const repository = new Repository(logger)
        const usecase = new Usecase(repository, logger)
        this.httpHandler = new HttpHandler(usecase, this.logger)
        this.loadHttp()
    }

    private loadHttp() {
        this.httpPublic()
        this.httpCms()
    }

    private httpPublic() {
        const verifySettingByDomain = VerifySettingByDomain(this.config.db.name)

        this.http.app.get(
            '/v1/public/pages/:slug',
            verifySettingByDomain,
            this.httpHandler.FindBySlug()
        )
    }

    private httpCms() {
        const verifyAuth = VerifyAuth(this.config.jwt.access_key)
        const verifySettingById = VerifySettingById(this.config.db.name)

        this.http.app.post(
            '/v1/pages/:idSetting',
            verifyAuth,
            verifySettingById,
            this.httpHandler.Store()
        )
        this.http.app.get(
            '/v1/pages/:idSetting',
            verifyAuth,
            verifySettingById,
            this.httpHandler.FindAll()
        )
        this.http.app.get(
            '/v1/pages/:idSetting/:idPage',
            verifyAuth,
            verifySettingById,
            this.httpHandler.Show()
        )
    }
}

export default Pages
