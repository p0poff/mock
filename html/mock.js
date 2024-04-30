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

// Usage example
const appData = [
    { name: 'John', age: 25, gender: 'Male' },
    { name: 'Jane', age: 30, gender: 'Female' },
    { name: 'Bob', age: 35, gender: 'Male' }
];

const tableConstructor = new TableConstructor(appData);
tableConstructor.fillTable();