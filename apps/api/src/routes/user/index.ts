import { validateRequest } from "@/src/lib/response/validate-request";
import { createUser } from "@/src/lib/user-creation";
import { createUserRequestSchema, type CreateUserResponse } from "@workspace/types/api/user";
import { Hono } from "hono";

const user = new Hono();

user.post("/", async (c) => {
  const body = await c.req.json();
  const request = validateRequest(createUserRequestSchema, body);
  const newUser = await createUser(request);

  const response: CreateUserResponse = {
    status: "ok",
    data: { id: newUser.id },
  };

  return c.json(response);
});

export default user;
