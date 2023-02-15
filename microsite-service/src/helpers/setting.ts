import Mongo from '../database/mongo/mongo'
import error from '../pkg/error'
import statusCode from '../pkg/statusCode'

export const Setting = async (database: string, id: string) => {
    const item = await Mongo.FindByIdSetting(database, id)

    if (!item) {
        throw new error(statusCode.NOT_FOUND, statusCode[statusCode.NOT_FOUND])
    }

    return item
}
