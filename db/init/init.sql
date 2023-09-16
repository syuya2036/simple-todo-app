CREATE TABLE tasks (
  id serial PRIMARY KEY,
  title varchar(255) NOT NULL,
  completed boolean NOT NULL
);
--テーブルにデータを挿入
INSERT INTO tasks (title, completed) VALUES ('task1', false);
INSERT INTO tasks (title, completed) VALUES ('task2', true);