class Project {
    constructor(dataProject) {
        if (massProjects.length == 0) {
            this.Id = 1;
        } else {
            this.Id = massProjects[massProjects.length - 1].Id + 1;
        }

        this.Name = dataProject.Name;
        this.Date = dataProject.Date;
        this.tasks = [];
    }
}

let massProjects = [],
    requestProjects = new XMLHttpRequest();
    

function xhrGetProjects() {
    requestProjects.open('GET', 'getProject');

    requestProjects.onload = () => {
        try {
            let jsonProjects = JSON.parse(requestProjects.responseText);

            for (let i = 0; i < jsonProjects.length; i++) {
                massProjects.push(jsonProjects[i]);
                $$('tableActiveProjects').add(massProjects[i]);
            }
        } catch (err) {
            new Error(err);
        }
    };

    requestProjects.send();
    console.log(massProjects);
    return massProjects;
}



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
            select: true,
            autoConfig: true,
            // data: webix.ajax('project').then(function(data) {
            //     this.ClearAll();
            //     return data = data.json();
            // }),
            // url:{
            //     $proxy: true,
            //     load: function(view,params){
            //         return webix.ajax("public/json/1.json");
            //     },
            //     save: function(){
            //         // ...
            //     }
            // },
            // url: function(params) {
            //     return webix.ajax('getProject');
            // },
            data: xhrGetProjects(),
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
        view: 'popup',
        id: 'newProjects',
        position: 'center',
        body: {
            view: 'form', 
            id: 'newProject',
            width: 300,
            elements: [
                { view: 'text', label: 'Проект', name: 'Name', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Создан', labelWidth: 81, name: 'Date', validate: webix.rules.isNumber },
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

            for (let i = 0; i < massProjects.length; i++) {
                if (dataProject.Name == massProjects[i].Name) {
                    webix.message('Такой проект уже существует');
                    return;
                }
            }

            const project = new Project(dataProject);
            massProjects.push(project);
            
            $$('tableActiveProjects').add(project);
            webix.message('Новый проект создан');
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

// Удаление проекта. Для вызова нужно щёлкнуть на нужный проект, после этого нажать кнопку "Удалить проект"
function deleteProject() {
    if ($$('tableActiveProjects').getSelectedId() !== undefined) {
    webix.confirm({
            title: 'Проект будет удалён',
            text: 'Уверены, что хотите удалить проект?'
        }).then( () => {
            let dataProject = $$('tableActiveProjects').getSelectedId(),
            idProject = dataProject.Id;

            for ( let i = 0; i < massProjects.length; i++ ) {
                if ( massProjects[i].Id == IdProject ) {
                    let deleteProject = Object.assign({}, massProjects[i]);
                    deletedProject.push(deleteProject);
                    massProjects.splice(i, 1);
                }
            }

            $$('tableActiveProjects').remove(idProject);
            webix.message('Проект удалён');
        });
    }
}

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