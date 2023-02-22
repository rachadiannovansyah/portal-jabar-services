import winston from 'winston'
import Page from '../../../../database/mongo/schemas/page'
import { PropPaginate } from '../../../../helpers/paginate'
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

    public async FindAll({ offset, limit }: PropPaginate, database: string) {
        const page = Page(database)
        return page.find().skip(offset).limit(limit)
    }

    public async GetCount(database: string) {
        const page = Page(database)
        return page.count()
    }
}

export default Repository
