DROP TABLE IF EXISTS todo_lists;
CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    date timestamptz not null,
    status varchar(255) not null
);

DROP INDEX IF EXISTS idx_date_pagination;
CREATE INDEX idx_date_pagination ON todo_lists (date, id);

DROP INDEX IF EXISTS idx_date_pagination;
CREATE INDEX idx_status_pagination ON todo_lists (status, id);

