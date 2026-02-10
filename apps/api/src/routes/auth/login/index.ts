import { validateRequest } from "@/src/lib/response/validate-request";
import { userLogin } from "@/src/lib/user-login";
import {
  loginRequestSchema,
  type CreateUserResponse,
} from "@workspace/types/api/user";
import { Hono } from "hono";

const login = new Hono();

login.post("/", async (c) => {
  const body = await c.req.json();
  const request = validateRequest(loginRequestSchema, body);
  const verify = await userLogin(request);
  // FIX: Still need to return the proper response type, and add in logical error handling
  const response: CreateUserResponse = {
    status: "ok",
    data: { id: verify.id },
  };
  return c.json(response);
});

export default login;
