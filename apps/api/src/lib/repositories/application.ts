import { PrismaClient } from "../../../generated/prisma/client.js";
import { Repository } from "./repository.js";

// TODO: Move to shared types file
interface PaginationRequest {
  skip: number;
  take: number;
}
export class ApplicationRepository extends Repository {
  // TODO: we are just using default offset pagination for now -> eventually, we will want to move to cursor based pagination once we introduce caching
  async getApplicationUsers(id: string, pagination?: PaginationRequest): Promise<User[]> {
    return this.db.application.findMany({
      { pagination && }
  include: {
    users: true,
  },
  where: {
    id
  }
})
  }
}
