import type { User } from "@workspace/database";
import type {
  EmailLoginRequest,
  LoginRequest,
  OAuthLoginRequest,
} from "@workspace/types/api/user";
import { Database } from "../repositories/database";
import { Password } from "../password";

// FIX: Should return an active session
export async function userLogin(request: LoginRequest): Promise<User> {
  switch (request.type) {
    case "EMAIL":
      return emailUserLogin(request);
    case "OAUTH":
      return oauthUserLogin(request);
    default:
      throw new Error(`unsupported login type: ${request.type}`);
  }
}

async function emailUserLogin(request: EmailLoginRequest): Promise<User> {
  const { email, password } = request;
  const db = new Database();
  const { id, passwordHash } = await db.credential.getByEmail(email);
  const passValid = await Password.verify(password, passwordHash);

  if (passValid) {
    return { id };
  }

  throw new Error("INVALID_PASSWORD");
}

async function oauthUserLogin(requst: OAuthLoginRequest): Promise<User> {
  throw new Error("Not implemented");
}
