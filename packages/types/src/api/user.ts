import z from "zod";
import { apiResponseSchema } from "./response.js";

export const oauthProviderSchema = z.enum(["GOOGLE", "AZURE", "GITHUB"]);
export type OAuthProvider = z.infer<typeof oauthProviderSchema>;

export const emailCreateUserRequestSchema = z.object({
  type: z.literal("EMAIL"),
  email: z.string().email(),
  password: z.string(),
});

export type EmailCreateUserRequest = z.infer<
  typeof emailCreateUserRequestSchema
>;

export const oauthCreateUserRequestSchema = z.object({
  type: z.literal("OAUTH"),
  provider: oauthProviderSchema,
  // FIX: This is some placeholder code from the robot. We will need to update this accordingly depending on the provider -> which will probably entail extending this type to be discriminated between each provider
  accessToken: z.string(),
});

export type OAuthCreateUserRequest = z.infer<
  typeof oauthCreateUserRequestSchema
>;

export type CreateUserRequest = z.infer<typeof createUserRequestSchema>;

export const createUserResponseSchema = apiResponseSchema(
  z.object({ id: z.string() }),
);

export type CreateUserResponse = z.infer<typeof createUserResponseSchema>;

export const createUserRequestSchema = z.discriminatedUnion("type", [
  emailCreateUserRequestSchema,
  oauthCreateUserRequestSchema,
]);

// LOGIN

export const emailLoginRequestSchema = z.object({
  type: z.literal("EMAIL"),
  email: z.string().email(),
  password: z.string(),
});

export type EmailLoginRequest = z.infer<typeof emailLoginRequestSchema>;

export const oAuthLoginRequestSchema = z.object({
  type: z.literal("OAUTH"),
  provider: oauthProviderSchema,
  // FIX: This is some placeholder code from the robot. We will need to update this accordingly depending on the provider -> which will probably entail extending this type to be discriminated between each provider
  accessToken: z.string(),
});

export type OAuthLoginRequest = z.infer<typeof oAuthLoginRequestSchema>;

export const loginRequestSchema = z.discriminatedUnion("type", [
  emailCreateUserRequestSchema,
  oauthCreateUserRequestSchema,
]);

export type LoginRequest = z.infer<typeof loginRequestSchema>;

// FIX: Eventually this will need to return a valid session
export const loginResponseSchema = apiResponseSchema(
  z.object({ id: z.string() }),
);
export type LoginResponse = z.infer<typeof loginResponseSchema>;
