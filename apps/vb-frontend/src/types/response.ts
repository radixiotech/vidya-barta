export enum ApiStatus {
  OK = 'OK',
  ERROR = 'ERROR',
}

/**
 * @description App level error. Should be used within the application.
 */
export type ApplicationError = APIError & {
  statusCode: number;
};

/**
 * @description Concrete error structure coming from API.
 */
export type APIError = {
  message: string;
  fields?: Record<string, string>;
};

/**
 * @description API success response.
 */
export type APISuccessResponse<T = unknown> = T & {
  status: ApiStatus.OK;
};

/**
 * @description API error response.
 */
export type APIErrorResponse<T = { error: string }> = T & {
  status: ApiStatus.ERROR;
};

/**
 * @description Logical app level API response.
 */
export type APIResponse<TSuccessResponse = unknown> =
  | APISuccessResponse<TSuccessResponse>
  | APIErrorResponse<{ error: APIError }>;
