import { Entity, PrimaryKey, Property, Unique } from '@mikro-orm/core';

@Entity({ collection: 'User' })
export class User {
  @PrimaryKey()
  id!: number;

  @Unique({ name: 'User_email_key' })
  @Property({ columnType: 'text' })
  email!: string;

  @Property({ columnType: 'text', nullable: true })
  name?: string;
}
