import { Entity, PrimaryKey } from '@mikro-orm/core';

@Entity()
export class RoleAction {

  @PrimaryKey({ columnType: 'text' })
  name!: string;

}
