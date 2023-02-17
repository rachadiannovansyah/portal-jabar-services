import winston from 'winston'
import { status } from '../../../../database/constant/setting'
import Setting from '../../../../database/mongo/schemas/setting'
import { RemoveProcotol } from '../../../../helpers/http'
import { PropPaginate } from '../../../../helpers/paginate'
import { Store } from '../../entity/interface'

class Repository {
    private setting
    constructor(private logger: winston.Logger, private database: string) {
        this.setting = Setting(database)
    }

    public async Store(body: Store) {
        const settingNew = new this.setting({
            ...body,
            status: status.DRAFT,
            is_active: false,
        })

        return settingNew.save()
    }

    public async FindByDomain(domain: string) {
        return this.setting.findOne({ domain: RemoveProcotol(domain) })
    }

    public async FindByID(id: string) {
        return this.setting.findById(id)
    }

    public async FindAll({ limit, offset }: PropPaginate) {
        return this.setting.find().skip(offset).limit(limit)
    }

    public async GetCount() {
        return this.setting.count()
    }
}

export default Repository
