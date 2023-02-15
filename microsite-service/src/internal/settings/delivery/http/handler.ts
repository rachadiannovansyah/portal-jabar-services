import { NextFunction, Request, Response } from 'express'
import winston from 'winston'
import Usecase from '../../usecase/usecase'
import { Store } from '../../entity/schema'
import { validateFormRequest } from '../../../../helpers/validate'
import statusCode from '../../../../pkg/statusCode'

class Handler {
    constructor(private usecase: Usecase, private logger: winston.Logger) {}
    public Store() {
        return async (req: Request, res: Response, next: NextFunction) => {
            try {
                const value = validateFormRequest(Store, req.body)

                await this.usecase.Store(value)
                return res.status(statusCode.OK).json({ message: 'CREATED' })
            } catch (error) {
                return next(error)
            }
        }
    }
    public Show() {
        return async (req: Request, res: Response, next: NextFunction) => {
            try {
                const result = await this.usecase.Show(req.params.idSetting)
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
