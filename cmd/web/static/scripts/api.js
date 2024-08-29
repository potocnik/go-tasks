const apiUrl = 'http://localhost:10000/v1/api';

var _api = {
    getAll: function () { return httpGet(apiUrl + '/tasks') }
}

function httpGet (endpoint) {
    // Make a GET request
    fetch(endpoint)
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
}