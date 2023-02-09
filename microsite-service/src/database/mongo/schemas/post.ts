import { Schema } from 'mongoose'
import Mongo from '../mongo'

const schema = new Schema(
    {
        content: {
            type: String,
            required: true,
        },
        except: String,
        title: {
            type: String,
            required: true,
            index: true,
            unique: true,
        },
        author: {
            type: String,
            required: true,
        },
        slug: {
            type: String,
            index: true,
        },
        published_at: Date,
        views: {
            type: Number,
            default: 0,
            index: true,
        },
        shared: {
            type: Number,
            index: true,
            default: 0,
        },
        category: {
            type: String,
            required: true,
            index: true,
        },
        image: {
            type: String,
            required: true,
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
    return Mongo.Model(database, 'posts', schema)
}
