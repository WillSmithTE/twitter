const BASE_URL = process.env.NODE_ENV === 'development' ?
    'http://localhost:5000' : 'someprodurl'

export const api = {
    getAll: async () => {
        const response = await fetch(`${BASE_URL}/api/areas`);
        const jsonResponse = await response.json();
        return jsonResponse;
    }
}