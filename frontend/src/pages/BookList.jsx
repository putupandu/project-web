import React, { useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import { SlidersHorizontal, Grid, List as ListIcon } from 'lucide-react';
import { useBooks } from '../hooks/useBooks';
import BookGrid from '../components/books/BookGrid';
import CategoryFilter from '../components/filters/CategoryFilter';
import Pagination from '../components/common/Pagination';
import Loading from '../components/common/Loading';
import ErrorMessage from '../components/common/ErrorMessage';

const BookList = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [showFilters, setShowFilters] = useState(false);
  const [viewMode, setViewMode] = useState('grid'); // 'grid' or 'list'

  // initial filter
  const initialFilters = {
    page: parseInt(searchParams.get('page')) || 1,
    per_page: 12,
    category_id: searchParams.get('category_id') || null,
    sort: searchParams.get('sort') || 'created_at',
    order: searchParams.get('order') || 'DESC',
  };

  // FIX: tambahkan filters di sini
  const { books, loading, error, meta, filters, updateFilters, refetch } =
    useBooks(initialFilters);

  // Category Change
  const handleCategoryChange = (categoryId) => {
    const newFilters = { category_id: categoryId, page: 1 };
    updateFilters(newFilters);

    if (categoryId) {
      searchParams.set('category_id', categoryId);
    } else {
      searchParams.delete('category_id');
    }
    searchParams.set('page', '1');
    setSearchParams(searchParams);
  };
//
  // Page Change
  const handlePageChange = (page) => {
    updateFilters({ page });
    searchParams.set('page', page.toString());
    setSearchParams(searchParams);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  // Sort Change (FIX)
  const handleSortChange = (e) => {
    const newFilters = { sort: e.target.value, page: 1 };
    updateFilters(newFilters);

    searchParams.set('sort', e.target.value);
    searchParams.set('page', '1');
    setSearchParams(searchParams);
  };

  if (error) {
    return (
      <div className="container mx-auto px-4 py-8">
        <ErrorMessage message={error} onRetry={refetch} />
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-800 mb-2">Koleksi Buku</h1>
        <p className="text-gray-600">
          Jelajahi {meta?.total || 0} buku dalam perpustakaan digital kami
        </p>
      </div>

      {/* Toolbar */}
      <div className="bg-white rounded-lg shadow-md p-4 mb-6">
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
          {/* Filter Toggle */}
          <button
            onClick={() => setShowFilters(!showFilters)}
            className="flex items-center space-x-2 text-gray-700 hover:text-blue-600 transition"
          >
            <SlidersHorizontal size={20} />
            <span className="font-semibold">
              {showFilters ? 'Sembunyikan' : 'Tampilkan'} Filter
            </span>
          </button>

          <div className="flex items-center space-x-4 w-full md:w-auto">
            {/* Sort */}
            <select
              value={filters.sort}
              onChange={handleSortChange}
              className="px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-300"
            >
              <option value="created_at">Terbaru</option>
              <option value="title">Judul (A-Z)</option>
              <option value="year">Tahun</option>
              <option value="downloads">Paling Banyak Diunduh</option>
              <option value="views">Paling Banyak Dilihat</option>
            </select>

            {/* View Mode */}
            <div className="flex space-x-2">
              <button
                onClick={() => setViewMode('grid')}
                className={`p-2 rounded ${
                  viewMode === 'grid'
                    ? 'bg-blue-600 text-white'
                    : 'bg-gray-200 text-gray-600 hover:bg-gray-300'
                }`}
              >
                <Grid size={20} />
              </button>

              <button
                onClick={() => setViewMode('list')}
                className={`p-2 rounded ${
                  viewMode === 'list'
                    ? 'bg-blue-600 text-white'
                    : 'bg-gray-200 text-gray-600 hover:bg-gray-300'
                }`}
              >
                <ListIcon size={20} />
              </button>
            </div>
          </div>
        </div>
      </div>

      {/* Filters */}
      {showFilters && (
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <CategoryFilter
            selectedCategory={filters.category_id}
            onCategoryChange={handleCategoryChange}
          />
        </div>
      )}

      {/* Content */}
      {loading ? (
        <Loading message="Memuat koleksi buku..." />
      ) : (
        <>
          <BookGrid books={books} />

          {/* Pagination */}
          {meta && meta.total_pages > 1 && (
            <Pagination
              currentPage={filters.page}
              totalPages={meta.total_pages}
              onPageChange={handlePageChange}
            />
          )}
        </>
      )}
    </div>
  );
};

export default BookList;
// page layout refinement 2
// page layout refinement 6
// page layout refinement 10
