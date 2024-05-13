const routers = {
    data: null,
    fGetAll: function() {
        apiUrl = './admin/get-routes'
        this.fRequestGet(apiUrl, function() {
            modal.fClose();
            table.fGetTable(routers.data);
        })
    },

    fGetLog: function() {
        apiUrl = './admin/log-route'
        this.fRequestGet(apiUrl, function() {
            modal.fRefreshLog(routers.data);
        })
    },
        
    fDelRoute: function(id) {
        apiUrl = './admin/delete-route'
        this.fRequestPost(apiUrl, {Id: id}, function() {routers.fGetAll();});
    },

    fSaveRoute: function(data) {
        apiUrl = './admin/save-route'
        this.fRequestPost(apiUrl, data, function() {routers.fGetAll();});
    },

    fGetRoute: function(id) {
        apiUrl = './admin/get-route'
        this.fRequestPost(apiUrl, {Id: id}, function() {
            modal.fOpen(routers.data);
        });
    },

    fRequestGet: function(uri, f) {
        fetch(uri) 
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            return response.json();
        })
        .then(data => {
            routers.data = data;
            f();
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
    },

    fRequestPost: function(uri, data, f) {
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
            if (response.status != 200) {
                return null;
            }

            return response.json(); // Parsing JSON from the response
        })
        .then(data => {
            // Use the JSON data from the server
            routers.data = data;
            f();
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
            routers.fGetRoute(row_data.Id)
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

    log_dialog: document.getElementById("modal_dialog_log"),
    log_area: document.getElementById("modal_log_area"),
    log_close_btn: document.getElementById("modal_log_close_btn"),
    log_refresh_btn: document.getElementById("modal_log_refresh_btn"),
    log_open_btn: document.getElementById("open_log_btn"),
    log_template: document.getElementById("log_row"),

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
            routers.fSaveRoute(modal.fGetData())
        });

        this.log_open_btn.addEventListener('click', function() {
            modal.fLogOpen()
        });

        this.log_close_btn.addEventListener('click', function() {
            modal.fLogClose()
        });

        this.log_refresh_btn.addEventListener('click', function() {
            routers.fGetLog()
        });
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
            this.headers.value = JSON.stringify(data.Headers);
            this.body.value = data.Body;
        }
    },

    fGetData: function() {
        try {
            data =  {
                Url: this.route.value,
                Method: this.method.value,
                status_code: parseInt(this.code.value, 10),
                Headers: JSON.parse(this.headers.value || '{}'),
                Body: this.body.value
            };
        } catch {
            this.fMarkHeaders()
            throw new Error('Headers JSON is wrong');
        }

        if (this.id.value != '') {
            data.Id = parseInt(this.id.value, 10);
        }

        return data;
    },

    fMarkHeaders: function() {
        if (!this.headers.hasAttribute('aria-invalid')) {
            this.headers.setAttribute('aria-invalid', 'true');
        }
    },

    fUnmarkHeaders: function() {
        if (this.headers.hasAttribute('aria-invalid')) {
            this.headers.removeAttribute('aria-invalid');
        }
    },

    fOpen: function(data) {
        this.fUnmarkHeaders();
        this.fSetData(data);
        if (!this.dialog.hasAttribute('open')) {
            this.dialog.setAttribute('open', '');
        }
    },

    fClose: function() {
        if (this.dialog.hasAttribute('open')) {
            this.dialog.removeAttribute('open');
        }
    },

    fLogOpen: function(data) {
        routers.fGetLog()
        if (!this.log_dialog.hasAttribute('open')) {
            this.log_dialog.setAttribute('open', '');
        }
    },

    fLogClose: function() {
        if (this.log_dialog.hasAttribute('open')) {
            this.log_dialog.removeAttribute('open');
        }
    },

    fGetRow: function(row_data) {
        const instance = document.importNode(this.log_template.content, true);
        instance.querySelector(".log_row_date").textContent = row_data.Date;
        instance.querySelector(".log_row_url").textContent = row_data.Url;
        instance.querySelector(".log_row_method").textContent = row_data.Method;

        return instance;
    },

    fRefreshLog: function(data) {
        this.fClearLog();
        data.forEach(function(row, index) {
            modal.log_area.appendChild(modal.fGetRow(row));
        });
    },

    fClearLog: function() {
        this.log_area.innerHTML = '';
    }
}

modal.fInit();
routers.fGetAll();
