import { z } from "zod";

export const UserSchema = z.object({
  id: z.number().optional(),
  username: z.string(),
  email: z.string(),
  password: z.string(),
  role: z.string(),
  department: z.string(),
  profileImage: z.string().optional(),
  isActive: z.boolean(),
  position: z.string(),
  updatedAt: z.string().optional(),
  createdAt: z.string().optional(),
});

export type User = z.infer<typeof UserSchema>;
