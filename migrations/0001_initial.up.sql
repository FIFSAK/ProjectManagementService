DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_priority') THEN
            CREATE TYPE task_priority AS ENUM ('low', 'medium', 'high');
        END IF;
    END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
            CREATE TYPE task_status AS ENUM ('new', 'in_progress', 'done');
        END IF;
    END $$;


create table if not exists users(
    id serial primary key,
    name varchar(255),
    email varchar(255),
    registration_date date default current_date,
    role varchar(255)
);

create table IF NOT EXISTS projects(
    id serial primary key,
    title varchar(255),
    description text,
    creation_date date default current_date,
    completion_date date,
    manager_id int references users(id)
);


create table IF NOT EXISTS tasks(
    id serial primary key,
    title varchar(255),
    description text,
    priority task_priority,
    status task_status,
    responsible_user_id int references users(id),
    project_id int references projects(id),
    creation_date date default current_date,
    completion_date date
    );