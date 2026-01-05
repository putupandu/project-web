import api from "../services/api";

export const saveBook = (bookId) => {
  return api.post("/saved-books", {
    book_id: bookId,
  });
};

export const getSavedBooks = () => {
  return api.get("/saved-books");
};

export const removeSavedBook = (bookId) => {
  return api.delete(`/saved-books/${bookId}`);
};
