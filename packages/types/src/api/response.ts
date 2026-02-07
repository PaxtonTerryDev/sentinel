import z from "zod";

export const apiStatusSchema = z.enum(["ok", "error"]);
export type APIStatus = z.infer<typeof apiStatusSchema>;

export const errorResponseSchema = z.object({
  status: z.literal("error"),
  error: z.object({
    code: z.string(),
    message: z.string(),
  }),
});

export type ErrorResponse = z.infer<typeof errorResponseSchema>;

export function successResponseSchema<T extends z.ZodTypeAny>(dataSchema: T) {
  return z.object({
    status: z.literal("ok"),
    data: dataSchema,
  });
}

export function apiResponseSchema<T extends z.ZodTypeAny>(dataSchema: T) {
  return z.discriminatedUnion("status", [
    errorResponseSchema,
    successResponseSchema(dataSchema),
  ]);
}

export type APIResponse<T> =
  | ErrorResponse
  | { status: "ok"; data: T };
