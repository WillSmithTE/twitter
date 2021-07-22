const BASE_URL = process.env.NODE_ENV === 'development' ?
    'http://localhost:5000' : 'someprodurl'

const cache = {};

export const api = {
    getAll: async () => {
        if (cache.getAll) {
            return cache.getAll;
        } else {
            const response = await fetch(`${BASE_URL}/api/years`);
            const jsonResponse = await response.json();
            cache.getAll = jsonResponse;
            return jsonResponse;
        }
    }
}