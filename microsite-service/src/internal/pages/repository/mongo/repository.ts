import winston from 'winston'
import page from '../../../../database/mongo/schemas/page'
import { Store } from '../../entity/interface'

class Repository {
    constructor(private logger: winston.Logger) {}

    public async store(body: Store) {
        const pageNew = new page(body)

        return pageNew.save()
    }

    public async findByTitle(title: string) {
        return page.findOne({ title })
    }

    public async findBySlug(id: string) {
        return page.findOne({ slug: id })
    }
}

export default Repository
