class Project {
    constructor(dataProject) {
        if (projects.length == 0) {
            this.id = 1;
        } else {
            this.id = projects[projects.length - 1].id + 1;
        }

        this.projectName = dataProject.projectName;
        this.date = dataProject.date;
        this.tasks = [];
    }
}

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

            {
                view: 'button', id: 'completeButton', value: 'Завершённые проекты', autowidth: true, 
                on: {
                    'onItemClick': completedProject
                }
            }
        ]
    }, 
        {
            view: 'datatable',
            id: 'tableActiveProjects',
            sort: 'multi',
            scroll: 'y',
            select: true,
            autoConfig: true,
            data: projects,
            columns: [
                { id: 'projectName', header: 'Название проекта', fillspace: true },
                { id: 'date', header: 'Дата создания', fillspace: true }
            ],
            on: {
                'onItemDblClick': showProject
            }
        }
    ]
};

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
                { view: 'text', label: 'Проект', name: 'projectName', validate: webix.rules.isNotEmpty },
                { view: 'text', label: 'Создан', labelWidth: 81, name: 'date', validate: webix.rules.isNumber },
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

            for (let i = 0; i < projects.length; i++) {
                if (dataProject.projectName == projects[i].projectName) {
                    webix.message('Такой проект уже существует');
                    return;
                }
            }

            const project = new Project(dataProject);
            projects.push(project);
            
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

function deleteProject() {
    if ($$('tableActiveProjects').getSelectedId() !== undefined) {
    webix.confirm({
            title: 'Проект будет удалён',
            text: 'Уверены, что хотите удалить проект?'
        }).then( () => {
            let dataProject = $$('tableActiveProjects').getSelectedId(),
            idProject = dataProject.id;

            for (let i = 0; i < projects.length; i++) {
                if ( projects[i].id == idProject ) {
                    let deletProject = Object.assign({}, projects[i]);
                    deletedProject.push(deletProject);
                    projects.splice(i, 1);
                }
            }

            $$('tableActiveProjects').remove(idProject);
            webix.message('Проект удалён');
        });
    }
}

function completedProject() {
    webix.ui({
        view: 'popup',
        id: 'completeProjects',
        position: 'center',
        body: {
            rows: [
                {
                view: 'datatable',
                id: 'completedProject',
                scroll: 'y',
                autoConfig: true,
                data: complitedProjects,
                columns: [
                    { id: 'projectName', header: 'Название проекта', fillspace: true },
                    { id: 'date', header: 'Дата создания', fillspace: true }
                ],
                },

                {
                view: 'button', 
                id: 'canselWindowsCompletedProject', 
                value: 'Закрыть', 
                css: 'webix_primary', 
                inputWidth: 300,
                on: {
                    'onItemClick': function () {
                        $$('completeProjects').hide();
                        }
                    }
                }
            ]
        }
    }).show();
}