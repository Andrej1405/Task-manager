class Employee {
    constructor(id, dataEmployee) {
        this.Id = id;
        this.Surname = dataEmployee.Surname;
        this.Name = dataEmployee.Name;
        this.Position = dataEmployee.Position;
    }
}

let massEmployees = [];

const xhrRequestEmployee = {
    xhrGetEmployees: function() {
        webix.ajax().get('getEmployee').then(function(data) {
            data = data.json();
            
            for ( let i = 0; i < data.length; i++ ) {
                let employee = new Employee(data[i].Id, data[i]);

                massEmployees.push(employee);
                $$('tableActiveEmployees').add(employee);
            }
        });
    },

    xhrAddEmployees: function(valueForm) {
        webix.ajax().post('/employee/add', valueForm).then(function(data){
            let id = data.json();
            const newEmploy = new Employee(id, valueForm);

            massEmployees.push(newEmploy);
            $$('tableActiveEmployees').add(newEmploy);
            webix.message('Сотрудник добавлен');
            return;
        });
    },

    xhrDelEmployees: function(id, idEmployee) {
        idEmployee = {Id: idEmployee};
        webix.ajax().post('/employee/:id/delete', idEmployee ).then(function() {
            for ( let i = 0; i < massEmployees.length; i++ ) {
                if (massEmployees[i].Id == idEmployee.Id) {
                    massEmployees.splice(i, 1);
                }
            }
            webix.message('Сотрудник удалён');
            $$('tableActiveEmployees').remove(id);
            console.log(massEmployees)
        });
    },

    xhrUpdateEmployees: function(valueForm) {
        webix.ajax().post('/employee/:id/update', valueForm).then(function(){
            for ( let i = 0; i < massEmployees.length; i++ ) {
                if (massEmployees[i].Id == valueForm.Id) {
                    massEmployees[i].Surname = valueForm.Surname;
                    massEmployees[i].Name = valueForm.Name;
                    massEmployees[i].Position = valueForm.Position;
                }
            }
            $$('tableActiveEmployees').updateItem(valueForm.id, valueForm);
            return;
        });
    },
};

// Блок сайта, отвечающий за представление списка сотрудников. Состоит из меню и таблицы
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
        columns: [
            { id: 'Surname', header: 'Фамилия', fillspace: true },
            { id: 'Name', header: 'Имя', fillspace: true },
            { id: 'Position', header: 'Должность', fillspace: true }
        ],
        on: {
                'onItemDblClick': showEmployeeCard
            }
        }
    ]
};

//Функция добавления новых сотрудников. При вызоые функции открывается всплывающее окно с формой внутри.
function addNewEmployee() {
    webix.ui({
    view: 'window',
    id: 'newEmployees',
    head: 'Новый сотрудник',
    close: true,
    modal: true,
    position: 'center',
    width: 350,
    body: {
        view: 'form', 
        id: 'newEmployee',
        elementsConfig: {
            labelWidth: 100
        },
        elements: [
            { view: 'text', label: 'Фамилия', name: 'Surname', validate: webix.rules.isNotEmpty },
            { view: 'text', label: 'Имя', name: 'Name', validate: webix.rules.isNotEmpty },
            { view: 'text', label: 'Должность', name: 'Position', validate: webix.rules.isNotEmpty },
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

    // Добавление нового сотрудника в базу данных и таблицу на сайте
    function addEmployee() {
        if ( $$('newEmployee').validate() ) {
            let dataEmployee = $$('newEmployee').getValues();

            for ( let i = 0; i < massEmployees.length; i++ ) {
                if ( (massEmployees[i].Name == dataEmployee.Name) && (massEmployees[i].Surname == dataEmployee.Surname) ) {
                    webix.message('Такой сотрудник уже создан');
                    return;
                }
            }
            xhrRequestEmployee.xhrAddEmployees(dataEmployee);

            $$('newEmployee').clear();
            $$('newEmployees').hide();
            return;
        }
    }

    // Закрытие окна без сохранения изменений.
    function canselEmployee() {
        $$('newEmployee').clear();
        $$('newEmployees').hide();
        return;
    }
}

function showEmployeeCard(id) {
    webix.ui({
        view: 'window',
        id: 'editEmployee',
        head: 'Карточка сотрудника',
        close: true,
        modal: true,
        position: 'center',
        width: 370,
        body: {
            view: 'form', 
            id: 'cardEmployee',
            elementsConfig: {
                labelWidth: 107
              },
            elements: [
                { view: 'text', label: 'Фамилия', name: 'Surname', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Имя', name: 'Name', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Должность', name: 'Position', validate: webix.rules.isNotEmpty },
                { margin: 5, cols: [
                { view: 'button', value: 'Сохранить' , minWidth: 70, css: 'webix_primary', click: saveEmployee },
                { view: 'button', value: 'Удалить сотрудника' , minWidth: 70, css: 'webix_primary', height: 45, click: deleteEmployee},
                { view: 'button', value: 'Закрыть' , minWidth: 70, click: canselEmployee },
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
            idEmployee = dataEmployee.Id;

            xhrRequestEmployee.xhrDelEmployees(id, idEmployee);

            $$('cardEmployee').clear();
            $$('editEmployee').hide();
            
        });
    }

    function saveEmployee() {
        if ( $$('cardEmployee').validate()) {
            let newValues = $$('cardEmployee').getValues();

            if ((values.Surname == newValues.Surname) && (values.Name == newValues.Name) && (values.Position == newValues.Position)) {
                $$('cardEmployee').clear();
                $$('editEmployee').hide();
                return;
            }
            
            xhrRequestEmployee.xhrUpdateEmployees(newValues);

            $$('cardEmployee').clear();
            $$('editEmployee').hide();
        }
    }
    
    function canselEmployee() {
        $$('editEmployee').hide();
        return;
    }
}

document.addEventListener('DOMContentLoaded', xhrRequestEmployee.xhrGetEmployees);