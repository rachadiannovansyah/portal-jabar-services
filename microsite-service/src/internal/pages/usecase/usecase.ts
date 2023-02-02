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

    public async store(body: Store) {
        const isExist = await this.repository.findByTitle(body.title)

        if (isExist)
            throw new error(
                statusCode.BAD_REQUEST,
                translate('exists', { attribute: 'title' })
            )

        const result = await this.repository.store(body)
        return result
    }

    public async show(id: string) {
        const item = await this.repository.findBySlug(id)

        if (!item)
            throw new error(
                statusCode.NOT_FOUND,
                statusCode[statusCode.NOT_FOUND]
            )

        return item
    }
}

export default Usecase
