import { useEffect, useState } from "react";
import { getSavedBooks, removeSavedBook } from "../api/savedBookApi";
import BookGrid from "../components/books/BookGrid";
import Loading from "../components/common/Loading";
import ErrMessage from "../components/common/ErrMessage";

export default function SavedBooks() {
  const [books, setBooks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    fetchSavedBooks();
  }, []);

  const fetchSavedBooks = async () => {
    try {
      const res = await getSavedBooks();
      setBooks(res.data.data);
    } catch (err) {
      setError("Gagal mengambil buku tersimpan");
    } finally {
      setLoading(false);
    }
  };
//
  const handleRemove = async (id) => {
    await removeSavedBook(id);
    setBooks((prev) => prev.filter((b) => b.id !== id));
  };

  if (loading) return <Loading />;
  if (error) return <ErrMessage message={error} />;

  return (
    <div>
      <h1>Buku Tersimpan</h1>

      {books.length === 0 ? (
        <p>Belum ada buku yang disimpan</p>
      ) : (
        <BookGrid books={books} onRemove={handleRemove} />
      )}
    </div>
  );
}
