import winston from 'winston'
import { Config } from '../../config/config.interface'
import Http from '../../transport/http/http'
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
        const verify = this.http.VerifyAuth(this.config.jwt.access_key)
        this.http.app.post('/v1/pages/:idSetting', verify, handler.Store())
        this.http.app.get(
            '/v1/pages/:idSetting/:idPages',
            verify,
            handler.Show()
        )
    }
}

export default Pages
