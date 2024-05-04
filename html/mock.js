function TableConstructor(app) {
    this.app = app;

    this.fillTable = function() {
        // Get the table element from the DOM
        const table = document.getElementById('myTable');

        // Clear the table
        table.innerHTML = '';

        // Loop through the app data and create table rows
        for (let i = 0; i < this.app.length; i++) {
            const row = table.insertRow();

            // Create table cells and populate them with data
            const cell1 = row.insertCell();
            const cell2 = row.insertCell();
            const cell3 = row.insertCell();

            cell1.innerHTML = this.app[i].name;
            cell2.innerHTML = this.app[i].age;
            cell3.innerHTML = this.app[i].gender;
        }
    };
}

const server_data = {
    data: [],
    fGetAll: function() {
        apiUrl = './admin/get-routers'
        return fetch(apiUrl)  // Возвращает Promise
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok ' + response.statusText);
                }
                return response.json();
            })
            .then(data => {
                this.data = data;
                return data;  // Важно возвращать данные для дальнейшей обработки
            })
            .catch(error => {
                this.data = [];
                console.error('There was a problem with the fetch operation:', error);
                return [];  // Возвращаем пустой массив при ошибке
            });
    }
};

const routers = {
    data: [],
    fUpdateData: function() {
        server_data.fGetAll().then(data => {
            this.data = data;  // Обновляем данные только после завершения промиса
            console.log(this.data);  // Перемещаем вывод в консоль сюда для корректного отображения данных
        });
    }
}

routers.fUpdateData();
