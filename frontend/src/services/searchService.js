import api from './api';
//
export const searchService = {
  search: async (query, params = {}) => {
    const response = await api.get('/search', {
      params: { q: query, ...params },
    });
    return response;
  },
};// logic refinement 5
