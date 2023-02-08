import { Schema } from 'mongoose'
import Mongo from '../mongo'

const schema = new Schema(
    {
        created_by: {
            type: String,
            required: false,
            index: true,
        },
        published_by: {
            type: String,
            required: false,
            index: true,
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
    },
    {
        timestamps: {
            createdAt: 'created_at',
            updatedAt: 'updated_at',
        },
        versionKey: false,
    }
)

export default (database: string) => {
    return Mongo.model(database, 'settings', schema)
}
