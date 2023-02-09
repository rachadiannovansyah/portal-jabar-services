import schemaSetting from '../database/mongo/schemas/setting'
import error from '../pkg/error'
import statusCode from '../pkg/statusCode'

const findByDomain = async (database: string, origin: string) => {
    const setting = schemaSetting(database)

    return setting.findOne({
        domain: origin,
    })
}

export const Setting = async (database: string, origin: string) => {
    origin = origin.replace('http://', '').replace('https://', '')

    if (!origin)
        throw new error(
            statusCode.BAD_GATEWAY,
            statusCode[statusCode.BAD_GATEWAY]
        )

    const item = await findByDomain(database, origin)

    if (!item) {
        throw new error(statusCode.NOT_FOUND, statusCode[statusCode.NOT_FOUND])
    }

    return item
}
