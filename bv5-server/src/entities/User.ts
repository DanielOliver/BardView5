import { Entity, PrimaryKey, Property, Unique } from '@mikro-orm/core';

@Entity()
export class User {

  @PrimaryKey({ columnType: 'int8' })
  id!: string;

  @Property({ columnType: 'timestamp', length: 6, defaultRaw: `timezone('utc'::text, now())` })
  createdAt!: Date;

  @Property({ columnType: 'timestamp', length: 6, defaultRaw: `timezone('utc'::text, now())` })
  effectiveDate!: Date;

  @Property({ columnType: 'timestamp', length: 6, nullable: true })
  endDate?: Date;

  @Property()
  isActive: boolean = true;

  @Unique({ name: 'user_email_uindex' })
  @Property({ columnType: 'text' })
  email!: string;

  @Property({ columnType: 'text' })
  name!: string;

}
