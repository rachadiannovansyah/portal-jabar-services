import winston from 'winston'
import Page from '../../../../database/mongo/schemas/page'
import { Store } from '../../entity/interface'

class Repository {
    private page
    constructor(private logger: winston.Logger, database: string) {
        this.page = Page(database)
    }

    public async store(body: Store) {
        const pageNew = new this.page(body)

        return pageNew.save()
    }

    public async findByTitle(title: string) {
        return this.page.findOne({ title })
    }

    public async findBySlug(id: string) {
        return this.page.findOne({ slug: id })
    }
}

export default Repository
