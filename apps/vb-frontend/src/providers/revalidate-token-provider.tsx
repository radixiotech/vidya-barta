import { api } from '@/lib/http';
import { TokenStorage } from '@/lib/token-storage';
import { AxiosError } from 'axios';
import { PropsWithChildren, useEffect } from 'react';

const RevalidateTokenProvider = ({ children }: PropsWithChildren) => {
  const handleLogout = () => {
    localStorage.clear();
    history.pushState('', '', '/login');
  };

  const handleRevalidateToken = async (): Promise<string | null> => {
    return '';
  };

  useEffect(() => {
    api.interceptors.response.use(
      (response) => response,
      async (error: AxiosError) => {
        if (
          error.code === 'ERR_NETWORK' ||
          (error.response && error.response.status !== 401)
        ) {
          return Promise.reject(error);
        }

        const token = await handleRevalidateToken();
        if (!token) {
          handleLogout();
          return Promise.reject(error);
        }

        const ogRequest = error.config;
        TokenStorage.accessToken = token;
        ogRequest?.headers.set('Authorization', `Bearer ${token}`);
        api.defaults.headers.common.Authorization = `Bearer ${token}`;

        return api({ ...ogRequest });
      }
    );
  }, []);

  return children;
};

export default RevalidateTokenProvider;
