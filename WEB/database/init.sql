-- ----------------------------------
-- 1. Buat database (jika belum ada)
-- ----------------------------------
-- CREATE DATABASE e_library;

-- ----------------------------------
-- 2. Gunakan database
-- ----------------------------------
-- \c e_library;

-- ----------------------------------
-- 3. Tabel categories
-- ----------------------------------
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    slug VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    icon VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------------
-- 4. Tabel books
-- ----------------------------------
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year INTEGER NOT NULL,

    -- Diperbaiki: tambah ON UPDATE CASCADE
    category_id INTEGER REFERENCES categories(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,

    -- Diperbaiki: isbn unique
    isbn VARCHAR(50) UNIQUE,

    cover TEXT,
    description TEXT,
    publisher VARCHAR(255),
    pages INTEGER,
    language VARCHAR(50) DEFAULT 'Indonesian',
    downloads INTEGER DEFAULT 0,
    views INTEGER DEFAULT 0,
    file_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------------
-- 5. Index
-- ----------------------------------
CREATE INDEX IF NOT EXISTS idx_books_category ON books(category_id);
CREATE INDEX IF NOT EXISTS idx_books_year ON books(year);
CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);
CREATE INDEX IF NOT EXISTS idx_books_author ON books(author);
CREATE INDEX IF NOT EXISTS idx_categories_slug ON categories(slug);

-- ----------------------------------
-- 6. Insert kategori default (tanpa emoji)
-- ----------------------------------
INSERT INTO categories (name, slug, description, icon)
VALUES
('Meteorologi', 'meteorologi', 'Ilmu cuaca dan atmosfer', NULL),
('Klimatologi', 'klimatologi', 'Ilmu iklim dan perubahan iklim', NULL),
('Geofisika', 'geofisika', 'Ilmu gempa bumi dan geofisika', NULL),
('Instrumentasi', 'instrumentasi', 'Alat dan teknologi meteorologi', NULL),
('Hidrologi', 'hidrologi', 'Ilmu air dan hidrologi', NULL),
('Maritim', 'maritim', 'Meteorologi maritim dan oseanografi', NULL)
ON CONFLICT (slug) DO NOTHING;

-- ----------------------------------
-- 7. Insert sample books
-- ----------------------------------
INSERT INTO books (title, author, year, category_id, isbn, description, publisher, pages, cover, file_url)
VALUES
('Meteorologi dan Klimatologi Indonesia', 'E-LIBRARY Research Team', 2024, 1, '978-602-1234-56-7', 
 'Kajian mendalam tentang kondisi meteorologi dan klimatologi di Indonesia dengan analisis data terkini.',
 'E-LIBRARY Press', 350, 'https://via.placeholder.com/200x280/8B5CF6/FFFFFF?text=Meteorologi', '/uploads/books/meteorologi-indonesia.pdf'),

('Analisis Gempa Bumi Tektonik', 'Dr. Ahmad Setiawan', 2024, 3, '978-602-1234-57-4',
 'Studi komprehensif tentang gempa bumi tektonik di wilayah Indonesia dan sistem peringatan dini.',
 'E-LIBRARY Press', 280, 'https://via.placeholder.com/200x280/3B82F6/FFFFFF?text=Gempa', '/uploads/books/gempa-tektonik.pdf'),

('Panduan Prakiraan Cuaca', 'E-LIBRARY Publication', 2023, 1, '978-602-1234-58-1',
 'Panduan lengkap untuk memahami dan memprediksi cuaca dengan teknologi modern.',
 'E-LIBRARY Press', 220, 'https://via.placeholder.com/200x280/10B981/FFFFFF?text=Cuaca', '/uploads/books/prakiraan-cuaca.pdf'),

('Perubahan Iklim Global dan Dampaknya', 'International Climate Team', 2024, 2, '978-602-1234-59-8',
 'Analisis dampak perubahan iklim global terhadap Indonesia dan strategi adaptasi.',
 'E-LIBRARY Press', 400, 'https://via.placeholder.com/200x280/EF4444/FFFFFF?text=Iklim', '/uploads/books/perubahan-iklim.pdf'),

('Instrumentasi Meteorologi Modern', 'Dr. Budi Hartono', 2023, 4, '978-602-1234-60-4',
 'Panduan penggunaan alat-alat meteorologi modern dan teknologi sensor terkini.',
 'E-LIBRARY Press', 180, 'https://via.placeholder.com/200x280/6366F1/FFFFFF?text=Instrumen', '/uploads/books/instrumentasi.pdf'),

('Hidrologi Terapan', 'Prof. Siti Nurhaliza', 2023, 5, '978-602-1234-61-1',
 'Aplikasi hidrologi dalam manajemen sumber daya air dan penanggulangan banjir.',
 'E-LIBRARY Press', 320, 'https://via.placeholder.com/200x280/F59E0B/FFFFFF?text=Hidrologi', '/uploads/books/hidrologi-terapan.pdf')
ON CONFLICT DO NOTHING;

-- ----------------------------------
-- 8. Update download & view (sample data)
-- ----------------------------------
UPDATE books
SET downloads = FLOOR(RANDOM() * 2000 + 500),
    views = FLOOR(RANDOM() * 5000 + 1000);

-- ----------------------------------
-- 9. Trigger update timestamp
-- ----------------------------------
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_timestamp
BEFORE UPDATE ON books
FOR EACH ROW EXECUTE FUNCTION update_timestamp();
