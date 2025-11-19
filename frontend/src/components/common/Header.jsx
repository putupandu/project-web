import React from 'react';
import { Link } from 'react-router-dom';
import { Book, Menu, X } from 'lucide-react';
import SearchBar from './SearchBar';

const Header = () => {
  const [mobileMenuOpen, setMobileMenuOpen] = React.useState(false);

  return (
    <header className="bg-gradient-to-r from-blue-600 to-blue-800 text-white shadow-lg sticky top-0 z-50">
      <div className="container mx-auto px-4">
        {/* Top Bar */}
        <div className="flex items-center justify-between py-4">
          <Link to="/" className="flex items-center space-x-3">
            <Book size={40} />
            <div>
              <h1 className="text-2xl font-bold">E-Library</h1>
              <p className="text-xs text-blue-100">Perpustakaan Digital</p>
            </div>
          </Link>

          {/* Desktop Navigation */}
          <nav className="hidden md:flex items-center space-x-6">
            <Link to="/" className="hover:text-blue-200 transition">
              Beranda
            </Link>
            <Link to="/books" className="hover:text-blue-200 transition">
              Koleksi Buku
            </Link>
            <Link to="/about" className="hover:text-blue-200 transition">
              Tentang
            </Link>
          </nav>

          {/* Mobile Menu Button */}
          <button
            className="md:hidden"
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
          >
            {mobileMenuOpen ? <X size={24} /> : <Menu size={24} />}
          </button>
        </div>

        {/* Search Bar */}
        <div className="pb-4">
          <SearchBar />
        </div>

        {/* Mobile Navigation */}
        {mobileMenuOpen && (
          <nav className="md:hidden pb-4 space-y-2">
            <Link
              to="/"
              className="block py-2 hover:bg-blue-700 px-4 rounded"
              onClick={() => setMobileMenuOpen(false)}
            >
              Beranda
            </Link>
            <Link
              to="/books"
              className="block py-2 hover:bg-blue-700 px-4 rounded"
              onClick={() => setMobileMenuOpen(false)}
            >
              Koleksi Buku
            </Link>
            <Link
              to="/about"
              className="block py-2 hover:bg-blue-700 px-4 rounded"
              onClick={() => setMobileMenuOpen(false)}
            >
              Tentang
            </Link>
          </nav>
        )}
      </div>
    </header>
  );
};

export default Header;