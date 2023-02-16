import { NextFunction, Request, Response } from 'express'
import winston from 'winston'
import Usecase from '../../usecase/usecase'
import { Store } from '../../entity/schema'
import {
    ValidateFormRequest,
    ValidateObjectId,
} from '../../../../helpers/validate'
import statusCode from '../../../../pkg/statusCode'
import { Setting } from '../../../../helpers/setting'

class Handler {
    constructor(
        private usecase: Usecase,
        private logger: winston.Logger,
        private database: string
    ) {}
    public Store() {
        return async (req: any, res: Response, next: NextFunction) => {
            try {
                const value = ValidateFormRequest(Store, req.body)
                const idSetting = ValidateObjectId(
                    req.params.idSetting,
                    'idSetting'
                )
                const setting = await Setting(this.database, idSetting)

                const result = await this.usecase.Store(value, setting.id)
                return res
                    .status(statusCode.OK)
                    .json({ data: result.toJSON(), message: 'CREATED' })
            } catch (error) {
                return next(error)
            }
        }
    }
    public Show() {
        return async (req: any, res: Response, next: NextFunction) => {
            try {
                const idSetting = ValidateObjectId(
                    req.params.idSetting,
                    'idSetting'
                )
                const idPage = ValidateObjectId(req.params.idPage, 'idPage')
                const setting = await Setting(this.database, idSetting)
                const result = await this.usecase.Show(idPage, setting.id)
                return res.json({
                    data: {
                        setting,
                        page: result,
                    },
                })
            } catch (error) {
                return next(error)
            }
        }
    }
}

export default Handler
