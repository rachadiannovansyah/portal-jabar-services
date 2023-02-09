import Joi from 'joi'

export const Store = Joi.object({
    sections: Joi.array().required().min(1),
    title: Joi.string().required(),
    banner: Joi.string().optional(),
})
