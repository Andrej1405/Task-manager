class Employee {
    constructor(dataEmployee) {
        if (employees.length == 0) {
            this.id = 1;
        } else {
            this.id = employees[employees.length - 1].id + 1;
        }
        this.name = dataEmployee.name;
        this.surname = dataEmployee.surname;
        this.position = dataEmployee.position;
    }
}

let activeEmployees = {
    rows: [
        {
        cols: [
                {
                view: 'button', id: 'addEmployee', value: 'Добавить сотрудника', autowidth: true,
                on: {
                    'onItemClick': addNewEmployee
                    }
                }
             ]
        }, 

        {
        view: 'datatable',
        id: 'tableActiveEmployees',
        sort:'multi',
        scroll: 'y',
        select: true,
        columnWidth: 470,
        autoConfig: true,
        data: employees,
        columns: [
            { id: 'surname', header: 'Фамилия', fillspace: true },
            { id: 'name', header: 'Имя', fillspace: true },
            { id: 'position', header: 'Должность', fillspace: true }
        ],
        on: {
                'onItemDblClick': showEmployeeCard
            }
        }
    ]
};

function addNewEmployee() {
    webix.ui({
    view: 'popup',
    id: 'newEmployees',
    position: 'center',
    body: {
        view: 'form', 
        id: 'newEmployee',
        width: 300,
        elements: [
            { view: 'text', label: 'Фамилия', name: 'surname', validate: webix.rules.isNotEmpty },
            { view: 'text', label: 'Имя', name: 'name', validate: webix.rules.isNotEmpty },
            { view: 'text', label: 'Должность', labelWidth: 81, name: 'position', validate: webix.rules.isNotEmpty },
            { margin: 5, cols: [
            { view: 'button', value: 'Сохранить' , minWidth: 65, css: 'webix_primary', click: addEmployee},
            { view: 'button', value: 'Отменить', minWidth: 65, click: canselEmployee }
        ]}
    ],
    on: {
        onValidationError: function (key, obj) {
            let textMessage = 'Некорретно введена информация';
            webix.message( { type:"error", text: textMessage } );
            }
        }
    }
}).show();

    function addEmployee() {
        if ( $$('newEmployee').validate() ) {
            let dataEmployee = $$('newEmployee').getValues();

            for (let i = 0; i < employees.length; i++) {
                if ( (employees[i].name == dataEmployee.name) && (employees[i].surname == dataEmployee.surname) ) {
                    webix.message('Такой сотрудник уже создан');
                    //$$('newEmployee').clear();
                    return;
                }
            }

            const newEmploy = new Employee(dataEmployee);        
            employees.push(newEmploy);

            $$('tableActiveEmployees').add(newEmploy);
            webix.message('Сотрудник добавлен');
            $$('newEmployee').clear();
            $$('newEmployees').hide();
            return;
        }
    }

    function canselEmployee() {
        $$('newEmployee').clear();
        $$('newEmployees').hide();
        return;
    }
}

function showEmployeeCard(id) {
    webix.ui({
        view: 'popup',
        id: 'editEmployee',
        position: 'center',
        body: {
            view: 'form', 
            id: 'cardEmployee',
            width: 330,
            elements: [
                { view: 'text', label: 'Фамилия', name: 'surname', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Имя', name: 'name', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Должность', labelWidth: 81, name: 'position', validate: webix.rules.isNotEmpty },
                { margin: 5, cols: [
                { view: 'button', value: 'Сохранить' , minWidth: 65, css: 'webix_primary', click: saveEmployee },
                { view: 'button', value: 'Удалить сотрудника' , minWidth: 65, css: 'webix_primary', height: 45, click: deleteEmployee}
            ]}
        ],
            on: {
                onValidationError: function (key, obj) {
                    textMessage = 'Некорретно введена информация';
                    webix.message( { type:"error", text: textMessage } );
                    }
                }
        }
    }).show();

    let values = $$('tableActiveEmployees').getItem(id);
    $$('cardEmployee').setValues(values);

        function deleteEmployee() {
        webix.confirm({
            title: 'Сотрудник будет удалён',
            text: 'Уверены, что хотите удалить сотрудника?'
        }).then( () => {
            let dataEmployee = $$('cardEmployee').getValues();
            idEmployee = dataEmployee.id;

            for (let i = 0; i < employees.length; i++) {
                if (employees[i].id == idEmployee) {
                    employees.splice(i, 1);
                }
            }

            webix.message('Сотрудник удалён');
            $$('tableActiveEmployees').remove(id)
            $$('cardEmployee').clear();
            $$('editEmployee').hide();
            
        })/*.finally( () => $$('editEmployee').hide())*/;
    }

    function saveEmployee() {
        if ($$('cardEmployee').validate()) {
            let newValues = $$('cardEmployee').getValues();
        
            $$('tableActiveEmployees').updateItem(newValues.id, newValues);
            $$('cardEmployee').clear();
            $$('editEmployee').hide();
        }
    }             
}