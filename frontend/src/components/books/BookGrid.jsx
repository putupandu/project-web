import React from 'react';
import BookCard from './BookCard';

const BookGrid = ({ books }) => {
  if (!books || books.length === 0) {
    return (
      <div className="text-center py-20">
        <Book size={64} className="mx-auto text-gray-400 mb-4" />
        <h3 className="text-2xl font-bold text-gray-600 mb-2">
          Tidak ada buku ditemukan
        </h3>
        <p className="text-gray-500">
          Coba ubah filter atau kata kunci pencarian Anda
        </p>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      {books.map((book) => (
        <BookCard key={book.id} book={book} />
      ))}
    </div>
  );
};

export default BookGrid;