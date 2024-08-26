import axios from 'axios';
import { env } from './config';
import {
  ACCESS_TOKEN,
  getDataFromLocalStorage,
  TokenStorage,
} from './token-storage';

export const BASE_URL = env.VB_BACKEND_URL + '/api/v1';

export const api = axios.create({
  timeout: 10000,
  baseURL: BASE_URL,
  headers: { 'Content-Type': 'application/json', Accept: 'application/json' },
});

api.interceptors.request.use(
  (config) => {
    const token =
      TokenStorage.accessToken || getDataFromLocalStorage(ACCESS_TOKEN) || '';

    if (token) config.headers.set('Authorization', `Bearer ${token}`);

    return config;
  },
  (e) => Promise.reject(e)
);
