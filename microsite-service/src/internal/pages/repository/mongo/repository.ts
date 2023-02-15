import winston from 'winston'
import Page from '../../../../database/mongo/schemas/page'
import { Store } from '../../entity/interface'

class Repository {
    constructor(private logger: winston.Logger) {}

    public async Store(body: Store, database: string) {
        const page = Page(database)
        const pageNew = new page(body)

        return pageNew.save()
    }

    public async FindByTitle(title: string, database: string) {
        const page = Page(database)
        return page.findOne({ title })
    }

    public async FindById(id: string, database: string) {
        const page = Page(database)
        return page.findById(id)
    }
}

export default Repository
