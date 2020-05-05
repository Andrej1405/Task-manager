class Task {
    constructor(idTask, dataTask) {
        this.Id_project = dataTask.Id_project;
        this.IdTask = idTask;
        this.Task = dataTask.Task;
        this.DesignatedEmployee = dataTask.DesignatedEmployee;
        this.Hours = dataTask.Hours;
        this.HoursSpent = dataTask.HoursSpent;
        this.StatusTask = dataTask.StatusTask;
        this.TaskDescription = dataTask.TaskDescription;
    }
}
// Массив для хранения существующих задач.
const massTasks = [];

// Объект, содержащий асинхронные запросы.
const xhrRequestTask = {
    xhrGetTask: function() {
        webix.ajax().get('getTask').then(function(data) {
            data = data.json();
            for ( let i = 0; i < data.length; i++ ) {
                let task = new Task(data[i].IdTask, data[i]);
                massTasks.push(task);
            }
            return;
        });
    },

    xhrAddTask: function(idActiveProject, valueForm) {
        let id_project = idActiveProject;
        valueForm.Id_project = id_project;
        
        webix.ajax().post('/task/add', valueForm).then(function(id) {
            let idTask = id.json();
            let task = new Task(idTask, valueForm);
            massTasks.push(task);
            
            $$('tableTasksProject').add(task);
            webix.message('Задача добавлена');
            return;
        });
    },
    ///////добавить случаи, когда запрос не удался и вернулся отрицательный ответ
    xhrUpdateTask: function(valueForm) {
        webix.ajax().post('/task/:id/update', valueForm).then(function(data) {
            for ( let i = 0; i < massEmployees.length; i++ ) {
                if ( valueForm.DesignatedEmployee == massEmployees[i].Id) {
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
                    massTasks[i].TaskDescription = valueForm.TaskDescription;
                }
            }
            
            $$('tableTasksProject').updateItem(valueForm.id, valueForm);
            return;
        });
    },
};

// Основное окно, отображающее задачи выбранного проекта.
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
                    tooltip: true,
                    select: true,
                    editaction: 'dblclick',
                    scroll: 'y',
                    autoConfig: true,
                    columns: [
                        { id: 'Task', header: 'Задача', fillspace: true, sort: 'string' },
                        { id: 'DesignatedEmployee', header: 'Назначенный сотрудник', fillspace: true, sort: 'string' },
                        { id: 'Hours', header: 'Оценка задачи в часах', fillspace: true, sort: 'int' },
                        { id: 'HoursSpent', header: 'Потраченные часы', fillspace: true, sort: 'int' },
                        { id: 'StatusTask', header: 'Статус', fillspace: true, sort: 'string' },
                        { id: 'TaskDescription', header: 'Описание задачи', fillspace: true },
                        ],
                        on: {
                            'onItemDblClick': showTask
                    },

                },

                {
                    cols: [
                        { view: 'button', value: 'Добавить задачу', click: addTask, minWidth: 65, css: 'webix_primary' },
                        { view: 'button', value: 'Показать / скрыть закрытые задачи', click: hideShow, minWidth: 65, css: 'webix_primary' },
                        { view: 'button', value: 'Вернуться на главную' , click: canselTasks, minWidth: 65, css: 'webix_primary' }
                    ]
                }
            ]
        }
    }).show();
    // Получение Id выбранного проекта для отображения задач по этому проекту.
    const activeProject = $$('tableActiveProjects').getSelectedItem(),
          idActiveProject = activeProject.Id;

    // Массив, хранящий в себе задачи со статусом "Закрыто" и массив сотрудников для удобного вывода Фамилии-Имени сотрудника в форме и окне задач.
    let massHide = [],
        employeesInvolved = [];
    
    for ( let i = 0; i < massEmployees.length; i++ ) {
        let objEmployee = {};

        objEmployee.Id = massEmployees[i].Id;
        objEmployee.value = `${massEmployees[i].Surname} ${massEmployees[i].Name}`;

        employeesInvolved.push(objEmployee);
    }

    for ( let i = 0; i < massTasks.length; i++ ) {
        if ( massTasks[i].Id_project == idActiveProject ) {
            $$('tableTasksProject').add(massTasks[i]);

            if (massTasks[i].StatusTask == 'Закрыто') {
                massHide.push(massTasks[i]);

                let item = $$('tableTasksProject').getItem(massTasks[i].id);
                item.hidden = true;

                $$('tableTasksProject').updateItem(massTasks[i].id, item);
                $$('tableTasksProject').filter(function(obj) {
                    return !obj.hidden;
                });
            }
        }
    }

    function canselTasks() {
        $$('tasksProject').hide();
        return;
    }

    // Отображение / скрытие задач со статусом "Закрыто".
    function hideShow() {
        for ( let i = 0; i < massHide.length; i++ ) {  
            let item = $$('tableTasksProject').getItem(massHide[i].id);
            item.hidden = item.hidden ? false : true;

            $$('tableTasksProject').updateItem(massHide[i].id, item);
            $$('tableTasksProject').filter(function(obj) {
                return !obj.hidden;
            });
        }
    }

    // Окно добавления новой задачи.
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
                    { view: 'text', label: 'Задача', name: 'Task', invalidMessage: 'Введите краткое наименование задачи', validate: webix.rules.isNotEmpty },
                    { view: 'richselect', label: 'Назначенный сотрудник', name: 'DesignatedEmployee', invalidMessage: 'Назначьте сотрудника на задачу', options: employeesInvolved, validate: webix.rules.isNotEmpty },
                    { view: 'text', label: 'Оценка задачи в часах', name: 'Hours', invalidMessage: 'Введите число больше нуля', validate: webix.rules.isNumber },
                    { view: 'text', label: 'Потраченные часы', name: 'HoursSpent' },
                    { view: 'richselect', label: 'Статус', name: 'StatusTask', invalidMessage: 'Выберите статус задачи', options: [{Id: 1, value: 'Назначено'}], validate: webix.rules.isNotEmpty },
                    { view: 'text', label: 'Описание задачи', name: 'TaskDescription', invalidMessage: 'Введите описание задачи', validate: webix.rules.isNotEmpty },
                    { margin: 5, cols: [
                    { view: 'button', value: 'Сохранить' , minWidth: 65, css: 'webix_primary', click: saveTask},
                    { view: 'button', value: 'Отменить', minWidth: 65, click: canselSaveTask }
                ]}
            ],
            rules: {
                Hours: function(value) {
                    return value > 0;
                }
            },
            on: {
                onValidationError: function (key, obj) {
                    let textMessage = 'Некорретно введена информация';
                    webix.message( { type:"error", text: textMessage } );
                    }
                }
            }
        }).show();

        // Сохранение новой задачи
        function saveTask() {
            if ( $$('newTask').validate() ) {
                let dataTask = $$('newTask').getValues();
                
                xhrRequestTask.xhrAddTask(idActiveProject, dataTask);

                $$('newTask').clear();
                $$('addNewTask').hide();
                return;
            }
        }

        // Закрытие формы без сохранения.
        function canselSaveTask() {
            $$('newTask').clear();
            $$('addNewTask').hide();
            return;
        }
        
    }
    
    // Получение информации по задаче для вывода её в форму.
    function showTask(id) {
        let values = $$('tableTasksProject').getItem(id),
            status = values.StatusTask,
            actualstatus;

        /* В зависимости от выбранного статуса будет определённый перечень статусов, доступных для дальнейшего выбора в задаче. Добавлено для последовательного
         изменения статусов задач.
        */
        switch(status) {
            case 'Назначено': 
                actualstatus = [
                    {value: 'Назначено'},
                    {value: 'В работе'},
                    {value: 'Отменено'}
                ];
            break;
            case 'В работе':
                actualstatus = [
                        {value: 'В работе'},
                        {value: 'Выполнено'},
                        {value: 'Отменено'}
                ];
            break;
            case 'Выполнено':
                actualstatus = [
                    {value: 'Выполнено'},
                    {value: 'Закрыто'}
                ];
            break;
            case 'Закрыто':
                actualstatus = [
                    {value: 'Закрыто'}
                ];
            break;
            case 'Отменено':
                actualstatus = [
                    {value: 'Отменено'},
                    {value: 'Закрыто'}
                ];
            break;
        }
        // Форма редактирования задачи.
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
                    { view: 'text', label: 'Оценка задачи в часах', name: 'Hours', invalidMessage: 'Введите число больше нуля', validate: webix.rules.isNumber },
                    { view: 'text', label: 'Потраченные часы', name: 'HoursSpent', invalidMessage: 'Введите потраченные часы' },
                    { view: 'richselect', label: 'Статус', name: 'StatusTask', options: actualstatus },
                    { view: 'text', label: 'Описание задачи', name: 'TaskDescription', validate: webix.rules.isNotEmpty},
                    { margin: 5, cols: [
                    { view: 'button', value: 'Сохранить' , minWidth: 65, css: 'webix_primary', click: saveEditTask},
                    { view: 'button', value: 'Отменить', minWidth: 65, click: canselEditTask }
                ]}
            ],
                rules: {
                    Hours: function(value) {
                        return value > 0;
                    },
                    HoursSpent: function(value) {
                        const valueTask = $$('cardTask').getValues();
                        
                        if (valueTask.StatusTask == 'Выполнено') {
                            return value > 0;
                        }
                        return value;
                    }
                },
                on: {
                    onValidationError: function (key, obj) {
                        textMessage = 'Некорретно введена информация';
                        webix.message( { type:"error", text: textMessage } );
                        $$('cardTask').clearValidation();
                    }
                }
            }
        }).show();
        
        $$('cardTask').setValues(values);

        function saveEditTask() {
            if ( $$('cardTask').validate() ) {
                let dataTask = $$('cardTask').getValues();
                
                xhrRequestTask.xhrUpdateTask(dataTask);
                // Если статус задачи "Закрыто", то задача добавляется в массив зада, которые скрываются / отображаются нажатием кнопки.
                if (dataTask.StatusTask == 'Закрыто') {
                    for ( let i = 0; i < massTasks.length; i++ ) {
                        if (massTasks[i].IdTask == dataTask.IdTask) {
                            massHide.push(massTasks[i]);
                        }
                    }
                }
                
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