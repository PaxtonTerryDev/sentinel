import { Hono } from "hono";

const session = new Hono();

session.get("/", (c) => c.json({ message: "session endpoint" }));

export default session;
