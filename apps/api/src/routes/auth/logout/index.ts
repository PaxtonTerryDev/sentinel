import { Hono } from "hono";

const logout = new Hono();

logout.get("/", (c) => c.json({ message: "logout endpoint" }));

export default logout;
