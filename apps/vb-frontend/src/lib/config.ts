import { Env } from '@/types/env';
import { z } from 'zod';

export const envSchema = z.object({
  VITE_VB_BACKEND_URL: z
    .string()
    .min(1, { message: 'VB_BACKEND_URL is required.' }),
});

const parsedEnv = envSchema.parse(import.meta.env);

export const env: Env<EnvSchema> = {
  VB_BACKEND_URL: parsedEnv.VITE_VB_BACKEND_URL,
};

export type EnvSchema = z.infer<typeof envSchema>;
