import { Hono } from 'hono'
import { etag } from 'hono/etag'
import { logger } from 'hono/logger'
import { Server } from './server.js'

const app = new Hono()

app.use(
  etag(),
  logger()
)

app.get('/', (c) => {
  return c.text('Hello Hono!')
})

const server = new Server(app)

server.start(3000);
