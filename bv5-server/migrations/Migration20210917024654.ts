import { Migration } from '@mikro-orm/migrations';

export class Migration20210917024654 extends Migration {
  async up(): Promise<void> {
    this.addSql(`create table "user"
(
 id bigserial constraint user_pk primary key,
 created_at timestamp without time zone default (now() at time zone 'utc') not null,
 email text not null,
 name text not null
);`);
    this.addSql('create unique index user_email_uindex on "user" (email);');
  }
}
