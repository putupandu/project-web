import React from "react";
import BookCard from "./BookCard";

export default function BookGrid({ books, onRemove }) {
  return (
    <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-6">
      {books.map((book) => (
        <BookCard
          key={book.id}
          book={book}
          onRemove={onRemove}
        />
      ))}
    </div>
  );
}
//// books tweak 3
// books tweak 6
// books tweak 9
