import pg from 'k6/x/pg';

const client = pg.newClient("postgres://localhost:5432/chrisg?sslmode=disable");

export function setup() {
	client.exec("create table if not exists all_the_data (id SERIAL UNIQUE NOT NULL, code VARCHAR(10) NOT NULL, article TEXT, name TEXT NOT NULL, department VARCHAR(4) NOT NULL);");
}

export default function() {
	pg.insert(client, "insert into all_the_data (code, article, name, department) select left(md5(i::text), 10),md5(random()::text),md5(random()::text),left(md5(random()::text), 4)from generate_series(1, 1000000) s(i);");
}

export function teardown() {
	client.exec("drop table all_the_data;");
	pg.close(client);
}
