import Joi from 'joi'

export const Store = Joi.object({
    organization: Joi.string().required(),
    domain: Joi.string().required(),
    favicon: Joi.string().uri().optional(),
    icon: Joi.string().uri().required(),
    logo: Joi.string().uri().optional(),
    color_pallate: Joi.string().required(),
    name: Joi.string().required(),
})
