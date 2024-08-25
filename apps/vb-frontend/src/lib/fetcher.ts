import {
  APIError,
  APIResponse,
  ApiStatus,
  ApplicationError,
} from '@/types/response';

class AppError extends Error {
  statusCode: number;
  fields: APIError['fields'];
  override name = 'ApplicationError';

  constructor({ message, statusCode, fields }: ApplicationError) {
    super(message);
    this.fields = fields;
    this.statusCode = statusCode;
  }
}

export const extractAPIError = (error: unknown): ApplicationError | null => {
  if (!error) return null;

  if (error instanceof AppError) {
    return {
      fields: error.fields,
      message: error.message,
      statusCode: error.statusCode,
    };
  }

  return { message: 'Internal Server Error', statusCode: 500 };
};

export async function fetcher<TData = unknown>(
  input: string | URL | globalThis.Request,
  init?: RequestInit
): Promise<TData> {
  const response = await fetch(input, init);
  const data = (await response.json()) as APIResponse<{ data: TData }>;

  if (data.status === ApiStatus.OK) {
    return data.data;
  }

  throw new AppError({ ...data.error, statusCode: response.status });
}
