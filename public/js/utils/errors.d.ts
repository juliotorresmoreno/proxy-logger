

export type ErrorCode =
    | 400
    | 401
    | 402
    | 403
    | 404
    | 405
    | 406
    | 500
    | 501
    | 502
    | 503
    | 504
    | 505
    | 506;

export interface ResponseError {
    message: string;
}