import { ApiResponse, ApiStatus } from '@/types/response';

export class AppError<T> {
  name = 'ApplicationError';
  constructor(public error: T) {}
}

export async function fetcher<TData = unknown>(
  input: string | URL | globalThis.Request,
  init?: RequestInit
): Promise<TData> {
  const response = await fetch(input, init);
  const data = (await response.json()) as ApiResponse<
    { data: TData; message: string },
    { error: unknown }
  >;

  if (data.status === ApiStatus.OK) {
    return data.data;
  }

  throw new AppError(data.error);
}
