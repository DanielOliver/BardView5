export type Res<TSuccess, TError> =
  | {
      ok: true;
      value: TSuccess;
    }
  | {
      ok: false;
      err?: TError;
    };

export class Result<TSuccess, TError> {
  constructor(
    public readonly ok: boolean,
    public readonly value?: TSuccess,
    public readonly err?: TError
  ) {}

  map<TNext>(callback: (v: TSuccess) => TNext): Result<TNext, TError> {
    if ((!this.ok && !!this.err) || !this.value) {
      return new Result<TNext, TError>(false, undefined, this.err);
    }
    return new Result<TNext, TError>(true, callback(this.value), undefined);
  }

  as(): Res<TSuccess, TError> {
    if (this.ok && this.value) {
      return {
        ok: true,
        value: this.value,
      };
    }
    return {
      ok: false,
      err: this.err,
    };
  }

  chain(callback: (v: TSuccess) => void, fail?: (e: TError) => void) {
    if (this.ok && this.value) {
      callback(this.value);
    } else if (fail && this.err) {
      fail(this.err);
    }
  }
}

export function Ok<TSuccess, TError>(
  value: TSuccess
): Result<TSuccess, TError> {
  return new Result<TSuccess, TError>(true, value, undefined);
}

export function Err<TSuccess, TError>(value: TError): Result<TSuccess, TError> {
  return new Result<TSuccess, TError>(false, undefined, value);
}
