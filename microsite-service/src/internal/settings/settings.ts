import winston from 'winston'
import { Config } from '../../config/config.interface'
import Http from '../../transport/http/http'
import { VerifyAuth } from '../../transport/http/middleware/verifyAuth'
import Handler from './delivery/http/handler'
import Repository from './repository/mongo/repository'
import Usecase from './usecase/usecase'

class Settings {
    constructor(
        private http: Http,
        private logger: winston.Logger,
        private config: Config
    ) {
        const repository = new Repository(logger, config.db.name)
        const usecase = new Usecase(repository, logger)

        this.loadHttp(usecase)
    }

    private loadHttp(usecase: Usecase) {
        const handler = new Handler(usecase, this.logger)
        const verify = VerifyAuth(this.config.jwt.access_key)
        this.http.app.post('/v1/settings/', verify, handler.Store())
        this.http.app.get('/v1/settings/', verify, handler.FindAll())
        this.http.app.get('/v1/settings/:idSetting', verify, handler.Show())
    }
}

export default Settings
