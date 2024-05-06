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

const routers = {
    fGetAll: function() {
        apiUrl = './admin/get-routers'
        fetch(apiUrl) 
        .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok ' + response.statusText);
                }
                return response.json();
            })
            .then(data => {
                table.fGetTable(data);
            })
            .catch(error => {
                console.error('There was a problem with the fetch operation:', error);
            });
        },
        
    fDelRoute: function(id) {
        apiUrl = './admin/delete-route'
        fetch(apiUrl, {
            method: 'POST', // Specifying the HTTP method
            headers: {
                'Content-Type': 'application/json', // Indicating the media type of the resource
            },
            body: JSON.stringify({Id: id}), // Converting the JavaScript object into a JSON string
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return null; // Parsing JSON from the response
        })
        .then(data => {
            // Use the JSON data from the server
            routers.fGetAll();
        })
        .catch(error => {
            // Handle any errors
            console.error('There was a problem with the fetch operation:', error);
        });
    },

    fEditRoute: function(id) {
        console.log(id);
    }
};

const table = {
    template: document.getElementById("table_row"),
    container: document.getElementById("table_body"),
    fGetRow: function(row_data) {
        const instance = document.importNode(this.template.content, true);
        instance.querySelector(".row_id").textContent = row_data.Id;
        instance.querySelector(".row_route").textContent = row_data.Url;
        instance.querySelector(".row_method").textContent = row_data.Method;
        instance.querySelector(".row_code").textContent = row_data.status_code;

        const del_btn = instance.querySelector('.row_del_btn');
        const edit_btn = instance.querySelector('.row_edit_btn');
        
        del_btn.onclick = function() {
            if (confirm('Sure?')) {
                routers.fDelRoute(row_data.Id)
            }
        };

        edit_btn.onclick = function() {
            routers.fEditRoute(row_data.Id)
        };

        return instance;
    },

    fGetTable: function(data) {
        this.fClearTable();
        data.forEach(function(row, index) {
            table.container.appendChild(table.fGetRow(row));
        });
    },

    fClearTable: function() {
        this.container.innerHTML = '';
    }
}


routers.fGetAll();
