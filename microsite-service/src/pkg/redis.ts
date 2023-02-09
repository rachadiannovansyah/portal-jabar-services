import { createClient, RedisClientType } from 'redis'
import winston from 'winston'
import { Config } from '../config/config.interface'

class Redis {
    private client: RedisClientType
    constructor({ redis }: Config, logger: winston.Logger) {
        this.client = createClient({
            url: `redis://${redis.host}:${redis.port}`,
        })
        this.client.connect().then(() => {
            logger.info('redis connected')
        })
    }

    public async Store(key: string, value: any, ttl: number) {
        return this.client.set(key, value, {
            EX: ttl,
        })
    }

    public async Get(key: string) {
        return this.client.get(key)
    }
}

export default Redis
