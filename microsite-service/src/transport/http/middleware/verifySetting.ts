import { NextFunction, Response } from 'express'
import { Setting } from '../../../helpers/setting'
import { ValidateObjectId } from '../../../helpers/validate'

export const VerifySettingByParams = (database: string) => {
    return async (req: any, res: Response, next: NextFunction) => {
        try {
            const idSetting = ValidateObjectId(
                req.params.idSetting,
                'idSetting'
            )
            const setting = await Setting(database, idSetting)
            req['setting'] = setting
            return next()
        } catch (error) {
            return next(error)
        }
    }
}
