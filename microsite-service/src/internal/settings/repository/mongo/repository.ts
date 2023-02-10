import { randomUUID } from 'crypto'
import winston from 'winston'
import { status } from '../../../../database/constant/setting'
import Setting from '../../../../database/mongo/schemas/setting'
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
        return this.setting.findOne({ domain })
    }

    public async FindByID(id: string) {
        return this.setting.findById(id)
    }
}

export default Repository
