import { serve, type ServerType } from "@hono/node-server"
import type { Hono } from "hono"
import { log } from "@ptdlib/quicklog";

export class Server {
  private app: Hono;
  private server: ServerType | undefined;

  constructor(app: Hono) {
    this.app = app
  }

  start(port: number) {
    this.server = serve({
      fetch: this.app.fetch,
      port
    }, (_info) => {
      log.info(`[Server] Starting on port ${port}`)
      this.subscribeSignalHandlers();
    })
  }

  private subscribeSignalHandlers(): void {
    process.on('SIGINT', () => {
      log.warn(`[Server] Interrupt signal received`)
      this.runShutdown();
    })

    process.on("SIGTERM", () => {
      log.error(`[Server] Terminate signal received. Exiting`)
      this.server?.close((err) => {
        if (err) {
          log.error(String(err));
          process.exit(1);
        }
        process.exit(0)
      })
    })
  }

  private runShutdown() {
    log.warn("[Server] Running shutdown...");
    this.server?.close();
    process.exit(0);
  }
}
