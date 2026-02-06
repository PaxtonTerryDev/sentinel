import type { PrismaClient, User } from "@/generated/prisma/client";

export class UserRepository {
  private db: PrismaClient;

  constructor(db: PrismaClient) {
    this.db = db;
  }

  async create(): Promise<User> {
    return this.db.user.create({})
  }

  async get(id: string): Promise<User> {
    return this.db.user.findUniqueOrThrow({
      where: {
        id
      }
    })
  }
}
