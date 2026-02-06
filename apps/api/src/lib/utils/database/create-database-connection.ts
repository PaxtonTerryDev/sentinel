import "dotenv/config";
import { PrismaPg } from "@prisma/adapter-pg";
import { PrismaClient } from "@workspace/database"

export function createDatabaseConnection(): PrismaClient {
  const connectionString = `${process.env.DATABASE_URL}`;

  const adapter = new PrismaPg({ connectionString });
  return new PrismaClient({ adapter });
}
