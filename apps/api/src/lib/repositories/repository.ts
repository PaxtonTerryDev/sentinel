import { PrismaClient } from "../../../generated/prisma/client.js";

export class Repository {
  protected db: PrismaClient;

  constructor(db: PrismaClient) {
    this.db = db;
  }
}
