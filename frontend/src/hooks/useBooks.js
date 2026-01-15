// src/hooks/useBooks.js
import { useState, useEffect } from 'react';
import { bookService } from '../services/bookService';

export const useBooks = (initialFilters = {}) => {
  const [books, setBooks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [meta, setMeta] = useState(null);
  const [filters, setFilters] = useState(initialFilters);

  useEffect(() => {
    fetchBooks();
  }, [filters]);

  const fetchBooks = async () => {
    try {
      setLoading(true);
      setError(null);

      const response = await bookService.getAll(filters);

      // FIX PENTING — cover harus ada dan tidak boleh null
      const fixedBooks = (response.data || []).map((book) => {
        return {
          ...book,
          cover: book.cover ?? null,   // pastikan cover tetap ada
          file_url: book.file_url ?? null,
        };
      });

      setBooks(fixedBooks);
      setMeta(response.meta || null);
    } catch (err) {
      setError(err.message || "Failed to load books");
    } finally {
      setLoading(false);
    }
  };

  const updateFilters = (newFilters) => {
    setFilters((prev) => ({ ...prev, ...newFilters }));
  };

  const refetch = () => {
    fetchBooks();
  };

  return {
    books,
    loading,
    error,
    meta,
    filters,
    updateFilters,
    refetch,
  };
};
 // // logic refinement 1
