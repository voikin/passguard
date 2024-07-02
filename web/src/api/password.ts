import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8000',
});

export const getPasswordScore = async (password: string) => {
  const response = await apiClient.get('/api/evaluate', {
    params: { password },
  });
  return response.data;
};
