import { Hono } from "hono";

const login = new Hono();

login.get("/", (c) => c.json({ message: "login endpoint" }));

export default login;
