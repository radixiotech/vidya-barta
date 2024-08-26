import {
  APIError,
  APIErrorResponse,
  APIResponse,
  ApiStatus,
  ApplicationError,
} from '@/types/response';
import { AxiosError, AxiosRequestConfig, isAxiosError } from 'axios';
import { api } from './http';

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

export async function fetcher<TData = unknown, TBody = unknown>(
  config: AxiosRequestConfig<TBody>
): Promise<TData> {
  try {
    const response = await api<APIResponse<{ data: TData }>>(config);
    const data = response.data;

    if (data.status === ApiStatus.OK) {
      return data.data;
    }

    throw new AppError({ ...data.error, statusCode: response.status });
  } catch (error) {
    if (isAxiosError(error)) {
      const e = error as AxiosError<APIErrorResponse<{ error: APIError }>>;

      const statusCode = e.response?.status || 500;
      const fields = e.response?.data?.error.fields || {};
      const message =
        e.response?.data?.error.message || 'Internal Server Error';

      throw new AppError({ message, statusCode, fields });
    }

    throw new AppError({ message: 'Internal Server Error', statusCode: 500 });
  }
}
