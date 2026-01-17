import React, { useState, useEffect } from 'react';
import { categoryService } from '../../services/categoryService';

const CategoryFilter = ({ selectedCategory, onCategoryChange }) => {
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchCategories();
  }, []);
//
  const fetchCategories = async () => {
    try {
      const response = await categoryService.getAll();
      setCategories(response.data || []);
    } catch (error) {
      console.error('Error fetching categories:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <div className="animate-pulse h-10 bg-gray-200 rounded"></div>;

  return (
    <div className="mb-6">
      <h3 className="font-bold text-gray-800 mb-3">Kategori</h3>
      <div className="flex flex-wrap gap-2">
        <button
          onClick={() => onCategoryChange(null)}
          className={`px-4 py-2 rounded-lg font-semibold transition-colors ${
            !selectedCategory
              ? 'bg-blue-600 text-white'
              : 'bg-white text-gray-700 hover:bg-blue-50 border'
          }`}
        >
          Semua
        </button>
        {categories.map((category) => (
          <button
            key={category.id}
            onClick={() => onCategoryChange(category.id)}
            className={`px-4 py-2 rounded-lg font-semibold transition-colors flex items-center ${
              selectedCategory === category.id
                ? 'bg-blue-600 text-white'
                : 'bg-white text-gray-700 hover:bg-blue-50 border'
            }`}
          >
            <span className="mr-2">{category.icon}</span>
            <span>{category.name}</span>
            {category.book_count > 0 && (
              <span className="ml-2 text-xs bg-white/20 px-2 py-0.5 rounded-full">
                {category.book_count}
              </span>
            )}
          </button>
        ))}
      </div>
    </div>
  );
};

export default CategoryFilter;