import { QueryClientProvider } from '@tanstack/react-query';
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import './index.css';

import App from './App.tsx';
import { queryClient } from './lib/query-client.ts';
import RevalidateTokenProvider from './providers/revalidate-token-provider.tsx';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <RevalidateTokenProvider>
        <App />
      </RevalidateTokenProvider>
    </QueryClientProvider>
  </StrictMode>
);
