import { User } from "../../../generated/prisma/browser.js"
import { Repository } from "./repository.js"

export class UserRepository extends Repository {
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
