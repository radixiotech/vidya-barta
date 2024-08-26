import { useQuery } from '@tanstack/react-query';
import { Button } from '@vidya-barta/ui';
import { extractAPIError, fetcher } from './lib/fetcher';

function App() {
  const { error, refetch, data } = useQuery({
    retry: 0,
    queryKey: ['users'],
    queryFn: ({ signal }) => fetcher({ signal, url: '/users' }),
  });

  const apiError = extractAPIError(error);

  return (
    <div className="h-screen flex flex-col justify-center items-center">
      <pre>{!error && JSON.stringify(data, null, 2)}</pre>
      <pre>{apiError && JSON.stringify(apiError, null, 2)}</pre>
      <Button onClick={() => refetch()}>Refetch</Button>
    </div>
  );
}

export default App;
