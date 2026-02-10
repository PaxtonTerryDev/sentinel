import { Repository } from "./repository.js"

export class CredentialRepository extends Repository {
  // FIX: Should probably throw a "USER_NOT_FOUND" error or somthing if this fails
  async getByEmail(email: string) {
    return this.db.credential.findUniqueOrThrow({
      where: { email },
      include: {
        userCredential: {
          include: {
            user: true,
          },
        },
      },
    });
  }
}
