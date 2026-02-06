import type { PrismaClient } from "@workspace/database";
import { createDatabaseConnection } from "../utils/database/create-database-connection";

export class Database {
  private static conn: PrismaClient | undefined;
  protected db: PrismaClient;

  constructor() {
    if (!Database.conn) {
      Database.conn = createDatabaseConnection();
    }
    this.db = Database.conn;
  }
}
