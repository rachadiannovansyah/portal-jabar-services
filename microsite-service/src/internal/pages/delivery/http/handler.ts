import { NextFunction, Request, Response } from 'express'
import winston from 'winston'
import Usecase from '../../usecase/usecase'
import { Store } from '../../entity/schema'
import { validateFormRequest } from '../../../../helpers/validate'
import statusCode from '../../../../pkg/statusCode'

class Handler {
    constructor(private usecase: Usecase, private logger: winston.Logger) {}
    public store() {
        return async (req: Request, res: Response, next: NextFunction) => {
            try {
                const value = validateFormRequest(Store, req.body)

                await this.usecase.store(value)
                return res.status(statusCode.OK).json({ message: 'CREATED' })
            } catch (error) {
                return next(error)
            }
        }
    }
    public show() {
        return async (req: Request, res: Response, next: NextFunction) => {
            try {
                const result = await this.usecase.show(req.params.id)
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
