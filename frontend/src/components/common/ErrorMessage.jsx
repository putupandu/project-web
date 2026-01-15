import React from 'react';
import { AlertCircle } from 'lucide-react';

const ErrorMessage = ({ message, onRetry }) => {
  return (
    <div className="flex flex-col items-center justify-center py-20">
      <AlertCircle className="text-red-500 mb-4" size={48} />
      <h3 className="text-xl font-bold text-gray-800 mb-2">Terjadi Kesalahan</h3>
      <p className="text-gray-600 mb-4">{message}</p>
      {onRetry && (
        <button
          onClick={onRetry}
          className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          Coba Lagi
        </button>
      )}
    </div>
  );
};
//
export default ErrorMessage;// common component improvement 4
// common component improvement 8
