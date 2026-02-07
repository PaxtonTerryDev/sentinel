import type { User } from "@workspace/database";
import type { EmailCreateUserRequest, OauthCreateUserRequest } from "@workspace/types/api/user";
import { Password } from "../password";
import { Database } from "../repositories/database";

const db = new Database();

export async function createUser(request: EmailCreateUserRequest | OauthCreateUserRequest): Promise<User> {
  switch (request.type) {
    case "EMAIL":
      return createEmailUser(request);
    case "OAUTH":
      return createOAuthUser(request);
  }
}

async function createEmailUser(request: EmailCreateUserRequest): Promise<User> {
  const hashedPassword = await Password.hash(request.password);
  return db.user.createEmailUser(request.email, hashedPassword);
}

async function createOAuthUser(request: OauthCreateUserRequest): Promise<User> {
  // TODO: Implement OAuth user creation
  throw new Error("OAuth user creation not implemented");
}
