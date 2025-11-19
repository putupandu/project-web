import React, { useState, useEffect } from 'react';
import { useParams, useSearchParams } from 'react-router-dom';
import { bookService } from '../services/bookService';
import { categoryService } from '../services/categoryService';
import BookGrid from '../components/books/BookGrid';
import Pagination from '../components/common/Pagination';
import Loading from '../components/common/Loading';
import ErrorMessage from '../components/common/ErrorMessage';

const CategoryPage = () => {
  const { slug } = useParams();
  const [searchParams, setSearchParams] = useSearchParams();
  
  const [category, setCategory] = useState(null);
  const [books, setBooks] = useState([]);
  const [meta, setMeta] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [currentPage, setCurrentPage] = useState(
    parseInt(searchParams.get('page')) || 1
  );

  useEffect(() => {
    fetchData();
  }, [slug, currentPage]);

  const fetchData = async () => {
    try {
      setLoading(true);
      setError(null);

      // Fetch category details
      const categories = await categoryService.getAll();
      const foundCategory = categories.data.find((cat) => cat.slug === slug);
      
      if (!foundCategory) {
        throw new Error('Kategori tidak ditemukan');
      }
      
      setCategory(foundCategory);

      // Fetch books in this category
      const response = await bookService.getAll({
        category_id: foundCategory.id,
        page: currentPage,
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
    searchParams.set('page', page.toString());
    setSearchParams(searchParams);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  if (loading) return <Loading message="Memuat kategori..." />;
  if (error) return <ErrorMessage message={error} onRetry={fetchData} />;

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Category Header */}
      {category && (
        <div className="bg-white rounded-lg shadow-md p-8 mb-8">
          <div className="flex items-center space-x-4">
            <div className="text-5xl">{category.icon}</div>
            <div>
              <h1 className="text-3xl font-bold text-gray-800 mb-2">
                {category.name}
              </h1>
              <p className="text-gray-600">{category.description}</p>
              {category.book_count > 0 && (
                <p className="text-sm text-gray-500 mt-2">
                  {category.book_count} buku tersedia
                </p>
              )}
            </div>
          </div>
        </div>
      )}

      {/* Books */}
      {books.length === 0 ? (
        <div className="text-center py-20">
          <p className="text-xl text-gray-600">
            Belum ada buku dalam kategori ini
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

export default CategoryPage;