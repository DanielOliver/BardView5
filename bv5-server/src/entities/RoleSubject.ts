import { Entity, PrimaryKey } from '@mikro-orm/core';

@Entity()
export class RoleSubject {

  @PrimaryKey({ columnType: 'text' })
  name!: string;

}
