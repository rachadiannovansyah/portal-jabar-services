import { NextFunction, Response } from 'express'
import Mongo from '../../../database/mongo/mongo'
import { ValidateObjectId } from '../../../helpers/validate'
import error from '../../../pkg/error'
import statusCode from '../../../pkg/statusCode'

export const VerifySettingById = (database: string) => {
    return async (req: any, res: Response, next: NextFunction) => {
        try {
            const idSetting = ValidateObjectId(
                req.params.idSetting,
                'idSetting'
            )
            const setting = await Mongo.FindByIdSetting(database, idSetting)

            if (!setting) {
                throw new error(
                    statusCode.NOT_FOUND,
                    statusCode[statusCode.NOT_FOUND]
                )
            }

            req['setting'] = setting
            return next()
        } catch (error) {
            return next(error)
        }
    }
}

export const VerifySettingByDomain = (database: string) => {
    return async (req: any, res: Response, next: NextFunction) => {
        try {
            const setting = await Mongo.FindByDomainSetting(
                database,
                req.headers.origin
            )

            if (!setting) {
                throw new error(
                    statusCode.NOT_FOUND,
                    statusCode[statusCode.NOT_FOUND]
                )
            }

            req['setting'] = setting
            return next()
        } catch (error) {
            return next(error)
        }
    }
}
