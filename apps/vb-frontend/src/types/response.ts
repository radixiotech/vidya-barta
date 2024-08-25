export enum ApiStatus {
  OK = 'OK',
  ERROR = 'ERROR',
}

export type ApiSuccess<T = unknown> = T & {
  status: ApiStatus.OK;
};

export type ApiError<T = { error: string }> = T & {
  status: ApiStatus.ERROR;
};

export type ApiResponse<
  TSuccessResponse = unknown,
  TErrorResponse = { error: string },
> = ApiSuccess<TSuccessResponse> | ApiError<TErrorResponse>;
