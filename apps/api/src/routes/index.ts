import { Hono } from "hono";
import app from "./app/index.js";
import auth from "./auth/index.js";
import user from "./user/index.js";

const routes = new Hono();

routes.route("/app", app);
routes.route("/auth", auth);
routes.route("/user", user);

export default routes;
