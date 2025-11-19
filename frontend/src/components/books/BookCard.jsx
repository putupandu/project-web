import React from 'react';
import { Link } from 'react-router-dom';
import { Book, User, Calendar, Eye, Download } from 'lucide-react';

const BookCard = ({ book }) => {
  return (
    <Link
      to={`/books/${book.id}`}
      className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 group"
    >
      {/* Cover Image */}
      <div className="relative h-64 bg-gradient-to-br from-blue-400 to-purple-500 overflow-hidden">
        {book.cover ? (
          <img
            src={book.cover}
            alt={book.title}
            className="w-full h-full object-cover group-hover:scale-110 transition-transform duration-300"
          />
        ) : (
          <div className="flex items-center justify-center h-full">
            <Book size={80} className="text-white opacity-50" />
          </div>
        )}
        
        {/* Year Badge */}
        <div className="absolute top-2 right-2 bg-blue-600 text-white px-3 py-1 rounded-full text-xs font-semibold">
          {book.year}
        </div>

        {/* Category Badge */}
        {book.category && (
          <div className="absolute top-2 left-2 bg-white/90 backdrop-blur-sm text-gray-800 px-3 py-1 rounded-full text-xs font-semibold flex items-center">
            <span>{book.category.icon}</span>
            <span className="ml-1">{book.category.name}</span>
          </div>
        )}
      </div>

      {/* Content */}
      <div className="p-4">
        <h3 className="font-bold text-lg mb-2 text-gray-800 line-clamp-2 h-14 group-hover:text-blue-600 transition">
          {book.title}
        </h3>

        <div className="flex items-center text-sm text-gray-600 mb-2">
          <User size={14} className="mr-1 flex-shrink-0" />
          <span className="line-clamp-1">{book.author}</span>
        </div>

        {book.publisher && (
          <div className="flex items-center text-sm text-gray-600 mb-3">
            <Book size={14} className="mr-1 flex-shrink-0" />
            <span className="line-clamp-1">{book.publisher}</span>
          </div>
        )}

        {/* Stats */}
        <div className="flex justify-between items-center text-xs text-gray-500 pt-3 border-t">
          <div className="flex items-center">
            <Eye size={14} className="mr-1" />
            <span>{book.views || 0}</span>
          </div>
          <div className="flex items-center">
            <Download size={14} className="mr-1" />
            <span>{book.downloads || 0}</span>
          </div>
          {book.pages && (
            <div className="flex items-center">
              <Book size={14} className="mr-1" />
              <span>{book.pages} hal</span>
            </div>
          )}
        </div>
      </div>
    </Link>
  );
};

export default BookCard;