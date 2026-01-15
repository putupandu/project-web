// src/components/books/BookCard.jsx
import React from 'react';
import { Link } from 'react-router-dom';
import { Book as BookIcon, User, Eye, Download } from 'lucide-react';

const BookCard = ({ book }) => {
  return (
    <Link
      to={`/books/${book.id}`}
      className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 group"
    >
      {/* Cover Image */}
      <div className="relative h-64 bg-gray-200 overflow-hidden">
        {book.cover ? (
  <img
    src={book.cover} // backend sudah mengirim full URL
    alt={book.title}
    className="w-full h-full object-cover group-hover:scale-110 transition-transform duration-300"
    onError={(e) => {
      e.target.onerror = null;
      e.target.src = '/no-cover.png';
    }}
  />
) : (
  <div className="flex items-center justify-center h-full text-gray-500">
    No cover
  </div>
)}
//

        
        {/* Year Badge */}
        <div className="absolute top-2 right-2 bg-blue-600 text-white px-2 py-1 rounded-full text-xs font-semibold">
          {book.year}
        </div>
      </div>

      {/* Content */}
      <div className="p-3">
        <h3 className="font-bold text-sm mb-1 text-gray-800 line-clamp-2 h-10">
          {book.title}
        </h3>
        <div className="flex items-center text-xs text-gray-600 mb-1">
          <User size={10} className="mr-1" />
          <span className="line-clamp-1">{book.author}</span>
        </div>
        <div className="flex justify-between items-center text-xs text-gray-500 mt-2">
          <span><Eye size={12} className="mr-1 inline" /> {book.views || 0}</span>
          <span><Download size={12} className="mr-1 inline" /> {book.downloads || 0}</span>
        </div>
      </div>
    </Link>
  );
};

export default BookCard;// books tweak 2
// books tweak 5
// books tweak 8
// books tweak 11
