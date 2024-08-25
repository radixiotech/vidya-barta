import { useQuery } from '@tanstack/react-query';
import {
  Button,
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@vidya-barta/ui';
import { useState } from 'react';
import { fetcher } from './lib/fetcher';

function App() {
  const [open, setOpen] = useState(false);
  const { error, data, isFetching } = useQuery({
    retry: 0,
    queryKey: ['todo'],
    queryFn: ({ signal }) =>
      fetcher('https://jsonplaceholder.typicode.com/todos/1', {
        signal,
        method: 'GET',
      }),
  });

  return (
    <div className="h-screen flex flex-col justify-center items-center">
      {isFetching ? (
        <h1>Loading...</h1>
      ) : (
        <>
          <pre>
            {JSON.stringify(data, null, 2)}
            <br />
            {JSON.stringify(error, null, 2)}
          </pre>
          <Dialog open={open} defaultOpen={open} onOpenChange={setOpen}>
            <Button onClick={() => setOpen(true)}>Open Dialog</Button>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Are you absolutely sure?</DialogTitle>
                <DialogDescription>
                  This action cannot be undone. This will permanently delete
                  your account and remove your data from our servers.
                </DialogDescription>
              </DialogHeader>
            </DialogContent>
          </Dialog>
        </>
      )}
    </div>
  );
}

export default App;
