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
                modal.fClose();
                table.fGetTable(data);
            })
            .catch(error => {
                console.error('There was a problem with the fetch operation:', error);
            });
        },
        
    fDelRoute: function(id) {
        apiUrl = './admin/delete-route'
        this.fRequest(apiUrl, {Id: id});
    },

    fAddRoute: function(data) {
        apiUrl = './admin/add-route'
        this.fRequest(apiUrl, data);
    },

    fEditRoute: function(id) {
        console.log(id);
    },

    fRequest: function(uri, data) {
        fetch(uri, {
            method: 'POST', // Specifying the HTTP method
            headers: {
                'Content-Type': 'application/json', // Indicating the media type of the resource
            },
            body: JSON.stringify(data), // Converting the JavaScript object into a JSON string
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

const modal = {
    dialog: document.getElementById("modal_dialog"),    
    close_btn: document.getElementById("modal_close_btn"),    
    submit_btn: document.getElementById("modal_submit_btn"),   
    add_btn: document.getElementById("add_route_btn"),   


    title: document.getElementById("modal_title"),
    id: document.getElementById("modal_data_id"),
    route: document.getElementById("modal_data_route"),
    method: document.getElementById("modal_data_method"),
    code: document.getElementById("modal_data_code"),
    headers: document.getElementById("modal_data_headers"),
    body: document.getElementById("modal_data_body"),

    fInit: function() {
        this.close_btn.addEventListener('click', function() {
            modal.fClose()
        });

        this.add_btn.addEventListener('click', function() {
            modal.fOpen(null)
        });

        this.submit_btn.addEventListener('click', function() {
            routers.fAddRoute(modal.fGetData())
        });
    },
    
    fClose: function() {
        if (this.dialog.hasAttribute('open')) {
            this.dialog.removeAttribute('open');
        }
    },

    fSetData: function(data) {
        if (data == null) {
            this.id.value = null;
            this.route.value = '';
            this.method.value = '';
            this.code.value = '';
            this.headers.value = '';
            this.body.value = '';
        } else {
            this.id.value = data.Id;
            this.route.value = data.Url;
            this.method.value = data.Method;
            this.code.value = data.status_code;
            this.headers.value = data.Headers;
            this.body.value = data.Body;
        }
    },

    fGetData: function() {
        data =  {
            Url: this.route.value,
            Method: this.method.value,
            status_code: parseInt(this.code.value, 10),
            Headers: JSON.parse(this.headers.value),
            Body: this.body.value
        };

        if (this.id.value != '') {
            data.Id = parseInt(this.id.value, 10);
        }

        return data;
    },

    fOpen: function(data) {
        this.fSetData(data);
        if (!this.dialog.hasAttribute('open')) {
            this.dialog.setAttribute('open', '');
        }
    }
}

modal.fInit();
routers.fGetAll();
