import type { PrismaClient } from "@workspace/database";

export class Repository {
  protected db: PrismaClient;

  constructor(db: PrismaClient) {
    this.db = db;
  }
}
