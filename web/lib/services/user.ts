import apiClient from '../apiClient';

export const login = async (credentials: { email: string; password: string }) => {
    const response = await apiClient.post('/users/login', credentials);
    return response.data;
};

export const register = async (data: { username: string, email: string, password: string }) => {
    const response = await apiClient.post('/users/register', data);
    return response.data;
};
