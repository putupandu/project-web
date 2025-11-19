import React from 'react';
import { Book, Target, Users, Award, Mail, Phone, MapPin } from 'lucide-react';

const About = () => {
  return (
    <div className="container mx-auto px-4 py-8">
      {/* Hero */}
      <div className="bg-gradient-to-r from-blue-600 to-purple-600 rounded-2xl p-8 md:p-12 mb-12 text-white">
        <h1 className="text-4xl md:text-5xl font-bold mb-4">
          Tentang E-Library
        </h1>
        <p className="text-xl text-blue-100">
          Platform perpustakaan digital untuk mengakses publikasi ilmiah
        </p>
      </div>

      {/* Introduction */}
      <section className="mb-12">
        <div className="bg-white rounded-lg shadow-md p-8">
          <div className="flex items-start space-x-4 mb-6">
            <Book className="text-blue-600 flex-shrink-0" size={40} />
            <div>
              <h2 className="text-2xl font-bold text-gray-800 mb-3">
                Apa itu E-Library?
              </h2>
              <p className="text-gray-600 leading-relaxed mb-4">
                E-Library adalah perpustakaan digital yang dikembangkan oleh Badan Meteorologi, 
                Klimatologi, dan Geofisika (.) untuk memberikan akses mudah dan cepat terhadap 
                berbagai publikasi ilmiah, jurnal, makalah, dan dokumen teknis.
              </p>
              <p className="text-gray-600 leading-relaxed">
                Platform ini dirancang untuk mendukung peneliti, mahasiswa, akademisi, dan masyarakat 
                umum dalam mengakses informasi terkait meteorologi, klimatologi, geofisika, dan bidang 
                terkait lainnya.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Vision & Mission */}
      <section className="grid md:grid-cols-2 gap-6 mb-12">
        <div className="bg-white rounded-lg shadow-md p-8">
          <div className="flex items-center space-x-3 mb-4">
            <Target className="text-blue-600" size={32} />
            <h2 className="text-2xl font-bold text-gray-800">Visi</h2>
          </div>
          <p className="text-gray-600 leading-relaxed">
            Menjadi perpustakaan digital terdepan dalam menyediakan akses informasi ilmiah 
            di bidang meteorologi, klimatologi, dan geofisika untuk mendukung penelitian 
            dan pengembangan ilmu pengetahuan.
          </p>
        </div>

        <div className="bg-white rounded-lg shadow-md p-8">
          <div className="flex items-center space-x-3 mb-4">
            <Award className="text-purple-600" size={32} />
            <h2 className="text-2xl font-bold text-gray-800">Misi</h2>
          </div>
          <ul className="space-y-2 text-gray-600">
            <li className="flex items-start">
              <span className="text-purple-600 mr-2">•</span>
              <span>Menyediakan akses mudah ke publikasi ilmiah berkualitas</span>
            </li>
            <li className="flex items-start">
              <span className="text-purple-600 mr-2">•</span>
              <span>Mendukung kegiatan penelitian dan pendidikan</span>
            </li>
            <li className="flex items-start">
              <span className="text-purple-600 mr-2">•</span>
              <span>Menyebarluaskan hasil penelitian </span>
            </li>
          </ul>
        </div>
      </section>

      {/* Features */}
      <section className="mb-12">
        <h2 className="text-2xl font-bold text-gray-800 mb-6 text-center">
          Fitur Unggulan
        </h2>
        <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Book className="text-blue-600" size={32} />
            </div>
            <h3 className="font-bold text-gray-800 mb-2">Koleksi Lengkap</h3>
            <p className="text-gray-600 text-sm">
              Ribuan publikasi ilmiah dari berbagai kategori
            </p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Users className="text-green-600" size={32} />
            </div>
            <h3 className="font-bold text-gray-800 mb-2">Akses Mudah</h3>
            <p className="text-gray-600 text-sm">
              Dapat diakses kapan saja dan di mana saja
            </p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Award className="text-purple-600" size={32} />
            </div>
            <h3 className="font-bold text-gray-800 mb-2">Gratis</h3>
            <p className="text-gray-600 text-sm">
              Semua konten dapat diakses secara gratis
            </p>
          </div>
        </div>
      </section>

      {/* Contact */}
      <section className="bg-white rounded-lg shadow-md p-8">
        <h2 className="text-2xl font-bold text-gray-800 mb-6">Hubungi Kami</h2>
        <div className="grid md:grid-cols-2 gap-6">
          <div>
            <h3 className="font-bold text-gray-800 mb-4">Alamat</h3>
            <div className="space-y-3 text-gray-600">
              <div className="flex items-start space-x-3">
                <MapPin className="text-blue-600 flex-shrink-0 mt-1" size={20} />
                <div>
                  <p className="font-semibold">Kantor Pusat </p>
                  <p className="text-sm">
                    Jl. Angkasa I No.2, Kemayoran<br />
                    Jakarta Pusat 10720<br />
                    Indonesia
                  </p>
                </div>
              </div>

              <div className="flex items-center space-x-3">
                <Phone className="text-blue-600 flex-shrink-0" size={20} />
                <div>
                  <p className="font-semibold">Telepon</p>
                  <p className="text-sm">(021) 4246321</p>
                </div>
              </div>

              <div className="flex items-center space-x-3">
                <Mail className="text-blue-600 flex-shrink-0" size={20} />
                <div>
                  <p className="font-semibold">Email</p>
                  <p className="text-sm">info@bmkg.go.id</p>
                </div>
              </div>
            </div>
          </div>

          <div>
            <h3 className="font-bold text-gray-800 mb-4">Jam Operasional</h3>
            <div className="space-y-2 text-gray-600">
              <div className="flex justify-between">
                <span>Senin - Jumat</span>
                <span className="font-semibold">08:00 - 16:00 WIB</span>
              </div>
              <div className="flex justify-between">
                <span>Sabtu - Minggu</span>
                <span className="font-semibold">Tutup</span>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};

export default About;