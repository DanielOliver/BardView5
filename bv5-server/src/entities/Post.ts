import { Entity, ManyToOne, PrimaryKey, Property } from '@mikro-orm/core';
import { User } from './User';

@Entity({ collection: 'Post' })
export class Post {

  @PrimaryKey()
  id!: number;

  @Property({ columnType: 'text' })
  title!: string;

  @Property({ columnType: 'text', nullable: true })
  content?: string;

  @Property({ nullable: true })
  published?: boolean = false;

  @ManyToOne({ entity: () => User, fieldName: 'authorId', onUpdateIntegrity: 'cascade', onDelete: 'set null', nullable: true })
  authorId?: User;

}
