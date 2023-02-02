import Joi from 'joi'

export const Store = Joi.object({
    version: Joi.string().optional(),
    time: Joi.number().optional(),
    blocks: Joi.array().required().min(1),
    title: Joi.string().required(),
})
