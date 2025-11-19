import api from './api';

export const bookService = {
  // Get all books with filters
  getAll: async (params = {}) => {
    const response = await api.get('/books', { params });
    return response;
  },

  // Get single book
  getById: async (id) => {
    const response = await api.get(`/books/${id}`);
    return response;
  },

  // Create book
  create: async (data) => {
    const response = await api.post('/books', data);
    return response;
  },

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

  // Increment download
  download: async (id) => {
    const response = await api.post(`/books/${id}/download`);
    return response;
  },

  // Increment view
  view: async (id) => {
    const response = await api.post(`/books/${id}/view`);
    return response;
  },
};