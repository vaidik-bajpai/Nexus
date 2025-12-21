import apiClient from "../apiClient";

export const createBoard = async (board: { name: string, background: string, visibility: string }) => {
    try {
        const response = await apiClient.post("/boards/create", board);
        return response.data;
    } catch (error) {
        console.error("Error creating board:", error);
        throw error;
    }
}

export const listBoard = async (paginate: { page: number, size: number }) => {
    try {
        const response = await apiClient.get(`/boards/list?page=${paginate.page}&size=${paginate.size}`);
        return response.data;
    } catch (error) {
        console.error("Error listing boards:", error);
        throw error;
    }
}

export const getCardAndLists = async (id: string) => {
    try {
        const response = await apiClient.get(`/boards/${id}/cards-and-lists`);
        return response.data;
    } catch (error) {
        console.error("Error getting board cards and lists:", error);
        throw error;
    }
}

export const getBoardDetails = async (id: string) => {
    try {
        const response = await apiClient.get(`/boards/${id}/details`);
        return response.data;
    } catch (error) {
        console.error("Error getting board details:", error);
        throw error;
    }
}

export const updateList = async (list: { listID: string, boardID: string, name?: string, position?: number }) => {
    try {
        const response = await apiClient.put(`/boards/${list.boardID}/lists/${list.listID}/update`, { name: list.name, position: list.position });
        return response.data;
    } catch (error) {
        console.error("Error updating list:", error);
        throw error;
    }
}