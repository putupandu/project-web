import api from './api';

export const bookService = {
  // Get all books
  getAll: async (params = {}) => {
    const response = await api.get('/books', { params });
    return response;
  },

  // Get book by ID
  getById: async (id) => {
    const response = await api.get(`/books/${id}`);
    return response;
  },

  // Create new book
  create: async (data) => {
    const response = await api.post('/books', data);
    return response;
  },
//
  // Update book
  update: async (id, data) => {
    const response = await api.put(`/books/${id}`, data);
    return response;
  },

  // Delete book
  delete: async (id) => {
    const response = await api.delete(`/books/${id}`);
    return response;
  },

  // Increment download count
  download: async (id) => {
    const response = await api.post(`/books/${id}/download`);
    return response;
  },

  // Increment view count
  view: async (id) => {
    const response = await api.post(`/books/${id}/view`);
    return response;
  },

  // Get statistics (total books, total downloads)
  getStats: async () => {
    const response = await api.get('/books', {
      params: { per_page: 1 }
    });
    const total = response.meta?.total || 0;
    const totalDownloads = (response.data || []).reduce((sum, book) => sum + (book.downloads || 0), 0);
    return { total, totalDownloads };
  },

  // Saved books
  save: async (data) => {
    const response = await api.post('/saved-books', data);
    return response;
  },

  getSaved: async () => {
    const response = await api.get('/saved-books');
    return response;
  },

  deleteSaved: async (id) => {
    const response = await api.delete(`/saved-books/${id}`);
    return response;
  },
};// logic refinement 4
