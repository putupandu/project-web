import React, { useState, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { Search } from 'lucide-react';
import { searchService } from '../services/searchService';
import BookGrid from '../components/books/BookGrid';
import Pagination from '../components/common/Pagination';
import Loading from '../components/common/Loading';
import ErrorMessage from '../components/common/ErrorMessage';

const SearchResults = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const query = searchParams.get('q') || '';
  
  const [books, setBooks] = useState([]);
  const [meta, setMeta] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);

  useEffect(() => {
    if (query) {
      performSearch(currentPage);
    }
  }, [query, currentPage]);

  const performSearch = async (page = 1) => {
    try {
      setLoading(true);
      setError(null);
      
      const response = await searchService.search(query, {
        page,
        per_page: 12,
      });
      
      setBooks(response.data || []);
      setMeta(response.meta || null);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handlePageChange = (page) => {
    setCurrentPage(page);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  if (!query) {
    return (
      <div className="container mx-auto px-4 py-20 text-center">
        <Search className="mx-auto mb-4 text-gray-400" size={64} />
        <h2 className="text-2xl font-bold text-gray-800 mb-2">
          Silakan masukkan kata kunci pencarian
        </h2>
        <p className="text-gray-600">
          Gunakan kolom pencarian di atas untuk mencari buku
        </p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container mx-auto px-4 py-8">
        <ErrorMessage message={error} onRetry={() => performSearch(currentPage)} />
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-800 mb-2">
          Hasil Pencarian
        </h1>
        <p className="text-gray-600">
          Menampilkan hasil untuk: <span className="font-semibold">"{query}"</span>
        </p>
        {meta && (
          <p className="text-gray-500 mt-1">
            Ditemukan {meta.total} buku
          </p>
        )}
      </div>

      {/* Content */}
      {loading ? (
        <Loading message="Mencari buku..." />
      ) : books.length === 0 ? (
        <div className="text-center py-20">
          <Search className="mx-auto mb-4 text-gray-400" size={64} />
          <h3 className="text-2xl font-bold text-gray-600 mb-2">
            Tidak ada hasil ditemukan
          </h3>
          <p className="text-gray-500 mb-4">
            Coba gunakan kata kunci yang berbeda atau lebih umum
          </p>
        </div>
      ) : (
        <>
          <BookGrid books={books} />
          
          {/* Pagination */}
          {meta && meta.total_pages > 1 && (
            <Pagination
              currentPage={meta.page}
              totalPages={meta.total_pages}
              onPageChange={handlePageChange}
            />
          )}
        </>
      )}
    </div>
  );
};

export default SearchResults;