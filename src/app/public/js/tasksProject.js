class Task {
    constructor() {

    }
}

function showProject() {
    let employeesInvolved = {},
        idActiveProject = $$('tableActiveProjects').getSelectedId().id,
        tasksOfProject,
        number;

    for ( let i = 0; i < massProjects.length; i++ ) {
        if (massProjects[i].id == idActiveProject) {
            if (massProjects[i].tasks.length == 0) {
                number = 0;
                break;
            } else {
                let numberTask = massProjects[i].tasks;
                number = numberTask[numberTask.length - 1].id;
            }
        }
    }
    
    for ( let i = 0; i < massProjects.length; i++ ) {
        if (massProjects[i].id == idActiveProject) {
            tasksOfProject = massProjects[i].tasks;
            break;
        }
    }
    
    for ( let i = 0; i < employees.length; i++ ) {
        employeesInvolved[employees[i].id] = `${employees[i].surname} ${employees[i].name}`;
    }

    function addTask() {
        ++number;
        $$('tableTasksProject').add({'id' : number});
    }

    // function deleteTask() {
    //     let index = $$('tableTasksProject').getSelectedId();
    //     $$('tableTasksProject').remove(index);
    // }

    function updateTasks() {
        for (let i = 0; i < projects.length; i++) {
            if (projects[i].id == idActiveProject) {
                projects[i].tasks = [];
                let dataTasks = $$('tableTasksProject').data.pull;

                for (let obj in dataTasks) {
                    projects[i].tasks.push(dataTasks[obj]);
                }
                break;
            }
        }
    }

    // function completeProject() {
    //     webix.confirm({
    //         title: 'Закрытие проекта',
    //         text: 'Уверены, что хотите завершить проект?'
    //     }).then( () => {

    //         for ( let i = 0; i < projects.length; i++ ) {
    //             if ( projects[i].id == idActiveProject ) {
    //                 let complProject = Object.assign({}, projects[i]);
    //                 complitedProjects.push(complProject);
    //                 projects.splice(i, 1);
    //             }
    //         }

    //         $$('tableActiveProjects').remove(idActiveProject);
    //         $$('tasksProject').hide();
    //         webix.message('Проект закрыт');
    //     });
    // }

    function canselTasks() {
        $$('tasksProject').hide();
        return;
    }

    webix.ui({
        view: 'popup',
        id: 'tasksProject',
        position: 'center',
        body: {
            rows: [
                { view:"toolbar", elements:[
                    { view: 'button', value: 'Добавить задачу', click: addTask, minWidth: 65, css: 'webix_primary' },
                    //{ view: 'button', value: 'Удалить задачу', click: deleteTask, minWidth: 65, css: 'webix_primary' }
                ]},

                {
                view: 'datatable',
                id: 'tableTasksProject',
                select: true,
                width: 1150,
                height: 450,
                editable: true,
                editaction: 'dblclick',
                scroll: 'y',
                autoConfig: true,
                data: tasksOfProject,
                columns: [
                    { id: 'id' },
                    { id: 'task', header: 'Задача', fillspace: true, editor: 'text' },
                    { id: 'designatedEmployee', header: 'Назначенный сотрудник', fillspace: true, editor: 'select', options: employeesInvolved },
                    { id: 'hours', header: 'Часы', fillspace: true, editor: 'text' },
                    { id: 'hoursSpent', header: 'Потраченные часы', fillspace: true, editor: 'text' },
                    { id: 'statusTask', header: 'Статус', fillspace: true, editor: 'select', options: status },
                    //{ id:'', template: "<input class = 'delTask' type = 'button' value = 'Удалить задачу'>", fillspace: true},
                ],
                },
                {
                    cols: [
                        { view: 'button', value: 'Сохранить' , click: updateTasks, minWidth: 65, css: 'webix_primary' , },
                        //{ view: 'button', value: 'Завершить проект' , click: completeProject, minWidth: 65, css: 'webix_primary' , },
                        { view: 'button', value: 'Закрыть' , click: canselTasks, minWidth: 65, css: 'webix_primary'  }
                    ]
                }
            ]
        }
    }).show();

    // $$('tableTasksProject').on_click.delTask = function(e, id) {
    //     $$('tableTasksProject').remove(id)
    //     webix.message('Задача удалена');
    // };
}