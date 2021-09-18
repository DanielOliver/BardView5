import { Entity, PrimaryKey, Property } from '@mikro-orm/core';

@Entity()
export class RoleSubject {

  @PrimaryKey({ columnType: 'text' })
  name!: string;

  @Property({ columnType: 'timestamp', length: 6, defaultRaw: `timezone('utc'::text, now())` })
  createdAt!: Date;

}