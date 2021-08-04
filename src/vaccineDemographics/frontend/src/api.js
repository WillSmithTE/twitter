import database from './database.json';
 
const BASE_URL = process.env.NODE_ENV === 'development' ?
    'http://localhost:5000' : 'someprodurl'

const databaseFileName = 'database.json';
export const api = {
    getAll: async () => {
        if (isEmpty(database)) {
            console.error('getting new')
            const response = await fetch(`${BASE_URL}/api/areas`);
            const jsonResponse = await response.json();
            save(jsonResponse);
            return jsonResponse;
        } else {
            console.error('used saved')
            return database;
        }
    }
}

function isEmpty(database) {
    return Object.keys(database).length === 0;
}

function save(jsonData) {
    const fileData = JSON.stringify(jsonData);
    const blob = new Blob([fileData], { type: "text/plain" });
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.download = `${databaseFileName}`;
    link.href = url;
    link.click();
}
