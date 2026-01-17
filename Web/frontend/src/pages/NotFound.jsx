import React from 'react';
import { Link } from 'react-router-dom';
import { Home, Search, ArrowLeft } from 'lucide-react';

const NotFound = () => {
  return (
    <div className="container mx-auto px-4 py-20">
      <div className="max-w-2xl mx-auto text-center">
        {/* 404 Image/Icon */}
        <div className="mb-8">
          <h1 className="text-9xl font-bold text-gray-300 mb-4">404</h1>
          <div className="flex justify-center">
            <Search className="text-gray-400" size={80} />
          </div>
        </div>

        {/* Message */}
        <h2 className="text-3xl font-bold text-gray-800 mb-4">
          Halaman Tidak Ditemukan
        </h2>
        <p className="text-gray-600 mb-8 text-lg">
          Maaf, halaman yang Anda cari tidak dapat ditemukan atau mungkin telah dipindahkan.
        </p>
//
        {/* Action Buttons */}
        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <Link
            to="/"
            className="inline-flex items-center justify-center bg-blue-600 hover:bg-blue-700 text-white font-semibold px-6 py-3 rounded-lg transition"
          >
            <Home className="mr-2" size={20} />
            Kembali ke Beranda
          </Link>

          <Link
            to="/books"
            className="inline-flex items-center justify-center bg-white hover:bg-gray-50 text-blue-600 font-semibold px-6 py-3 rounded-lg border-2 border-blue-600 transition"
          >
            <Search className="mr-2" size={20} />
            Jelajahi Buku
          </Link>
        </div>

        {/* Helpful Links */}
        <div className="mt-12 pt-8 border-t">
          <p className="text-gray-600 mb-4">Atau kunjungi halaman berikut:</p>
          <div className="flex flex-wrap justify-center gap-4 text-sm">
            <Link to="/" className="text-blue-600 hover:text-blue-700 hover:underline">
              Beranda
            </Link>
            <Link to="/books" className="text-blue-600 hover:text-blue-700 hover:underline">
              Koleksi Buku
            </Link>
            <Link to="/about" className="text-blue-600 hover:text-blue-700 hover:underline">
              Tentang Kami
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default NotFound;