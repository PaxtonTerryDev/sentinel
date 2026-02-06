import z from "zod";

export const apiStatusSchema = z.enum(["OK", "ERROR"]);
export type APIStatus = z.infer<typeof apiStatusSchema>; 
