import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import {
  Book as BookIcon,
  User,
  Calendar,
  FileText,
  Download,
  Eye,
  ArrowLeft,
  Share2,
  Bookmark,
} from 'lucide-react';
import { bookService } from '../services/bookService';
import Loading from '../components/common/Loading';
import ErrorMessage from '../components/common/ErrorMessage';

const BookDetail = () => {
  const { id } = useParams();
  const [book, setBook] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [downloading, setDownloading] = useState(false);

  useEffect(() => {
    fetchBook();
  }, [id]);

  const fetchBook = async () => {
    try {
      setLoading(true);
      setError(null);
      
      // Increment view count
      await bookService.view(id);
      
      // Fetch book details
      const response = await bookService.getById(id);
      setBook(response.data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleDownload = async () => {
    try {
      setDownloading(true);
      await bookService.download(id);
      
      // Trigger actual download
      if (book.file_url) {
        window.open(book.file_url, '_blank');
      }
      
      // Update download count in UI
      setBook((prev) => ({
        ...prev,
        downloads: (prev.downloads || 0) + 1,
      }));
    } catch (err) {
      alert('Gagal mengunduh file');
    } finally {
      setDownloading(false);
    }
  };

  if (loading) return <Loading message="Memuat detail buku..." />;
  if (error) return <ErrorMessage message={error} onRetry={fetchBook} />;
  if (!book) return <ErrorMessage message="Buku tidak ditemukan" />;

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Back Button */}
      <Link
        to="/books"
        className="inline-flex items-center text-blue-600 hover:text-blue-700 mb-6 font-semibold"
      >
        <ArrowLeft className="mr-2" size={20} />
        Kembali ke Koleksi
      </Link>

      <div className="bg-white rounded-lg shadow-lg overflow-hidden">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 p-8">
          {/* Left Column - Book Cover */}
          <div className="lg:col-span-1">
            <div className="sticky top-8">
              <div className="aspect-[3/4] bg-gradient-to-br from-blue-400 to-purple-500 rounded-lg overflow-hidden shadow-xl mb-4">
                {book.cover ? (
                  <img
                    src={book.cover}
                    alt={book.title}
                    className="w-full h-full object-cover"
                  />
                ) : (
                  <div className="flex items-center justify-center h-full">
                    <BookIcon size={120} className="text-white opacity-50" />
                  </div>
                )}
              </div>

              {/* Action Buttons */}
              <div className="space-y-3">
                <button
                  onClick={handleDownload}
                  disabled={downloading || !book.file_url}
                  className="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-4 rounded-lg transition flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  <Download size={20} className="mr-2" />
                  {downloading ? 'Mengunduh...' : 'Download PDF'}
                </button>

                <button className="w-full bg-green-600 hover:bg-green-700 text-white font-semibold py-3 px-4 rounded-lg transition flex items-center justify-center">
                  <Eye size={20} className="mr-2" />
                  Baca Online
                </button>

                <div className="flex space-x-2">
                  <button className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-700 font-semibold py-2 px-4 rounded-lg transition flex items-center justify-center">
                    <Bookmark size={18} className="mr-1" />
                    Simpan
                  </button>
                  <button className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-700 font-semibold py-2 px-4 rounded-lg transition flex items-center justify-center">
                    <Share2 size={18} className="mr-1" />
                    Bagikan
                  </button>
                </div>
              </div>

              {/* Stats */}
              <div className="mt-6 bg-gray-50 rounded-lg p-4">
                <div className="grid grid-cols-2 gap-4 text-center">
                  <div>
                    <Eye className="mx-auto mb-1 text-blue-600" size={24} />
                    <p className="text-2xl font-bold text-gray-800">{book.views || 0}</p>
                    <p className="text-xs text-gray-600">Dilihat</p>
                  </div>
                  <div>
                    <Download className="mx-auto mb-1 text-green-600" size={24} />
                    <p className="text-2xl font-bold text-gray-800">{book.downloads || 0}</p>
                    <p className="text-xs text-gray-600">Diunduh</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Right Column - Book Info */}
          <div className="lg:col-span-2">
            <div className="mb-6">
              {book.category && (
                <Link
                  to={`/category/${book.category.slug}`}
                  className="inline-flex items-center bg-blue-100 text-blue-800 px-3 py-1 rounded-full text-sm font-semibold mb-4 hover:bg-blue-200 transition"
                >
                  <span className="mr-1">{book.category.icon}</span>
                  {book.category.name}
                </Link>
              )}

              <h1 className="text-3xl md:text-4xl font-bold text-gray-800 mb-4">
                {book.title}
              </h1>

              <div className="flex flex-wrap gap-4 text-gray-600 mb-6">
                <div className="flex items-center">
                  <User size={18} className="mr-2" />
                  <span className="font-semibold">{book.author}</span>
                </div>

                <div className="flex items-center">
                  <Calendar size={18} className="mr-2" />
                  <span>{book.year}</span>
                </div>

                {book.pages && (
                  <div className="flex items-center">
                    <FileText size={18} className="mr-2" />
                    <span>{book.pages} halaman</span>
                  </div>
                )}

                {book.language && (
                  <div className="flex items-center">
                    <BookIcon size={18} className="mr-2" />
                    <span>{book.language}</span>
                  </div>
                )}
              </div>
            </div>

            {/* Description */}
            <div className="mb-6">
              <h2 className="text-xl font-bold text-gray-800 mb-3">Deskripsi</h2>
              <p className="text-gray-600 leading-relaxed whitespace-pre-line">
                {book.description || 'Tidak ada deskripsi tersedia.'}
              </p>
            </div>

            {/* Details */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
              <div className="bg-gray-50 rounded-lg p-4">
                <h3 className="font-semibold text-gray-700 mb-2">Penerbit</h3>
                <p className="text-gray-800">{book.publisher || '-'}</p>
              </div>

              {book.isbn && (
                <div className="bg-gray-50 rounded-lg p-4">
                  <h3 className="font-semibold text-gray-700 mb-2">ISBN</h3>
                  <p className="text-gray-800 font-mono">{book.isbn}</p>
                </div>
              )}

              <div className="bg-gray-50 rounded-lg p-4">
                <h3 className="font-semibold text-gray-700 mb-2">Tahun Terbit</h3>
                <p className="text-gray-800">{book.year}</p>
              </div>

              <div className="bg-gray-50 rounded-lg p-4">
                <h3 className="font-semibold text-gray-700 mb-2">Bahasa</h3>
                <p className="text-gray-800">{book.language || 'Indonesian'}</p>
              </div>
            </div>

            {/* Related Info */}
            <div className="border-t pt-6">
              <p className="text-sm text-gray-500">
                Dipublikasikan pada {new Date(book.created_at).toLocaleDateString('id-ID', {
                  day: 'numeric',
                  month: 'long',
                  year: 'numeric',
                })}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BookDetail;