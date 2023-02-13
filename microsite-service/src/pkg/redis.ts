import * as redis from 'redis'
import winston from 'winston'
import { Config } from '../config/config.interface'

class Redis {
    private client: redis.RedisClientType
    constructor(config: Config, logger: winston.Logger) {
        this.client = redis.createClient({
            url: `redis://${config.redis.host}:${config.redis.port}`,
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
