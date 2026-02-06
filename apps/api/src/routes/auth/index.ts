import { Hono } from "hono";
import login from "./login/index.js";
import logout from "./logout/index.js";
import session from "./session/index.js";

const auth = new Hono();

auth.route("/login", login);
auth.route("/logout", logout);
auth.route("/session", session);

export default auth;
