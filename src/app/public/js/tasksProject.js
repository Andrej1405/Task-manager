class Task {
    constructor(idTask, dataTask) {
        this.Id_project = dataTask.Id_project;
        this.IdTask = idTask;
        this.Task = dataTask.Task;
        this.DesignatedEmployee = dataTask.DesignatedEmployee;
        this.Hours = dataTask.Hours;
        this.HoursSpent = dataTask.HoursSpent;
        this.StatusTask = dataTask.StatusTask;
    }
}

let massStatus = [
    {Id: 1, value: 'Назначено'},
    {Id: 2, value: 'В работе'},
    {Id: 3, value: 'Выполнено'},
    {Id: 4, value: 'Отменено'}
],
    massTasks = [];

const xhrRequestTask = {
    xhrGetTask: function() {
        webix.ajax().get('getTask').then(function(data) {
            data = data.json();
            for ( let i = 0; i < data.length; i++ ) {
                for ( let j = 0; j < massEmployees.length; j++ ) {
                    if (massEmployees[j].Id == data[i].DesignatedEmployee) {
                        data[i].DesignatedEmployee = `${massEmployees[j].Surname} ${massEmployees[j].Name}`;
                    }
                }
                let task = new Task(data[i].IdTask, data[i]);
                massTasks.push(task);
            }
            return;
        });
    },

    xhrAddTask: function(idActiveProject, valueForm) {
        let id_project = idActiveProject;
        valueForm.Id_project = id_project;

        for ( let i = 0; i < massEmployees.length; i++ ) {
            if( valueForm.DesignatedEmployee == `${massEmployees[i].Surname} ${massEmployees[i].Name}`) {
                valueForm.DesignatedEmployee = massEmployees[i].Id;
            }
        }

        webix.ajax().post('/task/add', valueForm).then(function(id) {
            for ( let i = 0; i < massEmployees.length; i++ ) {
                if( valueForm.DesignatedEmployee == massEmployees[i].Id) {
                    valueForm.DesignatedEmployee = `${massEmployees[i].Surname} ${massEmployees[i].Name}`;
                }
            }

            let idTask = id.json();
            
            let task = new Task(idTask, valueForm);
            massTasks.push(task);

            $$('tableTasksProject').add(task);
            
            webix.message('Задача добавлена');
            return;
        });
    },

    // xhrDelTask: function() {
    
    // },

    xhrUpdateTask: function(valueForm) {
        for ( let i = 0; i < massEmployees.length; i++ ) {
            if( valueForm.DesignatedEmployee == `${massEmployees[i].Surname} ${massEmployees[i].Name}`) {
                valueForm.DesignatedEmployee = massEmployees[i].Id;
            }
        }
        console.log(valueForm)
        webix.ajax().post('/task/:id/update', valueForm).then(function(data) {
            for ( let i = 0; i < massEmployees.length; i++ ) {
                if( valueForm.DesignatedEmployee == massEmployees[i].Id) {
                    valueForm.DesignatedEmployee = `${massEmployees[i].Surname} ${massEmployees[i].Name}`;
                }
            }
            for ( let i = 0; i < massTasks.length; i++ ) {
                if (massTasks[i].IdTask == valueForm.IdTask) {
                    massTasks[i].Task = valueForm.Task;
                    massTasks[i].DesignatedEmployee = valueForm.DesignatedEmployee;
                    massTasks[i].Hours = valueForm.Hours;
                    massTasks[i].HoursSpent = valueForm.HoursSpent;
                    massTasks[i].StatusTask = valueForm.StatusTask;
                }
            }
            $$('tableTasksProject').updateItem(valueForm.id, valueForm);
            return;
        });
    },
};

function showProject() {
    webix.ui({
        view: 'window',
        head: 'Задачи проекта',
        id: 'tasksProject',
        close: true,
        fullscreen: true,
        body: {
            rows: [
                {
                    view: 'datatable',
                    id: 'tableTasksProject',
                    select: true,
                    editable: true,
                    editaction: 'dblclick',
                    scroll: 'y',
                    autoConfig: true,
                    columns: [
                        { id: 'Task', header: 'Задача', fillspace: true },
                        { id: 'DesignatedEmployee', header: 'Назначенный сотрудник', fillspace: true },
                        { id: 'Hours', header: 'Часы', fillspace: true, },
                        { id: 'HoursSpent', header: 'Потраченные часы', fillspace: true },
                        { id: 'StatusTask', header: 'Статус', fillspace: true },
                        ],
                        on: {
                            'onItemDblClick': showTask
                    },

                },

                {
                    cols: [
                        { view: 'button', value: 'Добавить задачу', click: addTask, minWidth: 65, css: 'webix_primary' },
                        { view: 'button', value: 'Вернуться на главную' , click: canselTasks, minWidth: 65, css: 'webix_primary'  }
                    ]
                }
            ]
        }
    }).show();

    const activeProject = $$('tableActiveProjects').getSelectedItem(),
          idActiveProject = activeProject.Id;

    let massTasksOfProject = [],
        employeesInvolved = [];
    
    for ( let i = 0; i < massEmployees.length; i++ ) {
        let objEmployee = {};

        objEmployee.Id = massEmployees[i].Id; //?
        objEmployee.value = `${massEmployees[i].Surname} ${massEmployees[i].Name}`;

        employeesInvolved.push(objEmployee);
    }

    for ( let i = 0; i < massTasks.length; i++ ) {
        if (massTasks[i].Id_project == idActiveProject) {
            massTasksOfProject.push(massTasks[i]);
            $$('tableTasksProject').add(massTasks[i]);
        }
    }

    function canselTasks() {
        $$('tasksProject').hide();
        return;
    }
    //////////////////////////////////////////////////////////////
    function addTask() {
        webix.ui({
            view: 'window',
            id: 'addNewTask',
            head: 'Новая задача',
            close: true,
            modal: true,
            position: 'center',
            width: 500,
            body: {
                view: 'form', 
                id: 'newTask',
                elementsConfig: {
                    labelWidth: 180
                },
                elements: [
                    { view: 'text', label: 'Задача', name: 'Task', validate: webix.rules.isNotEmpty },
                    { view: 'richselect', label: 'Назначенный сотрудник', name: 'DesignatedEmployee', options: employeesInvolved, validate: webix.rules.isNotEmpty },
                    { view: 'text', label: 'Часы', name: 'Hours', validate: webix.rules.isNumber },
                    { view: 'text', label: 'Потраченные часы', name: 'HoursSpent' },
                    { view: 'richselect', label: 'Статус', name: 'StatusTask', value: 1, options: massStatus },
                    { margin: 5, cols: [
                    { view: 'button', value: 'Сохранить' , minWidth: 65, css: 'webix_primary', click: saveTask},
                    { view: 'button', value: 'Отменить', minWidth: 65, click: canselSaveTask }
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
    
        
        
        function saveTask() {
            if ( $$('newTask').validate() ) {
                let dataTask = $$('newTask').getValues();
                
                xhrRequestTask.xhrAddTask(idActiveProject, dataTask);

                $$('newTask').clear();
                $$('addNewTask').hide();
                return;
            }
        }
    
        function canselSaveTask() {
            $$('newTask').clear();
            $$('addNewTask').hide();
            return;
        }
        
    }
    //////////////////////////////////////////////
    function showTask(id) {
        webix.ui({
            view: 'window',
            id: 'showTask',
            head: 'Редактирование задачи',
            close: true,
            modal: true,
            width: 500,
            position: 'center',
            body: {
                view: 'form', 
                id: 'cardTask',
                elementsConfig: {
                    labelWidth: 180
                },
                elements: [
                    { view: 'text', label: 'Задача', name: 'Task', validate: webix.rules.isNotEmpty },
                    { view: 'richselect', label: 'Назначенный сотрудник', name: 'DesignatedEmployee', options: employeesInvolved, validate: webix.rules.isNotEmpty },
                    { view: 'text', label: 'Часы', name: 'Hours', validate: webix.rules.isNumber },
                    { view: 'text', label: 'Потраченные часы', name: 'HoursSpent' },
                    { view: 'richselect', label: 'Статус', name: 'StatusTask', options: massStatus },
                    { margin: 5, cols: [
                    { view: 'button', value: 'Сохранить' , minWidth: 65, css: 'webix_primary', click: saveEditTask},
                    { view: 'button', value: 'Отменить', minWidth: 65, click: canselEditTask }
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
        
        let values = $$('tableTasksProject').getItem(id);
        $$('cardTask').setValues(values);

        function saveEditTask() {
            if ( $$('cardTask').validate() ) {
                let dataTask = $$('cardTask').getValues();
                
                xhrRequestTask.xhrUpdateTask(dataTask);

                $$('cardTask').clear();
                $$('showTask').hide();
                return;
            }
            return;
        }
    
        function canselEditTask() {
            $$('showTask').hide();
            return;
        }
    }
}
//?
document.addEventListener('DOMContentLoaded', () => setTimeout(function() {
    xhrRequestTask.xhrGetTask(); 
}, 300));