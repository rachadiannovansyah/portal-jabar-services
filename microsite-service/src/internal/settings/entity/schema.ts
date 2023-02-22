import Joi from 'joi'

export const Store = Joi.object({
    organization: Joi.string().required(),
    domain: Joi.string().required(),
    favicon: Joi.string().uri().optional().default(null),
    icon: Joi.string().uri().required(),
    logo: Joi.string().uri().optional().default(null),
    color_palatte: Joi.string().required(),
    name: Joi.string().required(),
    navbar: Joi.object().optional().default(null),
    footer: Joi.object().optional().default(null),
    social_media: Joi.object().optional().default(null),
})
