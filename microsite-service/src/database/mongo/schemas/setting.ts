import { Schema } from 'mongoose'
import { RemoveProcotol } from '../../../helpers/http'
import Mongo from '../mongo'

const schema = new Schema(
    {
        created_by: {
            type: String,
            required: false,
            index: true,
        },
        name: {
            type: String,
            required: true,
            index: true,
        },
        published_by: {
            type: String,
            required: false,
            index: true,
        },
        favicon: {
            type: String,
            default: null,
        },
        color_palatte: {
            type: String,
            required: true,
        },
        domain: {
            type: String,
            required: true,
            index: true,
        },
        organization: {
            type: String,
            required: true,
            index: true,
        },
        published_at: Date,
        status: {
            type: String,
            index: true,
        },
        is_active: {
            type: Boolean,
            index: true,
        },
        icon: {
            type: String,
            required: true,
        },
        navbar: {
            type: Object,
            default: null,
        },
        footer: {
            type: Object,
            default: null,
        },
        social_media: {
            type: Object,
            default: null,
        },
    },
    {
        timestamps: {
            createdAt: 'created_at',
            updatedAt: 'updated_at',
        },
        versionKey: false,
    }
)

schema.pre('save', function (next) {
    this.domain = RemoveProcotol(this.domain)
    next()
})

export default (database: string) => {
    return Mongo.Model(database, 'settings', schema)
}
