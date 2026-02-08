import type { User } from "@workspace/database"
import { Repository } from "./repository.js";

// TODO: Move to shared types file
interface PaginationRequest {
  skip: number;
  take: number;
}

// TODO: we are just using default offset pagination for now -> eventually, we will want to move to cursor based pagination once we introduce caching
export class ApplicationRepository extends Repository {
  async getApplicationUsers(id: string, pagination?: PaginationRequest): Promise<User[]> {
    const result = await this.db.usersOnApplications.findMany({
      ...pagination,
      where: {
        applicationId: id
      },
      include: {
        user: true
      }
    })
    return result.map((r) => r.user)
  }
}
