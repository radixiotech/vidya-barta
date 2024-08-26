export const ACCESS_TOKEN = 'accessToken';
export const REFRESH_TOKEN = 'refreshToken';

export const getDataFromLocalStorage = <T>(
  key: string,
  parse?: boolean
): T | null => {
  try {
    const value = localStorage.getItem(key);
    return value ? (parse ? JSON.parse(value) : value) : null;
  } catch (e) {
    console.log(e);
    return null;
  }
};

export class TokenStorage {
  private static aToken: string = '';
  private static rToken: string = '';

  private constructor() {}

  static get accessToken() {
    return TokenStorage.aToken;
  }

  static get refreshToken() {
    return TokenStorage.rToken;
  }

  static set accessToken(token: string) {
    TokenStorage.aToken = token;
  }

  static set refreshToken(token: string) {
    TokenStorage.rToken = token;
  }
}
