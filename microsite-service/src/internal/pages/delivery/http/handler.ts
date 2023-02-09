import { NextFunction, Request, Response } from 'express'
import winston from 'winston'
import Usecase from '../../usecase/usecase'
import { Store } from '../../entity/schema'
import { validateFormRequest } from '../../../../helpers/validate'
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
                const value = validateFormRequest(Store, req.body)                
                const setting = await Setting(this.database, req.headers.origin)
                
                await this.usecase.Store(value, setting.id)
                return res.status(statusCode.OK).json({ message: 'CREATED' })
            } catch (error) {
                return next(error)
            }
        }
    }
    public Show() {
        return async (req: any, res: Response, next: NextFunction) => {
            try {
                const setting = await Setting(this.database, req.headers.origin)
                const result = await this.usecase.Show(
                    req.params.slug,
                    setting.id
                )
                return res.json({
                    data: result,
                })
            } catch (error) {
                return next(error)
            }
        }
    }
}

export default Handler
