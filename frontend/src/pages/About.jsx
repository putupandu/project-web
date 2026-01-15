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
                E-Library adalah platform perpustakaan digital yang dirancang untuk menyediakan akses mudah, 
                cepat, dan terorganisir terhadap berbagai jenis koleksi digital seperti jurnal ilmiah, 
                artikel, buku, laporan penelitian, modul pembelajaran, dan dokumen akademik lainnya.
              </p>
              <p className="text-gray-600 leading-relaxed">
                Tujuan utama E-Library adalah mendukung kebutuhan informasi bagi mahasiswa, peneliti, dosen, 
                pengajar, serta masyarakat umum yang membutuhkan sumber referensi terpercaya secara online.
                Dengan adanya E-Library, pengguna dapat mencari, membaca, 
                dan mengunduh berbagai materi digital kapan saja dan di mana saja tanpa batasan fisik.
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
            Menjadi platform perpustakaan digital yang modern, aksesibel, dan terpercaya dalam mendukung kegiatan pendidikan, 
            penelitian, serta pengembangan ilmu pengetahuan secara berkelanjutan.
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
              <span>Menyediakan akses cepat dan mudah terhadap berbagai koleksi digital yang berkualitas, lengkap, dan relevan.</span>
            </li>
            <li className="flex items-start">
              <span className="text-purple-600 mr-2">•</span>
              <span>Mendukung proses pembelajaran dan penelitian dengan menyediakan sumber informasi ilmiah yang terstruktur dan dapat diandalkan.</span>
            </li>
            <li className="flex items-start">
              <span className="text-purple-600 mr-2">•</span>
              <span>Membangun ekosistem digital yang efisien, sehingga pengguna dapat mencari, membaca, dan mengelola dokumen secara praktis melalui satu platform. </span>
            </li>
            <li className="flex items-start">
              <span className="text-purple-600 mr-2">•</span>
              <span>Mengoptimalkan teknologi informasi untuk meningkatkan pelayanan perpustakaan dan mempermudah distribusi pengetahuan secara luas. </span>
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
//
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
                  <p className="font-semibold"> Alamat Pusat </p>
                  <p className="text-sm">
                    Jl. ZA. Pagar Alam No.9 -11, Labuhan Ratu, Kec. Kedaton<br />
                    Bandar Lampung 35132<br />
                    Indonesia
                  </p>
                </div>  
              </div>

              <div className="flex items-center space-x-3">
                <Phone className="text-blue-600 flex-shrink-0" size={20} />
                <div>
                  <p className="font-semibold">Telepon</p>
                  <p className="text-sm">(021) 82828282</p>
                </div>
              </div>

              <div className="flex items-center space-x-3">
                <Mail className="text-blue-600 flex-shrink-0" size={20} />
                <div>
                  <p className="font-semibold">Email</p>
                  <p className="text-sm">uti@teknokrat.ac.id</p>
                </div>
              </div>
            </div>
          </div>

          <div>
            <h3 className="font-bold text-gray-800 mb-4">Jam Operasional</h3>
            <div className="space-y-2 text-gray-600">
              <div className="flex justify-between">
                <span>Senin - Jumat</span>
                <span className="font-semibold">07:30 - 21:00 WIB</span>
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

export default About;// page layout refinement 4
