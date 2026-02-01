import { PrismaClient } from "@/generated/prisma/client";
import { createDatabaseConnection } from "../utils/database/create-database-connection";

export class Database {
  private static conn: PrismaClient | undefined;
  private db: PrismaClient;

  constructor() {
    if (!Database.conn) {
      Database.conn = createDatabaseConnection();
    }
    this.db = Database.conn;
  }
}
