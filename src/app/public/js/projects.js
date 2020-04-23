class Project {
    constructor(id, dataProject) {
        this.Id = id;
        this.Name = dataProject.Name;
        this.Date = dataProject.Date;
    }
}

let massProjects = [];
    
const xhrRequestProject = {
    xhrGetProjects: function() {
        webix.ajax().get('getProject').then(function(data) {
            data = data.json();
            
            for ( let i = 0; i < data.length; i++ ) {
                let project = new Project(data[i].Id, data[i]);

                massProjects.push(project);
                $$('tableActiveProjects').add(project);
            }
        });
    },

    xhrUpdateProject: function(valueForm) {
        console.log(valueForm)
        webix.ajax().post('/project/:id/update', valueForm).then(function(){
            for ( let i = 0; i < massEmployees.length; i++ ) {
                if (massProjects[i].Id == valueForm.Id) {
                    massProjects[i].Name = valueForm.Name;
                    massProjects[i].Date = valueForm.Date;
                }
            }
            $$('tableActiveProjects').updateItem(valueForm.id, valueForm);
            return;
        });
    },

    xhrAddProject: function(valueForm) {
        webix.ajax().post('/project/add', valueForm).then(function(data) {
            let id = data.json();
            const newProject = new Project(id, valueForm);

            massProjects.push(newProject);
            $$('tableActiveProjects').add(newProject);
            webix.message('Проект создан');
            return;
        });
    },

    xhrDelProject: function(id, idProject) {
        idProject = {Id: idProject};
        webix.ajax().post('/project/:id/delete', idProject ).then(function() {
            for ( let i = 0; i < massProjects.length; i++ ) {
                if (massProjects[i].Id == idProject.Id) {
                    massProjects.splice(i, 1);
                }
            }

            webix.message('Проект удалён');
            $$('tableActiveProjects').remove(id);
        });
    }
};

// Основная компонента проектов. Состоит из меню для управления и таблицы для вывода данных.
let activeProjects = {
    rows: [
    {
        cols: [
            {
                view: 'button', id: 'addButton', value: 'Добавить проект', autowidth: true, 
                on: {
                    'onItemClick': addProject
                }
            },

            {
                view: 'button', id: 'editButton', value: 'Изменить проект', autowidth: true, 
                on: {
                    'onItemClick': editProject
                }
            },

            {
                view: 'button', id: 'delProject', value: 'Удалить проект', autowidth: true, 
                on: {
                    'onItemClick': deleteProject
                }
            },

            // Потенциальная кнопка для завершения проекта, перенесения его в отдельную таблицу с возможностью дальнейшего просмотра информации по проекту
            // {
            //     view: 'button', id: 'completeButton', value: 'Завершённые проекты', autowidth: true, 
            //     on: {
            //         'onItemClick': completedProject
            //     }
            // }
        ]
    }, 
        {
            view: 'datatable',
            id: 'tableActiveProjects',
            sort: 'multi',
            scroll: 'y',
            //tooltip:true,
            select: true,
            autoConfig: true,
            columns: [
                { id: 'Name', header: 'Название проекта', fillspace: true },
                { id: 'Date', header: 'Дата создания', fillspace: true }
            ],
            on: {
                'onItemDblClick': showProject
            }
        }
    ]
};

// Функция для добавления нового проекта. При вызове появляется всплывающее окно с формой для ввода данных внутри.
function addProject() {
    webix.ui({
        view: 'window',
        id: 'newProjects',
        head: 'Новый проект',
        modal: true,
        close: true,
        position: 'center',
        body: {
            view: 'form', 
            id: 'newProject',
            width: 300,
            elements: [
                { view: 'text', label: 'Проект', name: 'Name', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Создан', labelWidth: 81, name: 'Date', validate: webix.rules.isNotEmpty },
                { margin: 5, cols: [
                { view: 'button', value: 'Создать' , minWidth: 65, css: 'webix_primary', click: addNewProject },
                { view: 'button', value: 'Отмена', minWidth: 65, click: canselAddProject }
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

    function addNewProject() {
        if ( $$('newProject').validate() ) {
            let dataProject = $$('newProject').getValues();

            for ( let i = 0; i < massProjects.length; i++ ) {
                if (dataProject.Name == massProjects[i].Name) {
                    webix.message('Такой проект уже существует');
                    return;
                }
            }

            xhrRequestProject.xhrAddProject(dataProject);

            $$('newProject').clear();
            $$('newProjects').hide();
            return;
        }
    }

    function canselAddProject() {
        $$('newProject').clear();
        $$('newProjects').hide();
        return;
    }

}

function editProject() {
    if ($$('tableActiveProjects').getSelectedItem() !== undefined) {
        const id = $$('tableActiveProjects').getSelectedId();
        webix.ui({
            view: 'window',
            id: 'editProject',
            head: 'Изменение проекта',
            modal: true,
            close: true,
            position: 'center',
            body: {
                view: 'form', 
                id: 'cardProject',
                width: 300,
                elements: [
                    { view: 'text', label: 'Проект', name: 'Name', validate: webix.rules.isNotEmpty },
                    { view: 'text', label: 'Создан', labelWidth: 81, name: 'Date', validate: webix.rules.isNotEmpty },
                    { margin: 5, cols: [
                    { view: 'button', value: 'Сохранить' , minWidth: 65, css: 'webix_primary', click: editProject },
                    { view: 'button', value: 'Отмена', minWidth: 65, click: canselEditProject }
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
        
        let values = $$('tableActiveProjects').getItem(id);
        $$('cardProject').setValues(values);

        function editProject() {
            if ( $$('cardProject').validate()) {
                let newValues = $$('cardProject').getValues();
    
                if ((values.Name == newValues.Name) && (values.Date == newValues.Date)) {
                    $$('cardProject').clear();
                    $$('editProject').hide();
                    return;
                }
                
                xhrRequestProject.xhrUpdateProject(newValues);
    
                $$('cardProject').clear();
                $$('editProject').hide();
            }
        }

        function canselEditProject() {
            $$('editProject').hide();
            return;
        }
    }
}

// Удаление проекта. Для вызова нужно щёлкнуть на нужный проект, после этого нажать кнопку "Удалить проект"
function deleteProject() {
    if ($$('tableActiveProjects').getSelectedItem() !== undefined) {
    webix.confirm({
            title: 'Проект будет удалён',
            text: 'Уверены, что хотите удалить проект?'
        }).then( () => {
            let dataProject = $$('tableActiveProjects').getSelectedItem(),
                idProject = dataProject.Id,
                id = dataProject.id;

            xhrRequestProject.xhrDelProject(id, idProject);
        });
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

// Потенциальный функционал для отслеживания уже завершённых проектов
// function completedProject() {
//     webix.ui({
//         view: 'popup',
//         id: 'completeProjects',
//         position: 'center',
//         body: {
//             rows: [
//                 {
//                 view: 'datatable',
//                 id: 'completedProject',
//                 scroll: 'y',
//                 autoConfig: true,
//                 data: complitedProjects,
//                 columns: [
//                     { id: 'Name', header: 'Название проекта', fillspace: true },
//                     { id: 'Date', header: 'Дата создания', fillspace: true }
//                 ],
//                 },

//                 {
//                 view: 'button', 
//                 id: 'canselWindowsCompletedProject', 
//                 value: 'Закрыть', 
//                 css: 'webix_primary', 
//                 inputWidth: 300,
//                 on: {
//                     'onItemClick': function () {
//                         $$('completeProjects').hide();
//                         }
//                     }
//                 }
//             ]
//         }
//     }).show();
// }

document.addEventListener('DOMContentLoaded', xhrRequestProject.xhrGetProjects);