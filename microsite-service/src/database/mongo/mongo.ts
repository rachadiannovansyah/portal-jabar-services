import mongoose, { Schema } from 'mongoose'
import winston from 'winston'
import { Config } from '../../config/config.interface'
import { RemoveProcotol } from '../../helpers/http'
import Setting from './schemas/setting'

class Mongo {
    public static async Connect(logger: winston.Logger, { db }: Config) {
        mongoose.set('strictQuery', false)
        return mongoose
            .connect(`mongodb://${db.host}:${db.port}/${db.name}`, {
                authSource: db.auth_source,
                pass: db.password,
                user: db.username,
            })
            .then(() => {
                logger.info('Connection to database established')
            })
            .catch((e) => {
                logger.error(e.message)
                process.exit(-1)
            })
    }

    public static Model(database: string, collection: string, schema: Schema) {
        return mongoose.connection
            .useDb(database, {
                useCache: true,
            })
            .model(collection, schema)
    }

    public static FindByIdSetting(database: string, id: string) {
        const setting = Setting(database)

        return setting.findById(id)
    }

    public static FindByDomainSetting(database: string, domain: string) {
        const setting = Setting(database)

        return setting.findOne({
            domain: RemoveProcotol(domain),
        })
    }
}

export default Mongo
