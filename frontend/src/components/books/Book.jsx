
export default function Book({ book }) {
  return (
    <div className="book-card">
      <h3>{book.title}</h3>
      <p>by {book.author}</p>
      <p>Year: {book.year}</p>
    </div>
  );
}