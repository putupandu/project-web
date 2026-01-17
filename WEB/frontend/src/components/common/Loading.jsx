import React from 'react';
import { Loader2 } from 'lucide-react';

const Loading = ({ message = 'Loading...' }) => {
  return (
    <div className="flex flex-col items-center justify-center py-20">
      <Loader2 className="animate-spin text-blue-600 mb-4" size={48} />
      <p className="text-gray-600 text-lg">{message}</p>
    </div>
  );
};
//
export default Loading;// common component improvement 3
// common component improvement 7
// common component improvement 11
// common component improvement 15
