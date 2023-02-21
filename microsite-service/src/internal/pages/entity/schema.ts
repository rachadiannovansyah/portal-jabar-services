import Joi from 'joi'

export const Store = Joi.object({
    created_by: Joi.string().default('').forbidden(),
    is_active: Joi.boolean().default(false).forbidden(),
    slug: Joi.string().default('').forbidden(),
    sections: Joi.object().required(),
    title: Joi.string().required(),
    banner: Joi.string().optional(),
})
