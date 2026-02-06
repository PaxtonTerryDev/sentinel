import type { ZodType } from "zod";
import { errorResponse } from "./error-response";

export function validateRequest<T>(schema: ZodType<T>, body: unknown): T {
  const result = schema.safeParse(body);
  if (!result.success) {
    throw errorResponse("VALIDATION_ERROR", result.error.message);
  }
  return result.data;
}
