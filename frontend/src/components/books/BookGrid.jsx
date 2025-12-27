import Book from './book'; 

export default function BookGrid({ books }) {
  if (!books || books.length === 0) {
    return <p>No books available.</p>;
  }

  return (
    <div className="book-grid">
      {books.map(book => (
        <Book key={book.id} book={book} />
      ))}
    </div>
  );
}