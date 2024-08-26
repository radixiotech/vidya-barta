import { useQuery } from '@tanstack/react-query';
import { Button } from '@vidya-barta/ui';
import { BASE_URL, extractAPIError, fetcher } from './lib/fetcher';

function App() {
  const { error, refetch } = useQuery({
    retry: 0,
    queryKey: ['users'],
    queryFn: ({ signal }) => fetcher(`${BASE_URL}/users`, { signal }),
  });

  const apiError = extractAPIError(error);
  console.log(apiError);

  return (
    <div className="h-screen flex flex-col justify-center items-center">
      <Button onClick={() => refetch()}>Refetch</Button>
    </div>
  );
}

export default App;
