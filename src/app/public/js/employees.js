class Employee {
    constructor(id, dataEmployee) {
        this.Id = id;
        this.Surname = dataEmployee.Surname;
        this.Name = dataEmployee.Name;
        this.Position = dataEmployee.Position;
    }
}

// Массив для хранения существующих сотрудников.
const massEmployees = [];

// Объект, содеражщий асинхронные запросы.
const xhrRequestEmployee = {
    xhrGetEmployees: function() {
        webix.ajax().get('getEmployee').then(function(data) {
            data = data.json();
            
            for ( let i = 0; i < data.length; i++ ) {
                let employee = new Employee(data[i].Id, data[i]);

                massEmployees.push(employee);
                $$('tableActiveEmployees').add(employee);
            }
        }).finally( () => xhrRequestTask.xhrGetTask() )
          .catch( error => showError(error) );
    },

    xhrAddEmployees: function(valueForm) {
        webix.ajax().post('/employee/add', valueForm).then(function(data){
            let id = data.json();
            const newEmploy = new Employee(id, valueForm);

            massEmployees.push(newEmploy);
            $$('tableActiveEmployees').add(newEmploy);
            webix.message('Сотрудник добавлен');
            return;
        }).catch( error => showError(error) );
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
        }).catch( error => showError(error) );
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
        }).catch( error => showError(error) );
    },
};

function showError(err) {
    webix.message({
        text: err,
        type: 'error', 
        expire: 10000,
        id: 'message4'
    });
}

// Блок сайта, отвечающий за представление списка сотрудников.
let activeEmployees = {
    rows: [
        {
        cols: [
                { view: 'button', id: 'addEmployee', value: 'Добавить сотрудника', autowidth: true, click: addNewEmployee }
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
            { view: 'text', label: 'Фамилия', name: 'Surname', invalidMessage: 'Введите фамилию сотрудника', validate: webix.rules.isNotEmpty },
            { view: 'text', label: 'Имя', name: 'Name', invalidMessage: 'Введите имя сотрудника', validate: webix.rules.isNotEmpty },
            { view: 'text', label: 'Должность', name: 'Position', invalidMessage: 'Введите должность сотрудника', validate: webix.rules.isNotEmpty },
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

    // Добавление нового сотрудника.
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

    // Закрытие окна без сохранения.
    function canselEmployee() {
        $$('newEmployee').clear();
        $$('newEmployees').hide();
        return;
    }
}

// Окно просмотра карточки сотрудника для возможного редактировния данных и удаления сотрудника.
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
                { view: 'text', label: 'Фамилия', name: 'Surname', invalidMessage: 'Введите фамилию сотрудника', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Имя', name: 'Name', invalidMessage: 'Введите имя сотрудника', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Должность', name: 'Position', invalidMessage: 'Введите должность сотрудника', validate: webix.rules.isNotEmpty },
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
    
    // Удаление сотрудника.
    function deleteEmployee() {
        webix.confirm({
            title: 'Сотрудник будет удалён',
            text: 'Уверены, что хотите удалить сотрудника?'
        }).then( () => {
            let dataEmployee = $$('cardEmployee').getValues();
            idEmployee = dataEmployee.Id;

            for ( let i = 0; i < massTasks.length; i++ ) {
               if ( massTasks[i].DesignatedEmployee == `${dataEmployee.Surname} ${dataEmployee.Name}`) {
                    textMessage = 'Данный сотрудник назначен на задачу. Удалить можно после снятия сотрудника со всех задач';
                    webix.message( { type:"error", text: textMessage } );

                    $$('cardEmployee').clear();
                    $$('editEmployee').hide();
                    return;
               }
            }
            
            xhrRequestEmployee.xhrDelEmployees(id, idEmployee);

            $$('cardEmployee').clear();
            $$('editEmployee').hide();
            
        });
    }

    // Сохранение внесённых изменения.
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
    
    // Закрытие карточки сотрудника без сохранения изменений.
    function canselEmployee() {
        $$('editEmployee').hide();
        return;
    }
}

// Асинхронный запрос к базе для получения списка сотрудников после окончания построения DOM-дерева.
document.addEventListener('DOMContentLoaded', xhrRequestEmployee.xhrGetEmployees);