import { Schema } from 'mongoose'
import slugify from '../../../pkg/slug'
import Mongo from '../mongo'

const schema = new Schema(
    {
        created_by: {
            type: String,
            required: false,
        },
        title: {
            type: String,
            required: true,
            index: true,
            unique: true,
        },
        slug: {
            type: String,
            index: true,
        },
        blocks: {
            type: Array,
            required: true,
        },
        time: Number,
        version: String,
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
    this.slug = slugify(this.title)
    next()
})

export default (database: string) => {
    return Mongo.switch(database).model('pages', schema)
}
