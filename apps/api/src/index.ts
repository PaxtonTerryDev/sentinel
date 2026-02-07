import { Hono } from "hono";
import { etag } from "hono/etag";
import { logger } from "hono/logger";
import { Server } from "./server.js";
import { config } from "./lib/config/config.js";
import routes from "./routes/index.js";

const app = new Hono();

app.use(etag(), logger());

app.get("/", (c) => {
  return c.text("Hello Hono!");
});

app.route("/api", routes);

const server = new Server(app);

server.start(config.port);
