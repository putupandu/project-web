
import React from 'react';
import { Book, Mail, Phone, MapPin } from 'lucide-react';
import { Link } from 'react-router-dom';
//
const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white mt-12">
      <div className="container mx-auto px-4 py-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {/* About */}
          <div>
            <div className="flex items-center space-x-2 mb-4">
              <Book size={32} />
              <h3 className="text-xl font-bold">E-Library</h3>
            </div>
            <p className="text-gray-400 text-sm">
              Perpustakaan digital Badan Meteorologi, Klimatologi, dan Geofisika
              yang menyediakan akses ke berbagai publikasi ilmiah dan dokumen teknis.
            </p>
          </div>

          {/* Quick Links */}
          <div>
            <h3 className="text-lg font-bold mb-4">Tautan Cepat</h3>
            <ul className="space-y-2 text-gray-400 text-sm">
              <li>
                <Link to="/" className="hover:text-white transition">
                  Beranda
                </Link>
              </li>
              <li>
                <Link to="/books" className="hover:text-white transition">
                  Koleksi Buku
                </Link>
              </li>
              <li>
                <Link to="/about" className="hover:text-white transition">
                  Tentang Kami
                </Link>
              </li>
            </ul>
          </div>

          {/* Contact */}
          <div>
            <h3 className="text-lg font-bold mb-4">Kontak</h3>
            <ul className="space-y-3 text-gray-400 text-sm">
              <li className="flex items-start space-x-2">
                <MapPin size={18} className="mt-1 flex-shrink-0" />
                <span>Jl. ZA. Pagar Alam No.9 -11, Bandar Lampung 35132</span>
              </li>
              <li className="flex items-center space-x-2">
                <Phone size={18} />
                <span>(021) 82828282</span>
              </li>
              <li className="flex items-center space-x-2">
                <Mail size={18} />
                <span>uti@teknokrat.ac.id</span>
              </li>
            </ul>
          </div>
        </div>

        <div className="border-t border-gray-700 mt-8 pt-6 text-center text-gray-400 text-sm">
          <p>&copy; {new Date().getFullYear()} E-Library. All rights reserved.</p>
          <p className="mt-1">Universitas Teknokrat Indonesia</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;// common component improvement 2
// common component improvement 6
