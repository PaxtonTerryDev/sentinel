import type { PrismaClient } from "@workspace/database";
import { createDatabaseConnection } from "@workspace/database/connection";
import { CredentialRepository } from "./credential";
import { UserRepository } from "./user";

export class Database {
  private static conn: PrismaClient | undefined;
  protected db: PrismaClient;

  constructor() {
    if (!Database.conn) {
      Database.conn = createDatabaseConnection();
    }
    this.db = Database.conn;
    this.credential = new CredentialRepository(this.db);
    this.user = new UserRepository(this.db);
  }

  credential: CredentialRepository
  user: UserRepository
}
