import { HTTPException } from "hono/http-exception";
import type { ContentfulStatusCode } from "hono/utils/http-status";
import type { ErrorResponse } from "@workspace/types/api/response";

type ErrorDefinition = {
  statusCode: ContentfulStatusCode;
  message: string;
};

const ERROR_DEFINITIONS = {
  USER_EXISTS: {
    statusCode: 409,
    message: "A user with this email already exists",
  },
  VALIDATION_ERROR: {
    statusCode: 400,
    message: "Request validation failed",
  },
} as const satisfies Record<string, ErrorDefinition>;

export type ErrorType = keyof typeof ERROR_DEFINITIONS;

export function errorResponse(
  errorType: ErrorType,
  messageOverride?: string
): HTTPException {
  const { statusCode, message } = ERROR_DEFINITIONS[errorType];

  const body: ErrorResponse = {
    status: "error",
    error: {
      code: errorType,
      message: messageOverride ?? message,
    },
  };

  return new HTTPException(statusCode, {
    res: new Response(JSON.stringify(body), {
      status: statusCode,
      headers: { "Content-Type": "application/json" },
    }),
  });
}
