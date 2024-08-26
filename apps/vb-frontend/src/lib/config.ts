import { Env } from '@/types/env';
import { z, ZodError } from 'zod';

type EnvSchema = z.infer<typeof envSchema>;

export const envSchema = z.object({
  VITE_VB_BACKEND_URL: z
    .string({
      required_error: 'VB_BACKEND_URL is required.',
      invalid_type_error: 'VB_BACKEND_URL must be a string.',
    })
    .url('VB_BACKEND_URL must be a URL.')
    .min(1, { message: 'VB_BACKEND_URL is required.' }),
});

const parseEnv = () => {
  try {
    return envSchema.parse(import.meta.env);
  } catch (error) {
    if (error instanceof ZodError) throw error.flatten();
    else throw error;
  }
};

let parsedEnv: EnvSchema | undefined = undefined;
if (!parsedEnv) parsedEnv = (parseEnv() || {}) as EnvSchema;

export const env: Env<EnvSchema> = {
  VB_BACKEND_URL: parsedEnv.VITE_VB_BACKEND_URL,
};
