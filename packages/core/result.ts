// export interface Result<TSuccess, TError> {
//   ok: boolean;
//   success?: TSuccess;
//   err?: TError;
// }

export type Result<TSuccess, TError> =
  | {
      ok: true;
      success: TSuccess;
    }
  | {
      ok: false;
      err: TError;
    };

export function Ok<TSuccess, TError>(
  value: TSuccess
): Result<TSuccess, TError> {
  return {
    ok: true,
    success: value,
  };
}

export function Err<TSuccess, TError>(value: TError): Result<TSuccess, TError> {
  return {
    ok: false,
    err: value,
  };
}
