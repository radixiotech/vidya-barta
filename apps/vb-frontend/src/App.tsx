import {
  Button,
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@vidya-barta/ui';
import { useState } from 'react';

function App() {
  const [open, setOpen] = useState(false);

  return (
    <div className="h-screen flex flex-col justify-center items-center">
      <Dialog open={open} defaultOpen={open} onOpenChange={setOpen}>
        <DialogTrigger>
          <Button onClick={() => setOpen(true)}>Open Dialog</Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Are you absolutely sure?</DialogTitle>
            <DialogDescription>
              This action cannot be undone. This will permanently delete your
              account and remove your data from our servers.
            </DialogDescription>
          </DialogHeader>
        </DialogContent>
      </Dialog>
    </div>
  );
}

export default App;
