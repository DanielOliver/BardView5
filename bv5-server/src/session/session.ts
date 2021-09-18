export class SessionScope {
  userId!: string;

  constructor(props: Partial<SessionScope>) {
    Object.assign(this, props);
  }
}
