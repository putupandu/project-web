import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Book, TrendingUp, Clock, Star, ArrowRight } from 'lucide-react';
import { bookService } from '../services/bookService';
import { categoryService } from '../services/categoryService';
import BookCard from '../components/books/BookCard';
import Loading from '../components/common/Loading';

const Home = () => {
  const [latestBooks, setLatestBooks] = useState([]);
  const [popularBooks, setPopularBooks] = useState([]);
  const [categories, setCategories] = useState([]);
  const [totalBooks, setTotalBooks] = useState(0);
  const [totalDownloads, setTotalDownloads] = useState(0);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      setLoading(true);
      
      // Fetch stats
      const stats = await bookService.getStats();
      setTotalBooks(stats.total);
      setTotalDownloads(stats.totalDownloads);

      // Fetch latest books
      const latestResponse = await bookService.getAll({
        sort: 'created_at',
        order: 'DESC',
        per_page: 8,
      });
      setLatestBooks(latestResponse.data || []);

      // Fetch popular books
      const popularResponse = await bookService.getAll({
        sort: 'downloads',
        order: 'DESC',
        per_page: 4,
      });
      setPopularBooks(popularResponse.data || []);

      // Fetch categories
      const categoriesResponse = await categoryService.getAll();
      setCategories(categoriesResponse.data || []);
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <Loading message="Memuat halaman beranda..." />;

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Hero Section */}
      <section className="bg-gradient-to-r from-blue-600 to-purple-600 rounded-2xl p-8 md:p-12 mb-12 text-white">
        <div className="max-w-3xl">
          <h1 className="text-4xl md:text-5xl font-bold mb-4">
            Selamat Datang di E-Library
          </h1>
          <p className="text-xl mb-6 text-blue-100">
            Jelajahi berbagai koleksi publikasi ilmiah, jurnal, makalah, dan dokumen teknis 
            yang diterbitkan oleh serta lembaga terkait lainnya.
          </p>
          <Link
            to="/books"
            className="inline-flex items-center bg-white text-blue-600 px-6 py-3 rounded-lg font-semibold hover:bg-blue-50 transition"
          >
            Jelajahi Koleksi
            <ArrowRight className="ml-2" size={20} />
          </Link>
        </div>
      </section>

      {/* Statistics */}
      <section className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-12">
        <div className="bg-white rounded-lg p-6 shadow-md">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm mb-1">Total Buku</p>
              <p className="text-3xl font-bold text-blue-600">{totalBooks}</p>
            </div>
            <Book className="text-blue-600" size={40} />
          </div>
        </div>

        <div className="bg-white rounded-lg p-6 shadow-md">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm mb-1">Kategori</p>
              <p className="text-3xl font-bold text-green-600">{categories.length}</p>
            </div>
            <Star className="text-green-600" size={40} />
          </div>
        </div>

        <div className="bg-white rounded-lg p-6 shadow-md">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm mb-1">Total Unduhan</p>
              <p className="text-3xl font-bold text-purple-600">{totalDownloads}</p>
            </div>
            <TrendingUp className="text-purple-600" size={40} />
          </div>
        </div>

        <div className="bg-white rounded-lg p-6 shadow-md">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-600 text-sm mb-1">Buku Terbaru</p>
              <p className="text-3xl font-bold text-orange-600">{latestBooks.length}</p>
            </div>
            <Clock className="text-orange-600" size={40} />
          </div>
        </div>
      </section>

      {/* Categories */}
      <section className="mb-12">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-2xl font-bold text-gray-800">Kategori Utama</h2>
        </div>
        
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
          {categories.map((category) => (
            <Link
              key={category.id}
              to={`/category/${category.slug}`}
              className="bg-white rounded-lg p-6 shadow-md hover:shadow-xl transition-all text-center group"
            >
              <div className="text-4xl mb-3">{category.icon}</div>
              <h3 className="font-bold text-gray-800 group-hover:text-blue-600 transition">
                {category.name}
              </h3>
              {category.book_count > 0 && (
                <p className="text-sm text-gray-500 mt-1">
                  {category.book_count} buku
                </p>
              )}
            </Link>
          ))}
        </div>
      </section>

      {/* Popular Books */}
      {popularBooks.length > 0 && (
        <section className="mb-12">
          <div className="flex items-center justify-between mb-6">
            <div>
              <h2 className="text-2xl font-bold text-gray-800">Buku Populer</h2>
              <p className="text-gray-600">Buku yang paling banyak diunduh</p>
            </div>
            <Link
              to="/books?sort=downloads"
              className="text-blue-600 hover:text-blue-700 font-semibold flex items-center"
            >
              Lihat Semua
              <ArrowRight className="ml-1" size={16} />
            </Link>
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
            {popularBooks.map((book) => (
              <BookCard key={book.id} book={book} />
            ))}
          </div>
        </section>
      )}
//
      {/* Latest Books */}
      <section>
        <div className="flex items-center justify-between mb-6">
          <div>
            <h2 className="text-2xl font-bold text-gray-800">Buku Terbaru</h2>
            <p className="text-gray-600">Publikasi terkini dari </p>
          </div>
          <Link
            to="/books"
            className="text-blue-600 hover:text-blue-700 font-semibold flex items-center"
          >
            Lihat Semua
            <ArrowRight className="ml-1" size={16} />
          </Link>
        </div>
//
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {latestBooks.map((book) => (
            <BookCard key={book.id} book={book} />
          ))}
        </div>
      </section>
    </div>
  );
};

export default Home;// page layout refinement 1
// page layout refinement 5
// page layout refinement 9
