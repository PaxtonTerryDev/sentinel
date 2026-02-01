import { Hono } from 'hono'
import { etag } from 'hono/etag'
import { logger } from 'hono/logger'
import { Server } from './server.js'
import { config } from './lib/config/config.js'

const app = new Hono()

app.use(
  etag(),
  logger()
)

app.get('/', (c) => {
  return c.text('Hello Hono!')
})

const server = new Server(app)
console.dir(config, { depth: null}) 
server.start(config.port);
