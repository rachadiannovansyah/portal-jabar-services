import winston from 'winston'
import { translate } from '../../../helpers/translate'
import error from '../../../pkg/error'
import statusCode from '../../../pkg/statusCode'
import { Store } from '../entity/interface'
import Repository from '../repository/mongo/repository'

class Usecase {
    constructor(
        private repository: Repository,
        private logger: winston.Logger
    ) {}

    public async Store(body: Store) {
        const isExist = await this.repository.FindByDomain(body.domain)

        if (isExist)
            throw new error(
                statusCode.BAD_REQUEST,
                translate('exists', { attribute: 'domain' })
            )

        const result = await this.repository.Store(body)
        return result
    }

    public async Show(id: string) {
        const item = await this.repository.FindByID(id)

        if (!item)
            throw new error(
                statusCode.NOT_FOUND,
                statusCode[statusCode.NOT_FOUND]
            )

        return item
    }
}

export default Usecase
